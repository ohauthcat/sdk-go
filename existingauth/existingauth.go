// Package existingauth links a verified OhAuthCat profile to a user in an
// existing auth system (Auth0, Amazon Cognito, Firebase, Supabase, Clerk or a
// custom store) and starts a session there.
//
// Linking rule: a user is keyed by the profile's connection and IdP subject
// (connection_id + idp_id), never by email. A missing link is created only
// when the policy says OnMissingCreate. An existing provider account with the
// same email but no link is reported as CodeEmailConflict; link those accounts
// through an authenticated, explicit flow instead.
package existingauth

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	ohauthcat "ohauthcat.com/sdk-go"
)

// Provider names the existing auth system.
type Provider string

// Supported existing auth systems.
const (
	ProviderAuth0    Provider = "auth0"
	ProviderCognito  Provider = "cognito"
	ProviderFirebase Provider = "firebase"
	ProviderSupabase Provider = "supabase"
	ProviderClerk    Provider = "clerk"
	ProviderCustom   Provider = "custom"
)

// Code is a stable error code.
type Code string

// Stable error codes.
const (
	CodeInvalidInput        Code = "invalid_input"
	CodeNotLinked           Code = "not_linked"
	CodeEmailConflict       Code = "email_conflict"
	CodeProviderError       Code = "provider_error"
	CodeProviderUnavailable Code = "provider_unavailable"
	CodeInvalidResponse     Code = "invalid_response"
	CodeSessionRefused      Code = "session_refused"
)

// Error is the typed error every helper returns. Messages never carry
// tokens, secrets or response bodies.
type Error struct {
	Provider     Provider
	Code         Code
	Message      string
	Status       int
	ProviderCode string
}

func (e *Error) Error() string { return fmt.Sprintf("%s: %s: %s", e.Provider, e.Code, e.Message) }

var safeCode = regexp.MustCompile(`^[A-Za-z0-9_.:-]{1,64}$`)

func newError(p Provider, code Code, msg string, status int, providerCode string) *Error {
	if !safeCode.MatchString(providerCode) {
		providerCode = ""
	}
	return &Error{Provider: p, Code: code, Message: msg, Status: status, ProviderCode: providerCode}
}

func invalid(p Provider, msg string) *Error { return newError(p, CodeInvalidInput, msg, 0, "") }

// IsCode reports whether err is an *Error with the given code.
func IsCode(err error, code Code) bool {
	var e *Error
	return errors.As(err, &e) && e.Code == code
}

// ProfileInput holds the profile fields a link needs.
type ProfileInput struct {
	ID             string
	OrganizationID string
	ConnectionID   string
	IDPID          string
	Email          string
	EmailVerified  bool
	FirstName      string
	LastName       string
}

// FromProfile copies the needed fields from an SDK profile.
func FromProfile(p ohauthcat.Profile) ProfileInput {
	in := ProfileInput{
		ID: string(p.ID), OrganizationID: string(p.OrganizationID), ConnectionID: string(p.ConnectionID),
		IDPID: p.IDPID, Email: p.Email, EmailVerified: p.EmailVerified != nil && *p.EmailVerified,
	}
	if p.FirstName.Set && !p.FirstName.Null {
		in.FirstName = p.FirstName.Value
	}
	if p.LastName.Set && !p.LastName.Null {
		in.LastName = p.LastName.Value
	}
	return in
}

// LinkProfile is a validated profile plus its stable link key.
type LinkProfile struct {
	ProfileID      string
	OrganizationID string
	ConnectionID   string
	IDPID          string
	Email          string
	EmailVerified  bool
	FirstName      string
	LastName       string
	// Key is "oac_" + 40 hex characters of SHA-256(connection_id, idp_id).
	Key string
}

// OnMissing says what to do when no linked user exists.
type OnMissing string

// Behaviors when no linked user exists.
const (
	OnMissingCreate OnMissing = "create"
	OnMissingRefuse OnMissing = "refuse"
)

// Policy holds the explicit linking rules. There is no default.
type Policy struct {
	OnMissing OnMissing
}

// Result is a linked user and whether it was just created.
type Result[U any] struct {
	User    U
	Created bool
}

// Linker is the generic interface every provider helper implements.
// Find returns (zero, false, nil) when no user is linked.
type Linker[U any] interface {
	Provider() Provider
	Find(ctx context.Context, p LinkProfile) (U, bool, error)
	Create(ctx context.Context, p LinkProfile) (U, error)
}

const maxField = 1024

var emailPattern = regexp.MustCompile(`^[^\s@]+@[^\s@]+$`)

// LinkKey returns the stable link key for a connection and IdP subject.
func LinkKey(connectionID, idpID string) string {
	sum := sha256.Sum256([]byte(connectionID + "\n" + idpID))
	return "oac_" + hex.EncodeToString(sum[:])[:40]
}

// LinkUUID returns a UUID (version 8, RFC 9562) derived from the link key.
func LinkUUID(key string) string {
	sum := sha256.Sum256([]byte(key))
	b := sum[:16]
	b[6] = (b[6] & 0x0f) | 0x80
	b[8] = (b[8] & 0x3f) | 0x80
	h := hex.EncodeToString(b)
	return h[:8] + "-" + h[8:12] + "-" + h[12:16] + "-" + h[16:20] + "-" + h[20:]
}

func field(p Provider, value, name string, required bool) (string, error) {
	v := strings.TrimSpace(value)
	if v == "" || len(v) > maxField {
		if required {
			return "", invalid(p, fmt.Sprintf("profile.%s must be 1 to %d characters", name, maxField))
		}
		return "", nil
	}
	return v, nil
}

// NewLinkProfile validates the input and derives the link key.
func NewLinkProfile(in ProfileInput, p Provider) (LinkProfile, error) {
	var out LinkProfile
	var err error
	for _, f := range []struct {
		dst  *string
		src  string
		name string
		req  bool
	}{
		{&out.ConnectionID, in.ConnectionID, "connection_id", true},
		{&out.IDPID, in.IDPID, "idp_id", true},
		{&out.ProfileID, in.ID, "id", true},
		{&out.OrganizationID, in.OrganizationID, "organization_id", true},
		{&out.Email, in.Email, "email", false},
		{&out.FirstName, in.FirstName, "first_name", false},
		{&out.LastName, in.LastName, "last_name", false},
	} {
		if *f.dst, err = field(p, f.src, f.name, f.req); err != nil {
			return LinkProfile{}, err
		}
	}
	if out.Email != "" && !emailPattern.MatchString(out.Email) {
		return LinkProfile{}, invalid(p, "profile.email is not an email address")
	}
	out.Email = strings.ToLower(out.Email)
	out.EmailVerified = in.EmailVerified
	out.Key = LinkKey(out.ConnectionID, out.IDPID)
	return out, nil
}

// Link finds the user linked to this profile, or creates one when the policy
// allows it. A concurrent create is resolved by finding again.
func Link[U any](ctx context.Context, in ProfileInput, l Linker[U], policy Policy) (Result[U], error) {
	p := l.Provider()
	if policy.OnMissing != OnMissingCreate && policy.OnMissing != OnMissingRefuse {
		return Result[U]{}, invalid(p, `policy.OnMissing must be "create" or "refuse"`)
	}
	lp, err := NewLinkProfile(in, p)
	if err != nil {
		return Result[U]{}, err
	}
	user, ok, err := l.Find(ctx, lp)
	if err != nil {
		return Result[U]{}, err
	}
	if ok {
		return Result[U]{User: user}, nil
	}
	if policy.OnMissing == OnMissingRefuse {
		return Result[U]{}, newError(p, CodeNotLinked, "no user is linked to this profile", 0, "")
	}
	user, err = l.Create(ctx, lp)
	if err == nil {
		return Result[U]{User: user, Created: true}, nil
	}
	if IsCode(err, CodeEmailConflict) {
		if raced, ok, ferr := l.Find(ctx, lp); ferr == nil && ok {
			return Result[U]{User: raced}, nil
		}
	}
	return Result[U]{}, err
}

// ---------------------------------------------------------------- HTTP

const maxResponse = 1 << 20

type reply struct {
	status int
	body   json.RawMessage
}

func (r reply) decode(v any) bool {
	return len(r.body) > 0 && json.Unmarshal(r.body, v) == nil
}

func send(ctx context.Context, p Provider, client *http.Client, method, u string, headers map[string]string, body []byte) (reply, error) {
	if client == nil {
		client = &http.Client{Timeout: 10 * time.Second}
	}
	var reader io.Reader
	if body != nil {
		reader = bytes.NewReader(body)
	}
	req, err := http.NewRequestWithContext(ctx, method, u, reader)
	if err != nil {
		return reply{}, invalid(p, "request could not be built")
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	noRedirect := *client
	noRedirect.CheckRedirect = func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }
	resp, err := noRedirect.Do(req)
	if err != nil {
		return reply{}, newError(p, CodeProviderUnavailable, string(p)+" request failed", 0, "")
	}
	defer func() { _ = resp.Body.Close() }()
	raw, err := io.ReadAll(io.LimitReader(resp.Body, maxResponse+1))
	if err != nil {
		return reply{}, newError(p, CodeProviderUnavailable, string(p)+" response could not be read", resp.StatusCode, "")
	}
	if len(raw) > maxResponse {
		return reply{}, newError(p, CodeInvalidResponse, string(p)+" response is too large", resp.StatusCode, "")
	}
	if !json.Valid(raw) {
		raw = nil
	}
	return reply{status: resp.StatusCode, body: raw}, nil
}

func sendJSON(ctx context.Context, p Provider, client *http.Client, method, u string, headers map[string]string, body any) (reply, error) {
	var data []byte
	if body != nil {
		var err error
		if data, err = json.Marshal(body); err != nil {
			return reply{}, invalid(p, "request body could not be encoded")
		}
		headers["content-type"] = "application/json"
	}
	return send(ctx, p, client, method, u, headers, data)
}

func failure(p Provider, r reply, providerCode string) error {
	code := CodeProviderError
	if r.status >= 500 || r.status == http.StatusTooManyRequests {
		code = CodeProviderUnavailable
	}
	return newError(p, code, fmt.Sprintf("%s returned HTTP %d", p, r.status), r.status, providerCode)
}

func badResponse(p Provider) error {
	return newError(p, CodeInvalidResponse, string(p)+" returned an unexpected response", 0, "")
}

func httpsOrigin(p Provider, value, name string) (string, error) {
	u, err := url.Parse(value)
	if err != nil || u.Host == "" {
		return "", invalid(p, name+" must be a URL")
	}
	local := u.Hostname() == "127.0.0.1" || u.Hostname() == "localhost"
	if u.Scheme != "https" && (!local || u.Scheme != "http") {
		return "", invalid(p, name+" must use https")
	}
	return u.Scheme + "://" + u.Host, nil
}

func required(p Provider, value, name string) error {
	if strings.TrimSpace(value) == "" {
		return invalid(p, name+" is required")
	}
	return nil
}

type linkMetadata struct {
	Key            string `json:"key"`
	ConnectionID   string `json:"connection_id"`
	OrganizationID string `json:"organization_id"`
	IDPID          string `json:"idp_id"`
}

func metadata(p LinkProfile) map[string]linkMetadata {
	return map[string]linkMetadata{"ohauthcat": {Key: p.Key, ConnectionID: p.ConnectionID, OrganizationID: p.OrganizationID, IDPID: p.IDPID}}
}

func displayName(p LinkProfile) string {
	return strings.TrimSpace(p.FirstName + " " + p.LastName)
}
