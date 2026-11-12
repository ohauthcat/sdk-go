// Package frameworkhttp provides bounded process-local net/http SSO and webhook helpers.
package frameworkhttp

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"time"

	sdk "ohauthcat.com/sdk-go"
)

// SDK is the existing typed SSO resource contract.
type SDK interface {
	GetAuthorizationURL(sdk.AuthorizationOptions) (sdk.AuthorizationResult, error)
	GetProfileAndToken(context.Context, sdk.ExchangeOptions) (sdk.SSOToken, error)
	VerifyAccessToken(context.Context, string) (sdk.AccessTokenClaims, error)
	GetSSOProfile(context.Context, sdk.GetSSOProfileArgs) (sdk.Profile, error)
}

// Options binds every flow to fixed application and tenant configuration.
type Options struct {
	SSO                                                SDK
	Environment, Organization, Connection, RedirectURI string
	Now                                                func() time.Time
	Capacity                                           int
	WebhookSecrets                                     []string
	Queue                                              func(context.Context, sdk.Event) error
}

// Identity is the closed session identity; tokens and profile values are never returned.
type Identity struct {
	Environment  string `json:"environment"`
	Organization string `json:"organization"`
	Connection   string `json:"connection"`
	Profile      string `json:"profile"`
}

// ErrInvalid is a fixed configuration or request failure.
var ErrInvalid = errors.New("framework HTTP: invalid configuration or request")

// ErrClosed is returned after helper shutdown.
var ErrClosed = errors.New("framework HTTP: closed")

type flow struct {
	verifier string
	expires  time.Time
}
type session struct {
	identity Identity
	expires  time.Time
}

// Helper owns process-local flows and sessions. Closing never closes the caller's SDK.
type Helper struct {
	options  Options
	secure   bool
	root     context.Context
	cancel   context.CancelFunc
	mutex    sync.Mutex
	closed   bool
	flows    map[string]flow
	sessions map[string]session
}

var profileID = regexp.MustCompile(`^prof_[0-9A-Za-z]{16,64}$`)
var tenantID = regexp.MustCompile(`^(env|org|conn)_[0-9A-HJKMNP-TV-Z]{26}$`)

func validID(value, prefix string) bool {
	return strings.HasPrefix(value, prefix+"_") && tenantID.MatchString(value)
}
func sdkAbsent(client SDK) bool {
	if client == nil {
		return true
	}
	value := reflect.ValueOf(client)
	switch value.Kind() {
	case reflect.Pointer, reflect.Map, reflect.Slice, reflect.Func, reflect.Chan, reflect.Interface:
		return value.IsNil()
	default:
		return false
	}
}

// New validates and defensively copies fixed configuration before serving requests.
func New(options Options) (*Helper, error) {
	if sdkAbsent(options.SSO) || !validID(options.Environment, "env") || !validID(options.Organization, "org") || !validID(options.Connection, "conn") {
		return nil, ErrInvalid
	}
	redirect, err := url.Parse(options.RedirectURI)
	if err != nil || len(options.RedirectURI) > 2048 || strings.ContainsAny(options.RedirectURI, "\r\n\x00") || redirect.Hostname() == "" || redirect.User != nil || redirect.Fragment != "" || redirect.RawQuery != "" || (redirect.Scheme != "https" && (redirect.Scheme != "http" || (redirect.Hostname() != "localhost" && (net.ParseIP(redirect.Hostname()) == nil || !net.ParseIP(redirect.Hostname()).IsLoopback())))) {
		return nil, ErrInvalid
	}
	if options.Capacity == 0 {
		options.Capacity = 1024
	}
	if options.Capacity < 1 || options.Capacity > 10000 {
		return nil, ErrInvalid
	}
	if options.Now == nil {
		options.Now = time.Now
	}
	if len(options.WebhookSecrets) > 8 {
		return nil, ErrInvalid
	}
	for _, secret := range options.WebhookSecrets {
		if secret == "" || len(secret) > 4096 || strings.ContainsAny(secret, "\r\n\x00") {
			return nil, ErrInvalid
		}
	}
	options.WebhookSecrets = append([]string(nil), options.WebhookSecrets...)
	root, cancel := context.WithCancel(context.Background())
	return &Helper{options: options, secure: redirect.Scheme == "https", root: root, cancel: cancel, flows: map[string]flow{}, sessions: map[string]session{}}, nil
}

// Close cancels in-flight work and clears private flow and session state.
func (h *Helper) Close() error {
	if h == nil {
		return nil
	}
	h.mutex.Lock()
	defer h.mutex.Unlock()
	if !h.closed {
		h.closed = true
		h.cancel()
		clear(h.flows)
		clear(h.sessions)
		clear(h.options.WebhookSecrets)
	}
	return nil
}
func (h *Helper) request(r *http.Request) (context.Context, func(), bool) {
	if h == nil || r == nil {
		return nil, func() {}, false
	}
	h.mutex.Lock()
	closed := h.closed
	h.mutex.Unlock()
	if closed {
		return nil, func() {}, false
	}
	ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
	stop := context.AfterFunc(h.root, cancel)
	return ctx, func() { stop(); cancel() }, true
}
func method(w http.ResponseWriter, r *http.Request, want string) bool {
	if r == nil || r.Method != want {
		w.Header().Set("Allow", want)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return false
	}
	w.Header().Set("Cache-Control", "no-store")
	return true
}
func uniqueCookie(r *http.Request, name string) (*http.Cookie, error) {
	if r == nil {
		return nil, ErrInvalid
	}
	matches, total := 0, 0
	for _, header := range r.Header.Values("Cookie") {
		total += len(header)
		if total > 8192 {
			return nil, ErrInvalid
		}
		for _, part := range strings.Split(header, ";") {
			key, _, hasValue := strings.Cut(strings.TrimSpace(part), "=")
			if strings.TrimSpace(key) == name {
				matches++
				if !hasValue {
					return nil, ErrInvalid
				}
			}
		}
	}
	if matches != 1 {
		return nil, ErrInvalid
	}
	var found *http.Cookie
	for _, cookie := range r.Cookies() {
		if cookie.Name == name {
			if found != nil {
				return nil, ErrInvalid
			}
			found = cookie
		}
	}
	if found == nil {
		return nil, ErrInvalid
	}
	return found, nil
}
func opaque() (string, error) {
	var value [32]byte
	if _, err := rand.Read(value[:]); err != nil {
		return "", ErrInvalid
	}
	return base64.RawURLEncoding.EncodeToString(value[:]), nil
}
func (h *Helper) cookie(w http.ResponseWriter, name, value string, ttl time.Duration) {
	//nolint:gosec // Secure follows the helper option so plain-HTTP local development works.
	http.SetCookie(w, &http.Cookie{Name: name, Value: value, Path: "/", HttpOnly: true, Secure: h.secure, SameSite: http.SameSiteLaxMode, MaxAge: int(ttl.Seconds())})
}
func (h *Helper) prune(now time.Time) {
	for key, item := range h.flows {
		if !now.Before(item.expires) {
			delete(h.flows, key)
		}
	}
	for key, item := range h.sessions {
		if !now.Before(item.expires) {
			delete(h.sessions, key)
		}
	}
}
func authorizationValid(auth sdk.AuthorizationResult) bool {
	endpoint, err := url.Parse(auth.URL)
	if err != nil || endpoint.Scheme != "https" || endpoint.Hostname() == "" || endpoint.User != nil || endpoint.Fragment != "" || len(auth.URL) > 8192 {
		return false
	}
	_, pkceErr := sdk.PKCEChallenge(auth.CodeVerifier)
	return pkceErr == nil
}

// Start creates one bounded flow and redirects to the SDK authorization URL.
func (h *Helper) Start(w http.ResponseWriter, r *http.Request) {
	if !method(w, r, http.MethodGet) {
		return
	}
	ctx, done, ok := h.request(r)
	defer done()
	if !ok || ctx.Err() != nil {
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		return
	}
	state, err := opaque()
	if err != nil {
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		return
	}
	auth, err := h.options.SSO.GetAuthorizationURL(sdk.AuthorizationOptions{RedirectURI: h.options.RedirectURI, State: state, Organization: h.options.Organization, Connection: h.options.Connection})
	if err != nil || !authorizationValid(auth) || ctx.Err() != nil {
		http.Error(w, "Upstream unavailable", http.StatusBadGateway)
		return
	}
	h.mutex.Lock()
	h.prune(h.options.Now())
	if h.closed || len(h.flows) >= h.options.Capacity {
		h.mutex.Unlock()
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		return
	}
	h.flows[state] = flow{verifier: auth.CodeVerifier, expires: h.options.Now().Add(10 * time.Minute)}
	h.mutex.Unlock()
	h.cookie(w, "ohauthcat_flow", state, 10*time.Minute)
	http.Redirect(w, r, auth.URL, http.StatusSeeOther)
}

// Callback consumes state before exactly one code exchange and creates a private session.
func (h *Helper) Callback(w http.ResponseWriter, r *http.Request) {
	if !method(w, r, http.MethodGet) {
		return
	}
	ctx, done, ok := h.request(r)
	defer done()
	if !ok || ctx.Err() != nil {
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		return
	}
	if r.URL == nil || len(r.URL.RawQuery) > 8192 {
		http.Error(w, "Invalid callback", http.StatusBadRequest)
		return
	}
	query, parseErr := url.ParseQuery(r.URL.RawQuery)
	cookie, cookieErr := uniqueCookie(r, "ohauthcat_flow")
	state, code := query.Get("state"), query.Get("code")
	if parseErr != nil || cookieErr != nil || len(query) != 2 || len(query["state"]) != 1 || len(query["code"]) != 1 || len(state) != 43 || code == "" || len(code) > 4096 || strings.ContainsAny(code, "\r\n\x00") || subtle.ConstantTimeCompare([]byte(state), []byte(cookie.Value)) != 1 {
		http.Error(w, "Invalid callback", http.StatusBadRequest)
		return
	}
	h.mutex.Lock()
	h.prune(h.options.Now())
	pending, exists := h.flows[state]
	delete(h.flows, state)
	h.mutex.Unlock()
	if !exists {
		http.Error(w, "Restart sign-in", http.StatusConflict)
		return
	}
	result, exchangeErr := h.options.SSO.GetProfileAndToken(ctx, sdk.ExchangeOptions{Code: code, CodeVerifier: pending.verifier, RedirectURI: h.options.RedirectURI})
	if exchangeErr != nil {
		var restart *sdk.RestartSignIn
		if errors.As(exchangeErr, &restart) {
			http.Error(w, "Restart sign-in", http.StatusConflict)
		} else {
			http.Error(w, "Upstream unavailable", http.StatusBadGateway)
		}
		return
	}
	claims, verifyErr := h.options.SSO.VerifyAccessToken(ctx, result.AccessToken)
	if verifyErr != nil || claims.Environment != h.options.Environment || claims.Organization != h.options.Organization || claims.Connection != h.options.Connection || !profileID.MatchString(claims.Subject) {
		http.Error(w, "Upstream unavailable", http.StatusBadGateway)
		return
	}
	profile, profileErr := h.options.SSO.GetSSOProfile(ctx, sdk.GetSSOProfileArgs{AccessToken: result.AccessToken})
	if profileErr != nil || string(profile.ID) != claims.Subject || string(profile.OrganizationID) != claims.Organization || string(profile.ConnectionID) != claims.Connection || string(result.Profile.ID) != claims.Subject || string(result.Profile.OrganizationID) != claims.Organization || string(result.Profile.ConnectionID) != claims.Connection || ctx.Err() != nil {
		http.Error(w, "Upstream unavailable", http.StatusBadGateway)
		return
	}
	token, tokenErr := opaque()
	if tokenErr != nil {
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		return
	}
	identity := Identity{Environment: claims.Environment, Organization: claims.Organization, Connection: claims.Connection, Profile: claims.Subject}
	h.mutex.Lock()
	h.prune(h.options.Now())
	if h.closed || len(h.sessions) >= h.options.Capacity {
		h.mutex.Unlock()
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		return
	}
	h.sessions[token] = session{identity: identity, expires: h.options.Now().Add(time.Hour)}
	h.mutex.Unlock()
	h.cookie(w, "ohauthcat_session", token, time.Hour)
	h.cookie(w, "ohauthcat_flow", "", -time.Second)
	w.WriteHeader(http.StatusNoContent)
}

// Profile returns only the locally established closed identity.
func (h *Helper) Profile(w http.ResponseWriter, r *http.Request) {
	if !method(w, r, http.MethodGet) {
		return
	}
	ctx, done, ok := h.request(r)
	defer done()
	if !ok || ctx.Err() != nil {
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		return
	}
	cookie, err := uniqueCookie(r, "ohauthcat_session")
	if err != nil || len(cookie.Value) != 43 {
		http.Error(w, "Sign-in required", http.StatusUnauthorized)
		return
	}
	h.mutex.Lock()
	h.prune(h.options.Now())
	current, exists := h.sessions[cookie.Value]
	h.mutex.Unlock()
	if !exists {
		http.Error(w, "Sign-in required", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(current.identity)
}

// Webhook verifies exact bounded raw bytes before invoking the typed queue.
// Queue implementations must honor context cancellation and deduplicate by event ID.
func (h *Helper) Webhook(w http.ResponseWriter, r *http.Request) {
	if !method(w, r, http.MethodPost) {
		return
	}
	ctx, done, ok := h.request(r)
	defer done()
	if !ok || ctx.Err() != nil || h.options.Queue == nil || len(h.options.WebhookSecrets) == 0 {
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		return
	}
	if r.Body == nil {
		http.Error(w, "Invalid webhook", http.StatusBadRequest)
		return
	}
	defer func() { _ = r.Body.Close() }()
	stopClose := context.AfterFunc(ctx, func() { _ = r.Body.Close() })
	defer stopClose()
	body, readErr := io.ReadAll(io.LimitReader(r.Body, (1<<20)+1))
	if readErr != nil || len(body) > 1<<20 || ctx.Err() != nil {
		http.Error(w, "Invalid webhook", http.StatusBadRequest)
		return
	}
	h.mutex.Lock()
	secrets := append([]string(nil), h.options.WebhookSecrets...)
	h.mutex.Unlock()
	now := h.options.Now().Unix()
	event, verifyErr := sdk.VerifyWebhook(body, r.Header.Get("OhAuthCat-Signature"), secrets, sdk.VerifyOptions{Now: &now})
	clear(secrets)
	if verifyErr != nil || string(event.Environment) != h.options.Environment || (event.Organization != nil && string(*event.Organization) != h.options.Organization) {
		http.Error(w, "Invalid webhook", http.StatusBadRequest)
		return
	}
	h.mutex.Lock()
	closed := h.closed
	h.mutex.Unlock()
	if closed || ctx.Err() != nil {
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		return
	}
	if queueErr := h.options.Queue(ctx, event); queueErr != nil || ctx.Err() != nil {
		http.Error(w, "Upstream unavailable", http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
