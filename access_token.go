package ohauthcat

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

// AccessTokenClaims is the locally verified v1 access-token identity.
type AccessTokenClaims struct {
	Issuer       string `json:"iss"`
	Audience     string `json:"aud"`
	Subject      string `json:"sub"`
	Environment  string `json:"env"`
	Organization string `json:"org"`
	Connection   string `json:"conn"`
	IssuedAt     int64  `json:"iat"`
	NotBefore    int64  `json:"nbf"`
	ExpiresAt    int64  `json:"exp"`
	JWTID        string `json:"jti"`
}

// AccessTokenVerificationError reports a rejected access token without detail.
type AccessTokenVerificationError struct{}

func (*AccessTokenVerificationError) Error() string { return "Access token verification failed" }
func tokenFailure() error                           { return &AccessTokenVerificationError{} }

type jwtCache struct {
	mutex          sync.Mutex
	keys           map[string]*ecdsa.PublicKey
	expires        time.Time
	unknownRefresh time.Time
	failedRefresh  time.Time
	pending        chan struct{}
	clock          func() time.Time
}

func (cache *jwtCache) now() time.Time {
	if cache.clock != nil {
		return cache.clock()
	}
	return time.Now()
}

func decodeTokenSegment(segment string) ([]byte, error) {
	if segment == "" {
		return nil, tokenFailure()
	}
	decoded, err := base64.RawURLEncoding.DecodeString(segment)
	if err != nil || base64.RawURLEncoding.EncodeToString(decoded) != segment {
		return nil, tokenFailure()
	}
	return decoded, nil
}

func decodeTokenObject(segment string) (map[string]json.RawMessage, error) {
	decoded, err := decodeTokenSegment(segment)
	if err != nil || !utf8.Valid(decoded) {
		return nil, tokenFailure()
	}
	fields, err := objectFields(decoded)
	if err != nil {
		return nil, tokenFailure()
	}
	return fields, nil
}

func parseJWKS(data []byte) (map[string]*ecdsa.PublicKey, error) {
	if len(data) > 65536 || !utf8.Valid(data) {
		return nil, tokenFailure()
	}
	fields, err := objectFields(data)
	if err != nil {
		return nil, tokenFailure()
	}
	var entries []json.RawMessage
	if json.Unmarshal(fields["keys"], &entries) != nil || len(entries) == 0 || len(entries) > 32 {
		return nil, tokenFailure()
	}
	keys := make(map[string]*ecdsa.PublicKey)
	for _, entry := range entries {
		fields, err := objectFields(entry)
		if err != nil {
			return nil, tokenFailure()
		}
		var key struct {
			KTY       string `json:"kty"`
			Curve     string `json:"crv"`
			Algorithm string `json:"alg"`
			Use       string `json:"use"`
			ID        string `json:"kid"`
			X         string `json:"x"`
			Y         string `json:"y"`
		}
		if json.Unmarshal(entry, &key) != nil || key.KTY != "EC" || key.Curve != "P-256" || key.Algorithm != "ES256" || key.Use != "sig" || key.ID == "" || len(key.ID) > 128 || keys[key.ID] != nil {
			return nil, tokenFailure()
		}
		if _, exists := fields["d"]; exists {
			return nil, tokenFailure()
		}
		x, errX := decodeTokenSegment(key.X)
		y, errY := decodeTokenSegment(key.Y)
		if errX != nil || errY != nil || len(x) != 32 || len(y) != 32 {
			return nil, tokenFailure()
		}
		point := append(append([]byte{4}, x...), y...)
		public, err := ecdsa.ParseUncompressedPublicKey(elliptic.P256(), point)
		if err != nil {
			return nil, tokenFailure()
		}
		keys[key.ID] = public
	}
	return keys, nil
}

func jwksTTL(control string) time.Duration {
	ttl := int64(300)
	for _, raw := range strings.Split(control, ",") {
		directive := strings.ToLower(strings.TrimSpace(raw))
		if directive == "no-store" || directive == "no-cache" {
			return 0
		}
		if strings.HasPrefix(directive, "max-age=") {
			value, err := strconv.ParseUint(strings.TrimPrefix(directive, "max-age="), 10, 64)
			if err != nil {
				return 0
			}
			if value < 300 {
				ttl = int64(value)
			}
		}
	}
	return time.Duration(ttl) * time.Second
}

func (c *Client) fetchJWKS() (map[string]*ecdsa.PublicKey, time.Duration, error) {
	ctx, cancel := context.WithTimeout(c.root, c.config.Timeout)
	defer cancel()
	c.standby.probe(ctx, c.http)
	for attempt := 0; attempt < 2; attempt++ {
		host, onStandby, revision, _ := c.standby.snapshot(true)
		host, originErr := origin(host, true)
		if originErr != nil {
			return nil, 0, tokenFailure()
		}
		request, err := http.NewRequestWithContext(ctx, http.MethodGet, host+"/.well-known/jwks.json", http.NoBody)
		if err != nil {
			return nil, 0, tokenFailure()
		}
		response, sendErr := c.http.Do(request)
		if response != nil && response.StatusCode >= 400 && (response.StatusCode < 520 || response.StatusCode > 539) {
			if !onStandby {
				c.standby.success(revision)
			}
			if response.Body != nil {
				_ = response.Body.Close()
			}
			return nil, 0, tokenFailure()
		}
		edge := sendErr != nil && c.root.Err() == nil
		var data []byte
		var readErr, closeErr error
		if response != nil && response.Body != nil {
			if sendErr == nil {
				data, readErr = io.ReadAll(io.LimitReader(response.Body, 65537))
			}
			closeErr = response.Body.Close()
			edge = edge || (response.StatusCode >= 520 && response.StatusCode <= 539) || ((readErr != nil || closeErr != nil || ctx.Err() != nil) && c.root.Err() == nil)
		}
		if edge {
			if onStandby {
				c.standby.unavailable(revision)
				return nil, 0, tokenFailure()
			}
			result := rawResponse{}
			if response != nil {
				result.Status = response.StatusCode
				result.RequestID = safeRequestID(response.Header.Get("Request-Id"))
			}
			code := "transport_error"
			if ctx.Err() == context.DeadlineExceeded {
				code = "request_timeout"
			}
			switched := c.standby.edge(revision, responseFailure(code, result), result)
			if switched && attempt == 0 && ctx.Err() == nil {
				continue
			}
			return nil, 0, tokenFailure()
		}
		if sendErr != nil || response == nil || response.Body == nil || response.StatusCode != 200 || ctx.Err() != nil || len(data) > 65536 {
			return nil, 0, tokenFailure()
		}
		keys, err := parseJWKS(data)
		if err != nil {
			return nil, 0, err
		}
		if !onStandby {
			c.standby.success(revision)
		}
		return keys, jwksTTL(response.Header.Get("Cache-Control")), nil
	}
	return nil, 0, tokenFailure()
}

func (c *Client) tokenKey(ctx context.Context, kid string) (*ecdsa.PublicKey, error) {
	cache := &c.jwt
	cache.mutex.Lock()
	now := cache.now()
	key := cache.keys[kid]
	if cache.pending != nil {
		pending := cache.pending
		cache.mutex.Unlock()
		select {
		case <-pending:
		case <-ctx.Done():
			return nil, tokenFailure()
		case <-c.root.Done():
			return nil, tokenFailure()
		}
		cache.mutex.Lock()
		key = cache.keys[kid]
		failed := !cache.failedRefresh.IsZero()
		if key == nil {
			cache.unknownRefresh = cache.now()
		}
		cache.mutex.Unlock()
		if key == nil || failed {
			return nil, tokenFailure()
		}
		return key, nil
	}
	if key != nil && now.Before(cache.expires) {
		cache.mutex.Unlock()
		return key, nil
	}
	if key == nil && !cache.unknownRefresh.IsZero() && now.Sub(cache.unknownRefresh) < 30*time.Second {
		cache.mutex.Unlock()
		return nil, tokenFailure()
	}
	if !cache.failedRefresh.IsZero() && now.Sub(cache.failedRefresh) < 30*time.Second {
		cache.mutex.Unlock()
		return nil, tokenFailure()
	}
	if key == nil && cache.keys != nil {
		cache.unknownRefresh = now
	}
	pending := make(chan struct{})
	cache.pending = pending
	cache.mutex.Unlock()
	go func() {
		keys, ttl, err := c.fetchJWKS()
		cache.mutex.Lock()
		defer cache.mutex.Unlock()
		if err == nil && c.root.Err() == nil {
			cache.keys = keys
			cache.expires = cache.now().Add(ttl)
			cache.failedRefresh = time.Time{}
		} else {
			cache.failedRefresh = cache.now()
			cache.unknownRefresh = cache.now()
		}
		cache.pending = nil
		close(pending)
	}()
	select {
	case <-pending:
	case <-ctx.Done():
		return nil, tokenFailure()
	case <-c.root.Done():
		return nil, tokenFailure()
	}
	// The completed fetch's keys are usable for this verification even with a zero cache lifetime.
	cache.mutex.Lock()
	key = cache.keys[kid]
	failed := !cache.failedRefresh.IsZero()
	if key == nil {
		cache.unknownRefresh = cache.now()
	}
	cache.mutex.Unlock()
	if failed || key == nil {
		return nil, tokenFailure()
	}
	return key, nil
}

// VerifyAccessToken verifies ES256 against the fixed configured issuer and client/environment.
func (r *SSOResources) VerifyAccessToken(ctx context.Context, token string) (AccessTokenClaims, error) {
	var claims AccessTokenClaims
	if r == nil || r.client == nil || ctx == nil {
		return claims, tokenFailure()
	}
	c := r.client
	if c.config.ClientID == "" || c.config.EnvironmentID == "" || c.root.Err() != nil || ctx.Err() != nil || token == "" || len(token) > 4096 {
		return claims, tokenFailure()
	}
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return claims, tokenFailure()
	}
	header, err := decodeTokenObject(parts[0])
	if err != nil {
		return claims, err
	}
	var algorithm, typ, kid string
	if json.Unmarshal(header["alg"], &algorithm) != nil || algorithm != "ES256" || json.Unmarshal(header["typ"], &typ) != nil || typ != "at+jwt" || json.Unmarshal(header["kid"], &kid) != nil || kid == "" || len(kid) > 128 {
		return claims, tokenFailure()
	}
	if _, exists := header["crit"]; exists {
		return claims, tokenFailure()
	}
	if _, exists := header["b64"]; exists {
		return claims, tokenFailure()
	}
	signature, err := decodeTokenSegment(parts[2])
	if err != nil || len(signature) != 64 {
		return claims, tokenFailure()
	}
	key, err := c.tokenKey(ctx, kid)
	if err != nil {
		return claims, err
	}
	digest := sha256.Sum256([]byte(parts[0] + "." + parts[1]))
	if !ecdsa.Verify(key, digest[:], new(big.Int).SetBytes(signature[:32]), new(big.Int).SetBytes(signature[32:])) {
		return claims, tokenFailure()
	}
	payload, err := decodeTokenSegment(parts[1])
	if err != nil || !utf8.Valid(payload) {
		return claims, tokenFailure()
	}
	fields, err := objectFields(payload)
	if err != nil || json.Unmarshal(payload, &claims) != nil {
		return AccessTokenClaims{}, tokenFailure()
	}
	for _, name := range []string{"iss", "aud", "sub", "env", "org", "conn", "jti"} {
		var value string
		if json.Unmarshal(fields[name], &value) != nil || value == "" {
			return AccessTokenClaims{}, tokenFailure()
		}
	}
	for _, name := range []string{"iat", "nbf", "exp"} {
		var value int64
		raw := strings.TrimSpace(string(fields[name]))
		if raw == "" || raw[0] < '0' || raw[0] > '9' || json.Unmarshal(fields[name], &value) != nil || value < 0 || value > 9007199254740991 {
			return AccessTokenClaims{}, tokenFailure()
		}
	}
	now := c.jwt.now().Unix()
	if claims.Issuer != c.config.Issuer || claims.Audience != c.config.ClientID || claims.Environment != c.config.EnvironmentID || now >= claims.ExpiresAt+60 || now+60 < claims.NotBefore || now+60 < claims.IssuedAt || claims.ExpiresAt <= claims.IssuedAt || claims.NotBefore > claims.ExpiresAt {
		return AccessTokenClaims{}, tokenFailure()
	}
	if ctx.Err() != nil || c.root.Err() != nil {
		return AccessTokenClaims{}, tokenFailure()
	}
	return claims, nil
}
