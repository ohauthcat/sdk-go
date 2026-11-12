package ohauthcat

import (
	"context"
	"strings"
	"sync"
	"time"
)

// MigrationConfig binds helpers to one import. ResolveLegacyProfile authenticates
// UserInfo in the app and must honor its context; tokens never reach the API.
type MigrationConfig struct {
	ImportID             ImportID
	LegacyProvider       string
	ResolveLegacyProfile func(context.Context, string) (string, error)
	// CacheTTL defaults to 30 seconds; DisableCache bypasses reuse.
	CacheTTL     time.Duration
	DisableCache bool
}

// MigrationHelper shares no state with helpers bound to other imports.
type MigrationHelper struct {
	Routing    *MigrationHelper
	client     *Client
	config     MigrationConfig
	mutex      sync.Mutex
	cached     *Import
	expires    time.Time
	generation uint64
	now        func() time.Time
}

// Migration creates an import-bound helper without issuing requests.
func (c *Client) Migration(config MigrationConfig) (*MigrationHelper, error) {
	if c == nil || c.root == nil || c.Imports == nil || c.Sessions == nil || !migrationID(string(config.ImportID), "imp") || (config.LegacyProvider != "" && config.LegacyProvider != "polis" && config.LegacyProvider != "workos" && config.LegacyProvider != "generic") {
		return nil, requestFailure("configuration_invalid")
	}
	if config.CacheTTL == 0 {
		config.CacheTTL = 30 * time.Second
	}
	if config.CacheTTL < 0 || config.CacheTTL > time.Minute {
		return nil, requestFailure("configuration_invalid")
	}
	if config.DisableCache {
		config.CacheTTL = 0
	}
	helper := &MigrationHelper{client: c, config: config, now: time.Now}
	helper.Routing = helper
	return helper, nil
}

func migrationID(value, prefix string) bool {
	return len(value) == len(prefix)+27 && strings.HasPrefix(value, prefix+"_") && validULID(value[len(prefix)+1:])
}

func validULID(value string) bool {
	if len(value) != 26 || value[0] < '0' || value[0] > '7' {
		return false
	}
	for _, char := range value {
		if !strings.ContainsRune("0123456789ABCDEFGHJKMNPQRSTVWXYZ", char) {
			return false
		}
	}
	return true
}

// Invalidate immediately discards routing cache after a switch or rollback.
func (m *MigrationHelper) Invalidate() {
	if m == nil {
		return
	}
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.generation++
	m.cached = nil
	m.expires = time.Time{}
}

func (m *MigrationHelper) load(ctx context.Context, fresh bool) (Import, error) {
	var empty Import
	if m == nil || m.client == nil || m.client.root == nil {
		return empty, requestFailure("client_closed")
	}
	if ctx == nil {
		return empty, requestFailure("request_invalid")
	}
	if ctx.Err() != nil {
		return empty, contextFailure(ctx, rawResponse{})
	}
	m.client.mutex.Lock()
	closed := m.client.closed
	m.client.mutex.Unlock()
	if closed {
		return empty, requestFailure("client_closed")
	}
	m.mutex.Lock()
	if !fresh && m.cached != nil && m.now().Before(m.expires) {
		value := *m.cached
		m.mutex.Unlock()
		return value, nil
	}
	m.generation++
	revision := m.generation
	started := m.now()
	m.mutex.Unlock()
	value, err := m.client.Imports.GetImport(ctx, GetImportArgs{Path: GetImportPath{Imp: m.config.ImportID}})
	if err != nil {
		return empty, err
	}
	if value.ID != m.config.ImportID || value.Plan == nil || value.Plan.Organizations == nil {
		return empty, requestFailure("response_invalid")
	}
	seen := map[OrganizationID]bool{}
	for _, entry := range *value.Plan.Organizations {
		if entry.OrganizationID == nil {
			continue
		}
		if !migrationID(string(*entry.OrganizationID), "org") || seen[*entry.OrganizationID] || entry.Routing == nil || (*entry.Routing != "legacy" && *entry.Routing != "ohauthcat") {
			return empty, requestFailure("response_invalid")
		}
		seen[*entry.OrganizationID] = true
	}
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if revision == m.generation {
		m.cached = &value
		m.expires = started.Add(m.config.CacheTTL)
	}
	return value, nil
}

func migrationOrganization(value Import, org OrganizationID) (ImportPlanOrganizationsItem, error) {
	var empty ImportPlanOrganizationsItem
	if value.Plan == nil || value.Plan.Organizations == nil {
		return empty, requestFailure("response_invalid")
	}
	matches := 0
	for _, entry := range *value.Plan.Organizations {
		if entry.OrganizationID != nil && *entry.OrganizationID == org {
			matches++
			empty = entry
		}
	}
	if matches != 1 || empty.Routing == nil || (*empty.Routing != "legacy" && *empty.Routing != "ohauthcat") {
		return ImportPlanOrganizationsItem{}, requestFailure("response_invalid")
	}
	return empty, nil
}

// ForOrganization reads the importer-managed flag; it never defaults on errors.
func (m *MigrationHelper) ForOrganization(ctx context.Context, org OrganizationID) (string, error) {
	if m == nil || !migrationID(string(org), "org") {
		return "", requestFailure("request_invalid")
	}
	value, err := m.load(ctx, false)
	if err != nil {
		return "", err
	}
	entry, err := migrationOrganization(value, org)
	if err != nil {
		return "", err
	}
	return *entry.Routing, nil
}

// MigrateSession resolves an authenticated legacy session through the app's
// callback, then invokes the generated single-use migration operation.
func (m *MigrationHelper) MigrateSession(parent context.Context, token, idempotencyKey string) (MigrateSessionResult, error) {
	var empty MigrateSessionResult
	if m == nil || m.client == nil || m.client.root == nil || parent == nil || m.config.LegacyProvider == "" || m.config.ResolveLegacyProfile == nil || !headerText(token) || len(token) > 65536 {
		return empty, requestFailure("request_invalid")
	}
	ctx, cancel := context.WithTimeout(parent, m.client.config.Timeout)
	defer cancel()
	stop := context.AfterFunc(m.client.root, cancel)
	defer stop()
	imported, err := m.load(ctx, true)
	if err != nil {
		return empty, err
	}
	if imported.BundleSource == nil || *imported.BundleSource != m.config.LegacyProvider || (imported.State != "running" && imported.State != "completed") {
		return empty, requestFailure("migration_invalid")
	}
	if ctx.Err() != nil {
		return empty, contextFailure(ctx, rawResponse{})
	}
	legacy, resolveErr := m.config.ResolveLegacyProfile(ctx, token)
	if ctx.Err() != nil {
		return empty, contextFailure(ctx, rawResponse{})
	}
	if resolveErr != nil {
		return empty, requestFailure("legacy_identity_unavailable")
	}
	if !headerText(legacy) || legacy != strings.TrimSpace(legacy) || len(legacy) > 255 {
		return empty, requestFailure("legacy_identity_invalid")
	}
	value, err := m.client.Sessions.MigrateSession(ctx, MigrateSessionArgs{IdempotencyKey: idempotencyKey, Body: SessionMigrateRequest{Import: m.config.ImportID, LegacyProfileID: legacy}})
	if err != nil {
		return empty, err
	}
	entry, err := migrationOrganization(imported, value.Profile.OrganizationID)
	if err != nil {
		return empty, err
	}
	if (entry.State != nil && *entry.State == "rolled_back") || value.Profile.LegacyIds == nil || (*value.Profile.LegacyIds)[m.config.LegacyProvider] != legacy {
		return empty, requestFailure("response_invalid")
	}
	return value, nil
}
