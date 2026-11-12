package existingauth

import (
	"context"
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"
)

// ---------------------------------------------------------------- Auth0

// Auth0Options configures Auth0Linker. Set ManagementToken, or ClientID and
// ClientSecret for the client-credentials grant.
type Auth0Options struct {
	// Domain is the tenant domain (example.us.auth0.com) or an https origin.
	Domain string
	// Connection is the database connection linked users live in.
	Connection      string
	ManagementToken string
	ClientID        string
	ClientSecret    string
	HTTPClient      *http.Client
}

// Auth0User is a linked Auth0 user.
type Auth0User struct {
	UserID  string `json:"user_id"`
	Email   string `json:"email"`
	Blocked bool   `json:"blocked"`
}

// Auth0Linker uses the Auth0 Management API. Linked users are created in a
// database connection with user_id set to the link key, so the Auth0 user ID
// is "auth0|<key>" and lookups are exact. Scopes: read:users, create:users,
// update:users, delete:sessions.
type Auth0Linker struct {
	opts   Auth0Options
	origin string
	mu     sync.Mutex
	token  string
	expiry time.Time
}

// NewAuth0Linker validates the options.
func NewAuth0Linker(opts Auth0Options) (*Auth0Linker, error) {
	if err := required(ProviderAuth0, opts.Domain, "Domain"); err != nil {
		return nil, err
	}
	domain := opts.Domain
	if !strings.Contains(domain, "://") {
		domain = "https://" + domain
	}
	origin, err := httpsOrigin(ProviderAuth0, domain, "Domain")
	if err != nil {
		return nil, err
	}
	if err := required(ProviderAuth0, opts.Connection, "Connection"); err != nil {
		return nil, err
	}
	if opts.ManagementToken == "" {
		if err := required(ProviderAuth0, opts.ClientID, "ClientID"); err != nil {
			return nil, err
		}
		if err := required(ProviderAuth0, opts.ClientSecret, "ClientSecret"); err != nil {
			return nil, err
		}
	}
	return &Auth0Linker{opts: opts, origin: origin}, nil
}

// Provider implements Linker.
func (a *Auth0Linker) Provider() Provider { return ProviderAuth0 }

// Auth0UserID returns the Auth0 user ID a linked profile maps to.
func Auth0UserID(p LinkProfile) string { return "auth0|" + p.Key }

func (a *Auth0Linker) bearer(ctx context.Context) (string, error) {
	if a.opts.ManagementToken != "" {
		return a.opts.ManagementToken, nil
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.token != "" && time.Now().Before(a.expiry) {
		return a.token, nil
	}
	r, err := sendJSON(ctx, ProviderAuth0, a.opts.HTTPClient, http.MethodPost, a.origin+"/oauth/token", map[string]string{}, map[string]string{
		"grant_type": "client_credentials", "client_id": a.opts.ClientID,
		"client_secret": a.opts.ClientSecret, "audience": a.origin + "/api/v2/",
	})
	if err != nil {
		return "", err
	}
	if r.status != http.StatusOK {
		return "", failure(ProviderAuth0, r, "")
	}
	var body struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	if !r.decode(&body) || body.AccessToken == "" {
		return "", badResponse(ProviderAuth0)
	}
	a.token = body.AccessToken
	a.expiry = time.Now().Add(time.Duration(max(0, body.ExpiresIn-60)) * time.Second)
	return a.token, nil
}

func (a *Auth0Linker) call(ctx context.Context, method, path string, body any) (reply, error) {
	token, err := a.bearer(ctx)
	if err != nil {
		return reply{}, err
	}
	headers := map[string]string{"authorization": "Bearer " + token, "accept": "application/json"}
	return sendJSON(ctx, ProviderAuth0, a.opts.HTTPClient, method, a.origin+path, headers, body)
}

func auth0User(r reply) (Auth0User, error) {
	var u Auth0User
	if !r.decode(&u) || u.UserID == "" {
		return Auth0User{}, badResponse(ProviderAuth0)
	}
	return u, nil
}

// Find implements Linker.
func (a *Auth0Linker) Find(ctx context.Context, p LinkProfile) (Auth0User, bool, error) {
	r, err := a.call(ctx, http.MethodGet, "/api/v2/users/"+url.PathEscape(Auth0UserID(p)), nil)
	if err != nil {
		return Auth0User{}, false, err
	}
	if r.status == http.StatusNotFound {
		return Auth0User{}, false, nil
	}
	if r.status != http.StatusOK {
		return Auth0User{}, false, failure(ProviderAuth0, r, "")
	}
	u, err := auth0User(r)
	return u, err == nil, err
}

// Create implements Linker.
func (a *Auth0Linker) Create(ctx context.Context, p LinkProfile) (Auth0User, error) {
	if p.Email == "" {
		return Auth0User{}, invalid(ProviderAuth0, "an email is required to create an Auth0 database user")
	}
	random := make([]byte, 32)
	if _, err := rand.Read(random); err != nil {
		return Auth0User{}, invalid(ProviderAuth0, "random source failed")
	}
	body := map[string]any{
		"connection": a.opts.Connection, "user_id": p.Key, "email": p.Email,
		"email_verified": p.EmailVerified, "verify_email": false,
		// Random password nobody knows: sign-in happens only through SSO.
		"password":     base64.RawURLEncoding.EncodeToString(random) + "Aa1!",
		"app_metadata": metadata(p),
	}
	if name := displayName(p); name != "" {
		body["name"] = name
	}
	r, err := a.call(ctx, http.MethodPost, "/api/v2/users", body)
	if err != nil {
		return Auth0User{}, err
	}
	if r.status == http.StatusConflict {
		return Auth0User{}, newError(ProviderAuth0, CodeEmailConflict, "an Auth0 user with this email or ID already exists", r.status, "")
	}
	if r.status != http.StatusCreated && r.status != http.StatusOK {
		return Auth0User{}, failure(ProviderAuth0, r, "")
	}
	return auth0User(r)
}

// Revoke blocks the user and deletes their Auth0 sessions (for dsync.user.deleted).
func (a *Auth0Linker) Revoke(ctx context.Context, userID string) error {
	if err := required(ProviderAuth0, userID, "userID"); err != nil {
		return err
	}
	path := "/api/v2/users/" + url.PathEscape(userID)
	r, err := a.call(ctx, http.MethodPatch, path, map[string]bool{"blocked": true})
	if err != nil {
		return err
	}
	if r.status == http.StatusNotFound {
		return nil
	}
	if r.status != http.StatusOK {
		return failure(ProviderAuth0, r, "")
	}
	r, err = a.call(ctx, http.MethodDelete, path+"/sessions", nil)
	if err != nil {
		return err
	}
	if r.status != http.StatusOK && r.status != http.StatusAccepted && r.status != http.StatusNoContent {
		return failure(ProviderAuth0, r, "")
	}
	return nil
}

// ---------------------------------------------------------------- Cognito

// CognitoReply is the status and raw JSON of one Cognito action.
type CognitoReply struct {
	Status int
	Body   json.RawMessage
}

// CognitoTransport sends one Cognito JSON 1.1 action. Use SigV4Transport, or
// wrap the AWS SDK to reuse its credential chain.
type CognitoTransport func(ctx context.Context, action string, input map[string]any) (CognitoReply, error)

// AWSCredentials are static AWS credentials.
type AWSCredentials struct {
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
}

// SigV4Options configures SigV4Transport.
type SigV4Options struct {
	Region      string
	Credentials AWSCredentials
	// Endpoint overrides https://cognito-idp.<region>.amazonaws.com.
	Endpoint   string
	HTTPClient *http.Client
	Now        func() time.Time
}

var regionPattern = regexp.MustCompile(`^[a-z]{2}(-[a-z]+)+-\d$`)

func hmacSHA256(key []byte, data string) []byte {
	m := hmac.New(sha256.New, key)
	m.Write([]byte(data))
	return m.Sum(nil)
}

func sha256Hex(data []byte) string {
	s := sha256.Sum256(data)
	return hex.EncodeToString(s[:])
}

// SigV4Transport is a minimal AWS Signature Version 4 transport for Cognito,
// so the SDK needs no AWS module.
func SigV4Transport(opts SigV4Options) (CognitoTransport, error) {
	if !regionPattern.MatchString(opts.Region) {
		return nil, invalid(ProviderCognito, "region is invalid")
	}
	if err := required(ProviderCognito, opts.Credentials.AccessKeyID, "AccessKeyID"); err != nil {
		return nil, err
	}
	if err := required(ProviderCognito, opts.Credentials.SecretAccessKey, "SecretAccessKey"); err != nil {
		return nil, err
	}
	endpoint := opts.Endpoint
	if endpoint == "" {
		endpoint = "https://cognito-idp." + opts.Region + ".amazonaws.com"
	}
	origin, err := httpsOrigin(ProviderCognito, endpoint, "Endpoint")
	if err != nil {
		return nil, err
	}
	host := strings.SplitN(origin, "://", 2)[1]
	now := opts.Now
	if now == nil {
		now = time.Now
	}
	return func(ctx context.Context, action string, input map[string]any) (CognitoReply, error) {
		body, err := json.Marshal(input)
		if err != nil {
			return CognitoReply{}, invalid(ProviderCognito, "input could not be encoded")
		}
		amzDate := now().UTC().Format("20060102T150405Z")
		day := amzDate[:8]
		headers := map[string]string{
			"content-type": "application/x-amz-json-1.1",
			"host":         host,
			"x-amz-date":   amzDate,
			"x-amz-target": "AWSCognitoIdentityProviderService." + action,
		}
		if opts.Credentials.SessionToken != "" {
			headers["x-amz-security-token"] = opts.Credentials.SessionToken
		}
		names := make([]string, 0, len(headers))
		for n := range headers {
			names = append(names, n)
		}
		sort.Strings(names)
		var canonicalHeaders strings.Builder
		for _, n := range names {
			canonicalHeaders.WriteString(n + ":" + headers[n] + "\n")
		}
		signed := strings.Join(names, ";")
		canonical := strings.Join([]string{"POST", "/", "", canonicalHeaders.String(), signed, sha256Hex(body)}, "\n")
		scope := day + "/" + opts.Region + "/cognito-idp/aws4_request"
		toSign := strings.Join([]string{"AWS4-HMAC-SHA256", amzDate, scope, sha256Hex([]byte(canonical))}, "\n")
		k := hmacSHA256([]byte("AWS4"+opts.Credentials.SecretAccessKey), day)
		k = hmacSHA256(hmacSHA256(hmacSHA256(k, opts.Region), "cognito-idp"), "aws4_request")
		sig := hex.EncodeToString(hmacSHA256(k, toSign))
		delete(headers, "host")
		headers["authorization"] = "AWS4-HMAC-SHA256 Credential=" + opts.Credentials.AccessKeyID + "/" + scope +
			", SignedHeaders=" + signed + ", Signature=" + sig
		r, err := send(ctx, ProviderCognito, opts.HTTPClient, http.MethodPost, origin+"/", headers, body)
		if err != nil {
			return CognitoReply{}, err
		}
		return CognitoReply{Status: r.status, Body: r.body}, nil
	}, nil
}

// CognitoOptions configures CognitoLinker.
type CognitoOptions struct {
	UserPoolID string
	Transport  CognitoTransport
	// ClientID is the app client used for sessions (custom auth enabled).
	ClientID string
	// ClientSecret is sent as SECRET_HASH when the client has one.
	ClientSecret string
}

// CognitoUser is a linked Cognito user.
type CognitoUser struct {
	Username string
	Enabled  bool
}

// CognitoSession holds the tokens Cognito issued.
type CognitoSession struct {
	AccessToken  string
	IDToken      string
	RefreshToken string
	ExpiresIn    int
	TokenType    string
}

// CognitoLinker uses Amazon Cognito user pools. The Cognito username is the
// link key, so the pool must accept plain usernames (email may be an alias,
// not the username attribute). Sessions use admin-initiated CUSTOM_AUTH.
//
// IAM on the pool ARN: cognito-idp:AdminGetUser, AdminCreateUser,
// AdminInitiateAuth, AdminRespondToAuthChallenge, AdminUserGlobalSignOut,
// AdminDisableUser.
type CognitoLinker struct{ opts CognitoOptions }

var poolPattern = regexp.MustCompile(`^[\w-]+_[0-9a-zA-Z]+$`)

// NewCognitoLinker validates the options.
func NewCognitoLinker(opts CognitoOptions) (*CognitoLinker, error) {
	if !poolPattern.MatchString(opts.UserPoolID) {
		return nil, invalid(ProviderCognito, "UserPoolID is invalid")
	}
	if opts.Transport == nil {
		return nil, invalid(ProviderCognito, "Transport is required")
	}
	return &CognitoLinker{opts: opts}, nil
}

// Provider implements Linker.
func (c *CognitoLinker) Provider() Provider { return ProviderCognito }

func (c *CognitoLinker) call(ctx context.Context, action string, input map[string]any) (reply, string, error) {
	input["UserPoolId"] = c.opts.UserPoolID
	r, err := c.opts.Transport(ctx, action, input)
	if err != nil {
		var typed *Error
		if errors.As(err, &typed) {
			return reply{}, "", err
		}
		return reply{}, "", newError(ProviderCognito, CodeProviderUnavailable, "cognito request failed", 0, "")
	}
	out := reply{status: r.Status, body: r.Body}
	var e struct {
		Type string `json:"__type"`
	}
	out.decode(&e)
	parts := strings.Split(e.Type, "#")
	return out, parts[len(parts)-1], nil
}

func (c *CognitoLinker) secretHash(username string) map[string]string {
	if c.opts.ClientID == "" || c.opts.ClientSecret == "" {
		return map[string]string{}
	}
	return map[string]string{"SECRET_HASH": base64.StdEncoding.EncodeToString(hmacSHA256([]byte(c.opts.ClientSecret), username+c.opts.ClientID))}
}

// Find implements Linker.
func (c *CognitoLinker) Find(ctx context.Context, p LinkProfile) (CognitoUser, bool, error) {
	r, kind, err := c.call(ctx, "AdminGetUser", map[string]any{"Username": p.Key})
	if err != nil {
		return CognitoUser{}, false, err
	}
	if r.status == http.StatusBadRequest && kind == "UserNotFoundException" {
		return CognitoUser{}, false, nil
	}
	if r.status != http.StatusOK {
		return CognitoUser{}, false, failure(ProviderCognito, r, kind)
	}
	var body struct {
		Username string `json:"Username"`
		Enabled  *bool  `json:"Enabled"`
	}
	if !r.decode(&body) || body.Username == "" {
		return CognitoUser{}, false, badResponse(ProviderCognito)
	}
	return CognitoUser{Username: body.Username, Enabled: body.Enabled == nil || *body.Enabled}, true, nil
}

// Create implements Linker.
func (c *CognitoLinker) Create(ctx context.Context, p LinkProfile) (CognitoUser, error) {
	attrs := []map[string]string{}
	add := func(name, value string) {
		if value != "" {
			attrs = append(attrs, map[string]string{"Name": name, "Value": value})
		}
	}
	add("email", p.Email)
	if p.Email != "" && p.EmailVerified {
		add("email_verified", "true")
	}
	add("given_name", p.FirstName)
	add("family_name", p.LastName)
	r, kind, err := c.call(ctx, "AdminCreateUser", map[string]any{"Username": p.Key, "MessageAction": "SUPPRESS", "UserAttributes": attrs})
	if err != nil {
		return CognitoUser{}, err
	}
	if kind == "UsernameExistsException" || kind == "AliasExistsException" {
		return CognitoUser{}, newError(ProviderCognito, CodeEmailConflict, "a Cognito user with this username or email alias already exists", r.status, kind)
	}
	if r.status != http.StatusOK {
		return CognitoUser{}, failure(ProviderCognito, r, kind)
	}
	var body struct {
		User struct {
			Username string `json:"Username"`
			Enabled  *bool  `json:"Enabled"`
		} `json:"User"`
	}
	if !r.decode(&body) || body.User.Username == "" {
		return CognitoUser{}, badResponse(ProviderCognito)
	}
	return CognitoUser{Username: body.User.Username, Enabled: body.User.Enabled == nil || *body.User.Enabled}, nil
}

// StartSession starts a Cognito session through admin-initiated CUSTOM_AUTH.
// Your auth challenge Lambda triggers must accept challengeAnswer.
func (c *CognitoLinker) StartSession(ctx context.Context, user CognitoUser, challengeAnswer string) (CognitoSession, error) {
	for _, f := range [][2]string{{c.opts.ClientID, "ClientID"}, {challengeAnswer, "challengeAnswer"}, {user.Username, "Username"}} {
		if err := required(ProviderCognito, f[0], f[1]); err != nil {
			return CognitoSession{}, err
		}
	}
	params := c.secretHash(user.Username)
	params["USERNAME"] = user.Username
	r, kind, err := c.call(ctx, "AdminInitiateAuth", map[string]any{"ClientId": c.opts.ClientID, "AuthFlow": "CUSTOM_AUTH", "AuthParameters": params})
	if err != nil {
		return CognitoSession{}, err
	}
	if r.status != http.StatusOK {
		return CognitoSession{}, failure(ProviderCognito, r, kind)
	}
	type result struct {
		AuthenticationResult *struct {
			AccessToken  string `json:"AccessToken"`
			IDToken      string `json:"IdToken"`
			RefreshToken string `json:"RefreshToken"`
			ExpiresIn    int    `json:"ExpiresIn"`
			TokenType    string `json:"TokenType"`
		} `json:"AuthenticationResult"`
		ChallengeName string `json:"ChallengeName"`
		Session       string `json:"Session"`
	}
	var started result
	r.decode(&started)
	if started.ChallengeName != "" {
		if started.ChallengeName != "CUSTOM_CHALLENGE" || started.Session == "" {
			return CognitoSession{}, newError(ProviderCognito, CodeSessionRefused, "unexpected Cognito challenge", 0, "")
		}
		answers := c.secretHash(user.Username)
		answers["USERNAME"] = user.Username
		answers["ANSWER"] = challengeAnswer
		r, kind, err = c.call(ctx, "AdminRespondToAuthChallenge", map[string]any{
			"ClientId": c.opts.ClientID, "ChallengeName": "CUSTOM_CHALLENGE", "Session": started.Session, "ChallengeResponses": answers,
		})
		if err != nil {
			return CognitoSession{}, err
		}
		if kind == "NotAuthorizedException" {
			return CognitoSession{}, newError(ProviderCognito, CodeSessionRefused, "Cognito refused the challenge answer", r.status, kind)
		}
		if r.status != http.StatusOK {
			return CognitoSession{}, failure(ProviderCognito, r, kind)
		}
	}
	var done result
	if !r.decode(&done) || done.AuthenticationResult == nil {
		return CognitoSession{}, newError(ProviderCognito, CodeSessionRefused, "Cognito did not issue tokens", 0, "")
	}
	a := done.AuthenticationResult
	if a.AccessToken == "" || a.IDToken == "" {
		return CognitoSession{}, badResponse(ProviderCognito)
	}
	if a.TokenType == "" {
		a.TokenType = "Bearer"
	}
	return CognitoSession{AccessToken: a.AccessToken, IDToken: a.IDToken, RefreshToken: a.RefreshToken, ExpiresIn: a.ExpiresIn, TokenType: a.TokenType}, nil
}

// Revoke signs the user out everywhere and disables them (for dsync.user.deleted).
func (c *CognitoLinker) Revoke(ctx context.Context, username string) error {
	if err := required(ProviderCognito, username, "username"); err != nil {
		return err
	}
	for _, action := range []string{"AdminUserGlobalSignOut", "AdminDisableUser"} {
		r, kind, err := c.call(ctx, action, map[string]any{"Username": username})
		if err != nil {
			return err
		}
		if kind == "UserNotFoundException" {
			return nil
		}
		if r.status != http.StatusOK {
			return failure(ProviderCognito, r, kind)
		}
	}
	return nil
}

// ---------------------------------------------------------------- Firebase

// FirebaseUser is a linked Firebase user.
type FirebaseUser struct {
	UID      string
	Disabled bool
}

// FirebaseNewUser is what FirebaseUserStore.CreateUser receives.
type FirebaseNewUser struct {
	UID           string
	Email         string
	EmailVerified bool
	DisplayName   string
}

// ErrFirebaseConflict is returned by a FirebaseUserStore when the uid or
// email already exists.
var ErrFirebaseConflict = errors.New("firebase user already exists")

// FirebaseUserStore is the subset of the Firebase Admin SDK's auth client the
// helper uses. GetUser returns (zero, false, nil) when no user has the uid.
type FirebaseUserStore interface {
	GetUser(ctx context.Context, uid string) (FirebaseUser, bool, error)
	CreateUser(ctx context.Context, user FirebaseNewUser) (FirebaseUser, error)
}

// FirebaseOptions configures FirebaseLinker.
type FirebaseOptions struct {
	ClientEmail string
	// PrivateKeyPEM is the service account's private_key.
	PrivateKeyPEM string
	Users         FirebaseUserStore
	Now           func() time.Time
}

// FirebaseLinker uses Firebase Authentication. The uid is the link key.
// StartSession mints a custom token for signInWithCustomToken.
type FirebaseLinker struct {
	opts FirebaseOptions
	key  *rsa.PrivateKey
}

const firebaseAudience = "https://identitytoolkit.googleapis.com/google.identity.identitytoolkit.v1.IdentityToolkit"

var reservedClaims = map[string]bool{
	"acr": true, "amr": true, "at_hash": true, "aud": true, "auth_time": true, "azp": true, "cnf": true, "c_hash": true,
	"exp": true, "firebase": true, "iat": true, "iss": true, "jti": true, "nbf": true, "nonce": true, "sub": true,
}

var uidPattern = regexp.MustCompile(`^[\x21-\x7e]{1,128}$`)

// NewFirebaseLinker validates the options and parses the key.
func NewFirebaseLinker(opts FirebaseOptions) (*FirebaseLinker, error) {
	if err := required(ProviderFirebase, opts.ClientEmail, "ClientEmail"); err != nil {
		return nil, err
	}
	block, _ := pem.Decode([]byte(opts.PrivateKeyPEM))
	if block == nil {
		return nil, invalid(ProviderFirebase, "PrivateKeyPEM is not a private key")
	}
	var key *rsa.PrivateKey
	if parsed, err := x509.ParsePKCS8PrivateKey(block.Bytes); err == nil {
		key, _ = parsed.(*rsa.PrivateKey)
	} else if k, err := x509.ParsePKCS1PrivateKey(block.Bytes); err == nil {
		key = k
	}
	if key == nil {
		return nil, invalid(ProviderFirebase, "PrivateKeyPEM must be an RSA private key")
	}
	if opts.Users == nil {
		return nil, invalid(ProviderFirebase, "Users is required")
	}
	if opts.Now == nil {
		opts.Now = time.Now
	}
	return &FirebaseLinker{opts: opts, key: key}, nil
}

// Provider implements Linker.
func (f *FirebaseLinker) Provider() Provider { return ProviderFirebase }

func storeError(err error) error {
	var typed *Error
	switch {
	case errors.As(err, &typed):
		return err
	case errors.Is(err, ErrFirebaseConflict):
		return newError(ProviderFirebase, CodeEmailConflict, "a Firebase user with this email or uid already exists", 0, "")
	default:
		return newError(ProviderFirebase, CodeProviderError, "firebase user store failed", 0, "")
	}
}

// Find implements Linker.
func (f *FirebaseLinker) Find(ctx context.Context, p LinkProfile) (FirebaseUser, bool, error) {
	u, ok, err := f.opts.Users.GetUser(ctx, p.Key)
	if err != nil {
		return FirebaseUser{}, false, storeError(err)
	}
	return u, ok, nil
}

// Create implements Linker.
func (f *FirebaseLinker) Create(ctx context.Context, p LinkProfile) (FirebaseUser, error) {
	u, err := f.opts.Users.CreateUser(ctx, FirebaseNewUser{UID: p.Key, Email: p.Email, EmailVerified: p.EmailVerified, DisplayName: displayName(p)})
	if err != nil {
		return FirebaseUser{}, storeError(err)
	}
	return u, nil
}

// StartSession mints a Firebase custom token (RS256, one hour) for the uid.
func (f *FirebaseLinker) StartSession(user FirebaseUser, claims map[string]any) (string, error) {
	if !uidPattern.MatchString(user.UID) {
		return "", invalid(ProviderFirebase, "uid must be 1 to 128 printable characters")
	}
	if user.Disabled {
		return "", newError(ProviderFirebase, CodeSessionRefused, "the Firebase user is disabled", 0, "")
	}
	for name := range claims {
		if reservedClaims[name] {
			return "", invalid(ProviderFirebase, "claim "+name+" is reserved")
		}
	}
	now := f.opts.Now().Unix()
	payload := map[string]any{"iss": f.opts.ClientEmail, "sub": f.opts.ClientEmail, "aud": firebaseAudience, "iat": now, "exp": now + 3600, "uid": user.UID}
	if len(claims) > 0 {
		payload["claims"] = claims
	}
	head, _ := json.Marshal(map[string]string{"alg": "RS256", "typ": "JWT"})
	body, err := json.Marshal(payload)
	if err != nil {
		return "", invalid(ProviderFirebase, "claims could not be encoded")
	}
	input := base64.RawURLEncoding.EncodeToString(head) + "." + base64.RawURLEncoding.EncodeToString(body)
	digest := sha256.Sum256([]byte(input))
	sig, err := rsa.SignPKCS1v15(nil, f.key, crypto.SHA256, digest[:])
	if err != nil {
		return "", invalid(ProviderFirebase, "custom token could not be signed")
	}
	return input + "." + base64.RawURLEncoding.EncodeToString(sig), nil
}

// ---------------------------------------------------------------- Supabase

// SupabaseOptions configures SupabaseLinker.
type SupabaseOptions struct {
	// URL is the project URL, e.g. https://<ref>.supabase.co.
	URL string
	// ServiceRoleKey is the service role key. Server-side only.
	ServiceRoleKey string
	HTTPClient     *http.Client
}

// SupabaseUser is a linked Supabase user.
type SupabaseUser struct {
	ID     string
	Email  string
	Banned bool
}

// SupabaseSession is a Supabase Auth session.
type SupabaseSession struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

// SupabaseMagicLink is a generated magic link.
type SupabaseMagicLink struct {
	ActionLink string
	TokenHash  string
}

// SupabaseLinker uses the Supabase Auth admin API. The user ID is a UUID
// derived from the link key.
type SupabaseLinker struct {
	opts   SupabaseOptions
	origin string
}

// NewSupabaseLinker validates the options.
func NewSupabaseLinker(opts SupabaseOptions) (*SupabaseLinker, error) {
	origin, err := httpsOrigin(ProviderSupabase, opts.URL, "URL")
	if err != nil {
		return nil, err
	}
	if err := required(ProviderSupabase, opts.ServiceRoleKey, "ServiceRoleKey"); err != nil {
		return nil, err
	}
	return &SupabaseLinker{opts: opts, origin: origin}, nil
}

// Provider implements Linker.
func (s *SupabaseLinker) Provider() Provider { return ProviderSupabase }

func (s *SupabaseLinker) call(ctx context.Context, method, path string, body any) (reply, error) {
	headers := map[string]string{"apikey": s.opts.ServiceRoleKey, "authorization": "Bearer " + s.opts.ServiceRoleKey, "accept": "application/json"}
	return sendJSON(ctx, ProviderSupabase, s.opts.HTTPClient, method, s.origin+"/auth/v1"+path, headers, body)
}

func supabaseUser(r reply) (SupabaseUser, error) {
	var body struct {
		ID          string `json:"id"`
		Email       string `json:"email"`
		BannedUntil string `json:"banned_until"`
	}
	if !r.decode(&body) || body.ID == "" {
		return SupabaseUser{}, badResponse(ProviderSupabase)
	}
	banned := false
	if body.BannedUntil != "" {
		until, err := time.Parse(time.RFC3339Nano, body.BannedUntil)
		banned = err != nil || until.After(time.Now())
	}
	return SupabaseUser{ID: body.ID, Email: body.Email, Banned: banned}, nil
}

func supabaseCode(r reply) string {
	var body struct {
		ErrorCode string `json:"error_code"`
	}
	r.decode(&body)
	return body.ErrorCode
}

// Find implements Linker.
func (s *SupabaseLinker) Find(ctx context.Context, p LinkProfile) (SupabaseUser, bool, error) {
	r, err := s.call(ctx, http.MethodGet, "/admin/users/"+LinkUUID(p.Key), nil)
	if err != nil {
		return SupabaseUser{}, false, err
	}
	if r.status == http.StatusNotFound {
		return SupabaseUser{}, false, nil
	}
	if r.status != http.StatusOK {
		return SupabaseUser{}, false, failure(ProviderSupabase, r, supabaseCode(r))
	}
	u, err := supabaseUser(r)
	return u, err == nil, err
}

// Create implements Linker.
func (s *SupabaseLinker) Create(ctx context.Context, p LinkProfile) (SupabaseUser, error) {
	if p.Email == "" {
		return SupabaseUser{}, invalid(ProviderSupabase, "an email is required to create a Supabase user")
	}
	body := map[string]any{"id": LinkUUID(p.Key), "email": p.Email, "email_confirm": p.EmailVerified, "app_metadata": metadata(p)}
	if name := displayName(p); name != "" {
		body["user_metadata"] = map[string]string{"full_name": name}
	}
	r, err := s.call(ctx, http.MethodPost, "/admin/users", body)
	if err != nil {
		return SupabaseUser{}, err
	}
	code := supabaseCode(r)
	if r.status == http.StatusUnprocessableEntity || code == "email_exists" || code == "user_already_exists" {
		return SupabaseUser{}, newError(ProviderSupabase, CodeEmailConflict, "a Supabase user with this email or ID already exists", r.status, code)
	}
	if r.status != http.StatusOK && r.status != http.StatusCreated {
		return SupabaseUser{}, failure(ProviderSupabase, r, code)
	}
	return supabaseUser(r)
}

// MagicLink generates a magic link for the user's email.
func (s *SupabaseLinker) MagicLink(ctx context.Context, user SupabaseUser, redirectTo string) (SupabaseMagicLink, error) {
	if err := required(ProviderSupabase, user.Email, "user.Email"); err != nil {
		return SupabaseMagicLink{}, err
	}
	if user.Banned {
		return SupabaseMagicLink{}, newError(ProviderSupabase, CodeSessionRefused, "the Supabase user is banned", 0, "")
	}
	body := map[string]string{"type": "magiclink", "email": user.Email}
	if redirectTo != "" {
		body["redirect_to"] = redirectTo
	}
	r, err := s.call(ctx, http.MethodPost, "/admin/generate_link", body)
	if err != nil {
		return SupabaseMagicLink{}, err
	}
	if r.status != http.StatusOK {
		return SupabaseMagicLink{}, failure(ProviderSupabase, r, supabaseCode(r))
	}
	type link struct {
		ActionLink  string `json:"action_link"`
		HashedToken string `json:"hashed_token"`
	}
	var out struct {
		link
		Properties *link `json:"properties"`
	}
	r.decode(&out)
	l := out.link
	if out.Properties != nil {
		l = *out.Properties
	}
	if l.ActionLink == "" || l.HashedToken == "" {
		return SupabaseMagicLink{}, badResponse(ProviderSupabase)
	}
	return SupabaseMagicLink{ActionLink: l.ActionLink, TokenHash: l.HashedToken}, nil
}

// StartSession generates a magic link and verifies it server-side.
func (s *SupabaseLinker) StartSession(ctx context.Context, user SupabaseUser) (SupabaseSession, error) {
	link, err := s.MagicLink(ctx, user, "")
	if err != nil {
		return SupabaseSession{}, err
	}
	r, err := sendJSON(ctx, ProviderSupabase, s.opts.HTTPClient, http.MethodPost, s.origin+"/auth/v1/verify",
		map[string]string{"apikey": s.opts.ServiceRoleKey}, map[string]string{"type": "magiclink", "token_hash": link.TokenHash})
	if err != nil {
		return SupabaseSession{}, err
	}
	if r.status == http.StatusUnauthorized || r.status == http.StatusForbidden {
		return SupabaseSession{}, newError(ProviderSupabase, CodeSessionRefused, "Supabase refused the magic link", r.status, "")
	}
	if r.status != http.StatusOK {
		return SupabaseSession{}, failure(ProviderSupabase, r, supabaseCode(r))
	}
	var session SupabaseSession
	if !r.decode(&session) || session.AccessToken == "" || session.RefreshToken == "" {
		return SupabaseSession{}, badResponse(ProviderSupabase)
	}
	if session.TokenType == "" {
		session.TokenType = "bearer"
	}
	return session, nil
}

// Revoke bans the user so no session can be refreshed (for dsync.user.deleted).
func (s *SupabaseLinker) Revoke(ctx context.Context, userID string) error {
	if err := required(ProviderSupabase, userID, "userID"); err != nil {
		return err
	}
	r, err := s.call(ctx, http.MethodPut, "/admin/users/"+url.PathEscape(userID), map[string]string{"ban_duration": "876000h"})
	if err != nil {
		return err
	}
	if r.status == http.StatusNotFound || r.status == http.StatusOK {
		return nil
	}
	return failure(ProviderSupabase, r, supabaseCode(r))
}

// ---------------------------------------------------------------- Clerk

// ClerkOptions configures ClerkLinker.
type ClerkOptions struct {
	// SecretKey is the instance's secret key (sk_live_… / sk_test_…).
	SecretKey string
	// APIURL overrides https://api.clerk.com.
	APIURL     string
	HTTPClient *http.Client
}

// ClerkUser is a linked Clerk user.
type ClerkUser struct {
	ID     string `json:"id"`
	Banned bool   `json:"banned"`
}

// ClerkSignInToken is a sign-in token for the ticket strategy.
type ClerkSignInToken struct {
	Token string `json:"token"`
	URL   string `json:"url"`
}

// ClerkLinker uses the Clerk Backend API. The user's external_id is the link key.
type ClerkLinker struct {
	opts   ClerkOptions
	origin string
}

var clerkKeyPattern = regexp.MustCompile(`^sk_(live|test)_\S+$`)

// NewClerkLinker validates the options.
func NewClerkLinker(opts ClerkOptions) (*ClerkLinker, error) {
	if !clerkKeyPattern.MatchString(opts.SecretKey) {
		return nil, invalid(ProviderClerk, "SecretKey must be a Clerk secret key")
	}
	api := opts.APIURL
	if api == "" {
		api = "https://api.clerk.com"
	}
	origin, err := httpsOrigin(ProviderClerk, api, "APIURL")
	if err != nil {
		return nil, err
	}
	return &ClerkLinker{opts: opts, origin: origin}, nil
}

// Provider implements Linker.
func (c *ClerkLinker) Provider() Provider { return ProviderClerk }

func (c *ClerkLinker) call(ctx context.Context, method, path string, body any) (reply, string, error) {
	headers := map[string]string{"authorization": "Bearer " + c.opts.SecretKey, "accept": "application/json"}
	r, err := sendJSON(ctx, ProviderClerk, c.opts.HTTPClient, method, c.origin+"/v1"+path, headers, body)
	if err != nil {
		return reply{}, "", err
	}
	var e struct {
		Errors []struct {
			Code string `json:"code"`
		} `json:"errors"`
	}
	code := ""
	if r.decode(&e) && len(e.Errors) > 0 {
		code = e.Errors[0].Code
	}
	return r, code, nil
}

// Find implements Linker.
func (c *ClerkLinker) Find(ctx context.Context, p LinkProfile) (ClerkUser, bool, error) {
	r, code, err := c.call(ctx, http.MethodGet, "/users?external_id="+url.QueryEscape(p.Key)+"&limit=2", nil)
	if err != nil {
		return ClerkUser{}, false, err
	}
	if r.status != http.StatusOK {
		return ClerkUser{}, false, failure(ProviderClerk, r, code)
	}
	var users []ClerkUser
	if !r.decode(&users) || len(users) > 1 {
		return ClerkUser{}, false, badResponse(ProviderClerk)
	}
	if len(users) == 0 {
		return ClerkUser{}, false, nil
	}
	if users[0].ID == "" {
		return ClerkUser{}, false, badResponse(ProviderClerk)
	}
	return users[0], true, nil
}

// Create implements Linker.
func (c *ClerkLinker) Create(ctx context.Context, p LinkProfile) (ClerkUser, error) {
	body := map[string]any{"external_id": p.Key, "skip_password_requirement": true, "private_metadata": metadata(p)}
	if p.Email != "" {
		body["email_address"] = []string{p.Email}
	}
	if p.FirstName != "" {
		body["first_name"] = p.FirstName
	}
	if p.LastName != "" {
		body["last_name"] = p.LastName
	}
	r, code, err := c.call(ctx, http.MethodPost, "/users", body)
	if err != nil {
		return ClerkUser{}, err
	}
	if code == "form_identifier_exists" || code == "external_id_exists" {
		return ClerkUser{}, newError(ProviderClerk, CodeEmailConflict, "a Clerk user with this email or external ID already exists", r.status, code)
	}
	if r.status != http.StatusOK {
		return ClerkUser{}, failure(ProviderClerk, r, code)
	}
	var u ClerkUser
	if !r.decode(&u) || u.ID == "" {
		return ClerkUser{}, badResponse(ProviderClerk)
	}
	return u, nil
}

// StartSession creates a sign-in token; expiresIn 0 means 300 seconds.
func (c *ClerkLinker) StartSession(ctx context.Context, user ClerkUser, expiresIn time.Duration) (ClerkSignInToken, error) {
	if err := required(ProviderClerk, user.ID, "user.ID"); err != nil {
		return ClerkSignInToken{}, err
	}
	if user.Banned {
		return ClerkSignInToken{}, newError(ProviderClerk, CodeSessionRefused, "the Clerk user is banned", 0, "")
	}
	if expiresIn == 0 {
		expiresIn = 300 * time.Second
	}
	seconds := int(expiresIn / time.Second)
	if seconds < 1 || seconds > 2_592_000 {
		return ClerkSignInToken{}, invalid(ProviderClerk, "expiresIn must be 1 s to 30 days")
	}
	r, code, err := c.call(ctx, http.MethodPost, "/sign_in_tokens", map[string]any{"user_id": user.ID, "expires_in_seconds": seconds})
	if err != nil {
		return ClerkSignInToken{}, err
	}
	if r.status != http.StatusOK {
		return ClerkSignInToken{}, failure(ProviderClerk, r, code)
	}
	var t ClerkSignInToken
	if !r.decode(&t) || t.Token == "" {
		return ClerkSignInToken{}, badResponse(ProviderClerk)
	}
	return t, nil
}

// Revoke bans the user, which ends all their sessions (for dsync.user.deleted).
func (c *ClerkLinker) Revoke(ctx context.Context, userID string) error {
	if err := required(ProviderClerk, userID, "userID"); err != nil {
		return err
	}
	r, code, err := c.call(ctx, http.MethodPost, "/users/"+url.PathEscape(userID)+"/ban", nil)
	if err != nil {
		return err
	}
	if r.status == http.StatusNotFound || r.status == http.StatusOK {
		return nil
	}
	return failure(ProviderClerk, r, code)
}
