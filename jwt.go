package ohauthcat

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"math"
	"math/big"
	"net/http"
	"net/url"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"
)

// TokenVerificationError is every JWT verification failure; it never says
// which check failed.
type TokenVerificationError struct{}

func (*TokenVerificationError) Error() string { return "Token verification failed" }

var errToken = &TokenVerificationError{}

const (
	maxJWT              = 16384
	maxJWKS             = 65536
	maxJWKSCache        = 300 * time.Second
	unknownKidInterval  = 30 * time.Second
	defaultJWTClockSkew = 60 * time.Second
)

// JWTClaims are the verified claims of a token.
type JWTClaims map[string]JSONValue

// String returns a string claim, or "" when absent or not a string.
func (c JWTClaims) String(name string) string {
	var value string
	if raw, ok := c[name]; ok && json.Unmarshal(raw.data, &value) == nil {
		return value
	}
	return ""
}

// Strings returns a string-array claim and whether it is present and valid.
func (c JWTClaims) Strings(name string) ([]string, bool) {
	raw, ok := c[name]
	if !ok {
		return nil, false
	}
	var values []string
	if json.Unmarshal(raw.data, &values) != nil {
		return nil, false
	}
	return values, true
}

// Bool returns a boolean claim, false when absent.
func (c JWTClaims) Bool(name string) bool {
	var value bool
	if raw, ok := c[name]; ok && json.Unmarshal(raw.data, &value) == nil {
		return value
	}
	return false
}

// JWTVerifierConfig configures an ES256 JWT verifier over a cached JWKS.
type JWTVerifierConfig struct {
	Issuer         string
	Audience       []string
	JWKSURI        string
	Types          []string
	RequiredClaims []string
	// AllowMissingExpiry accepts tokens without exp (Security Event Tokens).
	AllowMissingExpiry bool
	HTTPClient         *http.Client
	Timeout            time.Duration
	ClockSkew          time.Duration
	Now                func() time.Time
}

// JWTVerifier verifies ES256 JWTs. Keys are cached for max-age (at most five
// minutes); an unknown key ID refetches at most every 30 seconds.
type JWTVerifier struct {
	config      JWTVerifierConfig
	mu          sync.Mutex
	keys        map[string]*ecdsa.PublicKey
	expires     time.Time
	lastUnknown time.Time
}

func secureURL(value string) error {
	parsed, err := url.Parse(value)
	if err != nil || parsed.Host == "" {
		return errors.New("URL must be absolute")
	}
	local := parsed.Scheme == "http" && (parsed.Hostname() == "localhost" || parsed.Hostname() == "127.0.0.1")
	if parsed.Scheme != "https" && !local {
		return errors.New("URL must use https")
	}
	return nil
}

// NewJWTVerifier validates the configuration and returns a verifier.
func NewJWTVerifier(config JWTVerifierConfig) (*JWTVerifier, error) {
	if strings.TrimSpace(config.Issuer) == "" {
		return nil, errors.New("issuer is required")
	}
	if len(config.Audience) == 0 || slices.Contains(config.Audience, "") {
		return nil, errors.New("audience is required")
	}
	if len(config.Types) == 0 {
		return nil, errors.New("types is required")
	}
	if err := secureURL(config.JWKSURI); err != nil {
		return nil, errors.New("jwks URI must be an https URL")
	}
	if config.HTTPClient == nil {
		config.HTTPClient = &http.Client{CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}
	}
	if config.Timeout <= 0 {
		config.Timeout = 5 * time.Second
	}
	if config.ClockSkew <= 0 {
		config.ClockSkew = defaultJWTClockSkew
	}
	if config.Now == nil {
		config.Now = time.Now
	}
	return &JWTVerifier{config: config}, nil
}

func decodeSegment(segment string) ([]byte, error) {
	raw, err := base64.RawURLEncoding.Strict().DecodeString(segment)
	if err != nil || base64.RawURLEncoding.EncodeToString(raw) != segment {
		return nil, errToken
	}
	return raw, nil
}

type es256JWK struct {
	Kty string  `json:"kty"`
	Crv string  `json:"crv"`
	Kid string  `json:"kid"`
	Alg *string `json:"alg"`
	Use *string `json:"use"`
	X   string  `json:"x"`
	Y   string  `json:"y"`
	D   *string `json:"d"`
}

// parseES256JWKS reads ES256 keys; a malformed or private key refuses the set.
func parseES256JWKS(data []byte) (map[string]*ecdsa.PublicKey, error) {
	var set struct {
		Keys []es256JWK `json:"keys"`
	}
	if json.Unmarshal(data, &set) != nil || set.Keys == nil || len(set.Keys) > 32 {
		return nil, errToken
	}
	result := map[string]*ecdsa.PublicKey{}
	for _, key := range set.Keys {
		if key.D != nil {
			return nil, errToken
		}
		if key.Kty != "EC" || key.Crv != "P-256" || (key.Use != nil && *key.Use != "sig") || (key.Alg != nil && *key.Alg != "ES256") {
			continue
		}
		if key.Kid == "" || len(key.Kid) > 128 || result[key.Kid] != nil {
			return nil, errToken
		}
		x, xerr := decodeSegment(key.X)
		y, yerr := decodeSegment(key.Y)
		if xerr != nil || yerr != nil || len(x) != 32 || len(y) != 32 {
			return nil, errToken
		}
		// ParseUncompressedPublicKey rejects points that are not on the curve.
		public, perr := ecdsa.ParseUncompressedPublicKey(elliptic.P256(), append(append([]byte{4}, x...), y...))
		if perr != nil {
			return nil, errToken
		}
		result[key.Kid] = public
	}
	return result, nil
}

var maxAgePattern = regexp.MustCompile(`(?:^|,)\s*max-age=(\d+)\s*(?:,|$)`)
var noStorePattern = regexp.MustCompile(`(?i)(?:^|,)\s*(?:no-store|no-cache)\s*(?:,|$)`)

func (v *JWTVerifier) load(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, v.config.Timeout)
	defer cancel()
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, v.config.JWKSURI, http.NoBody)
	if err != nil {
		return errToken
	}
	response, err := v.config.HTTPClient.Do(request)
	if err != nil {
		return errToken
	}
	defer func() { _ = response.Body.Close() }()
	body, err := io.ReadAll(io.LimitReader(response.Body, maxJWKS+1))
	if err != nil || response.StatusCode != http.StatusOK || len(body) > maxJWKS {
		return errToken
	}
	keys, err := parseES256JWKS(body)
	if err != nil {
		return err
	}
	ttl := maxJWKSCache
	control := response.Header.Get("Cache-Control")
	if noStorePattern.MatchString(control) {
		ttl = 0
	} else if match := maxAgePattern.FindStringSubmatch(control); match != nil {
		if seconds, perr := strconv.Atoi(match[1]); perr == nil && time.Duration(seconds)*time.Second < ttl {
			ttl = time.Duration(seconds) * time.Second
		}
	}
	v.keys, v.expires = keys, v.config.Now().Add(ttl)
	return nil
}

func (v *JWTVerifier) key(ctx context.Context, kid string) *ecdsa.PublicKey {
	v.mu.Lock()
	defer v.mu.Unlock()
	now := v.config.Now()
	if !now.Before(v.expires) {
		if v.load(ctx) != nil {
			return nil
		}
	} else if v.keys[kid] == nil && now.Sub(v.lastUnknown) >= unknownKidInterval {
		v.lastUnknown = now
		if v.load(ctx) != nil {
			return nil
		}
	}
	return v.keys[kid]
}

// Verify checks the token and returns its claims, or a
// *TokenVerificationError.
func (v *JWTVerifier) Verify(ctx context.Context, token string) (JWTClaims, error) {
	if v == nil || token == "" || len(token) > maxJWT {
		return nil, errToken
	}
	parts := strings.Split(token, ".")
	if len(parts) != 3 || parts[0] == "" || parts[1] == "" || parts[2] == "" {
		return nil, errToken
	}
	headerBytes, err := decodeSegment(parts[0])
	if err != nil {
		return nil, errToken
	}
	var header map[string]json.RawMessage
	if json.Unmarshal(headerBytes, &header) != nil {
		return nil, errToken
	}
	var alg, typ, kid string
	if json.Unmarshal(header["alg"], &alg) != nil || alg != "ES256" ||
		json.Unmarshal(header["typ"], &typ) != nil || !slices.Contains(v.config.Types, typ) ||
		json.Unmarshal(header["kid"], &kid) != nil || kid == "" {
		return nil, errToken
	}
	if _, crit := header["crit"]; crit {
		return nil, errToken
	}
	signature, err := decodeSegment(parts[2])
	if err != nil || len(signature) != 64 {
		return nil, errToken
	}
	key := v.key(ctx, kid)
	if key == nil {
		return nil, errToken
	}
	sum := sha256.Sum256([]byte(parts[0] + "." + parts[1]))
	if !ecdsa.Verify(key, sum[:], new(big.Int).SetBytes(signature[:32]), new(big.Int).SetBytes(signature[32:])) {
		return nil, errToken
	}
	payload, err := decodeSegment(parts[1])
	if err != nil {
		return nil, errToken
	}
	var claims JWTClaims
	if json.Unmarshal(payload, &claims) != nil || claims == nil || claims.String("iss") != v.config.Issuer {
		return nil, errToken
	}
	audiences, ok := claims.Strings("aud")
	if !ok {
		if single := claims.String("aud"); single != "" {
			audiences = []string{single}
		}
	}
	if !slices.ContainsFunc(audiences, func(value string) bool { return slices.Contains(v.config.Audience, value) }) {
		return nil, errToken
	}
	now := float64(v.config.Now().Unix())
	skew := v.config.ClockSkew.Seconds()
	moment := func(name string) (float64, bool, error) {
		raw, present := claims[name]
		if !present {
			return 0, false, nil
		}
		var value float64
		if json.Unmarshal(raw.data, &value) != nil || value < 0 || value != math.Trunc(value) {
			return 0, false, errToken
		}
		return value, true, nil
	}
	exp, hasExp, expErr := moment("exp")
	nbf, hasNbf, nbfErr := moment("nbf")
	iat, hasIat, iatErr := moment("iat")
	if expErr != nil || nbfErr != nil || iatErr != nil ||
		(!hasExp && !v.config.AllowMissingExpiry) ||
		(hasExp && now >= exp+skew) || (hasNbf && now+skew < nbf) || (hasIat && now+skew < iat) {
		return nil, errToken
	}
	for _, name := range v.config.RequiredClaims {
		if claims.String(name) == "" {
			return nil, errToken
		}
	}
	return claims, nil
}
