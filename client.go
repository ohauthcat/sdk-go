package ohauthcat

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

// ClientConfig is copied and validated by NewClient. Custom transports must honor
// request contexts and provide bodies whose Read and Close respect cancellation.
type ClientConfig struct {
	APIKey             string
	APIVersion         string
	APIBaseURL         string
	AuthBaseURL        string
	ClientID           string
	EnvironmentID      string
	Issuer             string
	Timeout            time.Duration
	MaxRetries         *int
	MaxPages           int
	Standby            *bool
	StandbyAPIBaseURL  string
	StandbyAuthBaseURL string
	StandbyThreshold   int
	StandbyCooldown    time.Duration
	HTTPClient         *http.Client
}

// Client owns request cancellation and immutable configuration; Close is idempotent.
type Client struct {
	Resources
	config   ClientConfig
	http     *http.Client
	contract wireContract
	retries  int
	root     context.Context
	cancel   context.CancelFunc
	mutex    sync.Mutex
	closed   bool
	owned    *http.Transport
	jwt      jwtCache
	standby  *standbyState
}

type routeParameter struct {
	Name     string `json:"name"`
	In       string `json:"in"`
	Required bool   `json:"required"`
	Schema   schema `json:"schema"`
	Style    string `json:"style"`
	Explode  *bool  `json:"explode"`
}

type operationRoute struct {
	Path         string           `json:"path"`
	Method       string           `json:"method"`
	Server       string           `json:"server"`
	Auth         string           `json:"auth"`
	Parameters   []routeParameter `json:"parameters"`
	Headers      []routeParameter `json:"headers"`
	Body         *schema          `json:"body"`
	BodyRequired bool             `json:"bodyRequired"`
	Status       int              `json:"status"`
	Media        string           `json:"media"`
	Response     *schema          `json:"response"`
}

type wireContract struct {
	Schemas map[string]schema         `json:"schemas"`
	Routes  map[string]operationRoute `json:"routes"`
}

type requestInput struct {
	Path           []byte
	Query          []byte
	Body           []byte
	IdempotencyKey string
	AccessToken    string
	// RawBody is a non-JSON request body sent with RawMedia; only the
	// operations in rawBodyMedia accept one.
	RawBody  []byte
	RawMedia string
}

// NewClient creates a management client. Access-token verification additionally
// requires explicitly configured ClientID and EnvironmentID.
func NewClient(config ClientConfig) (*Client, error) {
	if !headerText(config.APIKey) {
		return nil, requestFailure("configuration_invalid")
	}
	if config.APIVersion != "" && !regexp.MustCompile(`^1\.\d+$`).MatchString(config.APIVersion) {
		return nil, requestFailure("configuration_invalid")
	}
	for _, target := range []*string{&config.APIBaseURL, &config.AuthBaseURL} {
		if *target != "" {
			normalized, err := origin(*target, false)
			if err != nil {
				return nil, err
			}
			*target = normalized
		}
	}
	if config.Issuer == "" {
		config.Issuer = "https://auth.ohauthcat.com"
	}
	issuer, err := origin(config.Issuer, true)
	if err != nil {
		return nil, err
	}
	config.Issuer = issuer
	if config.ClientID != "" && (!strings.HasPrefix(config.ClientID, "pk_") || len(config.ClientID) <= 3 || !headerText(config.ClientID)) {
		return nil, requestFailure("configuration_invalid")
	}
	if config.Timeout == 0 {
		config.Timeout = 30 * time.Second
	}
	if config.Timeout < 0 || config.Timeout > 5*time.Minute {
		return nil, requestFailure("configuration_invalid")
	}
	retries := 2
	if config.MaxRetries != nil {
		retries = *config.MaxRetries
	}
	if retries < 0 || retries > 5 {
		return nil, requestFailure("configuration_invalid")
	}
	config.MaxRetries = nil
	if config.MaxPages == 0 {
		config.MaxPages = 1000
	}
	if config.MaxPages < 1 || config.MaxPages > 10000 {
		return nil, requestFailure("configuration_invalid")
	}
	var contract wireContract
	if json.Unmarshal([]byte(pinnedContract), &contract) != nil || len(contract.Routes) == 0 || len(contract.Schemas) == 0 {
		return nil, requestFailure("contract_invalid")
	}
	if config.EnvironmentID != "" {
		definition, exists := contract.Schemas["EnvironmentId"]
		encoded, _ := json.Marshal(config.EnvironmentID)
		if !exists || !matchesJSON(encoded, definition, contract.Schemas, false) {
			return nil, requestFailure("configuration_invalid")
		}
	}
	standby, err := newStandby(config)
	if err != nil {
		return nil, err
	}
	config.Standby = nil
	clientHTTP := http.Client{}
	var owned *http.Transport
	if config.HTTPClient != nil {
		clientHTTP = *config.HTTPClient
	}
	if clientHTTP.Transport == nil {
		standard, ok := http.DefaultTransport.(*http.Transport)
		if !ok {
			return nil, requestFailure("configuration_invalid")
		}
		owned = standard.Clone()
		clientHTTP.Transport = owned
	}
	clientHTTP.Jar = nil
	clientHTTP.Timeout = 0
	clientHTTP.CheckRedirect = func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }
	config.HTTPClient = nil
	root, cancel := context.WithCancel(context.Background())
	client := &Client{config: config, http: &clientHTTP, contract: contract, retries: retries, root: root, standby: standby, cancel: cancel, owned: owned}
	client.Resources = makeResources(client)
	return client, nil
}

func origin(value string, secure bool) (string, error) {
	if strings.IndexFunc(value, func(r rune) bool { return r <= 32 || r == 127 }) >= 0 || strings.ContainsAny(value, "?#") {
		return "", requestFailure("configuration_invalid")
	}
	parsed, err := url.Parse(value)
	if err != nil || parsed.Hostname() == "" || parsed.User != nil || (parsed.Scheme != "https" && parsed.Scheme != "http") || (parsed.Path != "" && parsed.Path != "/") || parsed.Opaque != "" || strings.HasSuffix(parsed.Host, ":") {
		return "", requestFailure("configuration_invalid")
	}
	port := parsed.Port()
	if port != "" {
		number, err := strconv.Atoi(port)
		if err != nil || number < 1 || number > 65535 {
			return "", requestFailure("configuration_invalid")
		}
	}
	if secure && parsed.Scheme != "https" {
		ip := net.ParseIP(parsed.Hostname())
		if parsed.Hostname() != "localhost" && (ip == nil || !ip.IsLoopback()) {
			return "", requestFailure("configuration_invalid")
		}
	}
	if (parsed.Scheme == "https" && port == "443") || (parsed.Scheme == "http" && port == "80") {
		parsed.Host = parsed.Hostname()
		if strings.Contains(parsed.Host, ":") {
			parsed.Host = "[" + parsed.Host + "]"
		}
	}
	parsed.Path = ""
	parsed.RawPath = ""
	parsed.Host = strings.ToLower(parsed.Host)
	return parsed.String(), nil
}

func headerText(value string) bool {
	return strings.TrimSpace(value) != "" && strings.IndexFunc(value, func(r rune) bool { return r < 32 || r == 127 }) < 0
}

// Close cancels active requests and releases owned idle connections.
func (c *Client) Close() error {
	if c == nil {
		return nil
	}
	c.mutex.Lock()
	if c.closed {
		c.mutex.Unlock()
		return nil
	}
	c.closed = true
	if c.cancel != nil {
		c.cancel()
	}
	c.mutex.Unlock()
	if c.owned != nil {
		c.owned.CloseIdleConnections()
	}
	return nil
}

func (c *Client) prepare(route operationRoute, input requestInput) (string, http.Header, error) {
	path := route.Path
	pathData, queryData := input.Path, input.Query
	if pathData == nil {
		pathData = []byte("{}")
	}
	if queryData == nil {
		queryData = []byte("{}")
	}
	paths, err := objectFields(pathData)
	if err != nil {
		return "", nil, requestFailure("request_invalid")
	}
	queries, err := objectFields(queryData)
	if err != nil {
		return "", nil, requestFailure("request_invalid")
	}
	allowedPath, allowedQuery := map[string]bool{}, map[string]bool{}
	values := url.Values{}
	for _, parameter := range route.Parameters {
		source := queries
		allowed := allowedQuery
		if parameter.In == "path" {
			source = paths
			allowed = allowedPath
		} else if parameter.In != "query" {
			return "", nil, requestFailure("contract_invalid")
		}
		allowed[parameter.Name] = true
		data, exists := source[parameter.Name]
		if !exists {
			if parameter.Required {
				return "", nil, requestFailure("request_invalid")
			}
			continue
		}
		if !matchesJSON(data, parameter.Schema, c.contract.Schemas, false) {
			return "", nil, requestFailure("request_invalid")
		}
		if parameter.In == "path" {
			text, scalarErr := scalar(data)
			if scalarErr != nil {
				return "", nil, scalarErr
			}
			path = strings.ReplaceAll(path, "{"+parameter.Name+"}", url.PathEscape(text))
		} else {
			items := []json.RawMessage{data}
			if bytes.TrimSpace(data)[0] == '[' {
				if json.Unmarshal(data, &items) != nil {
					return "", nil, requestFailure("request_invalid")
				}
			}
			var texts []string
			for _, item := range items {
				text, scalarErr := scalar(item)
				if scalarErr != nil {
					return "", nil, scalarErr
				}
				texts = append(texts, text)
			}
			if parameter.Style != "" && parameter.Style != "form" {
				return "", nil, requestFailure("contract_invalid")
			}
			if parameter.Explode != nil && !*parameter.Explode {
				values.Add(parameter.Name, strings.Join(texts, ","))
			} else {
				for _, text := range texts {
					values.Add(parameter.Name, text)
				}
			}
		}
	}
	for name := range paths {
		if !allowedPath[name] {
			return "", nil, requestFailure("request_invalid")
		}
	}
	for name := range queries {
		if !allowedQuery[name] {
			return "", nil, requestFailure("request_invalid")
		}
	}
	if strings.ContainsAny(path, "{}") || (route.BodyRequired && input.Body == nil && input.RawBody == nil) || (route.Body == nil && input.Body != nil) {
		return "", nil, requestFailure("request_invalid")
	}
	if input.Body != nil && (route.Body == nil || !matchesJSON(input.Body, *route.Body, c.contract.Schemas, false)) {
		return "", nil, requestFailure("request_invalid")
	}
	base := route.Server
	server, err := url.Parse(base)
	if err != nil {
		return "", nil, requestFailure("contract_invalid")
	}
	if strings.HasPrefix(server.Hostname(), "auth.") {
		if c.config.AuthBaseURL != "" {
			base = c.config.AuthBaseURL
		}
	} else if c.config.APIBaseURL != "" {
		base = c.config.APIBaseURL
	}
	target, err := url.Parse(strings.TrimSuffix(base, "/") + path)
	if err != nil {
		return "", nil, requestFailure("request_invalid")
	}
	target.RawQuery = values.Encode()
	token := c.config.APIKey
	if route.Auth == "accessToken" {
		token = input.AccessToken
		if !headerText(token) {
			return "", nil, requestFailure("request_invalid")
		}
	} else if route.Auth != "secretKey" {
		return "", nil, requestFailure("contract_invalid")
	}
	if input.AccessToken != "" && route.Auth != "accessToken" {
		return "", nil, requestFailure("request_invalid")
	}
	headers := http.Header{"Authorization": {"Bearer " + token}, "Accept": {route.Media}, "User-Agent": {"ohauthcat-sdk-go/0.0.2"}}
	if c.config.APIVersion != "" {
		headers.Set("OhAuthCat-Version", c.config.APIVersion)
	}
	if route.Method != "GET" {
		key := input.IdempotencyKey
		if key == "" {
			key, err = randomUUID()
			if err != nil {
				return "", nil, err
			}
		}
		if !headerText(key) {
			return "", nil, requestFailure("request_invalid")
		}
		headers.Set("Idempotency-Key", key)
	}
	for _, parameter := range route.Headers {
		value := headers.Get(parameter.Name)
		if value == "" {
			if parameter.Required {
				return "", nil, requestFailure("request_invalid")
			}
			continue
		}
		encoded, _ := json.Marshal(value)
		if !matchesJSON(encoded, parameter.Schema, c.contract.Schemas, false) {
			return "", nil, requestFailure("request_invalid")
		}
	}
	if input.Body != nil {
		headers.Set("Content-Type", "application/json")
	} else if input.RawBody != nil {
		headers.Set("Content-Type", input.RawMedia)
	}
	return target.String(), headers, nil
}

func scalar(data []byte) (string, error) {
	data = bytes.TrimSpace(data)
	if len(data) == 0 {
		return "", requestFailure("request_invalid")
	}
	if data[0] == '"' {
		var text string
		if json.Unmarshal(data, &text) == nil {
			return text, nil
		}
	}
	if matchesKind(data, "number") || matchesKind(data, "boolean") {
		return string(data), nil
	}
	return "", requestFailure("request_invalid")
}

func randomUUID() (string, error) {
	var data [16]byte
	if _, err := rand.Read(data[:]); err != nil {
		return "", requestFailure("transport_error")
	}
	data[6] = (data[6] & 15) | 64
	data[8] = (data[8] & 63) | 128
	encoded := hex.EncodeToString(data[:])
	return encoded[:8] + "-" + encoded[8:12] + "-" + encoded[12:16] + "-" + encoded[16:20] + "-" + encoded[20:], nil
}

func (c *Client) request(parent context.Context, operation string, input requestInput) (rawResponse, error) {
	var empty rawResponse
	if c == nil || parent == nil {
		return empty, requestFailure("request_invalid")
	}
	c.mutex.Lock()
	closed := c.closed
	c.mutex.Unlock()
	if closed {
		return empty, requestFailure("client_closed")
	}
	route, exists := c.contract.Routes[operation]
	if !exists {
		return empty, requestFailure("request_invalid")
	}
	if input.RawBody != nil && (input.Body != nil || input.RawMedia == "" || rawBodyMedia[operation] != input.RawMedia) {
		return empty, requestFailure("request_invalid")
	}
	if input.RawBody == nil && input.RawMedia != "" {
		return empty, requestFailure("request_invalid")
	}
	_, headers, err := c.prepare(route, input)
	if err != nil {
		return empty, err
	}
	input.IdempotencyKey = headers.Get("Idempotency-Key")
	ctx, cancel := context.WithTimeout(parent, c.config.Timeout)
	defer cancel()
	stop := context.AfterFunc(c.root, cancel)
	defer stop()
	c.standby.probe(ctx, c.http)
	var primaryError error
	var primaryResponse rawResponse
	for attempt := 0; attempt < 2; attempt++ {
		selected, onStandby, revision, retained := c.standby.snapshot(strings.Contains(route.Server, "auth."))
		result, err := c.requestOne(ctx, operation, input, selected, revision, onStandby)
		if err == nil {
			if !onStandby {
				c.standby.success(revision)
			}
			return result, nil
		}
		var apiError *APIError
		edge := errors.As(err, &apiError) && ((result.Status >= 520 && result.Status <= 539) || ((apiError.Code == "transport_error" || apiError.Code == "request_timeout") && result.Status != 429 && result.Status != 503))
		if !edge || parent.Err() != nil || c.root.Err() != nil {
			if !onStandby && !edge && result.Status >= 400 {
				c.standby.success(revision)
			}
			return result, err
		}
		if onStandby {
			c.standby.unavailable(revision)
			if primaryError != nil {
				return primaryResponse, primaryError
			}
			if retained != nil {
				var original *APIError
				if errors.As(retained, &original) {
					return rawResponse{Status: original.Status, RequestID: original.RequestID}, retained
				}
				return rawResponse{}, retained
			}
			return result, err
		}
		switched := c.standby.edge(revision, err, result)
		if primaryError == nil && switched && retryEligible(operation, route, headers) && ctx.Err() == nil {
			primaryError = err
			primaryResponse = result
			continue
		}
		return result, err
	}
	return empty, requestFailure("transport_error")
}

func (c *Client) requestOne(parent context.Context, operation string, input requestInput, selected string, revision uint64, onStandby bool) (rawResponse, error) {
	var result rawResponse
	if c == nil || parent == nil {
		return result, requestFailure("request_invalid")
	}
	c.mutex.Lock()
	closed := c.closed
	c.mutex.Unlock()
	if closed {
		return result, requestFailure("client_closed")
	}
	route, exists := c.contract.Routes[operation]
	if !exists {
		return result, requestFailure("request_invalid")
	}
	target, headers, err := c.prepare(route, input)
	if err != nil {
		return result, err
	}
	target = selectHost(target, selected)
	ctx, cancel := context.WithTimeout(parent, c.config.Timeout)
	defer cancel()
	stop := context.AfterFunc(c.root, cancel)
	defer stop()
	body := bytes.Clone(input.Body)
	if input.RawBody != nil {
		body = bytes.Clone(input.RawBody)
	}
	for attempt := 0; ; attempt++ {
		if ctx.Err() != nil {
			return result, contextFailure(ctx, result)
		}
		request, err := http.NewRequestWithContext(ctx, route.Method, target, bytes.NewReader(body))
		if err != nil {
			return result, requestFailure("request_invalid")
		}
		request.Header = headers.Clone()
		result = rawResponse{}
		response, err := c.http.Do(request)
		if response != nil {
			result.Status = response.StatusCode
			result.RequestID = safeRequestID(response.Header.Get("Request-Id"))
		}
		if err != nil {
			if response != nil && response.Body != nil {
				_ = response.Body.Close()
			}
			return result, contextFailure(ctx, result)
		}
		if ctx.Err() != nil {
			_ = response.Body.Close()
			return result, contextFailure(ctx, result)
		}
		if !onStandby && (response.StatusCode == 429 || response.StatusCode == 503) {
			c.standby.success(revision)
		}
		if retryEligible(operation, route, headers) && (response.StatusCode == 429 || response.StatusCode == 503) && attempt < c.retries {
			if response.Body.Close() != nil {
				return result, responseFailure("transport_error", result)
			}
			delay := retryDelay(response.Header.Get("Retry-After"), attempt, time.Now())
			deadline, _ := ctx.Deadline()
			if delay >= time.Until(deadline) {
				return result, responseFailure("request_timeout", result)
			}
			if err := waitContext(ctx, delay); err != nil {
				return result, contextFailure(ctx, result)
			}
			continue
		}
		data, readErr := io.ReadAll(io.LimitReader(response.Body, 1024*1024+1))
		closeErr := response.Body.Close()
		if ctx.Err() != nil {
			return result, contextFailure(ctx, result)
		}
		if readErr != nil || closeErr != nil {
			return result, responseFailure("transport_error", result)
		}
		if len(data) > 1024*1024 {
			return result, responseFailure("response_invalid", result)
		}
		result.Body = data
		if response.StatusCode < 200 || response.StatusCode >= 300 {
			return result, c.parseAPIError(result)
		}
		if route.Media == "application/json" && route.Response != nil {
			normalized, err := normalizeProfileFields(data, *route.Response, c.contract.Schemas)
			if err != nil || !matchesJSON(normalized, *route.Response, c.contract.Schemas, true) {
				return result, responseFailure("response_invalid", result)
			}
			result.Body = normalized
		}
		return result, nil
	}
}

func contextFailure(ctx context.Context, response rawResponse) *APIError {
	code := "transport_error"
	if ctx.Err() == context.DeadlineExceeded {
		code = "request_timeout"
	} else if ctx.Err() != nil {
		code = "request_aborted"
	}
	failure := responseFailure(code, response)
	failure.cause = ctx.Err()
	return failure
}

func safeRequestID(value string) string {
	if len(value) > 128 || strings.IndexFunc(value, func(r rune) bool { return r < 32 || r > 126 }) >= 0 {
		return ""
	}
	return strings.TrimSpace(value)
}

func retryEligible(operation string, route operationRoute, headers http.Header) bool {
	if operation == "exchangeSsoCode" {
		return false
	}
	if route.Method == "GET" || route.Method == "HEAD" {
		return true
	}
	for _, parameter := range route.Headers {
		if parameter.Name == "Idempotency-Key" && headers.Get(parameter.Name) != "" {
			return true
		}
	}
	return false
}

func waitContext(ctx context.Context, delay time.Duration) error {
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

func retryDelay(header string, attempt int, now time.Time) time.Duration {
	value := strings.TrimSpace(header)
	if regexp.MustCompile(`^\d+(?:\.\d+)?$`).MatchString(value) {
		seconds, err := strconv.ParseFloat(value, 64)
		if err == nil && !math.IsInf(seconds, 0) && seconds >= 0 {
			if seconds >= float64(math.MaxInt64)/float64(time.Second) {
				return time.Duration(math.MaxInt64)
			}
			return time.Duration(math.Ceil(seconds * float64(time.Second)))
		}
	}
	if date, err := http.ParseTime(value); err == nil && !date.Before(now) {
		return date.Sub(now)
	}
	if attempt < 0 {
		attempt = 0
	}
	if attempt > 4 {
		attempt = 4
	}
	limit := time.Second * time.Duration(1<<attempt)
	if limit > 10*time.Second {
		limit = 10 * time.Second
	}
	var random [1]byte
	_, _ = rand.Read(random[:])
	return limit/2 + time.Duration(float64(limit/2)*float64(random[0])/255)
}

// parseAPIError decodes an error body tolerantly: a field that drifts from the
// published pattern (for example a request ID in another format) must never
// hide the catalogue code, doc URL or request ID. The code still has to be a
// known catalogue API error whose status and type match the response.
func (c *Client) parseAPIError(response rawResponse) *APIError {
	var packet struct {
		Error struct {
			Code      string  `json:"code"`
			Type      string  `json:"type"`
			Param     *string `json:"param"`
			DocURL    string  `json:"doc_url"`
			RequestID string  `json:"request_id"`
		} `json:"error"`
	}
	if json.Unmarshal(response.Body, &packet) != nil {
		return responseFailure("unknown_error", response)
	}
	code := ErrorCode(packet.Error.Code)
	metadata, known := CatalogueDefinition(code)
	if !known || metadata.HTTPStatus == 0 || metadata.HTTPStatus != response.Status || metadata.Type != packet.Error.Type {
		return responseFailure("unknown_error", response)
	}
	failure := responseFailure(string(code), response)
	failure.Type = metadata.Type
	if definition, known := CatalogueDefinition(code); known {
		failure.Message = definition.DeveloperMessage
	}
	failure.DocURL = safeDocURL(packet.Error.DocURL)
	if packet.Error.Param != nil && len(*packet.Error.Param) <= 256 {
		failure.Param = *packet.Error.Param
	}
	if requestID := safeRequestID(packet.Error.RequestID); requestID != "" {
		failure.RequestID = requestID
	}
	return failure
}

// safeDocURL keeps an absolute https documentation link and drops anything else.
func safeDocURL(raw string) string {
	if len(raw) > 2048 {
		return ""
	}
	parsed, err := url.Parse(raw)
	if err != nil || parsed.Scheme != "https" || parsed.Host == "" || parsed.User != nil {
		return ""
	}
	return raw
}
