package ohauthcat

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"strings"
	"unicode/utf8"
)

// ErrPasskeyInput reports a missing user ID or a credential that is not a JSON object.
var ErrPasskeyInput = errors.New("passkey user ID or credential invalid")

// NativePasskeys is the server-side half of native passkeys (the iOS and
// Android PasskeyBackend). The mobile app never holds a secret key: it asks
// your backend for options, runs the platform ceremony and posts the
// credential back, and your backend calls these methods with its secret key.
// Bind the user ID to your own authenticated session before registering, and
// create your own session after Authenticate returns the user.
//
// Once the environment activates a custom sign-in host
// (PUT /v1/environment/signin_host), new passkeys use that host as their RP
// ID. Add it to the app's associated domains and Android asset links before
// switching; passkeys created earlier keep working only under
// signin.<domain>.
type NativePasskeys struct {
	client *Client
}

// NewNativePasskeys wraps a client configured with a secret key.
func NewNativePasskeys(client *Client) *NativePasskeys {
	return &NativePasskeys{client: client}
}

func passkeyUser(userID string) (EndUserID, error) {
	trimmed := strings.TrimSpace(userID)
	if trimmed == "" {
		return "", ErrPasskeyInput
	}
	return EndUserID(trimmed), nil
}

func passkeyCredential(raw json.RawMessage) (map[string]JSONValue, error) {
	trimmed := bytes.TrimSpace(raw)
	if len(trimmed) == 0 || trimmed[0] != '{' {
		return nil, ErrPasskeyInput
	}
	var credential map[string]JSONValue
	if json.Unmarshal(trimmed, &credential) != nil || credential == nil {
		return nil, ErrPasskeyInput
	}
	return credential, nil
}

func (p *NativePasskeys) resources() (*UsersResources, *UserManagementResources, error) {
	if p == nil || p.client == nil {
		return nil, nil, requestFailure("client_closed")
	}
	return p.client.Users, p.client.UserManagement, nil
}

// RegistrationOptions returns WebAuthn creation options (publicKey) for the user's new passkey.
func (p *NativePasskeys) RegistrationOptions(ctx context.Context, userID string) (map[string]JSONValue, error) {
	users, _, err := p.resources()
	if err != nil {
		return nil, err
	}
	user, err := passkeyUser(userID)
	if err != nil {
		return nil, err
	}
	result, err := users.StartEndUserPasskeyRegistration(ctx, StartEndUserPasskeyRegistrationArgs{Path: StartEndUserPasskeyRegistrationPath{User: user}})
	if err != nil {
		return nil, err
	}
	return result.PublicKey, nil
}

// Register verifies the attestation the device returned and stores the passkey.
// A blank name is left unset.
func (p *NativePasskeys) Register(ctx context.Context, userID string, attestation json.RawMessage, name string) (EndUserFactor, error) {
	users, _, err := p.resources()
	if err != nil {
		return EndUserFactor{}, err
	}
	user, err := passkeyUser(userID)
	if err != nil {
		return EndUserFactor{}, err
	}
	credential, err := passkeyCredential(attestation)
	if err != nil {
		return EndUserFactor{}, err
	}
	body := RegisterEndUserPasskeyBody{Credential: credential}
	if label := strings.TrimSpace(name); label != "" {
		body.Name = &label
	}
	return users.RegisterEndUserPasskey(ctx, RegisterEndUserPasskeyArgs{Path: RegisterEndUserPasskeyPath{User: user}, Body: body})
}

// AuthenticationOptions returns WebAuthn request options for a discoverable-credential sign-in.
func (p *NativePasskeys) AuthenticationOptions(ctx context.Context) (map[string]JSONValue, error) {
	_, management, err := p.resources()
	if err != nil {
		return nil, err
	}
	result, err := management.StartEndUserPasskeyAuthentication(ctx, StartEndUserPasskeyAuthenticationArgs{})
	if err != nil {
		return nil, err
	}
	return result.PublicKey, nil
}

// ErrPasskeySessionMissing reports a session sign-in answer without its token set.
var ErrPasskeySessionMissing = errors.New("passkey sign-in returned no session tokens")

// PasskeySession is the token set a session passkey sign-in returns, issued by
// the login host with the client ID as audience.
type PasskeySession struct {
	User         EndUser
	SessionID    string
	AccessToken  string
	RefreshToken string
	IDToken      string
	TokenType    string
	ExpiresIn    int64
}

func passkeyText(value string, maximum int) (string, error) {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" || utf8.RuneCountInString(trimmed) > maximum || !utf8.ValidString(trimmed) {
		return "", ErrPasskeyInput
	}
	for _, r := range trimmed {
		if r < 0x20 || r == 0x7f {
			return "", ErrPasskeyInput
		}
	}
	return trimmed, nil
}

func (p *NativePasskeys) authenticate(ctx context.Context, assertion json.RawMessage, body AuthenticateEndUserPasskeyBody) (EndUserPasskeyAuthentication, error) {
	_, management, err := p.resources()
	if err != nil {
		return EndUserPasskeyAuthentication{}, err
	}
	body.Credential, err = passkeyCredential(assertion)
	if err != nil {
		return EndUserPasskeyAuthentication{}, err
	}
	return management.AuthenticateEndUserPasskey(ctx, AuthenticateEndUserPasskeyArgs{Body: body})
}

// Authenticate verifies the assertion and returns the end user it signs in.
func (p *NativePasskeys) Authenticate(ctx context.Context, assertion json.RawMessage) (EndUserPasskeyAuthentication, error) {
	return p.authenticate(ctx, assertion, AuthenticateEndUserPasskeyBody{})
}

// AuthenticateSession verifies the assertion, creates a session and returns
// its token set for clientID (a publishable key of the environment), as
// /user_management/token would. When the environment requires MFA and the
// assertion was not user-verified, the API answers mfa_required instead.
func (p *NativePasskeys) AuthenticateSession(ctx context.Context, assertion json.RawMessage, clientID string) (PasskeySession, error) {
	client, err := passkeyText(clientID, 256)
	if err != nil {
		return PasskeySession{}, err
	}
	session := true
	result, err := p.authenticate(ctx, assertion, AuthenticateEndUserPasskeyBody{Session: &session, ClientID: &client})
	if err != nil {
		return PasskeySession{}, err
	}
	if result.SessionID == nil || result.AccessToken == nil || result.RefreshToken == nil || result.IDToken == nil || result.ExpiresIn == nil {
		return PasskeySession{}, ErrPasskeySessionMissing
	}
	return PasskeySession{
		User:         result.User,
		SessionID:    *result.SessionID,
		AccessToken:  *result.AccessToken,
		RefreshToken: *result.RefreshToken,
		IDToken:      *result.IDToken,
		TokenType:    "Bearer",
		ExpiresIn:    *result.ExpiresIn,
	}, nil
}

// RenamePasskey sets the display name of one of the user's passkeys
// (trimmed, 1 to 64 characters, no control characters).
func (p *NativePasskeys) RenamePasskey(ctx context.Context, userID, passkeyID, name string) (EndUserFactor, error) {
	users, _, err := p.resources()
	if err != nil {
		return EndUserFactor{}, err
	}
	user, err := passkeyUser(userID)
	if err != nil {
		return EndUserFactor{}, err
	}
	passkey, err := passkeyText(passkeyID, 64)
	if err != nil {
		return EndUserFactor{}, err
	}
	label, err := passkeyText(name, 64)
	if err != nil {
		return EndUserFactor{}, err
	}
	return users.UpdateEndUserPasskey(ctx, UpdateEndUserPasskeyArgs{
		Path: UpdateEndUserPasskeyPath{User: user, Passkey: EndUserFactorID(passkey)},
		Body: UpdateEndUserPasskeyBody{Name: label},
	})
}
