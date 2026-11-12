// Code generated from the pinned OpenAPI contract. DO NOT EDIT.
package ohauthcat

import (
	"context"
	"encoding/json"
	"iter"
)

type ActivateConnectionArgs struct {
	IdempotencyKey string                 `json:"idempotency_key"`
	Path           ActivateConnectionPath `json:"path"`
}

type ActivateConnectionClientKeyArgs struct {
	IdempotencyKey string                          `json:"idempotency_key"`
	Path           ActivateConnectionClientKeyPath `json:"path"`
}

type ActivateEncryptionKeyArgs struct {
	IdempotencyKey string                    `json:"idempotency_key"`
	Path           ActivateEncryptionKeyPath `json:"path"`
}

type AuthenticateEndUserPasskeyArgs struct {
	IdempotencyKey string                         `json:"idempotency_key"`
	Body           AuthenticateEndUserPasskeyBody `json:"body"`
}

type CancelImportArgs struct {
	IdempotencyKey string           `json:"idempotency_key"`
	Path           CancelImportPath `json:"path"`
}

type CheckAuthorizationArgs struct {
	IdempotencyKey string                 `json:"idempotency_key"`
	Body           CheckAuthorizationBody `json:"body"`
}

type CheckAuthorizationBulkArgs struct {
	IdempotencyKey string                     `json:"idempotency_key"`
	Body           CheckAuthorizationBulkBody `json:"body"`
}

type ChooseDirectorySourceOAuthCompanyArgs struct {
	IdempotencyKey string                                `json:"idempotency_key"`
	Path           ChooseDirectorySourceOAuthCompanyPath `json:"path"`
	Body           ChooseDirectorySourceOAuthCompanyBody `json:"body"`
}

type ClearEndUserPasswordArgs struct {
	IdempotencyKey string                   `json:"idempotency_key"`
	Path           ClearEndUserPasswordPath `json:"path"`
}

type CreateAuditLogEventArgs struct {
	IdempotencyKey string                  `json:"idempotency_key"`
	Body           CreateAuditLogEventBody `json:"body"`
}

type CreateAuditLogExportArgs struct {
	IdempotencyKey string                   `json:"idempotency_key"`
	Body           CreateAuditLogExportBody `json:"body"`
}

type CreateAuditLogSchemaArgs struct {
	IdempotencyKey string                   `json:"idempotency_key"`
	Body           CreateAuditLogSchemaBody `json:"body"`
}

type CreateConnectionArgs struct {
	IdempotencyKey string               `json:"idempotency_key"`
	Body           CreateConnectionBody `json:"body"`
}

type CreateConnectionCertificateArgs struct {
	IdempotencyKey string                          `json:"idempotency_key"`
	Path           CreateConnectionCertificatePath `json:"path"`
	Body           CreateConnectionCertificateBody `json:"body"`
}

type CreateConnectionClientKeyArgs struct {
	IdempotencyKey string                        `json:"idempotency_key"`
	Path           CreateConnectionClientKeyPath `json:"path"`
}

type CreateConnectionShareArgs struct {
	IdempotencyKey string                    `json:"idempotency_key"`
	Path           CreateConnectionSharePath `json:"path"`
	Body           CreateConnectionShareBody `json:"body"`
}

type CreateCustomDomainArgs struct {
	IdempotencyKey string                 `json:"idempotency_key"`
	Body           CreateCustomDomainBody `json:"body"`
}

type CreateDirectoryArgs struct {
	IdempotencyKey string              `json:"idempotency_key"`
	Body           CreateDirectoryBody `json:"body"`
}

type CreateDirectorySourceSshKeyArgs struct {
	IdempotencyKey string                          `json:"idempotency_key"`
	Path           CreateDirectorySourceSshKeyPath `json:"path"`
}

type CreateDomainArgs struct {
	IdempotencyKey string           `json:"idempotency_key"`
	Path           CreateDomainPath `json:"path"`
	Body           CreateDomainBody `json:"body"`
}

type CreateEncryptionKeyArgs struct {
	IdempotencyKey string                  `json:"idempotency_key"`
	Body           CreateEncryptionKeyBody `json:"body"`
}

type CreateEndUserArgs struct {
	IdempotencyKey string            `json:"idempotency_key"`
	Body           CreateEndUserBody `json:"body"`
}

type CreateEndUserInviteArgs struct {
	IdempotencyKey string                  `json:"idempotency_key"`
	Body           CreateEndUserInviteBody `json:"body"`
}

type CreateExportArgs struct {
	IdempotencyKey string            `json:"idempotency_key"`
	Body           *CreateExportBody `json:"body"`
}

type CreateFeatureFlagArgs struct {
	IdempotencyKey string                `json:"idempotency_key"`
	Body           CreateFeatureFlagBody `json:"body"`
}

type CreateImportArgs struct {
	IdempotencyKey string           `json:"idempotency_key"`
	Body           CreateImportBody `json:"body"`
}

type CreateItContactArgs struct {
	IdempotencyKey string              `json:"idempotency_key"`
	Path           CreateItContactPath `json:"path"`
	Body           CreateItContactBody `json:"body"`
}

type CreateLogStreamArgs struct {
	IdempotencyKey string              `json:"idempotency_key"`
	Body           CreateLogStreamBody `json:"body"`
}

type CreateM2MClientArgs struct {
	IdempotencyKey string              `json:"idempotency_key"`
	Body           CreateM2MClientBody `json:"body"`
}

type CreateMembershipArgs struct {
	IdempotencyKey string               `json:"idempotency_key"`
	Path           CreateMembershipPath `json:"path"`
	Body           CreateMembershipBody `json:"body"`
}

type CreateOAuthClientArgs struct {
	IdempotencyKey string                `json:"idempotency_key"`
	Body           CreateOAuthClientBody `json:"body"`
}

type CreateOAuthResourceArgs struct {
	IdempotencyKey string                  `json:"idempotency_key"`
	Body           CreateOAuthResourceBody `json:"body"`
}

type CreateOrganizationArgs struct {
	IdempotencyKey string                 `json:"idempotency_key"`
	Body           CreateOrganizationBody `json:"body"`
}

type CreateOrganizationAPIKeyArgs struct {
	IdempotencyKey string                       `json:"idempotency_key"`
	Path           CreateOrganizationAPIKeyPath `json:"path"`
	Body           CreateOrganizationAPIKeyBody `json:"body"`
}

type CreateOrganizationAPIKeySecretArgs struct {
	IdempotencyKey string                             `json:"idempotency_key"`
	Path           CreateOrganizationAPIKeySecretPath `json:"path"`
}

type CreateOrganizationExportArgs struct {
	IdempotencyKey string                       `json:"idempotency_key"`
	Body           CreateOrganizationExportBody `json:"body"`
}

type CreatePermissionArgs struct {
	IdempotencyKey string               `json:"idempotency_key"`
	Body           CreatePermissionBody `json:"body"`
}

type CreatePipeConnectionArgs struct {
	IdempotencyKey string                   `json:"idempotency_key"`
	Body           CreatePipeConnectionBody `json:"body"`
}

type CreatePolicyArgs struct {
	IdempotencyKey string           `json:"idempotency_key"`
	Body           CreatePolicyBody `json:"body"`
}

type CreatePortalInviteArgs struct {
	IdempotencyKey string                 `json:"idempotency_key"`
	Body           CreatePortalInviteBody `json:"body"`
}

type CreatePortalLinkArgs struct {
	IdempotencyKey string               `json:"idempotency_key"`
	Body           CreatePortalLinkBody `json:"body"`
}

type CreatePortalProfileArgs struct {
	IdempotencyKey string                  `json:"idempotency_key"`
	Body           CreatePortalProfileBody `json:"body"`
}

type CreateRedirectURIArgs struct {
	IdempotencyKey string                `json:"idempotency_key"`
	Body           CreateRedirectURIBody `json:"body"`
}

type CreateRoleArgs struct {
	IdempotencyKey string         `json:"idempotency_key"`
	Body           CreateRoleBody `json:"body"`
}

type CreateRoleAssignmentArgs struct {
	IdempotencyKey string                   `json:"idempotency_key"`
	Path           CreateRoleAssignmentPath `json:"path"`
	Body           CreateRoleAssignmentBody `json:"body"`
}

type CreateRoleRuleArgs struct {
	IdempotencyKey string             `json:"idempotency_key"`
	Body           CreateRoleRuleBody `json:"body"`
}

type CreateSignInHookArgs struct {
	IdempotencyKey string               `json:"idempotency_key"`
	Body           CreateSignInHookBody `json:"body"`
}

type CreateSsfStreamArgs struct {
	IdempotencyKey string              `json:"idempotency_key"`
	Body           CreateSsfStreamBody `json:"body"`
}

type CreateTrustedIssuerArgs struct {
	IdempotencyKey string                  `json:"idempotency_key"`
	Body           CreateTrustedIssuerBody `json:"body"`
}

type CreateUploadArgs struct {
	IdempotencyKey string           `json:"idempotency_key"`
	Body           CreateUploadBody `json:"body"`
}

type CreateVaultKeyArgs struct {
	IdempotencyKey string             `json:"idempotency_key"`
	Body           CreateVaultKeyBody `json:"body"`
}

type CreateWebhookEndpointArgs struct {
	IdempotencyKey string                    `json:"idempotency_key"`
	Body           CreateWebhookEndpointBody `json:"body"`
}

type DeactivateConnectionArgs struct {
	IdempotencyKey string                   `json:"idempotency_key"`
	Path           DeactivateConnectionPath `json:"path"`
}

type DeleteConnectionArgs struct {
	IdempotencyKey string                `json:"idempotency_key"`
	Path           DeleteConnectionPath  `json:"path"`
	Query          DeleteConnectionQuery `json:"query"`
}

type DeleteConnectionCertificateArgs struct {
	IdempotencyKey string                          `json:"idempotency_key"`
	Path           DeleteConnectionCertificatePath `json:"path"`
}

type DeleteConnectionSCIMBindingArgs struct {
	IdempotencyKey string                          `json:"idempotency_key"`
	Path           DeleteConnectionSCIMBindingPath `json:"path"`
}

type DeleteCustomDomainArgs struct {
	IdempotencyKey string                 `json:"idempotency_key"`
	Path           DeleteCustomDomainPath `json:"path"`
}

type DeleteDirectoryArgs struct {
	IdempotencyKey string                `json:"idempotency_key"`
	Path           DeleteDirectoryPath   `json:"path"`
	Query          *DeleteDirectoryQuery `json:"query"`
}

type DeleteDomainArgs struct {
	IdempotencyKey string           `json:"idempotency_key"`
	Path           DeleteDomainPath `json:"path"`
}

type DeleteEndUserArgs struct {
	IdempotencyKey string            `json:"idempotency_key"`
	Path           DeleteEndUserPath `json:"path"`
}

type DeleteEndUserFactorArgs struct {
	IdempotencyKey string                  `json:"idempotency_key"`
	Path           DeleteEndUserFactorPath `json:"path"`
}

type DeleteFeatureFlagArgs struct {
	IdempotencyKey string                `json:"idempotency_key"`
	Path           DeleteFeatureFlagPath `json:"path"`
}

type DeleteItEmailConfigurationArgs struct {
	IdempotencyKey string `json:"idempotency_key"`
}

type DeleteItContactArgs struct {
	IdempotencyKey string              `json:"idempotency_key"`
	Path           DeleteItContactPath `json:"path"`
}

type DeleteLogStreamArgs struct {
	IdempotencyKey string              `json:"idempotency_key"`
	Path           DeleteLogStreamPath `json:"path"`
}

type DeleteMembershipArgs struct {
	IdempotencyKey string               `json:"idempotency_key"`
	Path           DeleteMembershipPath `json:"path"`
}

type DeleteOAuthClientArgs struct {
	IdempotencyKey string                `json:"idempotency_key"`
	Path           DeleteOAuthClientPath `json:"path"`
}

type DeleteOAuthResourceArgs struct {
	IdempotencyKey string                  `json:"idempotency_key"`
	Path           DeleteOAuthResourcePath `json:"path"`
}

type DeleteOAuthScopeArgs struct {
	IdempotencyKey string               `json:"idempotency_key"`
	Path           DeleteOAuthScopePath `json:"path"`
}

type DeleteOrganizationArgs struct {
	IdempotencyKey string                 `json:"idempotency_key"`
	Path           DeleteOrganizationPath `json:"path"`
}

type DeleteOrganizationUserManagementSettingsArgs struct {
	IdempotencyKey string                                       `json:"idempotency_key"`
	Path           DeleteOrganizationUserManagementSettingsPath `json:"path"`
}

type DeletePermissionArgs struct {
	IdempotencyKey string               `json:"idempotency_key"`
	Path           DeletePermissionPath `json:"path"`
}

type DeletePipeProviderArgs struct {
	IdempotencyKey string                 `json:"idempotency_key"`
	Path           DeletePipeProviderPath `json:"path"`
}

type DeletePolicyArgs struct {
	IdempotencyKey string           `json:"idempotency_key"`
	Path           DeletePolicyPath `json:"path"`
}

type DeletePortalProfileArgs struct {
	IdempotencyKey string                  `json:"idempotency_key"`
	Path           DeletePortalProfilePath `json:"path"`
}

type DeleteRadarAlertSubscriptionArgs struct {
	IdempotencyKey string `json:"idempotency_key"`
}

type DeleteRedirectURIArgs struct {
	IdempotencyKey string                `json:"idempotency_key"`
	Path           DeleteRedirectURIPath `json:"path"`
}

type DeleteRoleArgs struct {
	IdempotencyKey string         `json:"idempotency_key"`
	Path           DeleteRolePath `json:"path"`
}

type DeleteRoleAssignmentArgs struct {
	IdempotencyKey string                   `json:"idempotency_key"`
	Path           DeleteRoleAssignmentPath `json:"path"`
}

type DeleteRoleRuleArgs struct {
	IdempotencyKey string             `json:"idempotency_key"`
	Path           DeleteRoleRulePath `json:"path"`
}

type DeleteSignInHookArgs struct {
	IdempotencyKey string               `json:"idempotency_key"`
	Path           DeleteSignInHookPath `json:"path"`
}

type DeleteSignInRiskPolicyArgs struct {
	IdempotencyKey string `json:"idempotency_key"`
}

type DeleteSocialProviderArgs struct {
	IdempotencyKey string                   `json:"idempotency_key"`
	Path           DeleteSocialProviderPath `json:"path"`
}

type DeleteSsfStreamArgs struct {
	IdempotencyKey string              `json:"idempotency_key"`
	Path           DeleteSsfStreamPath `json:"path"`
}

type DeleteTrustedIssuerArgs struct {
	IdempotencyKey string                  `json:"idempotency_key"`
	Path           DeleteTrustedIssuerPath `json:"path"`
}

type DeleteVaultObjectArgs struct {
	IdempotencyKey string                 `json:"idempotency_key"`
	Path           DeleteVaultObjectPath  `json:"path"`
	Query          DeleteVaultObjectQuery `json:"query"`
}

type DeleteWebhookEndpointArgs struct {
	IdempotencyKey string                    `json:"idempotency_key"`
	Path           DeleteWebhookEndpointPath `json:"path"`
}

type DeprecateEncryptionKeyArgs struct {
	IdempotencyKey string                     `json:"idempotency_key"`
	Path           DeprecateEncryptionKeyPath `json:"path"`
}

type DevInitArgs struct {
	IdempotencyKey string       `json:"idempotency_key"`
	Body           *DevInitBody `json:"body"`
}

type DevTriggerArgs struct {
	IdempotencyKey string         `json:"idempotency_key"`
	Body           DevTriggerBody `json:"body"`
}

type DisconnectPipeConnectionArgs struct {
	IdempotencyKey string                       `json:"idempotency_key"`
	Path           DisconnectPipeConnectionPath `json:"path"`
}

type EvaluateFeatureFlagArgs struct {
	IdempotencyKey string                  `json:"idempotency_key"`
	Body           EvaluateFeatureFlagBody `json:"body"`
}

type EvaluateFeatureFlagsArgs struct {
	IdempotencyKey string                   `json:"idempotency_key"`
	Body           EvaluateFeatureFlagsBody `json:"body"`
}

type ExchangeSSOCodeArgs struct {
	IdempotencyKey string              `json:"idempotency_key"`
	Body           ExchangeSSOCodeBody `json:"body"`
}

type ExchangeTrustedTokenArgs struct {
	IdempotencyKey string                   `json:"idempotency_key"`
	Body           ExchangeTrustedTokenBody `json:"body"`
}

type ExportAuditLogEventsArgs struct {
	IdempotencyKey string                     `json:"idempotency_key"`
	Query          *ExportAuditLogEventsQuery `json:"query"`
}

type ForwardSSOResponseArgs struct {
	IdempotencyKey string                 `json:"idempotency_key"`
	Body           ForwardSSOResponseBody `json:"body"`
}

type GetAttributeCatalogueArgs struct {
	IdempotencyKey string `json:"idempotency_key"`
}

type GetAuditLogEventArgs struct {
	IdempotencyKey string               `json:"idempotency_key"`
	Path           GetAuditLogEventPath `json:"path"`
}

type GetAuditLogExportArgs struct {
	IdempotencyKey string                `json:"idempotency_key"`
	Path           GetAuditLogExportPath `json:"path"`
}

type GetAuditLogSchemaArgs struct {
	IdempotencyKey string                  `json:"idempotency_key"`
	Path           GetAuditLogSchemaPath   `json:"path"`
	Query          *GetAuditLogSchemaQuery `json:"query"`
}

type GetBotProtectionArgs struct {
	IdempotencyKey string `json:"idempotency_key"`
}

type GetBrandingArgs struct {
	IdempotencyKey string `json:"idempotency_key"`
}

type GetConnectionArgs struct {
	IdempotencyKey string              `json:"idempotency_key"`
	Path           GetConnectionPath   `json:"path"`
	Query          *GetConnectionQuery `json:"query"`
}

type GetConnectionAttributeMapArgs struct {
	IdempotencyKey string                        `json:"idempotency_key"`
	Path           GetConnectionAttributeMapPath `json:"path"`
}

type GetConnectionAttributePreviewArgs struct {
	IdempotencyKey string                            `json:"idempotency_key"`
	Path           GetConnectionAttributePreviewPath `json:"path"`
}

type GetConnectionHealthArgs struct {
	IdempotencyKey string                  `json:"idempotency_key"`
	Path           GetConnectionHealthPath `json:"path"`
}

type GetConnectionSingleLogoutArgs struct {
	IdempotencyKey string                        `json:"idempotency_key"`
	Path           GetConnectionSingleLogoutPath `json:"path"`
}

type GetConnectionSPMetadataArgs struct {
	IdempotencyKey string                      `json:"idempotency_key"`
	Path           GetConnectionSPMetadataPath `json:"path"`
}

type GetConnectionTestResultArgs struct {
	IdempotencyKey string                      `json:"idempotency_key"`
	Path           GetConnectionTestResultPath `json:"path"`
}

type GetCustomDomainArgs struct {
	IdempotencyKey string              `json:"idempotency_key"`
	Path           GetCustomDomainPath `json:"path"`
}

type GetDirectoryArgs struct {
	IdempotencyKey string           `json:"idempotency_key"`
	Path           GetDirectoryPath `json:"path"`
}

type GetDirectoryAttributeMapArgs struct {
	IdempotencyKey string                       `json:"idempotency_key"`
	Path           GetDirectoryAttributeMapPath `json:"path"`
}

type GetDirectoryGroupArgs struct {
	IdempotencyKey string                `json:"idempotency_key"`
	Path           GetDirectoryGroupPath `json:"path"`
}

type GetDirectorySourceArgs struct {
	IdempotencyKey string                 `json:"idempotency_key"`
	Path           GetDirectorySourcePath `json:"path"`
}

type GetDirectorySourceOAuthArgs struct {
	IdempotencyKey string                      `json:"idempotency_key"`
	Path           GetDirectorySourceOAuthPath `json:"path"`
}

type GetDirectorySourceSshKeyArgs struct {
	IdempotencyKey string                       `json:"idempotency_key"`
	Path           GetDirectorySourceSshKeyPath `json:"path"`
}

type GetDirectoryUserArgs struct {
	IdempotencyKey string               `json:"idempotency_key"`
	Path           GetDirectoryUserPath `json:"path"`
}

type GetEncryptionKeyArgs struct {
	IdempotencyKey string               `json:"idempotency_key"`
	Path           GetEncryptionKeyPath `json:"path"`
}

type GetEncryptionKeySetupArgs struct {
	IdempotencyKey string                      `json:"idempotency_key"`
	Query          *GetEncryptionKeySetupQuery `json:"query"`
}

type GetEndUserArgs struct {
	IdempotencyKey string         `json:"idempotency_key"`
	Path           GetEndUserPath `json:"path"`
}

type GetEnvironmentArgs struct {
	IdempotencyKey string `json:"idempotency_key"`
}

type GetEnvironmentSigninHostArgs struct {
	IdempotencyKey string                         `json:"idempotency_key"`
	Query          *GetEnvironmentSigninHostQuery `json:"query"`
}

type GetEventArgs struct {
	IdempotencyKey string       `json:"idempotency_key"`
	Path           GetEventPath `json:"path"`
}

type GetExportArgs struct {
	IdempotencyKey string        `json:"idempotency_key"`
	Path           GetExportPath `json:"path"`
}

type GetFeatureFlagArgs struct {
	IdempotencyKey string             `json:"idempotency_key"`
	Path           GetFeatureFlagPath `json:"path"`
}

type GetItEmailConfigurationArgs struct {
	IdempotencyKey string `json:"idempotency_key"`
}

type GetImportArgs struct {
	IdempotencyKey string        `json:"idempotency_key"`
	Path           GetImportPath `json:"path"`
}

type GetItContactArgs struct {
	IdempotencyKey string           `json:"idempotency_key"`
	Path           GetItContactPath `json:"path"`
}

type GetM2MClientArgs struct {
	IdempotencyKey string           `json:"idempotency_key"`
	Path           GetM2MClientPath `json:"path"`
}

type GetOAuthClientArgs struct {
	IdempotencyKey string             `json:"idempotency_key"`
	Path           GetOAuthClientPath `json:"path"`
}

type GetOAuthProtectedResourceMetadataArgs struct {
	IdempotencyKey string                                `json:"idempotency_key"`
	Path           GetOAuthProtectedResourceMetadataPath `json:"path"`
}

type GetOAuthResourceArgs struct {
	IdempotencyKey string               `json:"idempotency_key"`
	Path           GetOAuthResourcePath `json:"path"`
}

type GetOrganizationArgs struct {
	IdempotencyKey string                `json:"idempotency_key"`
	Path           GetOrganizationPath   `json:"path"`
	Query          *GetOrganizationQuery `json:"query"`
}

type GetOrganizationAllowedIntentsArgs struct {
	IdempotencyKey string                            `json:"idempotency_key"`
	Path           GetOrganizationAllowedIntentsPath `json:"path"`
}

type GetOrganizationAPIKeyArgs struct {
	IdempotencyKey string                    `json:"idempotency_key"`
	Path           GetOrganizationAPIKeyPath `json:"path"`
}

type GetOrganizationExportArgs struct {
	IdempotencyKey string                    `json:"idempotency_key"`
	Path           GetOrganizationExportPath `json:"path"`
}

type GetOrganizationOAuthSettingsArgs struct {
	IdempotencyKey string                           `json:"idempotency_key"`
	Path           GetOrganizationOAuthSettingsPath `json:"path"`
}

type GetOrganizationProfileArgs struct {
	IdempotencyKey string                     `json:"idempotency_key"`
	Path           GetOrganizationProfilePath `json:"path"`
}

type GetOrganizationStatusArgs struct {
	IdempotencyKey string                    `json:"idempotency_key"`
	Path           GetOrganizationStatusPath `json:"path"`
}

type GetOrganizationUserManagementSettingsArgs struct {
	IdempotencyKey string                                    `json:"idempotency_key"`
	Path           GetOrganizationUserManagementSettingsPath `json:"path"`
}

type GetPermissionArgs struct {
	IdempotencyKey string            `json:"idempotency_key"`
	Path           GetPermissionPath `json:"path"`
}

type GetPipeAccessTokenArgs struct {
	IdempotencyKey string                 `json:"idempotency_key"`
	Path           GetPipeAccessTokenPath `json:"path"`
}

type GetPipeConnectionArgs struct {
	IdempotencyKey string                `json:"idempotency_key"`
	Path           GetPipeConnectionPath `json:"path"`
}

type GetPolicyArgs struct {
	IdempotencyKey string        `json:"idempotency_key"`
	Path           GetPolicyPath `json:"path"`
}

type GetPortalProfileArgs struct {
	IdempotencyKey string               `json:"idempotency_key"`
	Path           GetPortalProfilePath `json:"path"`
}

type GetProviderCatalogueArgs struct {
	IdempotencyKey string `json:"idempotency_key"`
}

type GetRadarAlertSubscriptionArgs struct {
	IdempotencyKey string `json:"idempotency_key"`
}

type GetRoleArgs struct {
	IdempotencyKey string      `json:"idempotency_key"`
	Path           GetRolePath `json:"path"`
}

type GetRoleRuleArgs struct {
	IdempotencyKey string          `json:"idempotency_key"`
	Path           GetRoleRulePath `json:"path"`
}

type GetSignInExplanationArgs struct {
	IdempotencyKey string                   `json:"idempotency_key"`
	Path           GetSignInExplanationPath `json:"path"`
}

type GetSignInHookArgs struct {
	IdempotencyKey string            `json:"idempotency_key"`
	Path           GetSignInHookPath `json:"path"`
}

type GetSignInRiskPolicyArgs struct {
	IdempotencyKey string `json:"idempotency_key"`
}

type GetSignInTimelineArgs struct {
	IdempotencyKey string                `json:"idempotency_key"`
	Path           GetSignInTimelinePath `json:"path"`
}

type GetSsfStreamArgs struct {
	IdempotencyKey string           `json:"idempotency_key"`
	Path           GetSsfStreamPath `json:"path"`
}

type GetSSOProfileArgs struct {
	IdempotencyKey string `json:"idempotency_key"`
	AccessToken    string `json:"access_token"`
}

type GetUserManagementSettingsArgs struct {
	IdempotencyKey string `json:"idempotency_key"`
}

type GetUserManagementUsageArgs struct {
	IdempotencyKey string                       `json:"idempotency_key"`
	Query          *GetUserManagementUsageQuery `json:"query"`
}

type GetVaultObjectArgs struct {
	IdempotencyKey string              `json:"idempotency_key"`
	Path           GetVaultObjectPath  `json:"path"`
	Query          GetVaultObjectQuery `json:"query"`
}

type GetWebhookEndpointArgs struct {
	IdempotencyKey string                 `json:"idempotency_key"`
	Path           GetWebhookEndpointPath `json:"path"`
}

type ImportDirectoryCsvArgs struct {
	IdempotencyKey string                   `json:"idempotency_key"`
	Path           ImportDirectoryCsvPath   `json:"path"`
	Query          *ImportDirectoryCsvQuery `json:"query"`
}

type ListAuditLogEventsArgs struct {
	IdempotencyKey string                   `json:"idempotency_key"`
	Query          *ListAuditLogEventsQuery `json:"query"`
}

type ListAuditLogSchemasArgs struct {
	IdempotencyKey string                    `json:"idempotency_key"`
	Query          *ListAuditLogSchemasQuery `json:"query"`
}

type ListAuditLogsArgs struct {
	IdempotencyKey string              `json:"idempotency_key"`
	Query          *ListAuditLogsQuery `json:"query"`
}

type ListBotDetectionsArgs struct {
	IdempotencyKey string                  `json:"idempotency_key"`
	Query          *ListBotDetectionsQuery `json:"query"`
}

type ListConnectionCertificatesArgs struct {
	IdempotencyKey string                           `json:"idempotency_key"`
	Path           ListConnectionCertificatesPath   `json:"path"`
	Query          *ListConnectionCertificatesQuery `json:"query"`
}

type ListConnectionClientKeysArgs struct {
	IdempotencyKey string                       `json:"idempotency_key"`
	Path           ListConnectionClientKeysPath `json:"path"`
}

type ListConnectionHealthArgs struct {
	IdempotencyKey string                     `json:"idempotency_key"`
	Query          *ListConnectionHealthQuery `json:"query"`
}

type ListConnectionSharesArgs struct {
	IdempotencyKey string                   `json:"idempotency_key"`
	Path           ListConnectionSharesPath `json:"path"`
}

type ListConnectionSignInsArgs struct {
	IdempotencyKey string                      `json:"idempotency_key"`
	Path           ListConnectionSignInsPath   `json:"path"`
	Query          *ListConnectionSignInsQuery `json:"query"`
}

type ListConnectionsArgs struct {
	IdempotencyKey string                `json:"idempotency_key"`
	Query          *ListConnectionsQuery `json:"query"`
}

type ListCustomDomainsArgs struct {
	IdempotencyKey string                  `json:"idempotency_key"`
	Query          *ListCustomDomainsQuery `json:"query"`
}

type ListDirectoriesArgs struct {
	IdempotencyKey string                `json:"idempotency_key"`
	Query          *ListDirectoriesQuery `json:"query"`
}

type ListDirectoryGroupMembersArgs struct {
	IdempotencyKey string                          `json:"idempotency_key"`
	Path           ListDirectoryGroupMembersPath   `json:"path"`
	Query          *ListDirectoryGroupMembersQuery `json:"query"`
}

type ListDirectoryGroupsArgs struct {
	IdempotencyKey string                    `json:"idempotency_key"`
	Path           ListDirectoryGroupsPath   `json:"path"`
	Query          *ListDirectoryGroupsQuery `json:"query"`
}

type ListDirectoryUserGroupsArgs struct {
	IdempotencyKey string                        `json:"idempotency_key"`
	Path           ListDirectoryUserGroupsPath   `json:"path"`
	Query          *ListDirectoryUserGroupsQuery `json:"query"`
}

type ListDirectoryUsersArgs struct {
	IdempotencyKey string                   `json:"idempotency_key"`
	Path           ListDirectoryUsersPath   `json:"path"`
	Query          *ListDirectoryUsersQuery `json:"query"`
}

type ListDomainsArgs struct {
	IdempotencyKey string            `json:"idempotency_key"`
	Path           ListDomainsPath   `json:"path"`
	Query          *ListDomainsQuery `json:"query"`
}

type ListEmailEventsArgs struct {
	IdempotencyKey string                `json:"idempotency_key"`
	Query          *ListEmailEventsQuery `json:"query"`
}

type ListEncryptionKeysArgs struct {
	IdempotencyKey string `json:"idempotency_key"`
}

type ListEndUserFactorsArgs struct {
	IdempotencyKey string                 `json:"idempotency_key"`
	Path           ListEndUserFactorsPath `json:"path"`
}

type ListEndUserIdentitiesArgs struct {
	IdempotencyKey string                    `json:"idempotency_key"`
	Path           ListEndUserIdentitiesPath `json:"path"`
}

type ListEndUserInvitesArgs struct {
	IdempotencyKey string                   `json:"idempotency_key"`
	Query          *ListEndUserInvitesQuery `json:"query"`
}

type ListEndUserSessionsArgs struct {
	IdempotencyKey string                  `json:"idempotency_key"`
	Path           ListEndUserSessionsPath `json:"path"`
}

type ListEndUsersArgs struct {
	IdempotencyKey string             `json:"idempotency_key"`
	Query          *ListEndUsersQuery `json:"query"`
}

type ListEventsArgs struct {
	IdempotencyKey string           `json:"idempotency_key"`
	Query          *ListEventsQuery `json:"query"`
}

type ListFeatureFlagsArgs struct {
	IdempotencyKey string                 `json:"idempotency_key"`
	Query          *ListFeatureFlagsQuery `json:"query"`
}

type ListImportsArgs struct {
	IdempotencyKey string            `json:"idempotency_key"`
	Query          *ListImportsQuery `json:"query"`
}

type ListItContactsArgs struct {
	IdempotencyKey string               `json:"idempotency_key"`
	Path           ListItContactsPath   `json:"path"`
	Query          *ListItContactsQuery `json:"query"`
}

type ListLogStreamsArgs struct {
	IdempotencyKey string `json:"idempotency_key"`
}

type ListM2MClientsArgs struct {
	IdempotencyKey string               `json:"idempotency_key"`
	Query          *ListM2MClientsQuery `json:"query"`
}

type ListMembershipPermissionsArgs struct {
	IdempotencyKey string                          `json:"idempotency_key"`
	Path           ListMembershipPermissionsPath   `json:"path"`
	Query          *ListMembershipPermissionsQuery `json:"query"`
}

type ListMembershipsArgs struct {
	IdempotencyKey string                `json:"idempotency_key"`
	Path           ListMembershipsPath   `json:"path"`
	Query          *ListMembershipsQuery `json:"query"`
}

type ListOAuthClientsArgs struct {
	IdempotencyKey string                 `json:"idempotency_key"`
	Query          *ListOAuthClientsQuery `json:"query"`
}

type ListOAuthGrantsArgs struct {
	IdempotencyKey string                `json:"idempotency_key"`
	Query          *ListOAuthGrantsQuery `json:"query"`
}

type ListOAuthResourcesArgs struct {
	IdempotencyKey string `json:"idempotency_key"`
}

type ListOAuthScopesArgs struct {
	IdempotencyKey string `json:"idempotency_key"`
}

type ListOrganizationAPIKeySecretsArgs struct {
	IdempotencyKey string                            `json:"idempotency_key"`
	Path           ListOrganizationAPIKeySecretsPath `json:"path"`
}

type ListOrganizationAPIKeysArgs struct {
	IdempotencyKey string                        `json:"idempotency_key"`
	Path           ListOrganizationAPIKeysPath   `json:"path"`
	Query          *ListOrganizationAPIKeysQuery `json:"query"`
}

type ListOrganizationEventsArgs struct {
	IdempotencyKey string                       `json:"idempotency_key"`
	Path           ListOrganizationEventsPath   `json:"path"`
	Query          *ListOrganizationEventsQuery `json:"query"`
}

type ListOrganizationsArgs struct {
	IdempotencyKey string                  `json:"idempotency_key"`
	Query          *ListOrganizationsQuery `json:"query"`
}

type ListPassportLinksArgs struct {
	IdempotencyKey string                  `json:"idempotency_key"`
	Query          *ListPassportLinksQuery `json:"query"`
}

type ListPermissionsArgs struct {
	IdempotencyKey string `json:"idempotency_key"`
}

type ListPipeConnectionsArgs struct {
	IdempotencyKey string                    `json:"idempotency_key"`
	Query          *ListPipeConnectionsQuery `json:"query"`
}

type ListPipeProvidersArgs struct {
	IdempotencyKey string `json:"idempotency_key"`
}

type ListPoliciesArgs struct {
	IdempotencyKey string `json:"idempotency_key"`
}

type ListPortalInvitesArgs struct {
	IdempotencyKey string                  `json:"idempotency_key"`
	Query          *ListPortalInvitesQuery `json:"query"`
}

type ListPortalLinksArgs struct {
	IdempotencyKey string                `json:"idempotency_key"`
	Query          *ListPortalLinksQuery `json:"query"`
}

type ListPortalProfilesArgs struct {
	IdempotencyKey string `json:"idempotency_key"`
}

type ListRedirectUrisArgs struct {
	IdempotencyKey string                 `json:"idempotency_key"`
	Query          *ListRedirectUrisQuery `json:"query"`
}

type ListRequestsArgs struct {
	IdempotencyKey string             `json:"idempotency_key"`
	Query          *ListRequestsQuery `json:"query"`
}

type ListRoleAssignmentsArgs struct {
	IdempotencyKey string                  `json:"idempotency_key"`
	Path           ListRoleAssignmentsPath `json:"path"`
}

type ListRoleRulesArgs struct {
	IdempotencyKey string `json:"idempotency_key"`
}

type ListRolesArgs struct {
	IdempotencyKey string          `json:"idempotency_key"`
	Query          *ListRolesQuery `json:"query"`
}

type ListSCIMRequestsArgs struct {
	IdempotencyKey string                 `json:"idempotency_key"`
	Path           ListSCIMRequestsPath   `json:"path"`
	Query          *ListSCIMRequestsQuery `json:"query"`
}

type ListSignInHookDeliveriesArgs struct {
	IdempotencyKey string                         `json:"idempotency_key"`
	Path           ListSignInHookDeliveriesPath   `json:"path"`
	Query          *ListSignInHookDeliveriesQuery `json:"query"`
}

type ListSignInHooksArgs struct {
	IdempotencyKey string `json:"idempotency_key"`
}

type ListSignInsArgs struct {
	IdempotencyKey string            `json:"idempotency_key"`
	Query          *ListSignInsQuery `json:"query"`
}

type ListSocialProvidersArgs struct {
	IdempotencyKey string `json:"idempotency_key"`
}

type ListSsfStreamsArgs struct {
	IdempotencyKey string `json:"idempotency_key"`
}

type ListTrustedIssuersArgs struct {
	IdempotencyKey string `json:"idempotency_key"`
}

type ListVaultKeysArgs struct {
	IdempotencyKey string             `json:"idempotency_key"`
	Query          ListVaultKeysQuery `json:"query"`
}

type ListVaultObjectVersionsArgs struct {
	IdempotencyKey string                       `json:"idempotency_key"`
	Path           ListVaultObjectVersionsPath  `json:"path"`
	Query          ListVaultObjectVersionsQuery `json:"query"`
}

type ListWebhookDeliveriesArgs struct {
	IdempotencyKey string                      `json:"idempotency_key"`
	Path           ListWebhookDeliveriesPath   `json:"path"`
	Query          *ListWebhookDeliveriesQuery `json:"query"`
}

type ListWebhookEndpointsArgs struct {
	IdempotencyKey string                     `json:"idempotency_key"`
	Query          *ListWebhookEndpointsQuery `json:"query"`
}

type MigrateSessionArgs struct {
	IdempotencyKey string             `json:"idempotency_key"`
	Body           MigrateSessionBody `json:"body"`
}

type PreviewAttributesArgs struct {
	IdempotencyKey string                `json:"idempotency_key"`
	Body           PreviewAttributesBody `json:"body"`
}

type PreviewDirectorySourceArgs struct {
	IdempotencyKey string                     `json:"idempotency_key"`
	Path           PreviewDirectorySourcePath `json:"path"`
}

type PreviewRoleRulesArgs struct {
	IdempotencyKey string               `json:"idempotency_key"`
	Body           PreviewRoleRulesBody `json:"body"`
}

type PutDirectorySourceArgs struct {
	IdempotencyKey string                 `json:"idempotency_key"`
	Path           PutDirectorySourcePath `json:"path"`
	Body           PutDirectorySourceBody `json:"body"`
}

type PutItEmailConfigurationArgs struct {
	IdempotencyKey string                      `json:"idempotency_key"`
	Body           PutItEmailConfigurationBody `json:"body"`
}

type PutOAuthScopeArgs struct {
	IdempotencyKey string            `json:"idempotency_key"`
	Path           PutOAuthScopePath `json:"path"`
	Body           PutOAuthScopeBody `json:"body"`
}

type PutPipeProviderArgs struct {
	IdempotencyKey string              `json:"idempotency_key"`
	Path           PutPipeProviderPath `json:"path"`
	Body           PutPipeProviderBody `json:"body"`
}

type PutRadarAlertSubscriptionArgs struct {
	IdempotencyKey string                        `json:"idempotency_key"`
	Body           PutRadarAlertSubscriptionBody `json:"body"`
}

type PutSignInRiskPolicyArgs struct {
	IdempotencyKey string                  `json:"idempotency_key"`
	Body           PutSignInRiskPolicyBody `json:"body"`
}

type PutVaultObjectArgs struct {
	IdempotencyKey string             `json:"idempotency_key"`
	Path           PutVaultObjectPath `json:"path"`
	Body           PutVaultObjectBody `json:"body"`
}

type RegisterEndUserPasskeyArgs struct {
	IdempotencyKey string                     `json:"idempotency_key"`
	Path           RegisterEndUserPasskeyPath `json:"path"`
	Body           RegisterEndUserPasskeyBody `json:"body"`
}

type ReplaceAttributeCatalogueArgs struct {
	IdempotencyKey string                        `json:"idempotency_key"`
	Body           ReplaceAttributeCatalogueBody `json:"body"`
}

type ReplaceConnectionAttributeMapArgs struct {
	IdempotencyKey string                            `json:"idempotency_key"`
	Path           ReplaceConnectionAttributeMapPath `json:"path"`
	Body           ReplaceConnectionAttributeMapBody `json:"body"`
}

type ReplaceDirectoryAttributeMapArgs struct {
	IdempotencyKey string                           `json:"idempotency_key"`
	Path           ReplaceDirectoryAttributeMapPath `json:"path"`
	Body           ReplaceDirectoryAttributeMapBody `json:"body"`
}

type ReplacePortalProfileArgs struct {
	IdempotencyKey string                   `json:"idempotency_key"`
	Path           ReplacePortalProfilePath `json:"path"`
	Body           ReplacePortalProfileBody `json:"body"`
}

type ReplayDirectorySCIMArgs struct {
	IdempotencyKey string                  `json:"idempotency_key"`
	Path           ReplayDirectorySCIMPath `json:"path"`
	Body           ReplayDirectorySCIMBody `json:"body"`
}

type ResendPortalInviteArgs struct {
	IdempotencyKey string                 `json:"idempotency_key"`
	Path           ResendPortalInvitePath `json:"path"`
}

type ResendWebhookDeliveryArgs struct {
	IdempotencyKey string                    `json:"idempotency_key"`
	Path           ResendWebhookDeliveryPath `json:"path"`
}

type RetireConnectionClientKeysArgs struct {
	IdempotencyKey string                         `json:"idempotency_key"`
	Path           RetireConnectionClientKeysPath `json:"path"`
}

type RevokeConnectionShareArgs struct {
	IdempotencyKey string                    `json:"idempotency_key"`
	Path           RevokeConnectionSharePath `json:"path"`
}

type RevokeEndUserInviteArgs struct {
	IdempotencyKey string                  `json:"idempotency_key"`
	Path           RevokeEndUserInvitePath `json:"path"`
}

type RevokeEndUserSessionArgs struct {
	IdempotencyKey string                   `json:"idempotency_key"`
	Path           RevokeEndUserSessionPath `json:"path"`
}

type RevokeEndUserSessionsArgs struct {
	IdempotencyKey string                    `json:"idempotency_key"`
	Path           RevokeEndUserSessionsPath `json:"path"`
}

type RevokeM2MClientArgs struct {
	IdempotencyKey string              `json:"idempotency_key"`
	Path           RevokeM2MClientPath `json:"path"`
}

type RevokeOAuthGrantArgs struct {
	IdempotencyKey string               `json:"idempotency_key"`
	Path           RevokeOAuthGrantPath `json:"path"`
}

type RevokeOrganizationAPIKeyArgs struct {
	IdempotencyKey string                       `json:"idempotency_key"`
	Path           RevokeOrganizationAPIKeyPath `json:"path"`
}

type RevokeOrganizationAPIKeySecretArgs struct {
	IdempotencyKey string                             `json:"idempotency_key"`
	Path           RevokeOrganizationAPIKeySecretPath `json:"path"`
}

type RevokeOrganizationOAuthGrantsArgs struct {
	IdempotencyKey string                            `json:"idempotency_key"`
	Path           RevokeOrganizationOAuthGrantsPath `json:"path"`
	Body           RevokeOrganizationOAuthGrantsBody `json:"body"`
}

type RevokePassportLinkArgs struct {
	IdempotencyKey string                 `json:"idempotency_key"`
	Path           RevokePassportLinkPath `json:"path"`
}

type RevokePortalInviteArgs struct {
	IdempotencyKey string                 `json:"idempotency_key"`
	Path           RevokePortalInvitePath `json:"path"`
}

type RevokePortalLinkArgs struct {
	IdempotencyKey string               `json:"idempotency_key"`
	Path           RevokePortalLinkPath `json:"path"`
}

type RollOrganizationAPIKeyArgs struct {
	IdempotencyKey string                     `json:"idempotency_key"`
	Path           RollOrganizationAPIKeyPath `json:"path"`
}

type RollbackImportOrganizationArgs struct {
	IdempotencyKey string                          `json:"idempotency_key"`
	Path           RollbackImportOrganizationPath  `json:"path"`
	Body           *RollbackImportOrganizationBody `json:"body"`
}

type RotateDirectoryTokenArgs struct {
	IdempotencyKey string                   `json:"idempotency_key"`
	Path           RotateDirectoryTokenPath `json:"path"`
}

type RotateM2MClientSecretArgs struct {
	IdempotencyKey string                    `json:"idempotency_key"`
	Path           RotateM2MClientSecretPath `json:"path"`
	Body           RotateM2MClientSecretBody `json:"body"`
}

type RotateOAuthClientSecretArgs struct {
	IdempotencyKey string                      `json:"idempotency_key"`
	Path           RotateOAuthClientSecretPath `json:"path"`
}

type RotateSignInHookSecretArgs struct {
	IdempotencyKey string                     `json:"idempotency_key"`
	Path           RotateSignInHookSecretPath `json:"path"`
}

type RotateVaultKeyArgs struct {
	IdempotencyKey string             `json:"idempotency_key"`
	Path           RotateVaultKeyPath `json:"path"`
}

type RotateWebhookSecretArgs struct {
	IdempotencyKey string                  `json:"idempotency_key"`
	Path           RotateWebhookSecretPath `json:"path"`
}

type RunImportArgs struct {
	IdempotencyKey string        `json:"idempotency_key"`
	Path           RunImportPath `json:"path"`
}

type SendTestEmailArgs struct {
	IdempotencyKey string            `json:"idempotency_key"`
	Body           SendTestEmailBody `json:"body"`
}

type SetConnectionSCIMBindingArgs struct {
	IdempotencyKey string                       `json:"idempotency_key"`
	Path           SetConnectionSCIMBindingPath `json:"path"`
	Body           SetConnectionSCIMBindingBody `json:"body"`
}

type SetConnectionSingleLogoutArgs struct {
	IdempotencyKey string                        `json:"idempotency_key"`
	Path           SetConnectionSingleLogoutPath `json:"path"`
	Body           SetConnectionSingleLogoutBody `json:"body"`
}

type SetEndUserPasswordArgs struct {
	IdempotencyKey string                 `json:"idempotency_key"`
	Path           SetEndUserPasswordPath `json:"path"`
	Body           SetEndUserPasswordBody `json:"body"`
}

type SetEnvironmentSigninHostArgs struct {
	IdempotencyKey string                       `json:"idempotency_key"`
	Body           SetEnvironmentSigninHostBody `json:"body"`
}

type SetOrganizationAllowedIntentsArgs struct {
	IdempotencyKey string                            `json:"idempotency_key"`
	Path           SetOrganizationAllowedIntentsPath `json:"path"`
	Body           SetOrganizationAllowedIntentsBody `json:"body"`
}

type StartDirectorySourceOAuthArgs struct {
	IdempotencyKey string                        `json:"idempotency_key"`
	Path           StartDirectorySourceOAuthPath `json:"path"`
	Body           StartDirectorySourceOAuthBody `json:"body"`
}

type StartEndUserPasskeyAuthenticationArgs struct {
	IdempotencyKey string `json:"idempotency_key"`
}

type StartEndUserPasskeyRegistrationArgs struct {
	IdempotencyKey string                              `json:"idempotency_key"`
	Path           StartEndUserPasskeyRegistrationPath `json:"path"`
}

type StartSSOLogoutArgs struct {
	IdempotencyKey string             `json:"idempotency_key"`
	Body           StartSSOLogoutBody `json:"body"`
}

type SwitchImportOrganizationArgs struct {
	IdempotencyKey string                       `json:"idempotency_key"`
	Path           SwitchImportOrganizationPath `json:"path"`
}

type SyncDirectorySourceArgs struct {
	IdempotencyKey string                  `json:"idempotency_key"`
	Path           SyncDirectorySourcePath `json:"path"`
}

type TestConnectionArgs struct {
	IdempotencyKey string              `json:"idempotency_key"`
	Path           TestConnectionPath  `json:"path"`
	Body           *TestConnectionBody `json:"body"`
}

type TestSignInHookArgs struct {
	IdempotencyKey string             `json:"idempotency_key"`
	Path           TestSignInHookPath `json:"path"`
}

type TestWebhookEndpointArgs struct {
	IdempotencyKey string                  `json:"idempotency_key"`
	Path           TestWebhookEndpointPath `json:"path"`
	Body           TestWebhookEndpointBody `json:"body"`
}

type UpdateBotProtectionArgs struct {
	IdempotencyKey string                  `json:"idempotency_key"`
	Body           UpdateBotProtectionBody `json:"body"`
}

type UpdateBrandingArgs struct {
	IdempotencyKey string             `json:"idempotency_key"`
	Body           UpdateBrandingBody `json:"body"`
}

type UpdateConnectionArgs struct {
	IdempotencyKey string               `json:"idempotency_key"`
	Path           UpdateConnectionPath `json:"path"`
	Body           UpdateConnectionBody `json:"body"`
}

type UpdateDirectoryArgs struct {
	IdempotencyKey string              `json:"idempotency_key"`
	Path           UpdateDirectoryPath `json:"path"`
	Body           UpdateDirectoryBody `json:"body"`
}

type UpdateDirectorySourceArgs struct {
	IdempotencyKey string                    `json:"idempotency_key"`
	Path           UpdateDirectorySourcePath `json:"path"`
	Body           UpdateDirectorySourceBody `json:"body"`
}

type UpdateDomainArgs struct {
	IdempotencyKey string           `json:"idempotency_key"`
	Path           UpdateDomainPath `json:"path"`
	Body           UpdateDomainBody `json:"body"`
}

type UpdateEndUserArgs struct {
	IdempotencyKey string            `json:"idempotency_key"`
	Path           UpdateEndUserPath `json:"path"`
	Body           UpdateEndUserBody `json:"body"`
}

type UpdateEndUserPasskeyArgs struct {
	IdempotencyKey string                   `json:"idempotency_key"`
	Path           UpdateEndUserPasskeyPath `json:"path"`
	Body           UpdateEndUserPasskeyBody `json:"body"`
}

type UpdateEnvironmentArgs struct {
	IdempotencyKey string                `json:"idempotency_key"`
	Body           UpdateEnvironmentBody `json:"body"`
}

type UpdateFeatureFlagArgs struct {
	IdempotencyKey string                `json:"idempotency_key"`
	Path           UpdateFeatureFlagPath `json:"path"`
	Body           UpdateFeatureFlagBody `json:"body"`
}

type UpdateItContactArgs struct {
	IdempotencyKey string              `json:"idempotency_key"`
	Path           UpdateItContactPath `json:"path"`
	Body           UpdateItContactBody `json:"body"`
}

type UpdateM2MClientArgs struct {
	IdempotencyKey string              `json:"idempotency_key"`
	Path           UpdateM2MClientPath `json:"path"`
	Body           UpdateM2MClientBody `json:"body"`
}

type UpdateMembershipArgs struct {
	IdempotencyKey string               `json:"idempotency_key"`
	Path           UpdateMembershipPath `json:"path"`
	Body           UpdateMembershipBody `json:"body"`
}

type UpdateOAuthClientArgs struct {
	IdempotencyKey string                `json:"idempotency_key"`
	Path           UpdateOAuthClientPath `json:"path"`
	Body           UpdateOAuthClientBody `json:"body"`
}

type UpdateOAuthResourceArgs struct {
	IdempotencyKey string                  `json:"idempotency_key"`
	Path           UpdateOAuthResourcePath `json:"path"`
	Body           UpdateOAuthResourceBody `json:"body"`
}

type UpdateOrganizationArgs struct {
	IdempotencyKey string                 `json:"idempotency_key"`
	Path           UpdateOrganizationPath `json:"path"`
	Body           UpdateOrganizationBody `json:"body"`
}

type UpdateOrganizationOAuthSettingsArgs struct {
	IdempotencyKey string                              `json:"idempotency_key"`
	Path           UpdateOrganizationOAuthSettingsPath `json:"path"`
	Body           UpdateOrganizationOAuthSettingsBody `json:"body"`
}

type UpdateOrganizationProfileArgs struct {
	IdempotencyKey string                        `json:"idempotency_key"`
	Path           UpdateOrganizationProfilePath `json:"path"`
	Body           UpdateOrganizationProfileBody `json:"body"`
}

type UpdateOrganizationUserManagementSettingsArgs struct {
	IdempotencyKey string                                       `json:"idempotency_key"`
	Path           UpdateOrganizationUserManagementSettingsPath `json:"path"`
	Body           UpdateOrganizationUserManagementSettingsBody `json:"body"`
}

type UpdatePermissionArgs struct {
	IdempotencyKey string               `json:"idempotency_key"`
	Path           UpdatePermissionPath `json:"path"`
	Body           UpdatePermissionBody `json:"body"`
}

type UpdatePolicyArgs struct {
	IdempotencyKey string           `json:"idempotency_key"`
	Path           UpdatePolicyPath `json:"path"`
	Body           UpdatePolicyBody `json:"body"`
}

type UpdateRoleArgs struct {
	IdempotencyKey string         `json:"idempotency_key"`
	Path           UpdateRolePath `json:"path"`
	Body           UpdateRoleBody `json:"body"`
}

type UpdateRoleRuleArgs struct {
	IdempotencyKey string             `json:"idempotency_key"`
	Path           UpdateRoleRulePath `json:"path"`
	Body           UpdateRoleRuleBody `json:"body"`
}

type UpdateSignInHookArgs struct {
	IdempotencyKey string               `json:"idempotency_key"`
	Path           UpdateSignInHookPath `json:"path"`
	Body           UpdateSignInHookBody `json:"body"`
}

type UpdateSocialProviderArgs struct {
	IdempotencyKey string                   `json:"idempotency_key"`
	Path           UpdateSocialProviderPath `json:"path"`
	Body           UpdateSocialProviderBody `json:"body"`
}

type UpdateSsfStreamArgs struct {
	IdempotencyKey string              `json:"idempotency_key"`
	Path           UpdateSsfStreamPath `json:"path"`
	Body           UpdateSsfStreamBody `json:"body"`
}

type UpdateUserManagementSettingsArgs struct {
	IdempotencyKey string                           `json:"idempotency_key"`
	Body           UpdateUserManagementSettingsBody `json:"body"`
}

type UpdateWebhookEndpointArgs struct {
	IdempotencyKey string                    `json:"idempotency_key"`
	Path           UpdateWebhookEndpointPath `json:"path"`
	Body           UpdateWebhookEndpointBody `json:"body"`
}

type UploadBrandingFaviconArgs struct {
	IdempotencyKey string                    `json:"idempotency_key"`
	Body           UploadBrandingFaviconBody `json:"body"`
}

type UploadBrandingFontArgs struct {
	IdempotencyKey string                 `json:"idempotency_key"`
	Body           UploadBrandingFontBody `json:"body"`
}

type UploadBrandingLogoArgs struct {
	IdempotencyKey string                 `json:"idempotency_key"`
	Body           UploadBrandingLogoBody `json:"body"`
}

type ValidateEncryptionKeyArgs struct {
	IdempotencyKey string                    `json:"idempotency_key"`
	Path           ValidateEncryptionKeyPath `json:"path"`
}

type VaultDecryptArgs struct {
	IdempotencyKey string           `json:"idempotency_key"`
	Body           VaultDecryptBody `json:"body"`
}

type VaultEncryptArgs struct {
	IdempotencyKey string           `json:"idempotency_key"`
	Body           VaultEncryptBody `json:"body"`
}

type VerifyCustomDomainArgs struct {
	IdempotencyKey string                 `json:"idempotency_key"`
	Path           VerifyCustomDomainPath `json:"path"`
}

type VerifyDomainArgs struct {
	IdempotencyKey string           `json:"idempotency_key"`
	Path           VerifyDomainPath `json:"path"`
}

type VerifySsfStreamArgs struct {
	IdempotencyKey string              `json:"idempotency_key"`
	Path           VerifySsfStreamPath `json:"path"`
	Body           VerifySsfStreamBody `json:"body"`
}

type AttributeCatalogueResources struct{ client *Client }

func (r *AttributeCatalogueResources) GetAttributeCatalogue(ctx context.Context, args GetAttributeCatalogueArgs) (GetAttributeCatalogueResult, error) {
	var result GetAttributeCatalogueResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	response, err := r.client.request(ctx, "getAttributeCatalogue", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *AttributeCatalogueResources) PreviewAttributes(ctx context.Context, args PreviewAttributesArgs) (PreviewAttributesResult, error) {
	var result PreviewAttributesResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "previewAttributes", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *AttributeCatalogueResources) ReplaceAttributeCatalogue(ctx context.Context, args ReplaceAttributeCatalogueArgs) (ReplaceAttributeCatalogueResult, error) {
	var result ReplaceAttributeCatalogueResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "replaceAttributeCatalogue", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type AuditLogEventsResources struct{ client *Client }

func (r *AuditLogEventsResources) CreateAuditLogEvent(ctx context.Context, args CreateAuditLogEventArgs) (CreateAuditLogEventResult, error) {
	var result CreateAuditLogEventResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createAuditLogEvent", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *AuditLogEventsResources) ExportAuditLogEvents(ctx context.Context, args ExportAuditLogEventsArgs) (ExportAuditLogEventsResult, error) {
	var result ExportAuditLogEventsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "exportAuditLogEvents", input)
	if err != nil {
		return result, err
	}
	return ExportAuditLogEventsResult(response.Body), nil
}

func (r *AuditLogEventsResources) GetAuditLogEvent(ctx context.Context, args GetAuditLogEventArgs) (GetAuditLogEventResult, error) {
	var result GetAuditLogEventResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getAuditLogEvent", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *AuditLogEventsResources) ListAuditLogEvents(ctx context.Context, args ListAuditLogEventsArgs) (ListAuditLogEventsResult, error) {
	var result ListAuditLogEventsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listAuditLogEvents", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *AuditLogEventsResources) ListAuditLogEventsAll(ctx context.Context, args ListAuditLogEventsArgs) iter.Seq2[ListAuditLogEventsItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListAuditLogEventsItem, error) bool) {
		var zero ListAuditLogEventsItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListAuditLogEventsArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listAuditLogEvents", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListAuditLogEventsItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

type AuditLogSchemasResources struct{ client *Client }

func (r *AuditLogSchemasResources) CreateAuditLogSchema(ctx context.Context, args CreateAuditLogSchemaArgs) (CreateAuditLogSchemaResult, error) {
	var result CreateAuditLogSchemaResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createAuditLogSchema", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *AuditLogSchemasResources) GetAuditLogSchema(ctx context.Context, args GetAuditLogSchemaArgs) (GetAuditLogSchemaResult, error) {
	var result GetAuditLogSchemaResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "getAuditLogSchema", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *AuditLogSchemasResources) ListAuditLogSchemas(ctx context.Context, args ListAuditLogSchemasArgs) (ListAuditLogSchemasResult, error) {
	var result ListAuditLogSchemasResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listAuditLogSchemas", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *AuditLogSchemasResources) ListAuditLogSchemasAll(ctx context.Context, args ListAuditLogSchemasArgs) iter.Seq2[ListAuditLogSchemasItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListAuditLogSchemasItem, error) bool) {
		var zero ListAuditLogSchemasItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListAuditLogSchemasArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listAuditLogSchemas", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListAuditLogSchemasItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

type AuditLogsResources struct{ client *Client }

func (r *AuditLogsResources) CreateAuditLogExport(ctx context.Context, args CreateAuditLogExportArgs) (CreateAuditLogExportResult, error) {
	var result CreateAuditLogExportResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createAuditLogExport", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *AuditLogsResources) GetAuditLogExport(ctx context.Context, args GetAuditLogExportArgs) (GetAuditLogExportResult, error) {
	var result GetAuditLogExportResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getAuditLogExport", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *AuditLogsResources) ListAuditLogs(ctx context.Context, args ListAuditLogsArgs) (ListAuditLogsResult, error) {
	var result ListAuditLogsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listAuditLogs", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *AuditLogsResources) ListAuditLogsAll(ctx context.Context, args ListAuditLogsArgs) iter.Seq2[ListAuditLogsItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListAuditLogsItem, error) bool) {
		var zero ListAuditLogsItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListAuditLogsArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listAuditLogs", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListAuditLogsItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

type AuthorizationResources struct{ client *Client }

func (r *AuthorizationResources) CheckAuthorization(ctx context.Context, args CheckAuthorizationArgs) (CheckAuthorizationResult, error) {
	var result CheckAuthorizationResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "checkAuthorization", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *AuthorizationResources) CheckAuthorizationBulk(ctx context.Context, args CheckAuthorizationBulkArgs) (CheckAuthorizationBulkResult, error) {
	var result CheckAuthorizationBulkResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "checkAuthorizationBulk", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *AuthorizationResources) CreatePermission(ctx context.Context, args CreatePermissionArgs) (CreatePermissionResult, error) {
	var result CreatePermissionResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createPermission", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *AuthorizationResources) CreateRole(ctx context.Context, args CreateRoleArgs) (CreateRoleResult, error) {
	var result CreateRoleResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createRole", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *AuthorizationResources) DeletePermission(ctx context.Context, args DeletePermissionArgs) (DeletePermissionResult, error) {
	var result DeletePermissionResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deletePermission", input)
	if err != nil {
		return result, err
	}
	_ = response
	return result, nil
}

func (r *AuthorizationResources) DeleteRole(ctx context.Context, args DeleteRoleArgs) (DeleteRoleResult, error) {
	var result DeleteRoleResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deleteRole", input)
	if err != nil {
		return result, err
	}
	_ = response
	return result, nil
}

func (r *AuthorizationResources) GetPermission(ctx context.Context, args GetPermissionArgs) (GetPermissionResult, error) {
	var result GetPermissionResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getPermission", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *AuthorizationResources) GetRole(ctx context.Context, args GetRoleArgs) (GetRoleResult, error) {
	var result GetRoleResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getRole", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *AuthorizationResources) ListPermissions(ctx context.Context, args ListPermissionsArgs) (ListPermissionsResult, error) {
	var result ListPermissionsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	response, err := r.client.request(ctx, "listPermissions", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *AuthorizationResources) ListPermissionsAll(ctx context.Context, args ListPermissionsArgs) iter.Seq2[ListPermissionsItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListPermissionsItem, error) bool) {
		var zero ListPermissionsItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListPermissionsArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		for item, err := range r.client.paginate(ctx, "listPermissions", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListPermissionsItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

func (r *AuthorizationResources) ListRoles(ctx context.Context, args ListRolesArgs) (ListRolesResult, error) {
	var result ListRolesResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listRoles", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *AuthorizationResources) ListRolesAll(ctx context.Context, args ListRolesArgs) iter.Seq2[ListRolesItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListRolesItem, error) bool) {
		var zero ListRolesItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListRolesArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listRoles", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListRolesItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

func (r *AuthorizationResources) UpdatePermission(ctx context.Context, args UpdatePermissionArgs) (UpdatePermissionResult, error) {
	var result UpdatePermissionResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "updatePermission", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *AuthorizationResources) UpdateRole(ctx context.Context, args UpdateRoleArgs) (UpdateRoleResult, error) {
	var result UpdateRoleResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "updateRole", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type BrandingResources struct{ client *Client }

func (r *BrandingResources) GetBranding(ctx context.Context, args GetBrandingArgs) (GetBrandingResult, error) {
	var result GetBrandingResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	response, err := r.client.request(ctx, "getBranding", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *BrandingResources) UpdateBranding(ctx context.Context, args UpdateBrandingArgs) (UpdateBrandingResult, error) {
	var result UpdateBrandingResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "updateBranding", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *BrandingResources) UploadBrandingFavicon(ctx context.Context, args UploadBrandingFaviconArgs) (UploadBrandingFaviconResult, error) {
	var result UploadBrandingFaviconResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "uploadBrandingFavicon", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *BrandingResources) UploadBrandingFont(ctx context.Context, args UploadBrandingFontArgs) (UploadBrandingFontResult, error) {
	var result UploadBrandingFontResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "uploadBrandingFont", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *BrandingResources) UploadBrandingLogo(ctx context.Context, args UploadBrandingLogoArgs) (UploadBrandingLogoResult, error) {
	var result UploadBrandingLogoResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "uploadBrandingLogo", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type ConnectionHealthResources struct{ client *Client }

func (r *ConnectionHealthResources) ListConnectionHealth(ctx context.Context, args ListConnectionHealthArgs) (ListConnectionHealthResult, error) {
	var result ListConnectionHealthResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listConnectionHealth", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ConnectionHealthResources) ListConnectionHealthAll(ctx context.Context, args ListConnectionHealthArgs) iter.Seq2[ListConnectionHealthItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListConnectionHealthItem, error) bool) {
		var zero ListConnectionHealthItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListConnectionHealthArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listConnectionHealth", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListConnectionHealthItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

type ConnectionsResources struct{ client *Client }

func (r *ConnectionsResources) ActivateConnection(ctx context.Context, args ActivateConnectionArgs) (ActivateConnectionResult, error) {
	var result ActivateConnectionResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "activateConnection", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ConnectionsResources) ActivateConnectionClientKey(ctx context.Context, args ActivateConnectionClientKeyArgs) (ActivateConnectionClientKeyResult, error) {
	var result ActivateConnectionClientKeyResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "activateConnectionClientKey", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ConnectionsResources) CreateConnection(ctx context.Context, args CreateConnectionArgs) (CreateConnectionResult, error) {
	var result CreateConnectionResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createConnection", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ConnectionsResources) CreateConnectionCertificate(ctx context.Context, args CreateConnectionCertificateArgs) (CreateConnectionCertificateResult, error) {
	var result CreateConnectionCertificateResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createConnectionCertificate", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ConnectionsResources) CreateConnectionClientKey(ctx context.Context, args CreateConnectionClientKeyArgs) (CreateConnectionClientKeyResult, error) {
	var result CreateConnectionClientKeyResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createConnectionClientKey", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ConnectionsResources) CreateConnectionShare(ctx context.Context, args CreateConnectionShareArgs) (CreateConnectionShareResult, error) {
	var result CreateConnectionShareResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createConnectionShare", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ConnectionsResources) DeactivateConnection(ctx context.Context, args DeactivateConnectionArgs) (DeactivateConnectionResult, error) {
	var result DeactivateConnectionResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deactivateConnection", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ConnectionsResources) DeleteConnection(ctx context.Context, args DeleteConnectionArgs) (DeleteConnectionResult, error) {
	var result DeleteConnectionResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Query, err = json.Marshal(args.Query)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deleteConnection", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ConnectionsResources) DeleteConnectionCertificate(ctx context.Context, args DeleteConnectionCertificateArgs) (DeleteConnectionCertificateResult, error) {
	var result DeleteConnectionCertificateResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deleteConnectionCertificate", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ConnectionsResources) DeleteConnectionSCIMBinding(ctx context.Context, args DeleteConnectionSCIMBindingArgs) (DeleteConnectionSCIMBindingResult, error) {
	var result DeleteConnectionSCIMBindingResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deleteConnectionSCIMBinding", input)
	if err != nil {
		return result, err
	}
	_ = response
	return result, nil
}

func (r *ConnectionsResources) GetConnection(ctx context.Context, args GetConnectionArgs) (GetConnectionResult, error) {
	var result GetConnectionResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "getConnection", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ConnectionsResources) GetConnectionAttributeMap(ctx context.Context, args GetConnectionAttributeMapArgs) (GetConnectionAttributeMapResult, error) {
	var result GetConnectionAttributeMapResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getConnectionAttributeMap", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ConnectionsResources) GetConnectionAttributePreview(ctx context.Context, args GetConnectionAttributePreviewArgs) (GetConnectionAttributePreviewResult, error) {
	var result GetConnectionAttributePreviewResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getConnectionAttributePreview", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ConnectionsResources) GetConnectionHealth(ctx context.Context, args GetConnectionHealthArgs) (GetConnectionHealthResult, error) {
	var result GetConnectionHealthResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getConnectionHealth", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ConnectionsResources) GetConnectionSingleLogout(ctx context.Context, args GetConnectionSingleLogoutArgs) (GetConnectionSingleLogoutResult, error) {
	var result GetConnectionSingleLogoutResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getConnectionSingleLogout", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ConnectionsResources) GetConnectionSPMetadata(ctx context.Context, args GetConnectionSPMetadataArgs) (GetConnectionSPMetadataResult, error) {
	var result GetConnectionSPMetadataResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getConnectionSpMetadata", input)
	if err != nil {
		return result, err
	}
	return GetConnectionSPMetadataResult(response.Body), nil
}

func (r *ConnectionsResources) GetConnectionTestResult(ctx context.Context, args GetConnectionTestResultArgs) (GetConnectionTestResultResult, error) {
	var result GetConnectionTestResultResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getConnectionTestResult", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ConnectionsResources) ListConnectionCertificates(ctx context.Context, args ListConnectionCertificatesArgs) (ListConnectionCertificatesResult, error) {
	var result ListConnectionCertificatesResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listConnectionCertificates", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ConnectionsResources) ListConnectionCertificatesAll(ctx context.Context, args ListConnectionCertificatesArgs) iter.Seq2[ListConnectionCertificatesItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListConnectionCertificatesItem, error) bool) {
		var zero ListConnectionCertificatesItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListConnectionCertificatesArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		input.Path, err = json.Marshal(args.Path)
		if err != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listConnectionCertificates", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListConnectionCertificatesItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

func (r *ConnectionsResources) ListConnectionClientKeys(ctx context.Context, args ListConnectionClientKeysArgs) (ListConnectionClientKeysResult, error) {
	var result ListConnectionClientKeysResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "listConnectionClientKeys", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ConnectionsResources) ListConnectionShares(ctx context.Context, args ListConnectionSharesArgs) (ListConnectionSharesResult, error) {
	var result ListConnectionSharesResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "listConnectionShares", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ConnectionsResources) ListConnectionSignIns(ctx context.Context, args ListConnectionSignInsArgs) (ListConnectionSignInsResult, error) {
	var result ListConnectionSignInsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listConnectionSignIns", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ConnectionsResources) ListConnectionSignInsAll(ctx context.Context, args ListConnectionSignInsArgs) iter.Seq2[ListConnectionSignInsItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListConnectionSignInsItem, error) bool) {
		var zero ListConnectionSignInsItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListConnectionSignInsArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		input.Path, err = json.Marshal(args.Path)
		if err != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listConnectionSignIns", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListConnectionSignInsItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

func (r *ConnectionsResources) ListConnections(ctx context.Context, args ListConnectionsArgs) (ListConnectionsResult, error) {
	var result ListConnectionsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listConnections", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ConnectionsResources) ListConnectionsAll(ctx context.Context, args ListConnectionsArgs) iter.Seq2[ListConnectionsItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListConnectionsItem, error) bool) {
		var zero ListConnectionsItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListConnectionsArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listConnections", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListConnectionsItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

func (r *ConnectionsResources) ReplaceConnectionAttributeMap(ctx context.Context, args ReplaceConnectionAttributeMapArgs) (ReplaceConnectionAttributeMapResult, error) {
	var result ReplaceConnectionAttributeMapResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "replaceConnectionAttributeMap", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ConnectionsResources) RetireConnectionClientKeys(ctx context.Context, args RetireConnectionClientKeysArgs) (RetireConnectionClientKeysResult, error) {
	var result RetireConnectionClientKeysResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "retireConnectionClientKeys", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ConnectionsResources) RevokeConnectionShare(ctx context.Context, args RevokeConnectionShareArgs) (RevokeConnectionShareResult, error) {
	var result RevokeConnectionShareResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "revokeConnectionShare", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ConnectionsResources) SetConnectionSCIMBinding(ctx context.Context, args SetConnectionSCIMBindingArgs) (SetConnectionSCIMBindingResult, error) {
	var result SetConnectionSCIMBindingResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "setConnectionSCIMBinding", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ConnectionsResources) SetConnectionSingleLogout(ctx context.Context, args SetConnectionSingleLogoutArgs) (SetConnectionSingleLogoutResult, error) {
	var result SetConnectionSingleLogoutResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "setConnectionSingleLogout", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ConnectionsResources) TestConnection(ctx context.Context, args TestConnectionArgs) (TestConnectionResult, error) {
	var result TestConnectionResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	if args.Body != nil {
		input.Body, err = json.Marshal(args.Body)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "testConnection", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ConnectionsResources) UpdateConnection(ctx context.Context, args UpdateConnectionArgs) (UpdateConnectionResult, error) {
	var result UpdateConnectionResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "updateConnection", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type CustomDomainsResources struct{ client *Client }

func (r *CustomDomainsResources) CreateCustomDomain(ctx context.Context, args CreateCustomDomainArgs) (CreateCustomDomainResult, error) {
	var result CreateCustomDomainResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createCustomDomain", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *CustomDomainsResources) DeleteCustomDomain(ctx context.Context, args DeleteCustomDomainArgs) (DeleteCustomDomainResult, error) {
	var result DeleteCustomDomainResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deleteCustomDomain", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *CustomDomainsResources) GetCustomDomain(ctx context.Context, args GetCustomDomainArgs) (GetCustomDomainResult, error) {
	var result GetCustomDomainResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getCustomDomain", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *CustomDomainsResources) ListCustomDomains(ctx context.Context, args ListCustomDomainsArgs) (ListCustomDomainsResult, error) {
	var result ListCustomDomainsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listCustomDomains", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *CustomDomainsResources) ListCustomDomainsAll(ctx context.Context, args ListCustomDomainsArgs) iter.Seq2[ListCustomDomainsItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListCustomDomainsItem, error) bool) {
		var zero ListCustomDomainsItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListCustomDomainsArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listCustomDomains", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListCustomDomainsItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

func (r *CustomDomainsResources) VerifyCustomDomain(ctx context.Context, args VerifyCustomDomainArgs) (VerifyCustomDomainResult, error) {
	var result VerifyCustomDomainResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "verifyCustomDomain", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type DevResources struct{ client *Client }

func (r *DevResources) DevInit(ctx context.Context, args DevInitArgs) (DevInitResult, error) {
	var result DevInitResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Body != nil {
		input.Body, err = json.Marshal(args.Body)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "devInit", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *DevResources) DevTrigger(ctx context.Context, args DevTriggerArgs) (DevTriggerResult, error) {
	var result DevTriggerResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "devTrigger", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type DirectoriesResources struct{ client *Client }

func (r *DirectoriesResources) ChooseDirectorySourceOAuthCompany(ctx context.Context, args ChooseDirectorySourceOAuthCompanyArgs) (ChooseDirectorySourceOAuthCompanyResult, error) {
	var result ChooseDirectorySourceOAuthCompanyResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "chooseDirectorySourceOAuthCompany", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *DirectoriesResources) CreateDirectory(ctx context.Context, args CreateDirectoryArgs) (CreateDirectoryResult, error) {
	var result CreateDirectoryResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createDirectory", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *DirectoriesResources) CreateDirectorySourceSshKey(ctx context.Context, args CreateDirectorySourceSshKeyArgs) (CreateDirectorySourceSshKeyResult, error) {
	var result CreateDirectorySourceSshKeyResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createDirectorySourceSSHKey", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *DirectoriesResources) DeleteDirectory(ctx context.Context, args DeleteDirectoryArgs) (DeleteDirectoryResult, error) {
	var result DeleteDirectoryResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "deleteDirectory", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *DirectoriesResources) GetDirectory(ctx context.Context, args GetDirectoryArgs) (GetDirectoryResult, error) {
	var result GetDirectoryResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getDirectory", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *DirectoriesResources) GetDirectoryAttributeMap(ctx context.Context, args GetDirectoryAttributeMapArgs) (GetDirectoryAttributeMapResult, error) {
	var result GetDirectoryAttributeMapResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getDirectoryAttributeMap", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *DirectoriesResources) GetDirectorySource(ctx context.Context, args GetDirectorySourceArgs) (GetDirectorySourceResult, error) {
	var result GetDirectorySourceResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getDirectorySource", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *DirectoriesResources) GetDirectorySourceOAuth(ctx context.Context, args GetDirectorySourceOAuthArgs) (GetDirectorySourceOAuthResult, error) {
	var result GetDirectorySourceOAuthResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getDirectorySourceOAuth", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *DirectoriesResources) GetDirectorySourceSshKey(ctx context.Context, args GetDirectorySourceSshKeyArgs) (GetDirectorySourceSshKeyResult, error) {
	var result GetDirectorySourceSshKeyResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getDirectorySourceSSHKey", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *DirectoriesResources) ImportDirectoryCsv(ctx context.Context, args ImportDirectoryCsvArgs) (ImportDirectoryCsvResult, error) {
	var result ImportDirectoryCsvResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "importDirectoryCSV", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *DirectoriesResources) ListDirectories(ctx context.Context, args ListDirectoriesArgs) (ListDirectoriesResult, error) {
	var result ListDirectoriesResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listDirectories", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *DirectoriesResources) ListDirectoriesAll(ctx context.Context, args ListDirectoriesArgs) iter.Seq2[ListDirectoriesItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListDirectoriesItem, error) bool) {
		var zero ListDirectoriesItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListDirectoriesArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listDirectories", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListDirectoriesItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

func (r *DirectoriesResources) ListDirectoryGroups(ctx context.Context, args ListDirectoryGroupsArgs) (ListDirectoryGroupsResult, error) {
	var result ListDirectoryGroupsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listDirectoryGroups", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *DirectoriesResources) ListDirectoryGroupsAll(ctx context.Context, args ListDirectoryGroupsArgs) iter.Seq2[ListDirectoryGroupsItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListDirectoryGroupsItem, error) bool) {
		var zero ListDirectoryGroupsItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListDirectoryGroupsArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		input.Path, err = json.Marshal(args.Path)
		if err != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listDirectoryGroups", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListDirectoryGroupsItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

func (r *DirectoriesResources) ListDirectoryUsers(ctx context.Context, args ListDirectoryUsersArgs) (ListDirectoryUsersResult, error) {
	var result ListDirectoryUsersResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listDirectoryUsers", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *DirectoriesResources) ListDirectoryUsersAll(ctx context.Context, args ListDirectoryUsersArgs) iter.Seq2[ListDirectoryUsersItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListDirectoryUsersItem, error) bool) {
		var zero ListDirectoryUsersItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListDirectoryUsersArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		input.Path, err = json.Marshal(args.Path)
		if err != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listDirectoryUsers", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListDirectoryUsersItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

func (r *DirectoriesResources) ListSCIMRequests(ctx context.Context, args ListSCIMRequestsArgs) (ListSCIMRequestsResult, error) {
	var result ListSCIMRequestsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listScimRequests", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *DirectoriesResources) ListSCIMRequestsAll(ctx context.Context, args ListSCIMRequestsArgs) iter.Seq2[ListSCIMRequestsItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListSCIMRequestsItem, error) bool) {
		var zero ListSCIMRequestsItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListSCIMRequestsArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		input.Path, err = json.Marshal(args.Path)
		if err != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listScimRequests", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListSCIMRequestsItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

func (r *DirectoriesResources) PreviewDirectorySource(ctx context.Context, args PreviewDirectorySourceArgs) (PreviewDirectorySourceResult, error) {
	var result PreviewDirectorySourceResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "previewDirectorySource", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *DirectoriesResources) PutDirectorySource(ctx context.Context, args PutDirectorySourceArgs) (PutDirectorySourceResult, error) {
	var result PutDirectorySourceResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "putDirectorySource", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *DirectoriesResources) ReplaceDirectoryAttributeMap(ctx context.Context, args ReplaceDirectoryAttributeMapArgs) (ReplaceDirectoryAttributeMapResult, error) {
	var result ReplaceDirectoryAttributeMapResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "replaceDirectoryAttributeMap", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *DirectoriesResources) ReplayDirectorySCIM(ctx context.Context, args ReplayDirectorySCIMArgs) (ReplayDirectorySCIMResult, error) {
	var result ReplayDirectorySCIMResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "replayDirectoryScim", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *DirectoriesResources) RotateDirectoryToken(ctx context.Context, args RotateDirectoryTokenArgs) (RotateDirectoryTokenResult, error) {
	var result RotateDirectoryTokenResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "rotateDirectoryToken", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *DirectoriesResources) StartDirectorySourceOAuth(ctx context.Context, args StartDirectorySourceOAuthArgs) (StartDirectorySourceOAuthResult, error) {
	var result StartDirectorySourceOAuthResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "startDirectorySourceOAuth", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *DirectoriesResources) SyncDirectorySource(ctx context.Context, args SyncDirectorySourceArgs) (SyncDirectorySourceResult, error) {
	var result SyncDirectorySourceResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "syncDirectorySource", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *DirectoriesResources) UpdateDirectory(ctx context.Context, args UpdateDirectoryArgs) (UpdateDirectoryResult, error) {
	var result UpdateDirectoryResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "updateDirectory", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *DirectoriesResources) UpdateDirectorySource(ctx context.Context, args UpdateDirectorySourceArgs) (UpdateDirectorySourceResult, error) {
	var result UpdateDirectorySourceResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "updateDirectorySource", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type DirectoryGroupsResources struct{ client *Client }

func (r *DirectoryGroupsResources) GetDirectoryGroup(ctx context.Context, args GetDirectoryGroupArgs) (GetDirectoryGroupResult, error) {
	var result GetDirectoryGroupResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getDirectoryGroup", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *DirectoryGroupsResources) ListDirectoryGroupMembers(ctx context.Context, args ListDirectoryGroupMembersArgs) (ListDirectoryGroupMembersResult, error) {
	var result ListDirectoryGroupMembersResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listDirectoryGroupMembers", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *DirectoryGroupsResources) ListDirectoryGroupMembersAll(ctx context.Context, args ListDirectoryGroupMembersArgs) iter.Seq2[ListDirectoryGroupMembersItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListDirectoryGroupMembersItem, error) bool) {
		var zero ListDirectoryGroupMembersItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListDirectoryGroupMembersArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		input.Path, err = json.Marshal(args.Path)
		if err != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listDirectoryGroupMembers", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListDirectoryGroupMembersItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

type DirectoryUsersResources struct{ client *Client }

func (r *DirectoryUsersResources) GetDirectoryUser(ctx context.Context, args GetDirectoryUserArgs) (GetDirectoryUserResult, error) {
	var result GetDirectoryUserResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getDirectoryUser", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *DirectoryUsersResources) ListDirectoryUserGroups(ctx context.Context, args ListDirectoryUserGroupsArgs) (ListDirectoryUserGroupsResult, error) {
	var result ListDirectoryUserGroupsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listDirectoryUserGroups", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *DirectoryUsersResources) ListDirectoryUserGroupsAll(ctx context.Context, args ListDirectoryUserGroupsArgs) iter.Seq2[ListDirectoryUserGroupsItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListDirectoryUserGroupsItem, error) bool) {
		var zero ListDirectoryUserGroupsItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListDirectoryUserGroupsArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		input.Path, err = json.Marshal(args.Path)
		if err != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listDirectoryUserGroups", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListDirectoryUserGroupsItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

type DomainsResources struct{ client *Client }

func (r *DomainsResources) CreateDomain(ctx context.Context, args CreateDomainArgs) (CreateDomainResult, error) {
	var result CreateDomainResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createDomain", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *DomainsResources) DeleteDomain(ctx context.Context, args DeleteDomainArgs) (DeleteDomainResult, error) {
	var result DeleteDomainResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deleteDomain", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *DomainsResources) ListDomains(ctx context.Context, args ListDomainsArgs) (ListDomainsResult, error) {
	var result ListDomainsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listDomains", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *DomainsResources) ListDomainsAll(ctx context.Context, args ListDomainsArgs) iter.Seq2[ListDomainsItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListDomainsItem, error) bool) {
		var zero ListDomainsItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListDomainsArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		input.Path, err = json.Marshal(args.Path)
		if err != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listDomains", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListDomainsItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

func (r *DomainsResources) UpdateDomain(ctx context.Context, args UpdateDomainArgs) (UpdateDomainResult, error) {
	var result UpdateDomainResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "updateDomain", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *DomainsResources) VerifyDomain(ctx context.Context, args VerifyDomainArgs) (VerifyDomainResult, error) {
	var result VerifyDomainResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "verifyDomain", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type EmailEventsResources struct{ client *Client }

func (r *EmailEventsResources) ListEmailEvents(ctx context.Context, args ListEmailEventsArgs) (ListEmailEventsResult, error) {
	var result ListEmailEventsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listEmailEvents", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *EmailEventsResources) ListEmailEventsAll(ctx context.Context, args ListEmailEventsArgs) iter.Seq2[ListEmailEventsItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListEmailEventsItem, error) bool) {
		var zero ListEmailEventsItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListEmailEventsArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listEmailEvents", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListEmailEventsItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

func (r *EmailEventsResources) SendTestEmail(ctx context.Context, args SendTestEmailArgs) (SendTestEmailResult, error) {
	var result SendTestEmailResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "sendTestEmail", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type EncryptionKeysResources struct{ client *Client }

func (r *EncryptionKeysResources) ActivateEncryptionKey(ctx context.Context, args ActivateEncryptionKeyArgs) (ActivateEncryptionKeyResult, error) {
	var result ActivateEncryptionKeyResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "activateEncryptionKey", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *EncryptionKeysResources) CreateEncryptionKey(ctx context.Context, args CreateEncryptionKeyArgs) (CreateEncryptionKeyResult, error) {
	var result CreateEncryptionKeyResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createEncryptionKey", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *EncryptionKeysResources) DeprecateEncryptionKey(ctx context.Context, args DeprecateEncryptionKeyArgs) (DeprecateEncryptionKeyResult, error) {
	var result DeprecateEncryptionKeyResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deprecateEncryptionKey", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *EncryptionKeysResources) GetEncryptionKey(ctx context.Context, args GetEncryptionKeyArgs) (GetEncryptionKeyResult, error) {
	var result GetEncryptionKeyResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getEncryptionKey", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *EncryptionKeysResources) GetEncryptionKeySetup(ctx context.Context, args GetEncryptionKeySetupArgs) (GetEncryptionKeySetupResult, error) {
	var result GetEncryptionKeySetupResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "getEncryptionKeySetup", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *EncryptionKeysResources) ListEncryptionKeys(ctx context.Context, args ListEncryptionKeysArgs) (ListEncryptionKeysResult, error) {
	var result ListEncryptionKeysResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	response, err := r.client.request(ctx, "listEncryptionKeys", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *EncryptionKeysResources) ValidateEncryptionKey(ctx context.Context, args ValidateEncryptionKeyArgs) (ValidateEncryptionKeyResult, error) {
	var result ValidateEncryptionKeyResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "validateEncryptionKey", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type EnvironmentResources struct{ client *Client }

func (r *EnvironmentResources) GetEnvironment(ctx context.Context, args GetEnvironmentArgs) (GetEnvironmentResult, error) {
	var result GetEnvironmentResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	response, err := r.client.request(ctx, "getEnvironment", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *EnvironmentResources) GetEnvironmentSigninHost(ctx context.Context, args GetEnvironmentSigninHostArgs) (GetEnvironmentSigninHostResult, error) {
	var result GetEnvironmentSigninHostResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "getEnvironmentSigninHost", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *EnvironmentResources) SetEnvironmentSigninHost(ctx context.Context, args SetEnvironmentSigninHostArgs) (SetEnvironmentSigninHostResult, error) {
	var result SetEnvironmentSigninHostResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "setEnvironmentSigninHost", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *EnvironmentResources) UpdateEnvironment(ctx context.Context, args UpdateEnvironmentArgs) (UpdateEnvironmentResult, error) {
	var result UpdateEnvironmentResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "updateEnvironment", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type EventsResources struct{ client *Client }

func (r *EventsResources) GetEvent(ctx context.Context, args GetEventArgs) (GetEventResult, error) {
	var result GetEventResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getEvent", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *EventsResources) ListEvents(ctx context.Context, args ListEventsArgs) (ListEventsResult, error) {
	var result ListEventsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listEvents", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *EventsResources) ListEventsAll(ctx context.Context, args ListEventsArgs) iter.Seq2[ListEventsItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListEventsItem, error) bool) {
		var zero ListEventsItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListEventsArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listEvents", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListEventsItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

type ExportsResources struct{ client *Client }

func (r *ExportsResources) CreateExport(ctx context.Context, args CreateExportArgs) (CreateExportResult, error) {
	var result CreateExportResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Body != nil {
		input.Body, err = json.Marshal(args.Body)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "createExport", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ExportsResources) GetExport(ctx context.Context, args GetExportArgs) (GetExportResult, error) {
	var result GetExportResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getExport", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type FeatureFlagsResources struct{ client *Client }

func (r *FeatureFlagsResources) CreateFeatureFlag(ctx context.Context, args CreateFeatureFlagArgs) (CreateFeatureFlagResult, error) {
	var result CreateFeatureFlagResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createFeatureFlag", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *FeatureFlagsResources) DeleteFeatureFlag(ctx context.Context, args DeleteFeatureFlagArgs) (DeleteFeatureFlagResult, error) {
	var result DeleteFeatureFlagResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deleteFeatureFlag", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *FeatureFlagsResources) EvaluateFeatureFlag(ctx context.Context, args EvaluateFeatureFlagArgs) (EvaluateFeatureFlagResult, error) {
	var result EvaluateFeatureFlagResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "evaluateFeatureFlag", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *FeatureFlagsResources) EvaluateFeatureFlags(ctx context.Context, args EvaluateFeatureFlagsArgs) (EvaluateFeatureFlagsResult, error) {
	var result EvaluateFeatureFlagsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "evaluateFeatureFlags", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *FeatureFlagsResources) GetFeatureFlag(ctx context.Context, args GetFeatureFlagArgs) (GetFeatureFlagResult, error) {
	var result GetFeatureFlagResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getFeatureFlag", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *FeatureFlagsResources) ListFeatureFlags(ctx context.Context, args ListFeatureFlagsArgs) (ListFeatureFlagsResult, error) {
	var result ListFeatureFlagsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listFeatureFlags", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *FeatureFlagsResources) ListFeatureFlagsAll(ctx context.Context, args ListFeatureFlagsArgs) iter.Seq2[ListFeatureFlagsItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListFeatureFlagsItem, error) bool) {
		var zero ListFeatureFlagsItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListFeatureFlagsArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listFeatureFlags", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListFeatureFlagsItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

func (r *FeatureFlagsResources) UpdateFeatureFlag(ctx context.Context, args UpdateFeatureFlagArgs) (UpdateFeatureFlagResult, error) {
	var result UpdateFeatureFlagResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "updateFeatureFlag", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type IDPRadarResources struct{ client *Client }

func (r *IDPRadarResources) DeleteRadarAlertSubscription(ctx context.Context, args DeleteRadarAlertSubscriptionArgs) (DeleteRadarAlertSubscriptionResult, error) {
	var result DeleteRadarAlertSubscriptionResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	response, err := r.client.request(ctx, "deleteRadarAlertSubscription", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *IDPRadarResources) GetRadarAlertSubscription(ctx context.Context, args GetRadarAlertSubscriptionArgs) (GetRadarAlertSubscriptionResult, error) {
	var result GetRadarAlertSubscriptionResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	response, err := r.client.request(ctx, "getRadarAlertSubscription", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *IDPRadarResources) PutRadarAlertSubscription(ctx context.Context, args PutRadarAlertSubscriptionArgs) (PutRadarAlertSubscriptionResult, error) {
	var result PutRadarAlertSubscriptionResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "putRadarAlertSubscription", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type ImportsResources struct{ client *Client }

func (r *ImportsResources) CancelImport(ctx context.Context, args CancelImportArgs) (CancelImportResult, error) {
	var result CancelImportResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "cancelImport", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ImportsResources) CreateImport(ctx context.Context, args CreateImportArgs) (CreateImportResult, error) {
	var result CreateImportResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createImport", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ImportsResources) GetImport(ctx context.Context, args GetImportArgs) (GetImportResult, error) {
	var result GetImportResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getImport", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ImportsResources) ListImports(ctx context.Context, args ListImportsArgs) (ListImportsResult, error) {
	var result ListImportsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listImports", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ImportsResources) ListImportsAll(ctx context.Context, args ListImportsArgs) iter.Seq2[ListImportsItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListImportsItem, error) bool) {
		var zero ListImportsItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListImportsArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listImports", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListImportsItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

func (r *ImportsResources) RollbackImportOrganization(ctx context.Context, args RollbackImportOrganizationArgs) (RollbackImportOrganizationResult, error) {
	var result RollbackImportOrganizationResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	if args.Body != nil {
		input.Body, err = json.Marshal(args.Body)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "rollbackImportOrganization", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ImportsResources) RunImport(ctx context.Context, args RunImportArgs) (RunImportResult, error) {
	var result RunImportResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "runImport", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ImportsResources) SwitchImportOrganization(ctx context.Context, args SwitchImportOrganizationArgs) (SwitchImportOrganizationResult, error) {
	var result SwitchImportOrganizationResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "switchImportOrganization", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type InvitesResources struct{ client *Client }

func (r *InvitesResources) CreateEndUserInvite(ctx context.Context, args CreateEndUserInviteArgs) (CreateEndUserInviteResult, error) {
	var result CreateEndUserInviteResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createEndUserInvite", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *InvitesResources) ListEndUserInvites(ctx context.Context, args ListEndUserInvitesArgs) (ListEndUserInvitesResult, error) {
	var result ListEndUserInvitesResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listEndUserInvites", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *InvitesResources) ListEndUserInvitesAll(ctx context.Context, args ListEndUserInvitesArgs) iter.Seq2[ListEndUserInvitesItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListEndUserInvitesItem, error) bool) {
		var zero ListEndUserInvitesItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListEndUserInvitesArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listEndUserInvites", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListEndUserInvitesItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

func (r *InvitesResources) RevokeEndUserInvite(ctx context.Context, args RevokeEndUserInviteArgs) (RevokeEndUserInviteResult, error) {
	var result RevokeEndUserInviteResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "revokeEndUserInvite", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type ItContactsResources struct{ client *Client }

func (r *ItContactsResources) CreateItContact(ctx context.Context, args CreateItContactArgs) (CreateItContactResult, error) {
	var result CreateItContactResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createItContact", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ItContactsResources) DeleteItContact(ctx context.Context, args DeleteItContactArgs) (DeleteItContactResult, error) {
	var result DeleteItContactResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deleteItContact", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ItContactsResources) GetItContact(ctx context.Context, args GetItContactArgs) (GetItContactResult, error) {
	var result GetItContactResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getItContact", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ItContactsResources) ListItContacts(ctx context.Context, args ListItContactsArgs) (ListItContactsResult, error) {
	var result ListItContactsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listItContacts", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ItContactsResources) ListItContactsAll(ctx context.Context, args ListItContactsArgs) iter.Seq2[ListItContactsItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListItContactsItem, error) bool) {
		var zero ListItContactsItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListItContactsArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		input.Path, err = json.Marshal(args.Path)
		if err != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listItContacts", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListItContactsItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

func (r *ItContactsResources) UpdateItContact(ctx context.Context, args UpdateItContactArgs) (UpdateItContactResult, error) {
	var result UpdateItContactResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "updateItContact", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type ItEmailConfigurationResources struct{ client *Client }

func (r *ItEmailConfigurationResources) DeleteItEmailConfiguration(ctx context.Context, args DeleteItEmailConfigurationArgs) (DeleteItEmailConfigurationResult, error) {
	var result DeleteItEmailConfigurationResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	response, err := r.client.request(ctx, "deleteITEmailConfiguration", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ItEmailConfigurationResources) GetItEmailConfiguration(ctx context.Context, args GetItEmailConfigurationArgs) (GetItEmailConfigurationResult, error) {
	var result GetItEmailConfigurationResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	response, err := r.client.request(ctx, "getITEmailConfiguration", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *ItEmailConfigurationResources) PutItEmailConfiguration(ctx context.Context, args PutItEmailConfigurationArgs) (PutItEmailConfigurationResult, error) {
	var result PutItEmailConfigurationResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "putITEmailConfiguration", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type LogStreamsResources struct{ client *Client }

func (r *LogStreamsResources) CreateLogStream(ctx context.Context, args CreateLogStreamArgs) (CreateLogStreamResult, error) {
	var result CreateLogStreamResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createLogStream", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *LogStreamsResources) DeleteLogStream(ctx context.Context, args DeleteLogStreamArgs) (DeleteLogStreamResult, error) {
	var result DeleteLogStreamResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deleteLogStream", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *LogStreamsResources) ListLogStreams(ctx context.Context, args ListLogStreamsArgs) (ListLogStreamsResult, error) {
	var result ListLogStreamsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	response, err := r.client.request(ctx, "listLogStreams", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type M2MClientsResources struct{ client *Client }

func (r *M2MClientsResources) CreateM2MClient(ctx context.Context, args CreateM2MClientArgs) (CreateM2MClientResult, error) {
	var result CreateM2MClientResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createM2MClient", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *M2MClientsResources) GetM2MClient(ctx context.Context, args GetM2MClientArgs) (GetM2MClientResult, error) {
	var result GetM2MClientResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getM2MClient", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *M2MClientsResources) ListM2MClients(ctx context.Context, args ListM2MClientsArgs) (ListM2MClientsResult, error) {
	var result ListM2MClientsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listM2MClients", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *M2MClientsResources) ListM2MClientsAll(ctx context.Context, args ListM2MClientsArgs) iter.Seq2[ListM2MClientsItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListM2MClientsItem, error) bool) {
		var zero ListM2MClientsItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListM2MClientsArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listM2MClients", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListM2MClientsItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

func (r *M2MClientsResources) RevokeM2MClient(ctx context.Context, args RevokeM2MClientArgs) (RevokeM2MClientResult, error) {
	var result RevokeM2MClientResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "revokeM2MClient", input)
	if err != nil {
		return result, err
	}
	_ = response
	return result, nil
}

func (r *M2MClientsResources) RotateM2MClientSecret(ctx context.Context, args RotateM2MClientSecretArgs) (RotateM2MClientSecretResult, error) {
	var result RotateM2MClientSecretResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "rotateM2MClientSecret", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *M2MClientsResources) UpdateM2MClient(ctx context.Context, args UpdateM2MClientArgs) (UpdateM2MClientResult, error) {
	var result UpdateM2MClientResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "updateM2MClient", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type OauthResources struct{ client *Client }

func (r *OauthResources) CreateOAuthClient(ctx context.Context, args CreateOAuthClientArgs) (CreateOAuthClientResult, error) {
	var result CreateOAuthClientResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createOAuthClient", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OauthResources) CreateOAuthResource(ctx context.Context, args CreateOAuthResourceArgs) (CreateOAuthResourceResult, error) {
	var result CreateOAuthResourceResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createOAuthResource", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OauthResources) DeleteOAuthClient(ctx context.Context, args DeleteOAuthClientArgs) (DeleteOAuthClientResult, error) {
	var result DeleteOAuthClientResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deleteOAuthClient", input)
	if err != nil {
		return result, err
	}
	_ = response
	return result, nil
}

func (r *OauthResources) DeleteOAuthResource(ctx context.Context, args DeleteOAuthResourceArgs) (DeleteOAuthResourceResult, error) {
	var result DeleteOAuthResourceResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deleteOAuthResource", input)
	if err != nil {
		return result, err
	}
	_ = response
	return result, nil
}

func (r *OauthResources) DeleteOAuthScope(ctx context.Context, args DeleteOAuthScopeArgs) (DeleteOAuthScopeResult, error) {
	var result DeleteOAuthScopeResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deleteOAuthScope", input)
	if err != nil {
		return result, err
	}
	_ = response
	return result, nil
}

func (r *OauthResources) GetOAuthClient(ctx context.Context, args GetOAuthClientArgs) (GetOAuthClientResult, error) {
	var result GetOAuthClientResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getOAuthClient", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OauthResources) GetOAuthProtectedResourceMetadata(ctx context.Context, args GetOAuthProtectedResourceMetadataArgs) (GetOAuthProtectedResourceMetadataResult, error) {
	var result GetOAuthProtectedResourceMetadataResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getOAuthProtectedResourceMetadata", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OauthResources) GetOAuthResource(ctx context.Context, args GetOAuthResourceArgs) (GetOAuthResourceResult, error) {
	var result GetOAuthResourceResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getOAuthResource", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OauthResources) ListOAuthClients(ctx context.Context, args ListOAuthClientsArgs) (ListOAuthClientsResult, error) {
	var result ListOAuthClientsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listOAuthClients", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OauthResources) ListOAuthGrants(ctx context.Context, args ListOAuthGrantsArgs) (ListOAuthGrantsResult, error) {
	var result ListOAuthGrantsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listOAuthGrants", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OauthResources) ListOAuthResources(ctx context.Context, args ListOAuthResourcesArgs) (ListOAuthResourcesResult, error) {
	var result ListOAuthResourcesResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	response, err := r.client.request(ctx, "listOAuthResources", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OauthResources) ListOAuthScopes(ctx context.Context, args ListOAuthScopesArgs) (ListOAuthScopesResult, error) {
	var result ListOAuthScopesResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	response, err := r.client.request(ctx, "listOAuthScopes", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OauthResources) PutOAuthScope(ctx context.Context, args PutOAuthScopeArgs) (PutOAuthScopeResult, error) {
	var result PutOAuthScopeResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "putOAuthScope", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OauthResources) RevokeOAuthGrant(ctx context.Context, args RevokeOAuthGrantArgs) (RevokeOAuthGrantResult, error) {
	var result RevokeOAuthGrantResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "revokeOAuthGrant", input)
	if err != nil {
		return result, err
	}
	_ = response
	return result, nil
}

func (r *OauthResources) RotateOAuthClientSecret(ctx context.Context, args RotateOAuthClientSecretArgs) (RotateOAuthClientSecretResult, error) {
	var result RotateOAuthClientSecretResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "rotateOAuthClientSecret", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OauthResources) UpdateOAuthClient(ctx context.Context, args UpdateOAuthClientArgs) (UpdateOAuthClientResult, error) {
	var result UpdateOAuthClientResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "updateOAuthClient", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OauthResources) UpdateOAuthResource(ctx context.Context, args UpdateOAuthResourceArgs) (UpdateOAuthResourceResult, error) {
	var result UpdateOAuthResourceResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "updateOAuthResource", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type OrganizationExportsResources struct{ client *Client }

func (r *OrganizationExportsResources) CreateOrganizationExport(ctx context.Context, args CreateOrganizationExportArgs) (CreateOrganizationExportResult, error) {
	var result CreateOrganizationExportResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createOrganizationExport", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OrganizationExportsResources) GetOrganizationExport(ctx context.Context, args GetOrganizationExportArgs) (GetOrganizationExportResult, error) {
	var result GetOrganizationExportResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getOrganizationExport", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type OrganizationsResources struct{ client *Client }

func (r *OrganizationsResources) CreateMembership(ctx context.Context, args CreateMembershipArgs) (CreateMembershipResult, error) {
	var result CreateMembershipResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createMembership", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OrganizationsResources) CreateOrganization(ctx context.Context, args CreateOrganizationArgs) (CreateOrganizationResult, error) {
	var result CreateOrganizationResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createOrganization", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OrganizationsResources) CreateOrganizationAPIKey(ctx context.Context, args CreateOrganizationAPIKeyArgs) (CreateOrganizationAPIKeyResult, error) {
	var result CreateOrganizationAPIKeyResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createOrganizationApiKey", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OrganizationsResources) CreateOrganizationAPIKeySecret(ctx context.Context, args CreateOrganizationAPIKeySecretArgs) (CreateOrganizationAPIKeySecretResult, error) {
	var result CreateOrganizationAPIKeySecretResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createOrganizationApiKeySecret", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OrganizationsResources) CreateRoleAssignment(ctx context.Context, args CreateRoleAssignmentArgs) (CreateRoleAssignmentResult, error) {
	var result CreateRoleAssignmentResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createRoleAssignment", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OrganizationsResources) DeleteMembership(ctx context.Context, args DeleteMembershipArgs) (DeleteMembershipResult, error) {
	var result DeleteMembershipResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deleteMembership", input)
	if err != nil {
		return result, err
	}
	_ = response
	return result, nil
}

func (r *OrganizationsResources) DeleteOrganization(ctx context.Context, args DeleteOrganizationArgs) (DeleteOrganizationResult, error) {
	var result DeleteOrganizationResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deleteOrganization", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OrganizationsResources) DeleteOrganizationUserManagementSettings(ctx context.Context, args DeleteOrganizationUserManagementSettingsArgs) (DeleteOrganizationUserManagementSettingsResult, error) {
	var result DeleteOrganizationUserManagementSettingsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deleteOrganizationUserManagementSettings", input)
	if err != nil {
		return result, err
	}
	_ = response
	return result, nil
}

func (r *OrganizationsResources) DeleteRoleAssignment(ctx context.Context, args DeleteRoleAssignmentArgs) (DeleteRoleAssignmentResult, error) {
	var result DeleteRoleAssignmentResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deleteRoleAssignment", input)
	if err != nil {
		return result, err
	}
	_ = response
	return result, nil
}

func (r *OrganizationsResources) GetOrganization(ctx context.Context, args GetOrganizationArgs) (GetOrganizationResult, error) {
	var result GetOrganizationResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "getOrganization", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OrganizationsResources) GetOrganizationAllowedIntents(ctx context.Context, args GetOrganizationAllowedIntentsArgs) (GetOrganizationAllowedIntentsResult, error) {
	var result GetOrganizationAllowedIntentsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getOrganizationAllowedIntents", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OrganizationsResources) GetOrganizationAPIKey(ctx context.Context, args GetOrganizationAPIKeyArgs) (GetOrganizationAPIKeyResult, error) {
	var result GetOrganizationAPIKeyResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getOrganizationApiKey", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OrganizationsResources) GetOrganizationOAuthSettings(ctx context.Context, args GetOrganizationOAuthSettingsArgs) (GetOrganizationOAuthSettingsResult, error) {
	var result GetOrganizationOAuthSettingsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getOrganizationOAuthSettings", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OrganizationsResources) GetOrganizationProfile(ctx context.Context, args GetOrganizationProfileArgs) (GetOrganizationProfileResult, error) {
	var result GetOrganizationProfileResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getOrganizationProfile", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OrganizationsResources) GetOrganizationStatus(ctx context.Context, args GetOrganizationStatusArgs) (GetOrganizationStatusResult, error) {
	var result GetOrganizationStatusResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getOrganizationStatus", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OrganizationsResources) GetOrganizationUserManagementSettings(ctx context.Context, args GetOrganizationUserManagementSettingsArgs) (GetOrganizationUserManagementSettingsResult, error) {
	var result GetOrganizationUserManagementSettingsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getOrganizationUserManagementSettings", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OrganizationsResources) ListMembershipPermissions(ctx context.Context, args ListMembershipPermissionsArgs) (ListMembershipPermissionsResult, error) {
	var result ListMembershipPermissionsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listMembershipPermissions", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OrganizationsResources) ListMemberships(ctx context.Context, args ListMembershipsArgs) (ListMembershipsResult, error) {
	var result ListMembershipsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listMemberships", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OrganizationsResources) ListMembershipsAll(ctx context.Context, args ListMembershipsArgs) iter.Seq2[ListMembershipsItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListMembershipsItem, error) bool) {
		var zero ListMembershipsItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListMembershipsArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		input.Path, err = json.Marshal(args.Path)
		if err != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listMemberships", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListMembershipsItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

func (r *OrganizationsResources) ListOrganizationAPIKeySecrets(ctx context.Context, args ListOrganizationAPIKeySecretsArgs) (ListOrganizationAPIKeySecretsResult, error) {
	var result ListOrganizationAPIKeySecretsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "listOrganizationApiKeySecrets", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OrganizationsResources) ListOrganizationAPIKeys(ctx context.Context, args ListOrganizationAPIKeysArgs) (ListOrganizationAPIKeysResult, error) {
	var result ListOrganizationAPIKeysResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listOrganizationApiKeys", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OrganizationsResources) ListOrganizationAPIKeysAll(ctx context.Context, args ListOrganizationAPIKeysArgs) iter.Seq2[ListOrganizationAPIKeysItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListOrganizationAPIKeysItem, error) bool) {
		var zero ListOrganizationAPIKeysItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListOrganizationAPIKeysArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		input.Path, err = json.Marshal(args.Path)
		if err != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listOrganizationApiKeys", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListOrganizationAPIKeysItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

func (r *OrganizationsResources) ListOrganizationEvents(ctx context.Context, args ListOrganizationEventsArgs) (ListOrganizationEventsResult, error) {
	var result ListOrganizationEventsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listOrganizationEvents", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OrganizationsResources) ListOrganizationEventsAll(ctx context.Context, args ListOrganizationEventsArgs) iter.Seq2[ListOrganizationEventsItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListOrganizationEventsItem, error) bool) {
		var zero ListOrganizationEventsItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListOrganizationEventsArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		input.Path, err = json.Marshal(args.Path)
		if err != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listOrganizationEvents", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListOrganizationEventsItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

func (r *OrganizationsResources) ListOrganizations(ctx context.Context, args ListOrganizationsArgs) (ListOrganizationsResult, error) {
	var result ListOrganizationsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listOrganizations", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OrganizationsResources) ListOrganizationsAll(ctx context.Context, args ListOrganizationsArgs) iter.Seq2[ListOrganizationsItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListOrganizationsItem, error) bool) {
		var zero ListOrganizationsItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListOrganizationsArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listOrganizations", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListOrganizationsItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

func (r *OrganizationsResources) ListRoleAssignments(ctx context.Context, args ListRoleAssignmentsArgs) (ListRoleAssignmentsResult, error) {
	var result ListRoleAssignmentsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "listRoleAssignments", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OrganizationsResources) ListRoleAssignmentsAll(ctx context.Context, args ListRoleAssignmentsArgs) iter.Seq2[ListRoleAssignmentsItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListRoleAssignmentsItem, error) bool) {
		var zero ListRoleAssignmentsItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListRoleAssignmentsArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		input.Path, err = json.Marshal(args.Path)
		if err != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		for item, err := range r.client.paginate(ctx, "listRoleAssignments", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListRoleAssignmentsItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

func (r *OrganizationsResources) RevokeOrganizationAPIKey(ctx context.Context, args RevokeOrganizationAPIKeyArgs) (RevokeOrganizationAPIKeyResult, error) {
	var result RevokeOrganizationAPIKeyResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "revokeOrganizationApiKey", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OrganizationsResources) RevokeOrganizationAPIKeySecret(ctx context.Context, args RevokeOrganizationAPIKeySecretArgs) (RevokeOrganizationAPIKeySecretResult, error) {
	var result RevokeOrganizationAPIKeySecretResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "revokeOrganizationApiKeySecret", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OrganizationsResources) RevokeOrganizationOAuthGrants(ctx context.Context, args RevokeOrganizationOAuthGrantsArgs) (RevokeOrganizationOAuthGrantsResult, error) {
	var result RevokeOrganizationOAuthGrantsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "revokeOrganizationOAuthGrants", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OrganizationsResources) RollOrganizationAPIKey(ctx context.Context, args RollOrganizationAPIKeyArgs) (RollOrganizationAPIKeyResult, error) {
	var result RollOrganizationAPIKeyResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "rollOrganizationApiKey", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OrganizationsResources) SetOrganizationAllowedIntents(ctx context.Context, args SetOrganizationAllowedIntentsArgs) (SetOrganizationAllowedIntentsResult, error) {
	var result SetOrganizationAllowedIntentsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "setOrganizationAllowedIntents", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OrganizationsResources) UpdateMembership(ctx context.Context, args UpdateMembershipArgs) (UpdateMembershipResult, error) {
	var result UpdateMembershipResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "updateMembership", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OrganizationsResources) UpdateOrganization(ctx context.Context, args UpdateOrganizationArgs) (UpdateOrganizationResult, error) {
	var result UpdateOrganizationResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "updateOrganization", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OrganizationsResources) UpdateOrganizationOAuthSettings(ctx context.Context, args UpdateOrganizationOAuthSettingsArgs) (UpdateOrganizationOAuthSettingsResult, error) {
	var result UpdateOrganizationOAuthSettingsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "updateOrganizationOAuthSettings", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OrganizationsResources) UpdateOrganizationProfile(ctx context.Context, args UpdateOrganizationProfileArgs) (UpdateOrganizationProfileResult, error) {
	var result UpdateOrganizationProfileResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "updateOrganizationProfile", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *OrganizationsResources) UpdateOrganizationUserManagementSettings(ctx context.Context, args UpdateOrganizationUserManagementSettingsArgs) (UpdateOrganizationUserManagementSettingsResult, error) {
	var result UpdateOrganizationUserManagementSettingsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "updateOrganizationUserManagementSettings", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type PassportResources struct{ client *Client }

func (r *PassportResources) ListPassportLinks(ctx context.Context, args ListPassportLinksArgs) (ListPassportLinksResult, error) {
	var result ListPassportLinksResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listPassportLinks", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *PassportResources) ListPassportLinksAll(ctx context.Context, args ListPassportLinksArgs) iter.Seq2[ListPassportLinksItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListPassportLinksItem, error) bool) {
		var zero ListPassportLinksItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListPassportLinksArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listPassportLinks", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListPassportLinksItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

func (r *PassportResources) RevokePassportLink(ctx context.Context, args RevokePassportLinkArgs) (RevokePassportLinkResult, error) {
	var result RevokePassportLinkResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "revokePassportLink", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type PipesResources struct{ client *Client }

func (r *PipesResources) CreatePipeConnection(ctx context.Context, args CreatePipeConnectionArgs) (CreatePipeConnectionResult, error) {
	var result CreatePipeConnectionResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createPipeConnection", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *PipesResources) DeletePipeProvider(ctx context.Context, args DeletePipeProviderArgs) (DeletePipeProviderResult, error) {
	var result DeletePipeProviderResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deletePipeProvider", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *PipesResources) DisconnectPipeConnection(ctx context.Context, args DisconnectPipeConnectionArgs) (DisconnectPipeConnectionResult, error) {
	var result DisconnectPipeConnectionResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "disconnectPipeConnection", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *PipesResources) GetPipeAccessToken(ctx context.Context, args GetPipeAccessTokenArgs) (GetPipeAccessTokenResult, error) {
	var result GetPipeAccessTokenResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getPipeAccessToken", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *PipesResources) GetPipeConnection(ctx context.Context, args GetPipeConnectionArgs) (GetPipeConnectionResult, error) {
	var result GetPipeConnectionResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getPipeConnection", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *PipesResources) ListPipeConnections(ctx context.Context, args ListPipeConnectionsArgs) (ListPipeConnectionsResult, error) {
	var result ListPipeConnectionsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listPipeConnections", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *PipesResources) ListPipeConnectionsAll(ctx context.Context, args ListPipeConnectionsArgs) iter.Seq2[ListPipeConnectionsItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListPipeConnectionsItem, error) bool) {
		var zero ListPipeConnectionsItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListPipeConnectionsArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listPipeConnections", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListPipeConnectionsItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

func (r *PipesResources) ListPipeProviders(ctx context.Context, args ListPipeProvidersArgs) (ListPipeProvidersResult, error) {
	var result ListPipeProvidersResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	response, err := r.client.request(ctx, "listPipeProviders", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *PipesResources) PutPipeProvider(ctx context.Context, args PutPipeProviderArgs) (PutPipeProviderResult, error) {
	var result PutPipeProviderResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "putPipeProvider", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type PoliciesResources struct{ client *Client }

func (r *PoliciesResources) CreatePolicy(ctx context.Context, args CreatePolicyArgs) (CreatePolicyResult, error) {
	var result CreatePolicyResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createPolicy", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *PoliciesResources) DeletePolicy(ctx context.Context, args DeletePolicyArgs) (DeletePolicyResult, error) {
	var result DeletePolicyResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deletePolicy", input)
	if err != nil {
		return result, err
	}
	_ = response
	return result, nil
}

func (r *PoliciesResources) GetPolicy(ctx context.Context, args GetPolicyArgs) (GetPolicyResult, error) {
	var result GetPolicyResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getPolicy", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *PoliciesResources) ListPolicies(ctx context.Context, args ListPoliciesArgs) (ListPoliciesResult, error) {
	var result ListPoliciesResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	response, err := r.client.request(ctx, "listPolicies", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *PoliciesResources) UpdatePolicy(ctx context.Context, args UpdatePolicyArgs) (UpdatePolicyResult, error) {
	var result UpdatePolicyResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "updatePolicy", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type PortalProfilesResources struct{ client *Client }

func (r *PortalProfilesResources) CreatePortalProfile(ctx context.Context, args CreatePortalProfileArgs) (CreatePortalProfileResult, error) {
	var result CreatePortalProfileResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createPortalProfile", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *PortalProfilesResources) DeletePortalProfile(ctx context.Context, args DeletePortalProfileArgs) (DeletePortalProfileResult, error) {
	var result DeletePortalProfileResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deletePortalProfile", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *PortalProfilesResources) GetPortalProfile(ctx context.Context, args GetPortalProfileArgs) (GetPortalProfileResult, error) {
	var result GetPortalProfileResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getPortalProfile", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *PortalProfilesResources) ListPortalProfiles(ctx context.Context, args ListPortalProfilesArgs) (ListPortalProfilesResult, error) {
	var result ListPortalProfilesResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	response, err := r.client.request(ctx, "listPortalProfiles", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *PortalProfilesResources) ReplacePortalProfile(ctx context.Context, args ReplacePortalProfileArgs) (ReplacePortalProfileResult, error) {
	var result ReplacePortalProfileResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "replacePortalProfile", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type PortalResources struct{ client *Client }

func (r *PortalResources) CreatePortalInvite(ctx context.Context, args CreatePortalInviteArgs) (CreatePortalInviteResult, error) {
	var result CreatePortalInviteResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createPortalInvite", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *PortalResources) CreatePortalLink(ctx context.Context, args CreatePortalLinkArgs) (CreatePortalLinkResult, error) {
	var result CreatePortalLinkResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createPortalLink", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *PortalResources) ListPortalInvites(ctx context.Context, args ListPortalInvitesArgs) (ListPortalInvitesResult, error) {
	var result ListPortalInvitesResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listPortalInvites", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *PortalResources) ListPortalInvitesAll(ctx context.Context, args ListPortalInvitesArgs) iter.Seq2[ListPortalInvitesItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListPortalInvitesItem, error) bool) {
		var zero ListPortalInvitesItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListPortalInvitesArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listPortalInvites", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListPortalInvitesItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

func (r *PortalResources) ListPortalLinks(ctx context.Context, args ListPortalLinksArgs) (ListPortalLinksResult, error) {
	var result ListPortalLinksResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listPortalLinks", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *PortalResources) ListPortalLinksAll(ctx context.Context, args ListPortalLinksArgs) iter.Seq2[ListPortalLinksItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListPortalLinksItem, error) bool) {
		var zero ListPortalLinksItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListPortalLinksArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listPortalLinks", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListPortalLinksItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

func (r *PortalResources) ResendPortalInvite(ctx context.Context, args ResendPortalInviteArgs) (ResendPortalInviteResult, error) {
	var result ResendPortalInviteResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "resendPortalInvite", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *PortalResources) RevokePortalInvite(ctx context.Context, args RevokePortalInviteArgs) (RevokePortalInviteResult, error) {
	var result RevokePortalInviteResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "revokePortalInvite", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *PortalResources) RevokePortalLink(ctx context.Context, args RevokePortalLinkArgs) (RevokePortalLinkResult, error) {
	var result RevokePortalLinkResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "revokePortalLink", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type ProvidersResources struct{ client *Client }

func (r *ProvidersResources) GetProviderCatalogue(ctx context.Context, args GetProviderCatalogueArgs) (GetProviderCatalogueResult, error) {
	var result GetProviderCatalogueResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	response, err := r.client.request(ctx, "getProviderCatalogue", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type RedirectUrisResources struct{ client *Client }

func (r *RedirectUrisResources) CreateRedirectURI(ctx context.Context, args CreateRedirectURIArgs) (CreateRedirectURIResult, error) {
	var result CreateRedirectURIResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createRedirectUri", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *RedirectUrisResources) DeleteRedirectURI(ctx context.Context, args DeleteRedirectURIArgs) (DeleteRedirectURIResult, error) {
	var result DeleteRedirectURIResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deleteRedirectUri", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *RedirectUrisResources) ListRedirectUris(ctx context.Context, args ListRedirectUrisArgs) (ListRedirectUrisResult, error) {
	var result ListRedirectUrisResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listRedirectUris", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *RedirectUrisResources) ListRedirectUrisAll(ctx context.Context, args ListRedirectUrisArgs) iter.Seq2[ListRedirectUrisItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListRedirectUrisItem, error) bool) {
		var zero ListRedirectUrisItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListRedirectUrisArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listRedirectUris", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListRedirectUrisItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

type RequestsResources struct{ client *Client }

func (r *RequestsResources) ListRequests(ctx context.Context, args ListRequestsArgs) (ListRequestsResult, error) {
	var result ListRequestsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listRequests", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *RequestsResources) ListRequestsAll(ctx context.Context, args ListRequestsArgs) iter.Seq2[ListRequestsItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListRequestsItem, error) bool) {
		var zero ListRequestsItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListRequestsArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listRequests", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListRequestsItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

type RoleRulesResources struct{ client *Client }

func (r *RoleRulesResources) CreateRoleRule(ctx context.Context, args CreateRoleRuleArgs) (CreateRoleRuleResult, error) {
	var result CreateRoleRuleResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createRoleRule", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *RoleRulesResources) DeleteRoleRule(ctx context.Context, args DeleteRoleRuleArgs) (DeleteRoleRuleResult, error) {
	var result DeleteRoleRuleResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deleteRoleRule", input)
	if err != nil {
		return result, err
	}
	_ = response
	return result, nil
}

func (r *RoleRulesResources) GetRoleRule(ctx context.Context, args GetRoleRuleArgs) (GetRoleRuleResult, error) {
	var result GetRoleRuleResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getRoleRule", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *RoleRulesResources) ListRoleRules(ctx context.Context, args ListRoleRulesArgs) (ListRoleRulesResult, error) {
	var result ListRoleRulesResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	response, err := r.client.request(ctx, "listRoleRules", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *RoleRulesResources) PreviewRoleRules(ctx context.Context, args PreviewRoleRulesArgs) (PreviewRoleRulesResult, error) {
	var result PreviewRoleRulesResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "previewRoleRules", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *RoleRulesResources) UpdateRoleRule(ctx context.Context, args UpdateRoleRuleArgs) (UpdateRoleRuleResult, error) {
	var result UpdateRoleRuleResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "updateRoleRule", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type SSOResources struct{ client *Client }

func (r *SSOResources) ExchangeSSOCode(ctx context.Context, args ExchangeSSOCodeArgs) (ExchangeSSOCodeResult, error) {
	var result ExchangeSSOCodeResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "exchangeSsoCode", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *SSOResources) ExchangeTrustedToken(ctx context.Context, args ExchangeTrustedTokenArgs) (ExchangeTrustedTokenResult, error) {
	var result ExchangeTrustedTokenResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "exchangeTrustedToken", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *SSOResources) ForwardSSOResponse(ctx context.Context, args ForwardSSOResponseArgs) (ForwardSSOResponseResult, error) {
	var result ForwardSSOResponseResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "forwardSsoResponse", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *SSOResources) GetSSOProfile(ctx context.Context, args GetSSOProfileArgs) (GetSSOProfileResult, error) {
	var result GetSSOProfileResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.AccessToken = args.AccessToken
	response, err := r.client.request(ctx, "getSsoProfile", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *SSOResources) StartSSOLogout(ctx context.Context, args StartSSOLogoutArgs) (StartSSOLogoutResult, error) {
	var result StartSSOLogoutResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "startSsoLogout", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type SessionsResources struct{ client *Client }

func (r *SessionsResources) MigrateSession(ctx context.Context, args MigrateSessionArgs) (MigrateSessionResult, error) {
	var result MigrateSessionResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "migrateSession", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type SignInHooksResources struct{ client *Client }

func (r *SignInHooksResources) CreateSignInHook(ctx context.Context, args CreateSignInHookArgs) (CreateSignInHookResult, error) {
	var result CreateSignInHookResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createSignInHook", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *SignInHooksResources) DeleteSignInHook(ctx context.Context, args DeleteSignInHookArgs) (DeleteSignInHookResult, error) {
	var result DeleteSignInHookResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deleteSignInHook", input)
	if err != nil {
		return result, err
	}
	_ = response
	return result, nil
}

func (r *SignInHooksResources) GetSignInHook(ctx context.Context, args GetSignInHookArgs) (GetSignInHookResult, error) {
	var result GetSignInHookResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getSignInHook", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *SignInHooksResources) ListSignInHookDeliveries(ctx context.Context, args ListSignInHookDeliveriesArgs) (ListSignInHookDeliveriesResult, error) {
	var result ListSignInHookDeliveriesResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listSignInHookDeliveries", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *SignInHooksResources) ListSignInHookDeliveriesAll(ctx context.Context, args ListSignInHookDeliveriesArgs) iter.Seq2[ListSignInHookDeliveriesItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListSignInHookDeliveriesItem, error) bool) {
		var zero ListSignInHookDeliveriesItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListSignInHookDeliveriesArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		input.Path, err = json.Marshal(args.Path)
		if err != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listSignInHookDeliveries", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListSignInHookDeliveriesItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

func (r *SignInHooksResources) ListSignInHooks(ctx context.Context, args ListSignInHooksArgs) (ListSignInHooksResult, error) {
	var result ListSignInHooksResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	response, err := r.client.request(ctx, "listSignInHooks", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *SignInHooksResources) RotateSignInHookSecret(ctx context.Context, args RotateSignInHookSecretArgs) (RotateSignInHookSecretResult, error) {
	var result RotateSignInHookSecretResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "rotateSignInHookSecret", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *SignInHooksResources) TestSignInHook(ctx context.Context, args TestSignInHookArgs) (TestSignInHookResult, error) {
	var result TestSignInHookResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "testSignInHook", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *SignInHooksResources) UpdateSignInHook(ctx context.Context, args UpdateSignInHookArgs) (UpdateSignInHookResult, error) {
	var result UpdateSignInHookResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "updateSignInHook", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type SignInRiskPolicyResources struct{ client *Client }

func (r *SignInRiskPolicyResources) DeleteSignInRiskPolicy(ctx context.Context, args DeleteSignInRiskPolicyArgs) (DeleteSignInRiskPolicyResult, error) {
	var result DeleteSignInRiskPolicyResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	response, err := r.client.request(ctx, "deleteSignInRiskPolicy", input)
	if err != nil {
		return result, err
	}
	_ = response
	return result, nil
}

func (r *SignInRiskPolicyResources) GetSignInRiskPolicy(ctx context.Context, args GetSignInRiskPolicyArgs) (GetSignInRiskPolicyResult, error) {
	var result GetSignInRiskPolicyResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	response, err := r.client.request(ctx, "getSignInRiskPolicy", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *SignInRiskPolicyResources) PutSignInRiskPolicy(ctx context.Context, args PutSignInRiskPolicyArgs) (PutSignInRiskPolicyResult, error) {
	var result PutSignInRiskPolicyResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "putSignInRiskPolicy", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type SignInsResources struct{ client *Client }

func (r *SignInsResources) GetSignInExplanation(ctx context.Context, args GetSignInExplanationArgs) (GetSignInExplanationResult, error) {
	var result GetSignInExplanationResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getSignInExplanation", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *SignInsResources) GetSignInTimeline(ctx context.Context, args GetSignInTimelineArgs) (GetSignInTimelineResult, error) {
	var result GetSignInTimelineResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getSignInTimeline", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *SignInsResources) ListSignIns(ctx context.Context, args ListSignInsArgs) (ListSignInsResult, error) {
	var result ListSignInsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listSignIns", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *SignInsResources) ListSignInsAll(ctx context.Context, args ListSignInsArgs) iter.Seq2[ListSignInsItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListSignInsItem, error) bool) {
		var zero ListSignInsItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListSignInsArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listSignIns", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListSignInsItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

type SsfStreamsResources struct{ client *Client }

func (r *SsfStreamsResources) CreateSsfStream(ctx context.Context, args CreateSsfStreamArgs) (CreateSsfStreamResult, error) {
	var result CreateSsfStreamResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createSsfStream", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *SsfStreamsResources) DeleteSsfStream(ctx context.Context, args DeleteSsfStreamArgs) (DeleteSsfStreamResult, error) {
	var result DeleteSsfStreamResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deleteSsfStream", input)
	if err != nil {
		return result, err
	}
	_ = response
	return result, nil
}

func (r *SsfStreamsResources) GetSsfStream(ctx context.Context, args GetSsfStreamArgs) (GetSsfStreamResult, error) {
	var result GetSsfStreamResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getSsfStream", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *SsfStreamsResources) ListSsfStreams(ctx context.Context, args ListSsfStreamsArgs) (ListSsfStreamsResult, error) {
	var result ListSsfStreamsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	response, err := r.client.request(ctx, "listSsfStreams", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *SsfStreamsResources) UpdateSsfStream(ctx context.Context, args UpdateSsfStreamArgs) (UpdateSsfStreamResult, error) {
	var result UpdateSsfStreamResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "updateSsfStream", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *SsfStreamsResources) VerifySsfStream(ctx context.Context, args VerifySsfStreamArgs) (VerifySsfStreamResult, error) {
	var result VerifySsfStreamResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "verifySsfStream", input)
	if err != nil {
		return result, err
	}
	_ = response
	return result, nil
}

type TrustedIssuersResources struct{ client *Client }

func (r *TrustedIssuersResources) CreateTrustedIssuer(ctx context.Context, args CreateTrustedIssuerArgs) (CreateTrustedIssuerResult, error) {
	var result CreateTrustedIssuerResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createTrustedIssuer", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *TrustedIssuersResources) DeleteTrustedIssuer(ctx context.Context, args DeleteTrustedIssuerArgs) (DeleteTrustedIssuerResult, error) {
	var result DeleteTrustedIssuerResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deleteTrustedIssuer", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *TrustedIssuersResources) ListTrustedIssuers(ctx context.Context, args ListTrustedIssuersArgs) (ListTrustedIssuersResult, error) {
	var result ListTrustedIssuersResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	response, err := r.client.request(ctx, "listTrustedIssuers", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type UploadsResources struct{ client *Client }

func (r *UploadsResources) CreateUpload(ctx context.Context, args CreateUploadArgs) (CreateUploadResult, error) {
	var result CreateUploadResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createUpload", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type UserManagementResources struct{ client *Client }

func (r *UserManagementResources) AuthenticateEndUserPasskey(ctx context.Context, args AuthenticateEndUserPasskeyArgs) (AuthenticateEndUserPasskeyResult, error) {
	var result AuthenticateEndUserPasskeyResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "authenticateEndUserPasskey", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *UserManagementResources) DeleteSocialProvider(ctx context.Context, args DeleteSocialProviderArgs) (DeleteSocialProviderResult, error) {
	var result DeleteSocialProviderResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deleteSocialProvider", input)
	if err != nil {
		return result, err
	}
	_ = response
	return result, nil
}

func (r *UserManagementResources) GetBotProtection(ctx context.Context, args GetBotProtectionArgs) (GetBotProtectionResult, error) {
	var result GetBotProtectionResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	response, err := r.client.request(ctx, "getBotProtection", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *UserManagementResources) GetUserManagementSettings(ctx context.Context, args GetUserManagementSettingsArgs) (GetUserManagementSettingsResult, error) {
	var result GetUserManagementSettingsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	response, err := r.client.request(ctx, "getUserManagementSettings", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *UserManagementResources) GetUserManagementUsage(ctx context.Context, args GetUserManagementUsageArgs) (GetUserManagementUsageResult, error) {
	var result GetUserManagementUsageResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "getUserManagementUsage", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *UserManagementResources) ListBotDetections(ctx context.Context, args ListBotDetectionsArgs) (ListBotDetectionsResult, error) {
	var result ListBotDetectionsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listBotDetections", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *UserManagementResources) ListBotDetectionsAll(ctx context.Context, args ListBotDetectionsArgs) iter.Seq2[ListBotDetectionsItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListBotDetectionsItem, error) bool) {
		var zero ListBotDetectionsItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListBotDetectionsArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listBotDetections", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListBotDetectionsItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

func (r *UserManagementResources) ListSocialProviders(ctx context.Context, args ListSocialProvidersArgs) (ListSocialProvidersResult, error) {
	var result ListSocialProvidersResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	response, err := r.client.request(ctx, "listSocialProviders", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *UserManagementResources) StartEndUserPasskeyAuthentication(ctx context.Context, args StartEndUserPasskeyAuthenticationArgs) (StartEndUserPasskeyAuthenticationResult, error) {
	var result StartEndUserPasskeyAuthenticationResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	response, err := r.client.request(ctx, "startEndUserPasskeyAuthentication", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *UserManagementResources) UpdateBotProtection(ctx context.Context, args UpdateBotProtectionArgs) (UpdateBotProtectionResult, error) {
	var result UpdateBotProtectionResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "updateBotProtection", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *UserManagementResources) UpdateSocialProvider(ctx context.Context, args UpdateSocialProviderArgs) (UpdateSocialProviderResult, error) {
	var result UpdateSocialProviderResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "updateSocialProvider", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *UserManagementResources) UpdateUserManagementSettings(ctx context.Context, args UpdateUserManagementSettingsArgs) (UpdateUserManagementSettingsResult, error) {
	var result UpdateUserManagementSettingsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "updateUserManagementSettings", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type UsersResources struct{ client *Client }

func (r *UsersResources) ClearEndUserPassword(ctx context.Context, args ClearEndUserPasswordArgs) (ClearEndUserPasswordResult, error) {
	var result ClearEndUserPasswordResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "clearEndUserPassword", input)
	if err != nil {
		return result, err
	}
	_ = response
	return result, nil
}

func (r *UsersResources) CreateEndUser(ctx context.Context, args CreateEndUserArgs) (CreateEndUserResult, error) {
	var result CreateEndUserResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createEndUser", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *UsersResources) DeleteEndUser(ctx context.Context, args DeleteEndUserArgs) (DeleteEndUserResult, error) {
	var result DeleteEndUserResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deleteEndUser", input)
	if err != nil {
		return result, err
	}
	_ = response
	return result, nil
}

func (r *UsersResources) DeleteEndUserFactor(ctx context.Context, args DeleteEndUserFactorArgs) (DeleteEndUserFactorResult, error) {
	var result DeleteEndUserFactorResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deleteEndUserFactor", input)
	if err != nil {
		return result, err
	}
	_ = response
	return result, nil
}

func (r *UsersResources) GetEndUser(ctx context.Context, args GetEndUserArgs) (GetEndUserResult, error) {
	var result GetEndUserResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getEndUser", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *UsersResources) ListEndUserFactors(ctx context.Context, args ListEndUserFactorsArgs) (ListEndUserFactorsResult, error) {
	var result ListEndUserFactorsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "listEndUserFactors", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *UsersResources) ListEndUserIdentities(ctx context.Context, args ListEndUserIdentitiesArgs) (ListEndUserIdentitiesResult, error) {
	var result ListEndUserIdentitiesResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "listEndUserIdentities", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *UsersResources) ListEndUserSessions(ctx context.Context, args ListEndUserSessionsArgs) (ListEndUserSessionsResult, error) {
	var result ListEndUserSessionsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "listEndUserSessions", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *UsersResources) ListEndUsers(ctx context.Context, args ListEndUsersArgs) (ListEndUsersResult, error) {
	var result ListEndUsersResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listEndUsers", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *UsersResources) ListEndUsersAll(ctx context.Context, args ListEndUsersArgs) iter.Seq2[ListEndUsersItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListEndUsersItem, error) bool) {
		var zero ListEndUsersItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListEndUsersArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listEndUsers", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListEndUsersItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

func (r *UsersResources) RegisterEndUserPasskey(ctx context.Context, args RegisterEndUserPasskeyArgs) (RegisterEndUserPasskeyResult, error) {
	var result RegisterEndUserPasskeyResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "registerEndUserPasskey", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *UsersResources) RevokeEndUserSession(ctx context.Context, args RevokeEndUserSessionArgs) (RevokeEndUserSessionResult, error) {
	var result RevokeEndUserSessionResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "revokeEndUserSession", input)
	if err != nil {
		return result, err
	}
	_ = response
	return result, nil
}

func (r *UsersResources) RevokeEndUserSessions(ctx context.Context, args RevokeEndUserSessionsArgs) (RevokeEndUserSessionsResult, error) {
	var result RevokeEndUserSessionsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "revokeEndUserSessions", input)
	if err != nil {
		return result, err
	}
	_ = response
	return result, nil
}

func (r *UsersResources) SetEndUserPassword(ctx context.Context, args SetEndUserPasswordArgs) (SetEndUserPasswordResult, error) {
	var result SetEndUserPasswordResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "setEndUserPassword", input)
	if err != nil {
		return result, err
	}
	_ = response
	return result, nil
}

func (r *UsersResources) StartEndUserPasskeyRegistration(ctx context.Context, args StartEndUserPasskeyRegistrationArgs) (StartEndUserPasskeyRegistrationResult, error) {
	var result StartEndUserPasskeyRegistrationResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "startEndUserPasskeyRegistration", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *UsersResources) UpdateEndUser(ctx context.Context, args UpdateEndUserArgs) (UpdateEndUserResult, error) {
	var result UpdateEndUserResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "updateEndUser", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *UsersResources) UpdateEndUserPasskey(ctx context.Context, args UpdateEndUserPasskeyArgs) (UpdateEndUserPasskeyResult, error) {
	var result UpdateEndUserPasskeyResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "updateEndUserPasskey", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type VaultResources struct{ client *Client }

func (r *VaultResources) CreateVaultKey(ctx context.Context, args CreateVaultKeyArgs) (CreateVaultKeyResult, error) {
	var result CreateVaultKeyResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createVaultKey", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *VaultResources) DeleteVaultObject(ctx context.Context, args DeleteVaultObjectArgs) (DeleteVaultObjectResult, error) {
	var result DeleteVaultObjectResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Query, err = json.Marshal(args.Query)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deleteVaultObject", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *VaultResources) GetVaultObject(ctx context.Context, args GetVaultObjectArgs) (GetVaultObjectResult, error) {
	var result GetVaultObjectResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Query, err = json.Marshal(args.Query)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getVaultObject", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *VaultResources) ListVaultKeys(ctx context.Context, args ListVaultKeysArgs) (ListVaultKeysResult, error) {
	var result ListVaultKeysResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Query, err = json.Marshal(args.Query)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "listVaultKeys", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *VaultResources) ListVaultObjectVersions(ctx context.Context, args ListVaultObjectVersionsArgs) (ListVaultObjectVersionsResult, error) {
	var result ListVaultObjectVersionsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Query, err = json.Marshal(args.Query)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "listVaultObjectVersions", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *VaultResources) PutVaultObject(ctx context.Context, args PutVaultObjectArgs) (PutVaultObjectResult, error) {
	var result PutVaultObjectResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "putVaultObject", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *VaultResources) RotateVaultKey(ctx context.Context, args RotateVaultKeyArgs) (RotateVaultKeyResult, error) {
	var result RotateVaultKeyResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "rotateVaultKey", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *VaultResources) VaultDecrypt(ctx context.Context, args VaultDecryptArgs) (VaultDecryptResult, error) {
	var result VaultDecryptResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "vaultDecrypt", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *VaultResources) VaultEncrypt(ctx context.Context, args VaultEncryptArgs) (VaultEncryptResult, error) {
	var result VaultEncryptResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "vaultEncrypt", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type WebhookDeliveriesResources struct{ client *Client }

func (r *WebhookDeliveriesResources) ResendWebhookDelivery(ctx context.Context, args ResendWebhookDeliveryArgs) (ResendWebhookDeliveryResult, error) {
	var result ResendWebhookDeliveryResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "resendWebhookDelivery", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type WebhookEndpointsResources struct{ client *Client }

func (r *WebhookEndpointsResources) CreateWebhookEndpoint(ctx context.Context, args CreateWebhookEndpointArgs) (CreateWebhookEndpointResult, error) {
	var result CreateWebhookEndpointResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "createWebhookEndpoint", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *WebhookEndpointsResources) DeleteWebhookEndpoint(ctx context.Context, args DeleteWebhookEndpointArgs) (DeleteWebhookEndpointResult, error) {
	var result DeleteWebhookEndpointResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "deleteWebhookEndpoint", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *WebhookEndpointsResources) GetWebhookEndpoint(ctx context.Context, args GetWebhookEndpointArgs) (GetWebhookEndpointResult, error) {
	var result GetWebhookEndpointResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "getWebhookEndpoint", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *WebhookEndpointsResources) ListWebhookDeliveries(ctx context.Context, args ListWebhookDeliveriesArgs) (ListWebhookDeliveriesResult, error) {
	var result ListWebhookDeliveriesResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listWebhookDeliveries", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *WebhookEndpointsResources) ListWebhookDeliveriesAll(ctx context.Context, args ListWebhookDeliveriesArgs) iter.Seq2[ListWebhookDeliveriesItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListWebhookDeliveriesItem, error) bool) {
		var zero ListWebhookDeliveriesItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListWebhookDeliveriesArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		input.Path, err = json.Marshal(args.Path)
		if err != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listWebhookDeliveries", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListWebhookDeliveriesItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

func (r *WebhookEndpointsResources) ListWebhookEndpoints(ctx context.Context, args ListWebhookEndpointsArgs) (ListWebhookEndpointsResult, error) {
	var result ListWebhookEndpointsResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	if args.Query != nil {
		input.Query, err = json.Marshal(args.Query)
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := r.client.request(ctx, "listWebhookEndpoints", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *WebhookEndpointsResources) ListWebhookEndpointsAll(ctx context.Context, args ListWebhookEndpointsArgs) iter.Seq2[ListWebhookEndpointsItem, error] {
	snapshot, snapshotErr := json.Marshal(args)
	return func(yield func(ListWebhookEndpointsItem, error) bool) {
		var zero ListWebhookEndpointsItem
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if snapshotErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var args ListWebhookEndpointsArgs
		if json.Unmarshal(snapshot, &args) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{IdempotencyKey: args.IdempotencyKey}
		var err error
		if args.Query != nil {
			input.Query, err = json.Marshal(args.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listWebhookEndpoints", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var value ListWebhookEndpointsItem
			if json.Unmarshal(data, &value) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(value, nil) {
				return
			}
		}
	}
}

func (r *WebhookEndpointsResources) RotateWebhookSecret(ctx context.Context, args RotateWebhookSecretArgs) (RotateWebhookSecretResult, error) {
	var result RotateWebhookSecretResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "rotateWebhookSecret", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *WebhookEndpointsResources) TestWebhookEndpoint(ctx context.Context, args TestWebhookEndpointArgs) (TestWebhookEndpointResult, error) {
	var result TestWebhookEndpointResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "testWebhookEndpoint", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

func (r *WebhookEndpointsResources) UpdateWebhookEndpoint(ctx context.Context, args UpdateWebhookEndpointArgs) (UpdateWebhookEndpointResult, error) {
	var result UpdateWebhookEndpointResult
	if r == nil || r.client == nil {
		return result, requestFailure("client_closed")
	}
	var err error
	input := requestInput{IdempotencyKey: args.IdempotencyKey}
	input.Path, err = json.Marshal(args.Path)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input.Body, err = json.Marshal(args.Body)
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	response, err := r.client.request(ctx, "updateWebhookEndpoint", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}

type Resources struct {
	AttributeCatalogue   *AttributeCatalogueResources
	AuditLogEvents       *AuditLogEventsResources
	AuditLogSchemas      *AuditLogSchemasResources
	AuditLogs            *AuditLogsResources
	Authorization        *AuthorizationResources
	Branding             *BrandingResources
	ConnectionHealth     *ConnectionHealthResources
	Connections          *ConnectionsResources
	CustomDomains        *CustomDomainsResources
	Dev                  *DevResources
	Directories          *DirectoriesResources
	DirectoryGroups      *DirectoryGroupsResources
	DirectoryUsers       *DirectoryUsersResources
	Domains              *DomainsResources
	EmailEvents          *EmailEventsResources
	EncryptionKeys       *EncryptionKeysResources
	Environment          *EnvironmentResources
	Events               *EventsResources
	Exports              *ExportsResources
	FeatureFlags         *FeatureFlagsResources
	IDPRadar             *IDPRadarResources
	Imports              *ImportsResources
	Invites              *InvitesResources
	ItContacts           *ItContactsResources
	ItEmailConfiguration *ItEmailConfigurationResources
	LogStreams           *LogStreamsResources
	M2MClients           *M2MClientsResources
	Oauth                *OauthResources
	OrganizationExports  *OrganizationExportsResources
	Organizations        *OrganizationsResources
	Passport             *PassportResources
	Pipes                *PipesResources
	Policies             *PoliciesResources
	PortalProfiles       *PortalProfilesResources
	Portal               *PortalResources
	Providers            *ProvidersResources
	RedirectUris         *RedirectUrisResources
	Requests             *RequestsResources
	RoleRules            *RoleRulesResources
	SSO                  *SSOResources
	Sessions             *SessionsResources
	SignInHooks          *SignInHooksResources
	SignInRiskPolicy     *SignInRiskPolicyResources
	SignIns              *SignInsResources
	SsfStreams           *SsfStreamsResources
	TrustedIssuers       *TrustedIssuersResources
	Uploads              *UploadsResources
	UserManagement       *UserManagementResources
	Users                *UsersResources
	Vault                *VaultResources
	WebhookDeliveries    *WebhookDeliveriesResources
	WebhookEndpoints     *WebhookEndpointsResources
}

func makeResources(client *Client) Resources {
	return Resources{AttributeCatalogue: &AttributeCatalogueResources{client: client}, AuditLogEvents: &AuditLogEventsResources{client: client}, AuditLogSchemas: &AuditLogSchemasResources{client: client}, AuditLogs: &AuditLogsResources{client: client}, Authorization: &AuthorizationResources{client: client}, Branding: &BrandingResources{client: client}, ConnectionHealth: &ConnectionHealthResources{client: client}, Connections: &ConnectionsResources{client: client}, CustomDomains: &CustomDomainsResources{client: client}, Dev: &DevResources{client: client}, Directories: &DirectoriesResources{client: client}, DirectoryGroups: &DirectoryGroupsResources{client: client}, DirectoryUsers: &DirectoryUsersResources{client: client}, Domains: &DomainsResources{client: client}, EmailEvents: &EmailEventsResources{client: client}, EncryptionKeys: &EncryptionKeysResources{client: client}, Environment: &EnvironmentResources{client: client}, Events: &EventsResources{client: client}, Exports: &ExportsResources{client: client}, FeatureFlags: &FeatureFlagsResources{client: client}, IDPRadar: &IDPRadarResources{client: client}, Imports: &ImportsResources{client: client}, Invites: &InvitesResources{client: client}, ItContacts: &ItContactsResources{client: client}, ItEmailConfiguration: &ItEmailConfigurationResources{client: client}, LogStreams: &LogStreamsResources{client: client}, M2MClients: &M2MClientsResources{client: client}, Oauth: &OauthResources{client: client}, OrganizationExports: &OrganizationExportsResources{client: client}, Organizations: &OrganizationsResources{client: client}, Passport: &PassportResources{client: client}, Pipes: &PipesResources{client: client}, Policies: &PoliciesResources{client: client}, PortalProfiles: &PortalProfilesResources{client: client}, Portal: &PortalResources{client: client}, Providers: &ProvidersResources{client: client}, RedirectUris: &RedirectUrisResources{client: client}, Requests: &RequestsResources{client: client}, RoleRules: &RoleRulesResources{client: client}, SSO: &SSOResources{client: client}, Sessions: &SessionsResources{client: client}, SignInHooks: &SignInHooksResources{client: client}, SignInRiskPolicy: &SignInRiskPolicyResources{client: client}, SignIns: &SignInsResources{client: client}, SsfStreams: &SsfStreamsResources{client: client}, TrustedIssuers: &TrustedIssuersResources{client: client}, Uploads: &UploadsResources{client: client}, UserManagement: &UserManagementResources{client: client}, Users: &UsersResources{client: client}, Vault: &VaultResources{client: client}, WebhookDeliveries: &WebhookDeliveriesResources{client: client}, WebhookEndpoints: &WebhookEndpointsResources{client: client}}
}
