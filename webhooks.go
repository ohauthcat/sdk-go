package ohauthcat

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// VerifyOptions controls webhook freshness. Zero tolerance requires explicit disabling.
type VerifyOptions struct {
	ToleranceSeconds *int64
	Now              *int64
	DisableTolerance bool
}

// WebhookVerificationError contains a fixed safe failure code.
type WebhookVerificationError struct{ Code string }

func (e *WebhookVerificationError) Error() string {
	if e == nil {
		return "Webhook verification failed"
	}
	return "Webhook verification failed: " + e.Code
}
func webhookFailure(code string) error { return &WebhookVerificationError{Code: code} }

// SignatureDigest signs the exact raw bytes with the exact UTF-8 secret.
func SignatureDigest(body []byte, timestamp int64, secret string) (string, error) {
	if timestamp <= 0 || timestamp > 9007199254740991 {
		return "", webhookFailure("header_invalid")
	}
	if !utf8.ValidString(secret) {
		return "", webhookFailure("secret_missing")
	}
	mac := hmac.New(sha256.New, []byte(secret))
	_, _ = mac.Write([]byte(strconv.FormatInt(timestamp, 10) + "."))
	_, _ = mac.Write(body)
	return hex.EncodeToString(mac.Sum(nil)), nil
}

// ParseSignatureHeader accepts multiple v1 values while refusing ambiguous timestamps.
func ParseSignatureHeader(header string) (timestamp int64, signatures []string, err error) {
	if strings.TrimSpace(header) == "" {
		return 0, nil, webhookFailure("header_missing")
	}
	if !utf8.ValidString(header) || len(header) > 16384 {
		return 0, nil, webhookFailure("header_invalid")
	}
	seen := false
	signatures = []string{}
	for _, part := range strings.Split(header, ",") {
		name, value, found := strings.Cut(strings.TrimSpace(part), "=")
		if !found {
			return 0, nil, webhookFailure("header_invalid")
		}
		if name == "t" {
			if seen || len(value) > 20 || value == "" {
				return 0, nil, webhookFailure("header_invalid")
			}
			seen = true
			digits := value
			if digits == "" {
				return 0, nil, webhookFailure("header_invalid")
			}
			for _, digit := range digits {
				if digit < '0' || digit > '9' {
					return 0, nil, webhookFailure("header_invalid")
				}
			}
			parsed, err := strconv.ParseInt(value, 10, 64)
			if err != nil || parsed <= 0 || parsed > 9007199254740991 {
				return 0, nil, webhookFailure("header_invalid")
			}
			timestamp = parsed
		} else if name == "v1" && len(value) == 64 {
			signatures = append(signatures, strings.ToLower(value))
		}
	}
	if !seen || len(signatures) == 0 {
		return 0, nil, webhookFailure("header_invalid")
	}
	return timestamp, signatures, nil
}

// VerifySignature authenticates bytes without interpreting their event payload.
func VerifySignature(body []byte, header string, secrets []string, options VerifyOptions) error {
	tolerance := int64(300)
	if options.ToleranceSeconds != nil {
		tolerance = *options.ToleranceSeconds
	}
	now := time.Now().Unix()
	if options.Now != nil {
		now = *options.Now
	}
	if tolerance < 0 || tolerance > 9007199254740991 || (tolerance == 0 && !options.DisableTolerance) || now < 0 || now > 9007199254740991 {
		return webhookFailure("options_invalid")
	}
	timestamp, signatures, err := ParseSignatureHeader(header)
	if err != nil {
		return err
	}
	difference := now - timestamp
	if difference < 0 {
		difference = -difference
	}
	if !options.DisableTolerance && difference > tolerance {
		return webhookFailure("timestamp_outside")
	}
	matched := 0
	nonempty := false
	for _, secret := range secrets {
		if !utf8.ValidString(secret) {
			return webhookFailure("secret_missing")
		}
		if secret == "" {
			continue
		}
		nonempty = true
		expected, err := SignatureDigest(body, timestamp, secret)
		if err != nil {
			return err
		}
		for _, signature := range signatures {
			matched |= subtle.ConstantTimeCompare([]byte(expected), []byte(signature))
		}
	}
	if !nonempty {
		return webhookFailure("secret_missing")
	}
	if matched == 0 {
		return webhookFailure("signature_invalid")
	}
	return nil
}

// VerifyWebhook verifies the signature and decodes the strict pinned Event envelope.
func VerifyWebhook(body []byte, header string, secrets []string, options VerifyOptions) (Event, error) {
	var event Event
	if err := VerifySignature(body, header, secrets, options); err != nil {
		return event, err
	}
	if !utf8.Valid(body) {
		return event, webhookFailure("event_invalid")
	}
	var contract wireContract
	if json.Unmarshal([]byte(pinnedContract), &contract) != nil {
		return event, webhookFailure("event_invalid")
	}
	definition, ok := contract.Schemas["Event"]
	normalized, err := normalizeProfileFields(body, schema{Ref: "#/components/schemas/Event"}, contract.Schemas)
	if !ok || err != nil || !matchesJSON(normalized, definition, contract.Schemas, false) || json.Unmarshal(normalized, &event) != nil {
		return Event{}, webhookFailure("event_invalid")
	}
	return event, nil
}

// VerifyWebhook authenticates and decodes an event using the client's public helper.
func (c *Client) VerifyWebhook(body []byte, header string, secrets []string, options VerifyOptions) (Event, error) {
	return VerifyWebhook(body, header, secrets, options)
}
