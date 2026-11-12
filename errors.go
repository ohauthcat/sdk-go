package ohauthcat

import (
	"context"
	"fmt"
)

// ErrorCode identifies an API catalogue code or a local SDK failure.
type ErrorCode string

// ErrorDefinition describes a pinned catalogue entry.
type ErrorDefinition struct {
	Type             string
	HTTPStatus       int
	Retryable        bool
	DeveloperMessage string
}

// APIError retains safe request context without rendering response payloads.
type APIError struct {
	Code      ErrorCode
	Status    int
	Type      string
	Param     string
	DocURL    string
	RequestID string
	Message   string
	cause     error
}

// Unwrap exposes only safe context cancellation causes, never network payloads.
func (e *APIError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.cause
}

func (e *APIError) Error() string {
	if e == nil {
		return "OhAuthCat request failed"
	}
	return fmt.Sprintf("OhAuthCat request failed (%s, HTTP %d)", e.Code, e.Status)
}

type rawResponse struct {
	Body      []byte
	Status    int
	RequestID string
}

func requestFailure(code string) *APIError {
	failure := &APIError{Code: ErrorCode(code), Type: "unavailable", Message: "Request could not be completed"}
	if code == "request_timeout" {
		failure.cause = context.DeadlineExceeded
	}
	return failure
}

func responseFailure(code string, response rawResponse) *APIError {
	failure := requestFailure(code)
	failure.Status = response.Status
	failure.RequestID = response.RequestID
	return failure
}

// PlanLimitMetric reports the enforced monthly ceiling a plan_limit_reached
// error names in Param: monthly_active_users, audit_log_events or
// pipes_token_fetches. It reports false for any other error, and for a
// plan_limit_reached refusal about a resource count or a missing feature.
func (e *APIError) PlanLimitMetric() (string, bool) {
	if e == nil || e.Code != CodePlanLimitReached {
		return "", false
	}
	switch e.Param {
	case "monthly_active_users", "audit_log_events", "pipes_token_fetches":
		return e.Param, true
	}
	return "", false
}
