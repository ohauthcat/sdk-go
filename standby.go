package ohauthcat

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type standbyState struct {
	mutex                                sync.Mutex
	enabled                              bool
	api, auth, fallbackAPI, fallbackAuth string
	threshold, failures                  int
	cooldown                             time.Duration
	now                                  func() time.Time
	active, probing                      bool
	blockedUntil, probeAfter             time.Time
	primaryFailure                       error
	primaryResponse                      rawResponse
	revision                             uint64
}

func newStandby(config ClientConfig) (*standbyState, error) {
	s := &standbyState{api: config.APIBaseURL, auth: config.AuthBaseURL, fallbackAPI: config.StandbyAPIBaseURL, fallbackAuth: config.StandbyAuthBaseURL, threshold: config.StandbyThreshold, cooldown: config.StandbyCooldown, now: time.Now}
	if s.api == "" {
		s.api = "https://api.ohauthcat.com"
	}
	if s.auth == "" {
		s.auth = config.Issuer
	}
	explicit := s.fallbackAPI != "" && s.fallbackAuth != ""
	custom := s.api != "https://api.ohauthcat.com" || s.auth != "https://auth.ohauthcat.com" || config.Issuer != "https://auth.ohauthcat.com"
	if config.Standby != nil && *config.Standby && custom && !explicit {
		return nil, requestFailure("configuration_invalid")
	}
	s.enabled = (config.Standby == nil || *config.Standby) && (!custom || explicit)
	if s.fallbackAPI == "" {
		s.fallbackAPI = "https://api.oauthcat.com"
	}
	if s.fallbackAuth == "" {
		s.fallbackAuth = "https://auth.oauthcat.com"
	}
	for _, target := range []*string{&s.fallbackAPI, &s.fallbackAuth} {
		value, err := origin(*target, true)
		if err != nil {
			return nil, err
		}
		parsed, _ := url.Parse(value)
		if parsed.Scheme == "http" && parsed.Hostname() != "localhost" && parsed.Hostname() != "127.0.0.1" && parsed.Hostname() != "::1" {
			return nil, requestFailure("configuration_invalid")
		}
		*target = value
	}
	if s.threshold == 0 {
		s.threshold = 3
	}
	if s.cooldown == 0 {
		s.cooldown = time.Minute
	}
	if s.threshold < 1 || s.threshold > 100 || s.cooldown < 0 || s.cooldown > time.Hour || (s.enabled && (s.api == s.fallbackAPI || s.auth == s.fallbackAuth)) {
		return nil, requestFailure("configuration_invalid")
	}
	return s, nil
}
func (s *standbyState) snapshot(auth bool) (selected string, failedOver bool, revision uint64, failure error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	host := s.api
	if auth {
		host = s.auth
	}
	if s.active {
		host = s.fallbackAPI
		if auth {
			host = s.fallbackAuth
		}
	}
	return host, s.active, s.revision, s.primaryFailure
}
func (s *standbyState) success(revision uint64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if revision == s.revision && !s.active {
		s.failures = 0
	}
}
func (s *standbyState) edge(revision uint64, failure error, response rawResponse) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if revision != s.revision || !s.enabled || s.active {
		return false
	}
	s.failures++
	if s.failures < s.threshold || s.now().Before(s.blockedUntil) {
		return false
	}
	s.active = true
	s.primaryFailure = failure
	s.primaryResponse = rawResponse{Status: response.Status, RequestID: response.RequestID}
	s.probeAfter = s.now().Add(time.Minute)
	s.revision++
	return true
}
func (s *standbyState) unavailable(revision uint64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if revision != s.revision {
		return
	}
	s.active = false
	s.failures = 0
	s.blockedUntil = s.now().Add(s.cooldown)
	s.revision++
}
func (s *standbyState) probe(ctx context.Context, client *http.Client) {
	s.mutex.Lock()
	if !s.active || s.probing || s.now().Before(s.probeAfter) || ctx.Err() != nil {
		s.mutex.Unlock()
		return
	}
	s.probing = true
	revision := s.revision
	s.probeAfter = s.now().Add(time.Minute)
	s.mutex.Unlock()
	healthy := false
	defer func() {
		s.mutex.Lock()
		defer s.mutex.Unlock()
		s.probing = false
		if healthy && revision == s.revision {
			s.active = false
			s.failures = 0
			s.revision++
		}
	}()
	probe, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	for _, host := range []string{s.api, s.auth} {
		request, err := http.NewRequestWithContext(probe, http.MethodGet, host+"/healthz", http.NoBody)
		if err != nil {
			return
		}
		response, err := client.Do(request)
		if err != nil {
			if response != nil && response.Body != nil {
				_ = response.Body.Close()
			}
			return
		}
		data, readErr := io.ReadAll(io.LimitReader(response.Body, 1025))
		closeErr := response.Body.Close()
		if readErr != nil || closeErr != nil || len(data) > 1024 || response.StatusCode != 200 || probe.Err() != nil {
			return
		}
		var body map[string]string
		if json.Unmarshal(data, &body) != nil || len(body) != 1 || body["status"] != "ok" {
			return
		}
	}
	healthy = probe.Err() == nil
}
func selectHost(target, host string) string {
	u, _ := url.Parse(target)
	base, _ := url.Parse(host)
	u.Scheme = base.Scheme
	u.Host = base.Host
	return u.String()
}
