package ohauthcat

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

// AuthorizationOptions configures an SSO authorization URL.
type AuthorizationOptions struct {
	RedirectURI string
	// State is the CSRF state. Leave it empty and a random one is generated
	// and returned in AuthorizationResult.State.
	State        string
	Organization string
	Connection   string
	Domain       string
	LoginHint    string
}

// AuthorizationResult holds the authorization URL and the PKCE verifier to keep.
type AuthorizationResult struct {
	URL          string
	CodeVerifier string
	// State is the state sent in the URL. Store it in the user's session and
	// compare it in constant time on the callback before exchanging the code.
	State string
}

// ExchangeOptions configures an authorization code exchange.
type ExchangeOptions struct {
	Code         string
	CodeVerifier string
	RedirectURI  string
}

// RestartSignIn asks callers to start a fresh authorization flow.
type RestartSignIn struct{ APIError }

func (e *RestartSignIn) Error() string { return "Restart the sign-in flow" }
func (e *RestartSignIn) Unwrap() error {
	if e == nil {
		return nil
	}
	return &e.APIError
}

func boundedSSO(value string, maximum int, allowEmpty bool) bool {
	return utf8.ValidString(value) && len(value) <= maximum && !strings.ContainsAny(value, "\r\n\x00") && (allowEmpty || strings.TrimSpace(value) != "")
}

// PKCEChallenge implements RFC 7636 S256 using a verifier of 43 to 128 characters.
func PKCEChallenge(verifier string) (string, error) {
	if len(verifier) < 43 || len(verifier) > 128 {
		return "", requestFailure("request_invalid")
	}
	for _, character := range verifier {
		unreserved := character >= 'A' && character <= 'Z' || character >= 'a' && character <= 'z' || character >= '0' && character <= '9' || strings.ContainsRune("._~-", character)
		if !unreserved {
			return "", requestFailure("request_invalid")
		}
	}
	digest := sha256.Sum256([]byte(verifier))
	return base64.RawURLEncoding.EncodeToString(digest[:]), nil
}

func validSSODomain(value string) bool {
	value = strings.TrimSuffix(strings.ToLower(strings.TrimSpace(value)), ".")
	if len(value) > 253 {
		return false
	}
	return regexp.MustCompile(`^(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?$`).MatchString(value)
}

// GetAuthorizationURL creates a fresh PKCE verifier and broker authorization URL.
func (r *SSOResources) GetAuthorizationURL(input AuthorizationOptions) (AuthorizationResult, error) {
	var result AuthorizationResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	c := r.client
	if c.config.ClientID == "" || !boundedSSO(input.RedirectURI, 2048, false) || !boundedSSO(input.State, 1024, true) {
		return result, requestFailure("request_invalid")
	}
	redirect, err := url.Parse(input.RedirectURI)
	if err != nil || redirect.Host == "" || (redirect.Scheme != "http" && redirect.Scheme != "https") || redirect.User != nil || strings.Contains(input.RedirectURI, "#") {
		return result, requestFailure("request_invalid")
	}
	for _, entry := range []struct{ value, schema string }{{input.Organization, "OrganizationId"}, {input.Connection, "ConnectionId"}} {
		if entry.value != "" {
			encoded, _ := json.Marshal(entry.value)
			definition, exists := c.contract.Schemas[entry.schema]
			if !exists || !matchesJSON(encoded, definition, c.contract.Schemas, false) {
				return result, requestFailure("request_invalid")
			}
		}
	}
	if input.Domain != "" && (!boundedSSO(strings.TrimSpace(input.Domain), 253, false) || !validSSODomain(input.Domain)) {
		return result, requestFailure("request_invalid")
	}
	if input.LoginHint != "" {
		hint := strings.TrimSpace(input.LoginHint)
		parts := strings.Split(hint, "@")
		if !boundedSSO(hint, 320, false) || len(parts) != 2 || parts[0] == "" || strings.ContainsAny(hint, " \t\r\n<>\"") || !validSSODomain(parts[1]) {
			return result, requestFailure("request_invalid")
		}
	}
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return result, requestFailure("request_invalid")
	}
	result.CodeVerifier = base64.RawURLEncoding.EncodeToString(bytes)
	result.State = input.State
	if result.State == "" {
		stateBytes := make([]byte, 32)
		if _, err := rand.Read(stateBytes); err != nil {
			return AuthorizationResult{}, requestFailure("request_invalid")
		}
		result.State = base64.RawURLEncoding.EncodeToString(stateBytes)
	}
	challenge, _ := PKCEChallenge(result.CodeVerifier)
	origin, _, _, _ := c.standby.snapshot(true)
	endpoint, _ := url.Parse(origin + "/sso/authorize")
	query := url.Values{"client_id": {c.config.ClientID}, "redirect_uri": {input.RedirectURI}, "response_type": {"code"}, "state": {result.State}, "code_challenge": {challenge}, "code_challenge_method": {"S256"}}
	for _, entry := range []struct{ name, value string }{{"organization", input.Organization}, {"connection", input.Connection}, {"domain", input.Domain}, {"login_hint", input.LoginHint}} {
		if entry.value != "" {
			query.Set(entry.name, entry.value)
		}
	}
	endpoint.RawQuery = query.Encode()
	result.URL = endpoint.String()
	return result, nil
}

// GetProfileAndToken performs one exchange; retrying consumed codes can revoke a prior grant.
func (r *SSOResources) GetProfileAndToken(ctx context.Context, input ExchangeOptions) (ExchangeSSOCodeResult, error) {
	var zero ExchangeSSOCodeResult
	if r == nil || r.client == nil {
		return zero, requestFailure("client_closed")
	}
	if r.client.config.ClientID == "" || !strings.HasPrefix(r.client.config.APIKey, "sk_") || !boundedSSO(input.Code, 4096, false) || !boundedSSO(input.RedirectURI, 2048, false) {
		return zero, requestFailure("request_invalid")
	}
	if _, err := PKCEChallenge(input.CodeVerifier); err != nil {
		return zero, err
	}
	result, err := r.ExchangeSSOCode(ctx, ExchangeSSOCodeArgs{Body: SSOTokenRequest{GrantType: "authorization_code", ClientID: r.client.config.ClientID, Code: input.Code, CodeVerifier: input.CodeVerifier, RedirectURI: input.RedirectURI}})
	var api *APIError
	if errors.As(err, &api) && (api.Code == CodeInvalidGrant || api.Code == CodeFlowInterruptedRetry || api.Code == CodePKCEVerificationFailed) {
		return zero, &RestartSignIn{APIError: *api}
	}
	return result, err
}

// TrustedTokenAccessOptions configures ExchangeTrustedTokenForAccess.
type TrustedTokenAccessOptions struct {
	// SubjectToken is the JWT or ID token from a registered trusted issuer.
	SubjectToken string
	// SubjectTokenType defaults to urn:ietf:params:oauth:token-type:jwt.
	SubjectTokenType string
	// ClientID is the publishable key the access token is for; it defaults
	// to the client's configured ClientID.
	ClientID string
}

// TrustedTokenAccess is a trusted-token profile with the OhAuthCat access
// token issued for it. The token's sub is the profile ID (tprof_...), not an
// end-user ID; its ext_iss and ext_sub claims carry the original issuer and
// subject.
type TrustedTokenAccess struct {
	Profile         TrustedTokenProfile
	AccessToken     string
	IssuedTokenType string
	TokenType       string
	ExpiresIn       int64
}

// Token type URNs of RFC 8693; they are identifiers, not credentials.
const (
	subjectTokenTypeJWT     = "urn:ietf:params:oauth:token-type:jwt"          //nolint:gosec // RFC 8693 identifier, not a credential.
	subjectTokenTypeIDToken = "urn:ietf:params:oauth:token-type:id_token"     //nolint:gosec // RFC 8693 identifier, not a credential.
	issuedTokenTypeAccess   = "urn:ietf:params:oauth:token-type:access_token" //nolint:gosec // RFC 8693 identifier, not a credential.
)

// ErrTrustedTokenAccessMissing reports an exchange answer without its access token.
var ErrTrustedTokenAccessMissing = errors.New("token exchange returned no access token")

// ExchangeTrustedTokenForAccess exchanges a trusted third-party token for its
// profile and an OhAuthCat access token whose audience is the client ID.
// Exchanged identities are not billable profiles or monthly active users.
func (r *SSOResources) ExchangeTrustedTokenForAccess(ctx context.Context, input TrustedTokenAccessOptions) (TrustedTokenAccess, error) {
	var zero TrustedTokenAccess
	if r == nil || r.client == nil {
		return zero, requestFailure("client_closed")
	}
	clientID := strings.TrimSpace(input.ClientID)
	if clientID == "" {
		clientID = r.client.config.ClientID
	}
	tokenType := input.SubjectTokenType
	if tokenType == "" {
		tokenType = subjectTokenTypeJWT
	}
	if clientID == "" || !strings.HasPrefix(r.client.config.APIKey, "sk_") || !boundedSSO(input.SubjectToken, 16384, false) || !boundedSSO(clientID, 256, false) ||
		(tokenType != subjectTokenTypeJWT && tokenType != subjectTokenTypeIDToken) {
		return zero, requestFailure("request_invalid")
	}
	profile, err := r.ExchangeTrustedToken(ctx, ExchangeTrustedTokenArgs{Body: TrustedTokenExchange{
		GrantType:        "urn:ietf:params:oauth:grant-type:token-exchange",
		SubjectToken:     input.SubjectToken,
		SubjectTokenType: tokenType,
		ClientID:         &clientID,
	}})
	if err != nil {
		return zero, err
	}
	if profile.AccessToken == nil || profile.ExpiresIn == nil {
		return zero, ErrTrustedTokenAccessMissing
	}
	return TrustedTokenAccess{ //nolint:gosec // Returns the token the API issued to the caller; nothing is hardcoded.
		Profile:         profile,
		AccessToken:     *profile.AccessToken,
		IssuedTokenType: issuedTokenTypeAccess,
		TokenType:       "Bearer",
		ExpiresIn:       *profile.ExpiresIn,
	}, nil
}
