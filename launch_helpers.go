package ohauthcat

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"
)

// DefaultLoginOrigin is the hosted login origin of the production cloud.
const DefaultLoginOrigin = "https://login.ohauthcat.com"

// DefaultSignalsURL is the SSF transmitter base of the production cloud.
const DefaultSignalsURL = "https://signals.ohauthcat.com"

// SSF delivery methods.
const (
	PushDelivery = "urn:ietf:rfc:8935"
	PollDelivery = "urn:ietf:rfc:8936"
)

func loginOrigin(value string) (string, error) {
	if value == "" {
		value = DefaultLoginOrigin
	}
	parsed, err := url.Parse(value)
	if err != nil || secureURL(value) != nil || (parsed.Path != "" && parsed.Path != "/") || parsed.RawQuery != "" || parsed.Fragment != "" {
		return "", errors.New("login origin must be an https origin")
	}
	return parsed.Scheme + "://" + parsed.Host, nil
}

// HostedLoginOptions build a hosted login URL.
type HostedLoginOptions struct {
	ClientID       string // the environment's publishable key
	RedirectURI    string
	State          string
	Nonce          string
	OrganizationID string
	Lang           string
	LoginOrigin    string
}

// HostedLoginURL returns the hosted login URL (code + PKCE S256) and the code
// verifier to keep server-side for the code exchange.
func HostedLoginURL(options HostedLoginOptions) (loginURL, codeVerifier string, err error) {
	if strings.TrimSpace(options.ClientID) == "" || strings.TrimSpace(options.RedirectURI) == "" || strings.TrimSpace(options.State) == "" {
		return "", "", errors.New("client ID, redirect URI and state are required")
	}
	if redirect, parseErr := url.Parse(options.RedirectURI); parseErr != nil || redirect.Fragment != "" || !redirect.IsAbs() {
		return "", "", errors.New("redirect URI must be absolute without a fragment")
	}
	base, err := loginOrigin(options.LoginOrigin)
	if err != nil {
		return "", "", err
	}
	raw := make([]byte, 32)
	if _, err := rand.Read(raw); err != nil {
		return "", "", errors.New("random source failed")
	}
	verifier := base64.RawURLEncoding.EncodeToString(raw)
	challenge := sha256.Sum256([]byte(verifier))
	query := url.Values{
		"response_type":         {"code"},
		"client_id":             {options.ClientID},
		"redirect_uri":          {options.RedirectURI},
		"state":                 {options.State},
		"code_challenge":        {base64.RawURLEncoding.EncodeToString(challenge[:])},
		"code_challenge_method": {"S256"},
	}
	for key, value := range map[string]string{"nonce": options.Nonce, "organization_id": options.OrganizationID, "lang": options.Lang} {
		if value != "" {
			query.Set(key, value)
		}
	}
	return base + "/user_management/authorize?" + query.Encode(), verifier, nil
}

// EndUserVerifierOptions configure end-user access token verification.
type EndUserVerifierOptions struct {
	ClientID      string // the publishable key: the tokens' audience
	EnvironmentID string
	LoginOrigin   string
	HTTPClient    *http.Client
	Now           func() time.Time
}

// EndUserTokenVerifier verifies hosted-login access tokens.
type EndUserTokenVerifier struct {
	jwt *JWTVerifier
	env string
}

// NewEndUserTokenVerifier verifies ES256 at+jwt tokens with the login origin
// as issuer, the publishable key as audience and the environment.
func NewEndUserTokenVerifier(options EndUserVerifierOptions) (*EndUserTokenVerifier, error) {
	issuer, err := loginOrigin(options.LoginOrigin)
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(options.EnvironmentID) == "" {
		return nil, errors.New("environment ID is required")
	}
	verifier, err := NewJWTVerifier(JWTVerifierConfig{Issuer: issuer, Audience: []string{options.ClientID}, JWKSURI: issuer + "/user_management/jwks",
		Types: []string{"at+jwt"}, RequiredClaims: []string{"sub", "sid", "env"}, HTTPClient: options.HTTPClient, Now: options.Now})
	if err != nil {
		return nil, err
	}
	return &EndUserTokenVerifier{jwt: verifier, env: options.EnvironmentID}, nil
}

// Verify returns the claims of a valid end-user access token.
func (v *EndUserTokenVerifier) Verify(ctx context.Context, token string) (JWTClaims, error) {
	if v == nil {
		return nil, errToken
	}
	claims, err := v.jwt.Verify(ctx, token)
	if err != nil || claims.String("env") != v.env {
		return nil, errToken
	}
	for _, name := range []string{"roles", "permissions"} {
		if _, present := claims[name]; present {
			if _, ok := claims.Strings(name); !ok {
				return nil, errToken
			}
		}
	}
	return claims, nil
}

// PermissionDecision is a local permission check result.
type PermissionDecision int

// Permission decisions. Unknown means the token's permission list was
// truncated: call the authorization check API.
const (
	PermissionDenied PermissionDecision = iota
	PermissionGranted
	PermissionUnknown
)

// HasPermission checks verified claims for a permission.
func HasPermission(claims JWTClaims, permission string) PermissionDecision {
	granted, _ := claims.Strings("permissions")
	if slices.Contains(granted, permission) {
		return PermissionGranted
	}
	if claims.Bool("permissions_truncated") {
		return PermissionUnknown
	}
	return PermissionDenied
}

// HasRole checks verified claims for a role (roles, or the single role).
func HasRole(claims JWTClaims, role string) bool {
	if roles, ok := claims.Strings("roles"); ok {
		return slices.Contains(roles, role)
	}
	return claims.String("role") == role
}

// MachineTokenError is a failed client credentials request.
type MachineTokenError struct{ Code string }

func (e *MachineTokenError) Error() string {
	return "Machine token request failed (" + e.Code + ")"
}

// MachineTokenOptions configure a client credentials token source.
type MachineTokenOptions struct {
	ClientID     string
	ClientSecret string
	Scope        string
	Resource     []string // RFC 8707 resource indicators
	LoginOrigin  string
	HTTPClient   *http.Client
	Timeout      time.Duration
	RefreshSkew  time.Duration // default 60 s
	Now          func() time.Time
}

// MachineTokenSource caches client credentials tokens until shortly before
// expiry; concurrent callers share one request.
type MachineTokenSource struct {
	options  MachineTokenOptions
	endpoint string
	mu       sync.Mutex
	token    string
	expires  time.Time
}

// NewMachineTokenSource validates the options.
func NewMachineTokenSource(options MachineTokenOptions) (*MachineTokenSource, error) {
	if strings.TrimSpace(options.ClientID) == "" || strings.TrimSpace(options.ClientSecret) == "" {
		return nil, errors.New("client ID and secret are required")
	}
	base, err := loginOrigin(options.LoginOrigin)
	if err != nil {
		return nil, err
	}
	if options.HTTPClient == nil {
		options.HTTPClient = &http.Client{CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}
	}
	if options.Timeout <= 0 {
		options.Timeout = 10 * time.Second
	}
	if options.RefreshSkew <= 0 {
		options.RefreshSkew = time.Minute
	}
	if options.Now == nil {
		options.Now = time.Now
	}
	return &MachineTokenSource{options: options, endpoint: base + "/user_management/token"}, nil
}

// Clear drops the cached token.
func (s *MachineTokenSource) Clear() {
	s.mu.Lock()
	s.token = ""
	s.mu.Unlock()
}

// Token returns a cached or fresh access token.
func (s *MachineTokenSource) Token(ctx context.Context) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.token != "" && s.options.Now().Before(s.expires.Add(-s.options.RefreshSkew)) {
		return s.token, nil
	}
	form := url.Values{"grant_type": {"client_credentials"}, "client_id": {s.options.ClientID}, "client_secret": {s.options.ClientSecret}}
	if s.options.Scope != "" {
		form.Set("scope", s.options.Scope)
	}
	for _, resource := range s.options.Resource {
		form.Add("resource", resource)
	}
	ctx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, s.endpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return "", &MachineTokenError{Code: "network_error"}
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Accept", "application/json")
	response, err := s.options.HTTPClient.Do(request)
	if err != nil {
		return "", &MachineTokenError{Code: "network_error"}
	}
	defer func() { _ = response.Body.Close() }()
	data, _ := io.ReadAll(io.LimitReader(response.Body, 65536))
	var body struct {
		AccessToken string   `json:"access_token"`
		ExpiresIn   *float64 `json:"expires_in"`
		Error       string   `json:"error"`
	}
	parsed := json.Unmarshal(data, &body) == nil
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		if parsed && body.Error != "" {
			return "", &MachineTokenError{Code: body.Error}
		}
		return "", &MachineTokenError{Code: "token_request_failed"}
	}
	if !parsed || body.AccessToken == "" || body.ExpiresIn == nil || *body.ExpiresIn <= 0 || math.IsInf(*body.ExpiresIn, 0) {
		return "", &MachineTokenError{Code: "invalid_response"}
	}
	s.token = body.AccessToken
	s.expires = s.options.Now().Add(time.Duration(*body.ExpiresIn * float64(time.Second)))
	return s.token, nil
}

// NewMachineTokenVerifier verifies machine-client access tokens (RFC 9068) at
// a resource server for the given audiences.
func NewMachineTokenVerifier(audience []string, origin string, client *http.Client) (*JWTVerifier, error) {
	issuer, err := loginOrigin(origin)
	if err != nil {
		return nil, err
	}
	return NewJWTVerifier(JWTVerifierConfig{Issuer: issuer, Audience: audience, JWKSURI: issuer + "/user_management/jwks",
		Types: []string{"at+jwt"}, RequiredClaims: []string{"sub", "client_id"}, HTTPClient: client})
}

// SecurityEvent is one event of a Security Event Token.
type SecurityEvent struct {
	Type    string
	Payload map[string]JSONValue
}

// SecurityEventToken is a verified SET.
type SecurityEventToken struct {
	JTI    string
	Issuer string
	IAT    int64
	TXN    string
	SubID  map[string]JSONValue
	Events []SecurityEvent
}

// SETVerifierOptions configure SET verification.
type SETVerifierOptions struct {
	Audience   string // the stream audience the transmitter reported
	SignalsURL string
	MaxAge     time.Duration // default 10 minutes
	// Seen records a jti and returns false on a replay.
	Seen       func(jti string) bool
	HTTPClient *http.Client
	Now        func() time.Time
}

// SETVerifier verifies SETs (RFC 8417) from the OhAuthCat transmitter.
type SETVerifier struct {
	jwt     *JWTVerifier
	options SETVerifierOptions
}

func signalsBase(value string) (string, error) {
	if value == "" {
		value = DefaultSignalsURL
	}
	if secureURL(value) != nil {
		return "", errors.New("signals URL must use https")
	}
	return strings.TrimSuffix(value, "/"), nil
}

// NewSETVerifier checks ES256, typ secevent+jwt, issuer, audience, a fresh
// iat, a jti and at least one event.
func NewSETVerifier(options SETVerifierOptions) (*SETVerifier, error) {
	issuer, err := signalsBase(options.SignalsURL)
	if err != nil {
		return nil, err
	}
	if options.MaxAge <= 0 {
		options.MaxAge = 10 * time.Minute
	}
	if options.Now == nil {
		options.Now = time.Now
	}
	verifier, err := NewJWTVerifier(JWTVerifierConfig{Issuer: issuer, Audience: []string{options.Audience}, JWKSURI: issuer + "/ssf/jwks",
		Types: []string{"secevent+jwt"}, RequiredClaims: []string{"jti"}, AllowMissingExpiry: true, HTTPClient: options.HTTPClient, Now: options.Now})
	if err != nil {
		return nil, err
	}
	return &SETVerifier{jwt: verifier, options: options}, nil
}

// Verify returns the parsed SET or a *TokenVerificationError.
func (v *SETVerifier) Verify(ctx context.Context, token string) (SecurityEventToken, error) {
	if v == nil {
		return SecurityEventToken{}, errToken
	}
	claims, err := v.jwt.Verify(ctx, token)
	if err != nil {
		return SecurityEventToken{}, errToken
	}
	var iat int64
	var events map[string]map[string]JSONValue
	rawIat, ok := claims["iat"]
	if !ok || json.Unmarshal(rawIat.data, &iat) != nil || v.options.Now().Sub(time.Unix(iat, 0)) > v.options.MaxAge {
		return SecurityEventToken{}, errToken
	}
	if raw, ok := claims["events"]; !ok || json.Unmarshal(raw.data, &events) != nil || len(events) == 0 {
		return SecurityEventToken{}, errToken
	}
	result := SecurityEventToken{JTI: claims.String("jti"), Issuer: claims.String("iss"), IAT: iat, TXN: claims.String("txn")}
	for kind, payload := range events {
		if payload == nil {
			return SecurityEventToken{}, errToken
		}
		result.Events = append(result.Events, SecurityEvent{Type: kind, Payload: payload})
	}
	slices.SortFunc(result.Events, func(a, b SecurityEvent) int { return strings.Compare(a.Type, b.Type) })
	if raw, ok := claims["sub_id"]; ok {
		_ = json.Unmarshal(raw.data, &result.SubID)
	}
	if v.options.Seen != nil && !v.options.Seen(result.JTI) {
		return SecurityEventToken{}, errToken
	}
	return result, nil
}

// SSFError is a failed SSF management request.
type SSFError struct{ Status int }

func (e *SSFError) Error() string { return fmt.Sprintf("SSF request failed (HTTP %d)", e.Status) }

// SSFStreams manages SSF transmitter streams (SSF 1.0 §8) with a secret key.
type SSFStreams struct {
	apiKey string
	root   string
	client *http.Client
}

// NewSSFStreams validates the key and base URL.
func NewSSFStreams(apiKey, signalsURL string, client *http.Client) (*SSFStreams, error) {
	if strings.TrimSpace(apiKey) == "" {
		return nil, errors.New("API key is required")
	}
	root, err := signalsBase(signalsURL)
	if err != nil {
		return nil, err
	}
	if client == nil {
		client = &http.Client{Timeout: 10 * time.Second, CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}
	}
	return &SSFStreams{apiKey: apiKey, root: root, client: client}, nil
}

// StreamConfiguration is a transmitter stream configuration.
type StreamConfiguration map[string]JSONValue

func (s *SSFStreams) call(ctx context.Context, method, path string, query url.Values, body any) ([]byte, error) {
	if s == nil {
		return nil, &SSFError{}
	}
	target := s.root + path
	if len(query) > 0 {
		target += "?" + query.Encode()
	}
	var reader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, &SSFError{}
		}
		reader = bytes.NewReader(data)
	}
	request, err := http.NewRequestWithContext(ctx, method, target, reader)
	if err != nil {
		return nil, &SSFError{}
	}
	request.Header.Set("Authorization", "Bearer "+s.apiKey)
	request.Header.Set("Accept", "application/json")
	if body != nil {
		request.Header.Set("Content-Type", "application/json")
	}
	response, err := s.client.Do(request)
	if err != nil {
		return nil, &SSFError{}
	}
	defer func() { _ = response.Body.Close() }()
	data, err := io.ReadAll(io.LimitReader(response.Body, 1<<20))
	if err != nil || response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, &SSFError{Status: response.StatusCode}
	}
	return data, nil
}

func streamOf(data []byte) (StreamConfiguration, error) {
	var stream StreamConfiguration
	var id string
	if json.Unmarshal(data, &stream) != nil || stream == nil || json.Unmarshal(stream["stream_id"].data, &id) != nil || id == "" {
		return nil, &SSFError{}
	}
	return stream, nil
}

// Create creates a stream.
func (s *SSFStreams) Create(ctx context.Context, input map[string]any) (StreamConfiguration, error) {
	data, err := s.call(ctx, http.MethodPost, "/ssf/streams", nil, input)
	if err != nil {
		return nil, err
	}
	return streamOf(data)
}

// Get reads one stream.
func (s *SSFStreams) Get(ctx context.Context, streamID string) (StreamConfiguration, error) {
	data, err := s.call(ctx, http.MethodGet, "/ssf/streams", url.Values{"stream_id": {streamID}}, nil)
	if err != nil {
		return nil, err
	}
	return streamOf(data)
}

// List reads every stream.
func (s *SSFStreams) List(ctx context.Context) ([]StreamConfiguration, error) {
	data, err := s.call(ctx, http.MethodGet, "/ssf/streams", nil, nil)
	if err != nil {
		return nil, err
	}
	var items []json.RawMessage
	if json.Unmarshal(data, &items) != nil {
		items = []json.RawMessage{data}
	}
	result := make([]StreamConfiguration, 0, len(items))
	for _, item := range items {
		stream, err := streamOf(item)
		if err != nil {
			return nil, err
		}
		result = append(result, stream)
	}
	return result, nil
}

// Update changes a stream.
func (s *SSFStreams) Update(ctx context.Context, streamID string, changes map[string]any) (StreamConfiguration, error) {
	body := map[string]any{"stream_id": streamID}
	for key, value := range changes {
		if key != "stream_id" {
			body[key] = value
		}
	}
	data, err := s.call(ctx, http.MethodPatch, "/ssf/streams", nil, body)
	if err != nil {
		return nil, err
	}
	return streamOf(data)
}

// Remove deletes a stream.
func (s *SSFStreams) Remove(ctx context.Context, streamID string) error {
	_, err := s.call(ctx, http.MethodDelete, "/ssf/streams", url.Values{"stream_id": {streamID}}, nil)
	return err
}

// SetStatus enables, pauses or disables a stream.
func (s *SSFStreams) SetStatus(ctx context.Context, streamID, status, reason string) error {
	body := map[string]any{"stream_id": streamID, "status": status}
	if reason != "" {
		body["reason"] = reason
	}
	_, err := s.call(ctx, http.MethodPost, "/ssf/status", nil, body)
	return err
}

// RequestVerification asks the transmitter to send a verification event.
func (s *SSFStreams) RequestVerification(ctx context.Context, streamID, state string) error {
	body := map[string]any{"stream_id": streamID}
	if state != "" {
		body["state"] = state
	}
	_, err := s.call(ctx, http.MethodPost, "/ssf/verify", nil, body)
	return err
}

// Poll polls a poll-delivery stream (RFC 8936) and returns raw SETs by jti.
func (s *SSFStreams) Poll(ctx context.Context, streamID string, ack []string, maxEvents int) (sets map[string]string, moreAvailable bool, err error) {
	body := map[string]any{"ack": append([]string{}, ack...), "returnImmediately": true}
	if maxEvents > 0 {
		body["maxEvents"] = maxEvents
	}
	data, err := s.call(ctx, http.MethodPost, "/ssf/poll/"+url.PathEscape(streamID), nil, body)
	if err != nil {
		return nil, false, err
	}
	var result struct {
		Sets          map[string]string `json:"sets"`
		MoreAvailable bool              `json:"moreAvailable"`
	}
	if json.Unmarshal(data, &result) != nil || result.Sets == nil {
		return nil, false, &SSFError{}
	}
	return result.Sets, result.MoreAvailable, nil
}

var environmentPattern = regexp.MustCompile(`^env_[0-9A-Za-z]+$`)

// OAuthIssuer is the OAuth provider issuer of one environment.
func OAuthIssuer(environmentID, origin string) (string, error) {
	if !environmentPattern.MatchString(environmentID) {
		return "", errors.New("environment ID is invalid")
	}
	base, err := loginOrigin(origin)
	if err != nil {
		return "", err
	}
	return base + "/oauth/" + environmentID, nil
}

// OAuthTokenVerifier verifies OAuth provider access tokens (RFC 9068) for a
// resource. Grant revocation is only visible through introspection.
type OAuthTokenVerifier struct {
	jwt *JWTVerifier
	env string
}

// NewOAuthTokenVerifier builds a verifier for the environment's issuer.
func NewOAuthTokenVerifier(environmentID string, audience []string, origin string, client *http.Client) (*OAuthTokenVerifier, error) {
	issuer, err := OAuthIssuer(environmentID, origin)
	if err != nil {
		return nil, err
	}
	verifier, err := NewJWTVerifier(JWTVerifierConfig{Issuer: issuer, Audience: audience, JWKSURI: issuer + "/jwks",
		Types: []string{"at+jwt"}, RequiredClaims: []string{"sub", "client_id", "env"}, HTTPClient: client})
	if err != nil {
		return nil, err
	}
	return &OAuthTokenVerifier{jwt: verifier, env: environmentID}, nil
}

// Verify returns the claims of a valid OAuth access token. A sid claim, when
// present, is the OhAuthCat session behind the grant and must be a string.
func (v *OAuthTokenVerifier) Verify(ctx context.Context, token string) (JWTClaims, error) {
	if v == nil {
		return nil, errToken
	}
	claims, err := v.jwt.Verify(ctx, token)
	if err != nil || claims.String("env") != v.env {
		return nil, errToken
	}
	for _, name := range []string{"scope", "sid"} {
		if raw, ok := claims[name]; ok {
			var value string
			if json.Unmarshal(raw.data, &value) != nil {
				return nil, errToken
			}
		}
	}
	return claims, nil
}

// TokenScopes splits the space-separated scope claim.
func TokenScopes(claims JWTClaims) []string { return strings.Fields(claims.String("scope")) }

// ProtectedResourceMetadata is RFC 9728 metadata for an MCP server or API.
func ProtectedResourceMetadata(resource, environmentID, origin string, scopes []string) (map[string]any, error) {
	parsed, err := url.Parse(resource)
	if err != nil || parsed.Scheme != "https" || parsed.Host == "" || parsed.Fragment != "" {
		return nil, errors.New("resource must be an https URL without a fragment")
	}
	issuer, err := OAuthIssuer(environmentID, origin)
	if err != nil {
		return nil, err
	}
	metadata := map[string]any{"resource": resource, "authorization_servers": []string{issuer}, "bearer_methods_supported": []string{"header"}}
	if len(scopes) > 0 {
		metadata["scopes_supported"] = scopes
	}
	return metadata, nil
}

// ProtectedResourceMetadataURL is the RFC 9728 §3.1 well-known URL.
func ProtectedResourceMetadataURL(resource string) string {
	parsed, err := url.Parse(resource)
	if err != nil {
		return ""
	}
	path := parsed.EscapedPath()
	if path == "/" {
		path = ""
	}
	return parsed.Scheme + "://" + parsed.Host + "/.well-known/oauth-protected-resource" + path
}

// BearerChallenge builds a WWW-Authenticate value pointing at the metadata.
func BearerChallenge(resource, errorCode, scope string) string {
	quote := func(value string) string {
		return `"` + strings.NewReplacer(`"`, "", `\`, "").Replace(value) + `"`
	}
	parts := []string{"resource_metadata=" + quote(ProtectedResourceMetadataURL(resource))}
	if errorCode != "" {
		parts = append(parts, "error="+quote(errorCode))
	}
	if scope != "" {
		parts = append(parts, "scope="+quote(scope))
	}
	return "Bearer " + strings.Join(parts, ", ")
}

// BearerResult is the outcome of validating a request's bearer token.
type BearerResult struct {
	Claims          JWTClaims
	Status          int
	WWWAuthenticate string
}

var bearerPattern = regexp.MustCompile(`^Bearer ([A-Za-z0-9._~+/-]+=*)$`)

// ValidateBearer validates an Authorization header for resource with every
// required scope; on failure Status and WWWAuthenticate say what to send.
func ValidateBearer(ctx context.Context, verifier *OAuthTokenVerifier, resource, authorization string, required []string) BearerResult {
	match := bearerPattern.FindStringSubmatch(authorization)
	if match == nil {
		return BearerResult{Status: http.StatusUnauthorized, WWWAuthenticate: BearerChallenge(resource, "", "")}
	}
	claims, err := verifier.Verify(ctx, match[1])
	if err != nil {
		return BearerResult{Status: http.StatusUnauthorized, WWWAuthenticate: BearerChallenge(resource, "invalid_token", "")}
	}
	granted := TokenScopes(claims)
	for _, scope := range required {
		if !slices.Contains(granted, scope) {
			return BearerResult{Status: http.StatusForbidden, WWWAuthenticate: BearerChallenge(resource, "insufficient_scope", strings.Join(required, " "))}
		}
	}
	return BearerResult{Claims: claims, Status: http.StatusOK}
}

// RolloutScale is the total weight of a flag rollout (100000 = 100%).
const RolloutScale = 100000

// FlagCondition matches one attribute of the context.
type FlagCondition struct {
	Attribute string   `json:"attribute"`
	Operator  string   `json:"operator"`
	Values    []string `json:"values"`
}

// FlagServe is a fixed variant or a percentage rollout.
type FlagServe struct {
	Variant string `json:"variant,omitempty"`
	Rollout []struct {
		Variant string `json:"variant"`
		Weight  int    `json:"weight"`
	} `json:"rollout,omitempty"`
	BucketBy string `json:"bucket_by,omitempty"`
}

// FlagDefinition is one flag of a flag export (a FeatureFlag object).
type FlagDefinition struct {
	Key      string `json:"key"`
	Enabled  bool   `json:"enabled"`
	Variants []struct {
		Key   string    `json:"key"`
		Value JSONValue `json:"value"`
	} `json:"variants"`
	OffVariant string `json:"off_variant"`
	Rules      []struct {
		Conditions []FlagCondition `json:"conditions"`
		Serve      FlagServe       `json:"serve"`
	} `json:"rules"`
	Fallthrough FlagServe `json:"fallthrough"`
}

// FlagContext is the subject of an evaluation.
type FlagContext struct {
	OrganizationID string
	UserID         string
	Attributes     map[string]string
}

// FlagResult is one local evaluation.
type FlagResult struct {
	Variant string
	Value   JSONValue
	Reason  string // disabled, rule_match or default
	Rule    int
}

// FlagBucket is the subject's stable bucket in [0, RolloutScale).
func FlagBucket(flagKey, subject string) int {
	sum := sha256.Sum256([]byte(flagKey + "\x00" + subject))
	return int(binary.BigEndian.Uint64(sum[:8]) % RolloutScale)
}

func (c FlagContext) attribute(name string) (string, bool) {
	switch name {
	case "organization_id":
		return c.OrganizationID, c.OrganizationID != ""
	case "user_id":
		return c.UserID, c.UserID != ""
	}
	value, ok := c.Attributes[name]
	return value, ok
}

func (c FlagContext) matches(condition FlagCondition) bool {
	value, ok := c.attribute(condition.Attribute)
	if condition.Operator == "not_in" {
		return !ok || !slices.Contains(condition.Values, value)
	}
	if !ok {
		return false
	}
	switch condition.Operator {
	case "in":
		return slices.Contains(condition.Values, value)
	case "contains":
		return slices.ContainsFunc(condition.Values, func(x string) bool { return strings.Contains(value, x) })
	case "starts_with":
		return slices.ContainsFunc(condition.Values, func(x string) bool { return strings.HasPrefix(value, x) })
	case "ends_with":
		return slices.ContainsFunc(condition.Values, func(x string) bool { return strings.HasSuffix(value, x) })
	case "greater_than", "less_than":
		if len(condition.Values) == 0 {
			return false
		}
		got, err := strconv.ParseFloat(value, 64)
		want, werr := strconv.ParseFloat(condition.Values[0], 64)
		if err != nil || werr != nil {
			return false
		}
		if condition.Operator == "greater_than" {
			return got > want
		}
		return got < want
	}
	return false
}

func (f *FlagDefinition) serve(serve FlagServe, c FlagContext) string {
	if serve.Variant != "" {
		return serve.Variant
	}
	subject := c.UserID
	if serve.BucketBy != "" {
		subject, _ = c.attribute(serve.BucketBy)
	} else if subject == "" {
		subject = c.OrganizationID
	}
	if subject == "" {
		for _, weight := range serve.Rollout {
			if weight.Weight > 0 {
				return weight.Variant
			}
		}
	}
	if len(serve.Rollout) == 0 {
		return ""
	}
	bucket, total := FlagBucket(f.Key, subject), 0
	for _, weight := range serve.Rollout {
		total += weight.Weight
		if bucket < total {
			return weight.Variant
		}
	}
	return serve.Rollout[len(serve.Rollout)-1].Variant
}

func (f *FlagDefinition) value(variant string) JSONValue {
	for _, item := range f.Variants {
		if item.Key == variant {
			return item.Value
		}
	}
	return JSONValue{}
}

// EvaluateFlag evaluates a flag locally with the server's rules.
func EvaluateFlag(flag *FlagDefinition, c FlagContext) FlagResult {
	if !flag.Enabled {
		return FlagResult{Variant: flag.OffVariant, Value: flag.value(flag.OffVariant), Reason: "disabled", Rule: -1}
	}
	for index, rule := range flag.Rules {
		if !slices.ContainsFunc(rule.Conditions, func(condition FlagCondition) bool { return !c.matches(condition) }) {
			variant := flag.serve(rule.Serve, c)
			return FlagResult{Variant: variant, Value: flag.value(variant), Reason: "rule_match", Rule: index}
		}
	}
	variant := flag.serve(flag.Fallthrough, c)
	return FlagResult{Variant: variant, Value: flag.value(variant), Reason: "default", Rule: -1}
}

// LocalFlags evaluates flags over a periodically refreshed flag export. A
// failed refresh keeps the last good snapshot.
type LocalFlags struct {
	load     func(context.Context) ([]json.RawMessage, error)
	interval time.Duration
	onError  func(error)
	mu       sync.RWMutex
	flags    map[string]*FlagDefinition
	stop     chan struct{}
	done     chan struct{}
}

// NewLocalFlags builds a local evaluator; load returns every flag as JSON
// (for example by paging ListFeatureFlags). The interval is at least 5 s.
func NewLocalFlags(load func(context.Context) ([]json.RawMessage, error), interval time.Duration, onError func(error)) *LocalFlags {
	if interval < 5*time.Second {
		interval = 30 * time.Second
	}
	return &LocalFlags{load: load, interval: interval, onError: onError, flags: map[string]*FlagDefinition{}}
}

// Refresh reloads the snapshot once.
func (l *LocalFlags) Refresh(ctx context.Context) {
	items, err := l.load(ctx)
	if err != nil {
		if l.onError != nil {
			l.onError(err)
		}
		return
	}
	next := map[string]*FlagDefinition{}
	for _, item := range items {
		var flag FlagDefinition
		if json.Unmarshal(item, &flag) == nil && flag.Key != "" && flag.OffVariant != "" {
			next[flag.Key] = &flag
		}
	}
	l.mu.Lock()
	l.flags = next
	l.mu.Unlock()
}

// Start loads once and refreshes in the background until Stop.
func (l *LocalFlags) Start(ctx context.Context) {
	l.Refresh(ctx)
	l.mu.Lock()
	if l.stop != nil {
		l.mu.Unlock()
		return
	}
	l.stop, l.done = make(chan struct{}), make(chan struct{})
	stop, done := l.stop, l.done
	l.mu.Unlock()
	go func() {
		defer close(done)
		ticker := time.NewTicker(l.interval)
		defer ticker.Stop()
		for {
			select {
			case <-stop:
				return
			case <-ctx.Done():
				return
			case <-ticker.C:
				l.Refresh(ctx)
			}
		}
	}()
}

// Stop ends background refreshes.
func (l *LocalFlags) Stop() {
	l.mu.Lock()
	stop, done := l.stop, l.done
	l.stop, l.done = nil, nil
	l.mu.Unlock()
	if stop != nil {
		close(stop)
		<-done
	}
}

// Len is the number of flags in the snapshot.
func (l *LocalFlags) Len() int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return len(l.flags)
}

// Evaluate evaluates a flag; ok is false for an unknown flag.
func (l *LocalFlags) Evaluate(key string, c FlagContext) (FlagResult, bool) {
	l.mu.RLock()
	flag := l.flags[key]
	l.mu.RUnlock()
	if flag == nil {
		return FlagResult{}, false
	}
	return EvaluateFlag(flag, c), true
}

// IsEnabled reports whether the flag serves true, or fallback when unknown.
func (l *LocalFlags) IsEnabled(key string, c FlagContext, fallback bool) bool {
	result, ok := l.Evaluate(key, c)
	if !ok {
		return fallback
	}
	return string(result.Value.data) == "true"
}

// AuditEventHints checks an audit log event body against the pinned schema
// and returns readable problems; empty means it is well formed. The server
// still validates the event against the action's registered schema.
func AuditEventHints(body []byte) []string {
	var contract wireContract
	if json.Unmarshal([]byte(pinnedContract), &contract) != nil {
		return []string{"pinned contract unavailable"}
	}
	hints := []string{}
	auditWalk(body, schema{Ref: "#/components/schemas/AuditLogEventCreate"}, "body", contract.Schemas, &hints)
	return hints
}

func resolveSchema(definition schema, definitions map[string]schema) schema {
	for range 16 {
		if definition.Ref == "" {
			break
		}
		next, ok := definitions[strings.TrimPrefix(definition.Ref, "#/components/schemas/")]
		if !ok {
			break
		}
		definition = next
	}
	return definition
}

func auditWalk(data []byte, definition schema, path string, definitions map[string]schema, hints *[]string) {
	if matchesJSON(data, definition, definitions, false) {
		return
	}
	shape := resolveSchema(definition, definitions)
	var object map[string]json.RawMessage
	if shape.Properties != nil && json.Unmarshal(data, &object) == nil && object != nil {
		before := len(*hints)
		for _, name := range shape.Required {
			if value, ok := object[name]; !ok || string(value) == "null" {
				*hints = append(*hints, path+"."+name+" is required")
			}
		}
		names := make([]string, 0, len(object))
		for name := range object {
			names = append(names, name)
		}
		slices.Sort(names)
		for _, name := range names {
			if string(object[name]) == "null" {
				continue
			}
			if child, ok := shape.Properties[name]; ok {
				auditWalk(object[name], child, path+"."+name, definitions, hints)
			} else if string(shape.Additional) == "false" {
				*hints = append(*hints, path+"."+name+" is not a known field")
			}
		}
		if len(*hints) == before {
			*hints = append(*hints, path+" does not match the schema")
		}
		return
	}
	var items []json.RawMessage
	if shape.Items != nil && json.Unmarshal(data, &items) == nil && items != nil {
		before := len(*hints)
		for index, item := range items {
			auditWalk(item, *shape.Items, fmt.Sprintf("%s[%d]", path, index), definitions, hints)
		}
		if len(*hints) == before {
			*hints = append(*hints, path+" has the wrong number of items")
		}
		return
	}
	*hints = append(*hints, path+" does not match the schema")
}
