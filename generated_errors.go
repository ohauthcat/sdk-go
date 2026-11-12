// Code generated from the pinned error catalogue. DO NOT EDIT.
package ohauthcat

const (
	CodeSAMLResponseMalformed                ErrorCode = "saml_response_malformed"
	CodeSAMLIDPStatusError                   ErrorCode = "saml_idp_status_error"
	CodeSAMLDtdForbidden                     ErrorCode = "saml_dtd_forbidden"
	CodeSAMLSignatureMissing                 ErrorCode = "saml_signature_missing"
	CodeSAMLSignatureInvalid                 ErrorCode = "saml_signature_invalid"
	CodeSAMLResponseAltered                  ErrorCode = "saml_response_altered"
	CodeSAMLSignatureAlgorithmWeak           ErrorCode = "saml_signature_algorithm_weak"
	CodeSAMLAssertionCountInvalid            ErrorCode = "saml_assertion_count_invalid"
	CodeSAMLWrappingDetected                 ErrorCode = "saml_wrapping_detected"
	CodeSAMLDestinationMismatch              ErrorCode = "saml_destination_mismatch"
	CodeSAMLAudienceMismatch                 ErrorCode = "saml_audience_mismatch"
	CodeSAMLIssuerMismatch                   ErrorCode = "saml_issuer_mismatch"
	CodeSAMLResponseExpired                  ErrorCode = "saml_response_expired"
	CodeSAMLResponseNotYetValid              ErrorCode = "saml_response_not_yet_valid"
	CodeSAMLInResponseToMismatch             ErrorCode = "saml_in_response_to_mismatch"
	CodeSAMLAssertionReplayed                ErrorCode = "saml_assertion_replayed"
	CodeSAMLIDPInitiatedDisabled             ErrorCode = "saml_idp_initiated_disabled"
	CodeSAMLDecryptionFailed                 ErrorCode = "saml_decryption_failed"
	CodeSAMLNameidMissing                    ErrorCode = "saml_nameid_missing"
	CodeSAMLEmailMissing                     ErrorCode = "saml_email_missing"
	CodeOIDCEmailMissing                     ErrorCode = "oidc_email_missing"
	CodeOIDCDiscoveryUnreachable             ErrorCode = "oidc_discovery_unreachable"
	CodeOIDCTokenExchangeFailed              ErrorCode = "oidc_token_exchange_failed"
	CodeOIDCIDPError                         ErrorCode = "oidc_idp_error"
	CodeOIDCIDTokenInvalid                   ErrorCode = "oidc_id_token_invalid"
	CodeOIDCEmailNotVerified                 ErrorCode = "oidc_email_not_verified"
	CodeEmailOutsideVerifiedDomain           ErrorCode = "email_outside_verified_domain"
	CodeConnectionUntested                   ErrorCode = "connection_untested"
	CodeConnectionInactive                   ErrorCode = "connection_inactive"
	CodeConnectionNotActivated               ErrorCode = "connection_not_activated"
	CodeEnvironmentDisabled                  ErrorCode = "environment_disabled"
	CodeAccountSuspended                     ErrorCode = "account_suspended"
	CodeAccountClosed                        ErrorCode = "account_closed"
	CodeOrganizationNotFound                 ErrorCode = "organization_not_found"
	CodeClientIDInvalid                      ErrorCode = "client_id_invalid"
	CodePKCERequired                         ErrorCode = "pkce_required"
	CodeAuthorizeRequestInvalid              ErrorCode = "authorize_request_invalid"
	CodeSSOForwardingDisabled                ErrorCode = "sso_forwarding_disabled"
	CodeConnectionNotFound                   ErrorCode = "connection_not_found"
	CodeRedirectURINotAllowed                ErrorCode = "redirect_uri_not_allowed"
	CodeInvalidState                         ErrorCode = "invalid_state"
	CodeFlowInterruptedRetry                 ErrorCode = "flow_interrupted_retry"
	CodeInvalidGrant                         ErrorCode = "invalid_grant"
	CodePKCEVerificationFailed               ErrorCode = "pkce_verification_failed"
	CodeConnectionStandbyUnavailable         ErrorCode = "connection_standby_unavailable"
	CodeConsentDeclined                      ErrorCode = "consent_declined"
	CodeSignInBufferFull                     ErrorCode = "sign_in_buffer_full"
	CodeSCIMTokenInvalid                     ErrorCode = "scim_token_invalid"
	CodeSCIMUserOutOfDomain                  ErrorCode = "scim_user_out_of_domain"
	CodeSCIMFilterInvalid                    ErrorCode = "scim_filter_invalid"
	CodeSCIMPayloadTooLarge                  ErrorCode = "scim_payload_too_large"
	CodeSCIMRateLimited                      ErrorCode = "scim_rate_limited"
	CodeSCIMTemporarilyUnavailable           ErrorCode = "scim_temporarily_unavailable"
	CodeSCIMConflict                         ErrorCode = "scim_conflict"
	CodeSCIMPreconditionFailed               ErrorCode = "scim_precondition_failed"
	CodeAuthenticationRequired               ErrorCode = "authentication_required"
	CodeAPIKeyInvalid                        ErrorCode = "api_key_invalid"
	CodeAccessTokenInvalid                   ErrorCode = "access_token_invalid"
	CodeCliTokenInvalid                      ErrorCode = "cli_token_invalid"
	CodeCsrfTokenInvalid                     ErrorCode = "csrf_token_invalid"
	CodeOtpInvalid                           ErrorCode = "otp_invalid"
	CodeRecoveryCodeInvalid                  ErrorCode = "recovery_code_invalid"
	CodeInviteInvalid                        ErrorCode = "invite_invalid"
	CodeWebauthnInvalid                      ErrorCode = "webauthn_invalid"
	CodeTotpInvalid                          ErrorCode = "totp_invalid"
	CodeOauthSignInFailed                    ErrorCode = "oauth_sign_in_failed"
	CodeMfaRequired                          ErrorCode = "mfa_required"
	CodeAuthorizationPending                 ErrorCode = "authorization_pending"
	CodeSlowDown                             ErrorCode = "slow_down"
	CodeExpiredToken                         ErrorCode = "expired_token"
	CodeAccessDenied                         ErrorCode = "access_denied"
	CodeCliRequestStale                      ErrorCode = "cli_request_stale"
	CodeStepUpRequired                       ErrorCode = "step_up_required"
	CodeSignUpPaused                         ErrorCode = "sign_up_paused"
	CodeSandboxCreationUnavailable           ErrorCode = "sandbox_creation_unavailable"
	CodeSandboxHandoffExpired                ErrorCode = "sandbox_handoff_expired"
	CodeDebugBundlesUnavailable              ErrorCode = "debug_bundles_unavailable"
	CodeEmailDisposable                      ErrorCode = "email_disposable"
	CodeSSORequired                          ErrorCode = "sso_required"
	CodeRoleInsufficient                     ErrorCode = "role_insufficient"
	CodeLastOwner                            ErrorCode = "last_owner"
	CodeLastBreakGlassCredential             ErrorCode = "last_break_glass_credential"
	CodeEnvironmentResetNotAllowed           ErrorCode = "environment_reset_not_allowed"
	CodeSandboxRequired                      ErrorCode = "sandbox_required"
	CodeSSONotReady                          ErrorCode = "sso_not_ready"
	CodeBreakGlassLimit                      ErrorCode = "break_glass_limit"
	CodeAPIKeyScopeMissing                   ErrorCode = "api_key_scope_missing"
	CodeStaffRoleInsufficient                ErrorCode = "staff_role_insufficient"
	CodeBufferEntryConflict                  ErrorCode = "buffer_entry_conflict"
	CodeStaffMemberDisabled                  ErrorCode = "staff_member_disabled"
	CodeStaffLastAdmin                       ErrorCode = "staff_last_admin"
	CodeStaffOriginRefused                   ErrorCode = "staff_origin_refused"
	CodeStaffUnsupportedMediaType            ErrorCode = "staff_unsupported_media_type"
	CodeStaffRoleImmutable                   ErrorCode = "staff_role_immutable"
	CodeStaffRoleHeld                        ErrorCode = "staff_role_held"
	CodeStaffRoleInUse                       ErrorCode = "staff_role_in_use"
	CodeStaffAccountProtected                ErrorCode = "staff_account_protected"
	CodeStaffSelfChange                      ErrorCode = "staff_self_change"
	CodeStaffGrantExceedsOwn                 ErrorCode = "staff_grant_exceeds_own"
	CodeStaffSessionRequired                 ErrorCode = "staff_session_required"
	CodeStaffSessionDeviceChanged            ErrorCode = "staff_session_device_changed"
	CodeStaffPasskeyMinimum                  ErrorCode = "staff_passkey_minimum"
	CodeStaffStepUpRequired                  ErrorCode = "staff_step_up_required"
	CodeStaffPasskeyEnrolmentClosed          ErrorCode = "staff_passkey_enrolment_closed"
	CodeStaffTotpInvalid                     ErrorCode = "staff_totp_invalid"
	CodeStaffSignInRefused                   ErrorCode = "staff_sign_in_refused"
	CodeStaffPasswordRejected                ErrorCode = "staff_password_rejected"
	CodeStaffPasswordCurrentInvalid          ErrorCode = "staff_password_current_invalid"
	CodeStaffTotpReplayed                    ErrorCode = "staff_totp_replayed"
	CodeStaffTotpRequiredForRoot             ErrorCode = "staff_totp_required_for_root"
	CodeStaffTotpEnrolmentRequired           ErrorCode = "staff_totp_enrolment_required"
	CodeStaffTotpNotEnrolled                 ErrorCode = "staff_totp_not_enrolled"
	CodeIPNotAllowed                         ErrorCode = "ip_not_allowed"
	CodeResourceNotFound                     ErrorCode = "resource_not_found"
	CodeRouteNotFound                        ErrorCode = "route_not_found"
	CodeResourceAlreadyExists                ErrorCode = "resource_already_exists"
	CodeParameterInvalid                     ErrorCode = "parameter_invalid"
	CodeParameterMissing                     ErrorCode = "parameter_missing"
	CodeIdempotencyKeyReused                 ErrorCode = "idempotency_key_reused"
	CodeIdempotencyKeyInProgress             ErrorCode = "idempotency_key_in_progress"
	CodeVersionUnsupported                   ErrorCode = "version_unsupported"
	CodeLogoInvalid                          ErrorCode = "logo_invalid"
	CodeFontInvalid                          ErrorCode = "font_invalid"
	CodePlanLimitReached                     ErrorCode = "plan_limit_reached"
	CodeRateLimited                          ErrorCode = "rate_limited"
	CodeHomeRegionUnavailable                ErrorCode = "home_region_unavailable"
	CodeServiceUnavailable                   ErrorCode = "service_unavailable"
	CodeDomainVerificationFailed             ErrorCode = "domain_verification_failed"
	CodeDomainPublicSuffix                   ErrorCode = "domain_public_suffix"
	CodeDomainAlreadyClaimed                 ErrorCode = "domain_already_claimed"
	CodeMetadataFetchBlocked                 ErrorCode = "metadata_fetch_blocked"
	CodeMetadataInvalid                      ErrorCode = "metadata_invalid"
	CodeMetadataFetchFailed                  ErrorCode = "metadata_fetch_failed"
	CodeOIDCDiscoveryInvalid                 ErrorCode = "oidc_discovery_invalid"
	CodeCertificateInvalid                   ErrorCode = "certificate_invalid"
	CodeCertificateLastRemaining             ErrorCode = "certificate_last_remaining"
	CodeOrganizationNotEmpty                 ErrorCode = "organization_not_empty"
	CodeConnectionStateConflict              ErrorCode = "connection_state_conflict"
	CodeDirectoryStateConflict               ErrorCode = "directory_state_conflict"
	CodeWebhookEndpointUnreachable           ErrorCode = "webhook_endpoint_unreachable"
	CodeWebhookEndpointDisabled              ErrorCode = "webhook_endpoint_disabled"
	CodeImportSourceUnreachable              ErrorCode = "import_source_unreachable"
	CodePayloadTooLarge                      ErrorCode = "payload_too_large"
	CodeImportStateConflict                  ErrorCode = "import_state_conflict"
	CodeCompatSupportExpired                 ErrorCode = "compat_support_expired"
	CodeInternalError                        ErrorCode = "internal_error"
	CodeNotImplemented                       ErrorCode = "not_implemented"
	CodeDomainPendingDns                     ErrorCode = "domain_pending_dns"
	CodeFirstSCIMRequestPending              ErrorCode = "first_scim_request_pending"
	CodeTestSignInPending                    ErrorCode = "test_sign_in_pending"
	CodeDomainUnverified                     ErrorCode = "domain_unverified"
	CodePortalLinkExpired                    ErrorCode = "portal_link_expired"
	CodePortalLinkAlreadyUsed                ErrorCode = "portal_link_already_used"
	CodePortalLinkRevoked                    ErrorCode = "portal_link_revoked"
	CodeSignInConfirmationAbandoned          ErrorCode = "sign_in_confirmation_abandoned"
	CodeSignInPolicyDenied                   ErrorCode = "sign_in_policy_denied"
	CodePortalProfileIDPNotAllowed           ErrorCode = "portal_profile_idp_not_allowed"
	CodePortalProfileFeatureNotAllowed       ErrorCode = "portal_profile_feature_not_allowed"
	CodePortalProfileAttributesMissing       ErrorCode = "portal_profile_attributes_missing"
	CodePortalSetupFinished                  ErrorCode = "portal_setup_finished"
	CodePortalProfileInUse                   ErrorCode = "portal_profile_in_use"
	CodeCustomAttributeInUse                 ErrorCode = "custom_attribute_in_use"
	CodeSPKeyRotationConflict                ErrorCode = "sp_key_rotation_conflict"
	CodeSPKeyRotationNotReady                ErrorCode = "sp_key_rotation_not_ready"
	CodeSPKeyRotationKeyInUse                ErrorCode = "sp_key_rotation_key_in_use"
	CodeSPKeyRotationUnsupported             ErrorCode = "sp_key_rotation_unsupported"
	CodeSPKeyRotationRetirementWarning       ErrorCode = "sp_key_rotation_retirement_warning"
	CodeMcpKeyTooBroad                       ErrorCode = "mcp_key_too_broad"
	CodeMcpOriginRefused                     ErrorCode = "mcp_origin_refused"
	CodeBillingProviderUnavailable           ErrorCode = "billing_provider_unavailable"
	CodeBillingWebhookSignatureInvalid       ErrorCode = "billing_webhook_signature_invalid"
	CodeBillingSubscriptionRequired          ErrorCode = "billing_subscription_required"
	CodeBillingHandoverUnsupported           ErrorCode = "billing_handover_unsupported"
	CodeGstinInvalid                         ErrorCode = "gstin_invalid"
	CodeStaffApprovalRequired                ErrorCode = "staff_approval_required"
	CodeStaffApprovalSelf                    ErrorCode = "staff_approval_self"
	CodeStaffApprovalExpired                 ErrorCode = "staff_approval_expired"
	CodeStaffApprovalNotPending              ErrorCode = "staff_approval_not_pending"
	CodeStaffApprovalDuplicate               ErrorCode = "staff_approval_duplicate"
	CodeStaffApprovalStale                   ErrorCode = "staff_approval_stale"
	CodeOIDCTenantMismatch                   ErrorCode = "oidc_tenant_mismatch"
	CodeOIDCTenantClaimMissing               ErrorCode = "oidc_tenant_claim_missing"
	CodeOIDCPersonalAccountRefused           ErrorCode = "oidc_personal_account_refused"
	CodeManagedOIDCUnavailable               ErrorCode = "managed_oidc_unavailable"
	CodeConnectionShareInvalid               ErrorCode = "connection_share_invalid"
	CodeOrganizationSlugTaken                ErrorCode = "organization_slug_taken"
	CodeAPIKeyOrganizationForbidden          ErrorCode = "api_key_organization_forbidden"
	CodeAPIKeySecretLimit                    ErrorCode = "api_key_secret_limit"
	CodeSAMLSloDisabled                      ErrorCode = "saml_slo_disabled"
	CodeSAMLLogoutSessionNotFound            ErrorCode = "saml_logout_session_not_found"
	CodeTrustedTokenInvalid                  ErrorCode = "trusted_token_invalid"
	CodeTrustedTokenReplayed                 ErrorCode = "trusted_token_replayed"
	CodeDirectorySourceMismatch              ErrorCode = "directory_source_mismatch"
	CodeDirectoryCsvInvalid                  ErrorCode = "directory_csv_invalid"
	CodeDirectorySourceCredentialsInvalid    ErrorCode = "directory_source_credentials_invalid"
	CodeDirectorySourceMethodUnavailable     ErrorCode = "directory_source_method_unavailable"
	CodeDirectorySourceWorkspaceRequired     ErrorCode = "directory_source_workspace_required"
	CodeDirectorySourceAdminRequired         ErrorCode = "directory_source_admin_required"
	CodeDirectorySourceScopeMissing          ErrorCode = "directory_source_scope_missing"
	CodeDirectorySourceUnavailable           ErrorCode = "directory_source_unavailable"
	CodeDirectorySourceMassDeleteRefused     ErrorCode = "directory_source_mass_delete_refused"
	CodeSignInHookDenied                     ErrorCode = "sign_in_hook_denied"
	CodeSignInHookUnavailable                ErrorCode = "sign_in_hook_unavailable"
	CodeSignInRiskDenied                     ErrorCode = "sign_in_risk_denied"
	CodeEncryptionKeyUnavailable             ErrorCode = "encryption_key_unavailable"
	CodeEncryptionKeyInvalid                 ErrorCode = "encryption_key_invalid"
	CodeEncryptionKeyPolicyTooBroad          ErrorCode = "encryption_key_policy_too_broad"
	CodeVaultDecryptFailed                   ErrorCode = "vault_decrypt_failed"
	CodeEncryptionKeyStateConflict           ErrorCode = "encryption_key_state_conflict"
	CodeAuditLogSchemaNotFound               ErrorCode = "audit_log_schema_not_found"
	CodeFeatureFlagRevisionStale             ErrorCode = "feature_flag_revision_stale"
	CodePipeProviderNotConfigured            ErrorCode = "pipe_provider_not_configured"
	CodePipeConnectionNotConnected           ErrorCode = "pipe_connection_not_connected"
	CodeDirectorySourceOauthDomainUnverified ErrorCode = "directory_source_oauth_domain_unverified"
	CodeDirectorySourceOauthConfirmRequired  ErrorCode = "directory_source_oauth_confirm_required"
	CodePipeProviderHasConnections           ErrorCode = "pipe_provider_has_connections"
	CodePipeConnectBrowserMismatch           ErrorCode = "pipe_connect_browser_mismatch"
	CodePasswordTooShort                     ErrorCode = "password_too_short"
	CodePasswordTooLong                      ErrorCode = "password_too_long"
	CodePasswordInvalid                      ErrorCode = "password_invalid"
	CodePasswordContainsAccountDetails       ErrorCode = "password_contains_account_details"
	CodePasswordBreached                     ErrorCode = "password_breached"
	CodePasswordReused                       ErrorCode = "password_reused"
	CodeUserCredentialsInvalid               ErrorCode = "user_credentials_invalid"
	CodeUserCodeInvalid                      ErrorCode = "user_code_invalid"
	CodeUserBlocked                          ErrorCode = "user_blocked"
	CodeUserFactorInvalid                    ErrorCode = "user_factor_invalid"
	CodeUserSocialSignInFailed               ErrorCode = "user_social_sign_in_failed"
	CodeUserSocialDemoUnavailable            ErrorCode = "user_social_demo_unavailable"
	CodeUserLinkRefused                      ErrorCode = "user_link_refused"
	CodeRefreshTokenReused                   ErrorCode = "refresh_token_reused"
	CodeSignInBlocked                        ErrorCode = "sign_in_blocked"
	CodeBotChallengeFailed                   ErrorCode = "bot_challenge_failed"
	CodeOauthResourceExists                  ErrorCode = "oauth_resource_exists"
	CodeOauthThirdPartyDenied                ErrorCode = "oauth_third_party_denied"
	CodePassportCodeInvalid                  ErrorCode = "passport_code_invalid"
	CodePassportRotationNotPermitted         ErrorCode = "passport_rotation_not_permitted"
	CodePassportConsentRequired              ErrorCode = "passport_consent_required"
	CodePassportDomainProofFailed            ErrorCode = "passport_domain_proof_failed"
	CodePassportHandoffInvalid               ErrorCode = "passport_handoff_invalid"
	CodeLoginMethodDisabled                  ErrorCode = "login_method_disabled"
	CodeNativeOriginInvalid                  ErrorCode = "native_origin_invalid"
	CodeExplanationBudgetExhausted           ErrorCode = "explanation_budget_exhausted"
	CodeExplanationRateLimited               ErrorCode = "explanation_rate_limited"
	CodeLoginRequired                        ErrorCode = "login_required"
	CodeUserLastFactor                       ErrorCode = "user_last_factor"
	CodeDirectorySourceHostKeyMismatch       ErrorCode = "directory_source_host_key_mismatch"
	CodeDirectorySourceFileNotFound          ErrorCode = "directory_source_file_not_found"
	CodeDirectorySourceOauthStateInvalid     ErrorCode = "directory_source_oauth_state_invalid"
	CodeDirectorySourceOauthExpired          ErrorCode = "directory_source_oauth_expired"
	CodeDirectorySourceCompanyRequired       ErrorCode = "directory_source_company_required"
	CodeEndUserSSORequired                   ErrorCode = "end_user_sso_required"
	CodeSigninHostnameLimitReached           ErrorCode = "signin_hostname_limit_reached"
	CodeSignUpNotPending                     ErrorCode = "sign_up_not_pending"
	CodeProviderComingSoon                   ErrorCode = "provider_coming_soon"
	CodeProviderNotComingSoon                ErrorCode = "provider_not_coming_soon"
)

var errorCatalogue = map[ErrorCode]ErrorDefinition{
	CodeSAMLResponseMalformed:                {Type: "idp_error", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The SAMLResponse could not be parsed (size, encoding or XML structure)"},
	CodeSAMLIDPStatusError:                   {Type: "idp_error", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The IdP returned a non-Success SAML status; its top-level and second-level StatusCode values are in `param`"},
	CodeSAMLDtdForbidden:                     {Type: "idp_error", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The SAMLResponse contained a DOCTYPE/ENTITY, which is refused"},
	CodeSAMLSignatureMissing:                 {Type: "idp_error", HTTPStatus: 400, Retryable: false, DeveloperMessage: "Neither the assertion nor the response containing exactly one assertion is signed"},
	CodeSAMLSignatureInvalid:                 {Type: "idp_error", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The signature didn't verify against any configured IdP certificate"},
	CodeSAMLResponseAltered:                  {Type: "idp_error", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The signature verified with a configured IdP certificate, but the signed content's digest doesn't match: the response was changed after it was signed"},
	CodeSAMLSignatureAlgorithmWeak:           {Type: "idp_error", HTTPStatus: 400, Retryable: false, DeveloperMessage: "SHA-1 signature without the legacy flag on this connection"},
	CodeSAMLAssertionCountInvalid:            {Type: "idp_error", HTTPStatus: 400, Retryable: false, DeveloperMessage: "Zero or several assertions in the response"},
	CodeSAMLWrappingDetected:                 {Type: "idp_error", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The signed element isn't the element being used (signature-wrapping pattern)"},
	CodeSAMLDestinationMismatch:              {Type: "idp_error", HTTPStatus: 400, Retryable: false, DeveloperMessage: "`Destination` ≠ the ACS URL the response arrived on, or not registered"},
	CodeSAMLAudienceMismatch:                 {Type: "idp_error", HTTPStatus: 400, Retryable: false, DeveloperMessage: "`Audience` ≠ the connection's SP entity ID"},
	CodeSAMLIssuerMismatch:                   {Type: "idp_error", HTTPStatus: 400, Retryable: false, DeveloperMessage: "`Issuer` ≠ the configured IdP entity ID"},
	CodeSAMLResponseExpired:                  {Type: "idp_error", HTTPStatus: 400, Retryable: false, DeveloperMessage: "`NotOnOrAfter` passed (±2 min skew)"},
	CodeSAMLResponseNotYetValid:              {Type: "idp_error", HTTPStatus: 400, Retryable: false, DeveloperMessage: "`NotBefore` in the future beyond the skew"},
	CodeSAMLInResponseToMismatch:             {Type: "idp_error", HTTPStatus: 400, Retryable: false, DeveloperMessage: "`InResponseTo` doesn't match the flow's AuthnRequest"},
	CodeSAMLAssertionReplayed:                {Type: "authentication", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The assertion was already used"},
	CodeSAMLIDPInitiatedDisabled:             {Type: "permission", HTTPStatus: 400, Retryable: false, DeveloperMessage: "IdP-initiated sign-in is off for this connection"},
	CodeSAMLDecryptionFailed:                 {Type: "idp_error", HTTPStatus: 400, Retryable: false, DeveloperMessage: "Encrypted assertion couldn't be decrypted with our SP encryption key"},
	CodeSAMLNameidMissing:                    {Type: "idp_error", HTTPStatus: 400, Retryable: false, DeveloperMessage: "No NameID / subject"},
	CodeSAMLEmailMissing:                     {Type: "idp_error", HTTPStatus: 400, Retryable: false, DeveloperMessage: "No email attribute or NameID email"},
	CodeOIDCEmailMissing:                     {Type: "idp_error", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The OIDC IdP returned no email claim"},
	CodeOIDCDiscoveryUnreachable:             {Type: "idp_error", HTTPStatus: 502, Retryable: false, DeveloperMessage: "The discovery document couldn't be fetched (network, TLS, or blocked address)"},
	CodeOIDCTokenExchangeFailed:              {Type: "idp_error", HTTPStatus: 502, Retryable: false, DeveloperMessage: "The IdP rejected the code exchange (client ID/secret, redirect URI)"},
	CodeOIDCIDPError:                         {Type: "idp_error", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The IdP returned an OAuth error to the callback instead of a code; its `error` value is in the log. access_denied (the user canceled or is not assigned) is returned to the app's redirect URI"},
	CodeOIDCIDTokenInvalid:                   {Type: "idp_error", HTTPStatus: 400, Retryable: false, DeveloperMessage: "ID token failed iss/aud/exp/iat/nonce/signature/alg checks"},
	CodeOIDCEmailNotVerified:                 {Type: "idp_error", HTTPStatus: 400, Retryable: false, DeveloperMessage: "`email_verified` is false and the connection doesn't opt out"},
	CodeEmailOutsideVerifiedDomain:           {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The email domain isn't a verified domain of this organization"},
	CodeConnectionUntested:                   {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The connection hasn't passed a test sign-in"},
	CodeConnectionInactive:                   {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The connection is deactivated"},
	CodeConnectionNotActivated:               {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The connection passed its test sign-in but was never activated: activate it with POST /v1/connections/{id}/activate or finish setup in the portal"},
	CodeEnvironmentDisabled:                  {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "Sign-ins and SCIM are switched off for this environment (environment kill switch); the management API keeps working"},
	CodeAccountSuspended:                     {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The account is suspended by OhAuthCat; sign-ins, SCIM, portal links and custom portal domains are offline until the suspension is lifted"},
	CodeAccountClosed:                        {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The account is closed and scheduled for deletion; sign-ins, SCIM, portal links and API keys are refused, and the account is read-only apart from exports"},
	CodeOrganizationNotFound:                 {Type: "not_found", HTTPStatus: 404, Retryable: false, DeveloperMessage: "No organization matches the given organization/domain/login_hint"},
	CodeClientIDInvalid:                      {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "`client_id` must be a publishable key (pk_...) of an environment"},
	CodePKCERequired:                         {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "`code_challenge` (43 base64url characters) with `code_challenge_method=S256` is required; plain is refused"},
	CodeAuthorizeRequestInvalid:              {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "An authorize parameter is invalid or repeated, or the routing hints contradict each other; `param` names it"},
	CodeSSOForwardingDisabled:                {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "Forwarding (POST /v1/sso/forwarded) is off for this connection, or its window has ended; set `forwarding` on the connection"},
	CodeConnectionNotFound:                   {Type: "not_found", HTTPStatus: 404, Retryable: false, DeveloperMessage: "The `connection` is not a connection of this environment"},
	CodeRedirectURINotAllowed:                {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The `redirect_uri` isn't on the environment's allow-list (exact match)"},
	CodeInvalidState:                         {Type: "authentication", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The state token is missing, tampered with, expired or already used"},
	CodeFlowInterruptedRetry:                 {Type: "unavailable", HTTPStatus: 503, Retryable: true, DeveloperMessage: "The flow's origin cell couldn't confirm single use; fail closed"},
	CodeInvalidGrant:                         {Type: "authentication", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The code is invalid, expired, already used, or bound to another client/redirect"},
	CodePKCEVerificationFailed:               {Type: "authentication", HTTPStatus: 400, Retryable: false, DeveloperMessage: "`code_verifier` doesn't match the challenge (S256 only)"},
	CodeConnectionStandbyUnavailable:         {Type: "unavailable", HTTPStatus: 503, Retryable: true, DeveloperMessage: "During a Cloudflare outage, this connection didn't register the standby ACS"},
	CodeConsentDeclined:                      {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The user declined the consent screen"},
	CodeSignInBufferFull:                     {Type: "unavailable", HTTPStatus: 503, Retryable: true, DeveloperMessage: "Database unavailable and the fallback buffer is at its bound; fail closed"},
	CodeSCIMTokenInvalid:                     {Type: "authentication", HTTPStatus: 401, Retryable: false, DeveloperMessage: "The bearer token is unknown, revoked or rotated out"},
	CodeSCIMUserOutOfDomain:                  {Type: "permission", HTTPStatus: 200, Retryable: false, DeveloperMessage: "The user's email isn't in a verified domain; stored as `out_of_domain` and not emitted"},
	CodeSCIMFilterInvalid:                    {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "Unsupported or malformed SCIM filter"},
	CodeSCIMPayloadTooLarge:                  {Type: "invalid_request", HTTPStatus: 413, Retryable: false, DeveloperMessage: "The request body exceeds the cap"},
	CodeSCIMRateLimited:                      {Type: "rate_limited", HTTPStatus: 429, Retryable: true, DeveloperMessage: "Per-directory rate or concurrency limit"},
	CodeSCIMTemporarilyUnavailable:           {Type: "unavailable", HTTPStatus: 503, Retryable: true, DeveloperMessage: "The database is unavailable or re-homing; retry after `Retry-After`"},
	CodeSCIMConflict:                         {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "A user/group with this `userName`/`externalId` exists"},
	CodeSCIMPreconditionFailed:               {Type: "conflict", HTTPStatus: 412, Retryable: true, DeveloperMessage: "The resource changed since the `If-Match` ETag was read"},
	CodeAuthenticationRequired:               {Type: "authentication", HTTPStatus: 401, Retryable: false, DeveloperMessage: "Missing or malformed API key / session"},
	CodeAPIKeyInvalid:                        {Type: "authentication", HTTPStatus: 401, Retryable: false, DeveloperMessage: "Unknown, revoked or expired key"},
	CodeAccessTokenInvalid:                   {Type: "authentication", HTTPStatus: 401, Retryable: false, DeveloperMessage: "The SSO access token is malformed, expired, revoked or from another environment"},
	CodeCliTokenInvalid:                      {Type: "authentication", HTTPStatus: 401, Retryable: false, DeveloperMessage: "The CLI access token is malformed, expired or revoked, or its member or environment no longer has access; refresh it with the refresh token or run `ohauthcat login` again"},
	CodeCsrfTokenInvalid:                     {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The X-CSRF-Token header is missing or doesn't match the session"},
	CodeOtpInvalid:                           {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The email code is wrong, expired or out of attempts (deliberately not told apart; request a new code)"},
	CodeRecoveryCodeInvalid:                  {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The email and recovery code don't match an unused owner recovery code (deliberately not told apart)"},
	CodeInviteInvalid:                        {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The invite is unknown, used, revoked, expired or addressed to another email"},
	CodeWebauthnInvalid:                      {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The passkey ceremony failed (expired or unknown challenge, wrong origin or RP ID, bad signature, unknown or cloned credential)"},
	CodeTotpInvalid:                          {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The authenticator code is wrong, outside the time window or already used"},
	CodeOauthSignInFailed:                    {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "GitHub/Google sign-in failed: bad or reused state, provider error, no verified email or no account for it (sent as `error` on the redirect)"},
	CodeMfaRequired:                          {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The session needs a second factor first: verify a passkey or TOTP code, or enrol one (owners: a backup passkey)"},
	CodeAuthorizationPending:                 {Type: "invalid_request", HTTPStatus: 400, Retryable: true, DeveloperMessage: "RFC 8628 §3.5: the member hasn't approved the CLI sign-in yet; poll again after `interval` seconds"},
	CodeSlowDown:                             {Type: "invalid_request", HTTPStatus: 400, Retryable: true, DeveloperMessage: "RFC 8628 §3.5: polled faster than the interval; the interval grew by 5 seconds"},
	CodeExpiredToken:                         {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "RFC 8628 §3.5: the device code expired; start the CLI sign-in again"},
	CodeAccessDenied:                         {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "RFC 8628 §3.5 / RFC 6749 §4.1.2.1: the member denied the CLI sign-in"},
	CodeCliRequestStale:                      {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The CLI sign-in request is more than 5 minutes old and can no longer be approved; start the sign-in again from the CLI"},
	CodeStepUpRequired:                       {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "This action needs a recent second factor (passkey or TOTP within the last 10 minutes)"},
	CodeSignUpPaused:                         {Type: "unavailable", HTTPStatus: 503, Retryable: true, DeveloperMessage: "New sign-ups are paused while the sign-up challenge is unreachable (e.g. during a Cloudflare outage, on the standby host); existing accounts can still sign in"},
	CodeSandboxCreationUnavailable:           {Type: "unavailable", HTTPStatus: 503, Retryable: false, DeveloperMessage: "Anonymous sandbox creation is not configured on this deployment (no sign-up challenge verifier); sign up for an account instead"},
	CodeSandboxHandoffExpired:                {Type: "invalid_request", HTTPStatus: 410, Retryable: false, DeveloperMessage: "The sandbox handoff expired, was already used or is unknown; run ohauthcat init again"},
	CodeDebugBundlesUnavailable:              {Type: "unavailable", HTTPStatus: 503, Retryable: false, DeveloperMessage: "Debug bundles are not configured on this deployment: they are kept in the export object store, which needs EXPORT_S3_BUCKET, EXPORT_S3_REGION and EXPORT_S3_KMS_KEY_ID"},
	CodeEmailDisposable:                      {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "Sign-up refuses disposable email domains; use a work or personal address you keep"},
	CodeSSORequired:                          {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The account requires dashboard SSO; only SSO, owner backup passkeys (which sign in directly, never through SSO) and owner recovery codes sign in"},
	CodeRoleInsufficient:                     {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The member's role doesn't allow this action (the required role is in `param`)"},
	CodeLastOwner:                            {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The account's last owner can't be removed or demoted"},
	CodeLastBreakGlassCredential:             {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "An owner's last usable backup passkey can't be removed; register another backup passkey first"},
	CodeEnvironmentResetNotAllowed:           {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "Only sandbox environments can be reset; production environments and OhAuthCat's internal dashboard SSO environment never are"},
	CodeSandboxRequired:                      {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "Developer tools (the webhook relay, `trigger`, `init` and the SCIM simulator replay) work in sandbox environments only; use a sandbox API key or CLI login"},
	CodeSSONotReady:                          {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "Dashboard SSO can't be enforced yet: set it up, activate a connection, and hold a backup passkey and unused recovery codes (the missing step is in `message`)"},
	CodeBreakGlassLimit:                      {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The owner already holds the maximum number of backup passkeys (2)"},
	CodeAPIKeyScopeMissing:                   {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The key lacks the scope `<resource>:<action>` (in `param`)"},
	CodeStaffRoleInsufficient:                {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The staff member's role is below the role the staff route needs (in `param`)"},
	CodeBufferEntryConflict:                  {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The fallback-buffer entry can't take this action: it still decodes, the quarantine stream is full, or the record is already buffered or recorded (detail in `message`)"},
	CodeStaffMemberDisabled:                  {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The staff member is disabled; an admin can enable them again"},
	CodeStaffLastAdmin:                       {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The change would leave no active staff admin; promote another member first"},
	CodeStaffOriginRefused:                   {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "State-changing staff requests must carry an Origin header equal to the Mission HQ console origin"},
	CodeStaffUnsupportedMediaType:            {Type: "invalid_request", HTTPStatus: 415, Retryable: false, DeveloperMessage: "Staff write requests with a body must be Content-Type: application/json"},
	CodeStaffRoleImmutable:                   {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "Built-in staff roles are defined in code and cannot be changed or deleted"},
	CodeStaffRoleHeld:                        {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "Staff members cannot edit or delete a role they hold"},
	CodeStaffRoleInUse:                       {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The staff role is assigned to members; reassign them before deleting it"},
	CodeStaffAccountProtected:                {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The target is protected: the root staff member, the house account or a platform system account cannot be suspended, terminated, demoted or disabled"},
	CodeStaffSelfChange:                      {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "Staff members cannot change their own role, geo scope or status; another member must"},
	CodeStaffGrantExceedsOwn:                 {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The role or geo scope grants more than the granting member holds (permissions or countries)"},
	CodeStaffSessionRequired:                 {Type: "authentication", HTTPStatus: 401, Retryable: false, DeveloperMessage: "This staff route needs a live staff session: sign in with the staff email sign-in (POST /v1/session/sign_in) and send the session cookie with a request signature"},
	CodeStaffSessionDeviceChanged:            {Type: "authentication", HTTPStatus: 401, Retryable: false, DeveloperMessage: "The staff session was revoked: the request signature, its device key or the browser details no longer match the browser the session was opened in"},
	CodeStaffPasskeyMinimum:                  {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "Removing this passkey would leave the staff member without one (everyone, root included, keeps at least 1)"},
	CodeStaffStepUpRequired:                  {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "This action needs a staff passkey or authenticator assertion on this session in the last 5 minutes (2 minutes, and not already spent, for a password change) (POST /v1/session/step_up or /v1/session/step_up/totp)"},
	CodeStaffPasskeyEnrolmentClosed:          {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The member has no staff passkey and its enrolment window is closed or was never opened; an admin must reopen it"},
	CodeStaffTotpInvalid:                     {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The staff authenticator code is wrong or outside the time window (one 30-second step either side)"},
	CodeStaffSignInRefused:                   {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The staff email and password sign-in was refused: wrong password, unknown or disabled email, no password or authenticator app, or a locked account (one answer for all)"},
	CodeStaffPasswordRejected:                {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The new staff password breaks a password rule (12 to 256 characters, an uppercase letter, a lowercase letter, a number and a symbol, not part of the email, not in a breach corpus); the message names the rule and error.details lists any missing character type"},
	CodeStaffPasswordCurrentInvalid:          {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "Changing a staff password needs the current password in current_password; it was missing or wrong, or the password is locked after too many wrong tries"},
	CodeStaffTotpReplayed:                    {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "This staff authenticator code (or a later one) was already used; each 30-second code is accepted once"},
	CodeStaffTotpRequiredForRoot:             {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The root staff member must keep its authenticator app (TOTP); replace it by enrolling a new one instead"},
	CodeStaffTotpEnrolmentRequired:           {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The root staff member must finish setup by enrolling an authenticator app (POST /v1/me/totp/enroll/options); until then only GET /v1/me and the setup routes are allowed"},
	CodeStaffTotpNotEnrolled:                 {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The staff member has no authenticator app (TOTP), or no pending enrolment to confirm"},
	CodeIPNotAllowed:                         {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The key's IP allow-list doesn't include the caller"},
	CodeResourceNotFound:                     {Type: "not_found", HTTPStatus: 404, Retryable: false, DeveloperMessage: "The object doesn't exist in this environment (also returned for other tenants' IDs)"},
	CodeRouteNotFound:                        {Type: "not_found", HTTPStatus: 404, Retryable: false, DeveloperMessage: "No such endpoint. See the API reference on the docs site (doc_url names it)"},
	CodeResourceAlreadyExists:                {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "An object with this unique value (`param`) already exists in this environment"},
	CodeParameterInvalid:                     {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "A value sent isn't valid; `param` names its field when it is known"},
	CodeParameterMissing:                     {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "A required parameter is missing; `param` names it"},
	CodeIdempotencyKeyReused:                 {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "Same key with a different request body"},
	CodeIdempotencyKeyInProgress:             {Type: "conflict", HTTPStatus: 409, Retryable: true, DeveloperMessage: "A request with this `Idempotency-Key` is still being processed; retry after it finishes"},
	CodeVersionUnsupported:                   {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "Unknown `OhAuthCat-Version`, or a major past its sunset"},
	CodeLogoInvalid:                          {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The logo must be a PNG (16 to 2048 pixels per side) or an SVG with a viewBox of at most 2048 units, no larger than 256 KiB. The type is read from the file's bytes; `message` names the problem"},
	CodeFontInvalid:                          {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The font must be a single WOFF2 file (not a collection) of at most 512 KiB with a valid header and table directory; `message` names the problem"},
	CodePlanLimitReached:                     {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The project's plan doesn't include this feature or allows no more of it. The message names the feature or limit and the plan that unlocks it, and `param` the field or limit when there is one (for the enforced monthly limits monthly_active_users, audit_log_events or pipes_token_fetches); `doc_url` explains upgrades. Existing resources and existing users keep working"},
	CodeRateLimited:                          {Type: "rate_limited", HTTPStatus: 429, Retryable: true, DeveloperMessage: "With `Retry-After`"},
	CodeHomeRegionUnavailable:                {Type: "unavailable", HTTPStatus: 503, Retryable: true, DeveloperMessage: "The environment's home region cannot be reached; the request was refused rather than served from another region"},
	CodeServiceUnavailable:                   {Type: "unavailable", HTTPStatus: 503, Retryable: true, DeveloperMessage: "With `Retry-After`; writes paused (DB failover)"},
	CodeDomainVerificationFailed:             {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "TXT record not found or wrong value"},
	CodeDomainPublicSuffix:                   {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "Public email or public-suffix domain can't be claimed (inv. 20)"},
	CodeDomainAlreadyClaimed:                 {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "Verified by another organization in this environment"},
	CodeMetadataFetchBlocked:                 {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The URL failed the SSRF-safe fetcher (non-https, private address, too many redirects)"},
	CodeMetadataInvalid:                      {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The IdP metadata XML is invalid or has no usable signing cert"},
	CodeMetadataFetchFailed:                  {Type: "invalid_request", HTTPStatus: 400, Retryable: true, DeveloperMessage: "The metadata URL could not be fetched (network error, timeout or a non-2xx status)"},
	CodeOIDCDiscoveryInvalid:                 {Type: "invalid_request", HTTPStatus: 400, Retryable: true, DeveloperMessage: "The issuer's /.well-known/openid-configuration couldn't be fetched, or its issuer, endpoints or jwks_uri are missing, not https or don't match the issuer"},
	CodeCertificateInvalid:                   {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "`certificate_pem` isn't a single X.509 certificate (PEM or base64 DER), or it has expired"},
	CodeCertificateLastRemaining:             {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The certificate is the connection's only IdP certificate; add the new one before removing it"},
	CodeOrganizationNotEmpty:                 {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The organization still has connections, directories or live API keys; delete or revoke them first (`param` lists what is left, e.g. `connections:2,directories:1,api_keys:1`, or names the table of another child). Revoked keys and the organization's settings are deleted with it"},
	CodeConnectionStateConflict:              {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The connection's state doesn't allow this change (e.g. deactivating a connection that was never active)"},
	CodeDirectoryStateConflict:               {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The directory's state doesn't allow this change (e.g. deleting from the portal a directory that already received a request or has a source)"},
	CodeWebhookEndpointUnreachable:           {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The test delivery failed (status in `param`)"},
	CodeWebhookEndpointDisabled:              {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The webhook endpoint is disabled; enable it before sending or resending"},
	CodeImportSourceUnreachable:              {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The import source rejected the credentials or couldn't be reached"},
	CodePayloadTooLarge:                      {Type: "invalid_request", HTTPStatus: 413, Retryable: false, DeveloperMessage: "The request body exceeds the size limit for this route (1 MiB by default; larger for imports)"},
	CodeImportStateConflict:                  {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The import or organization's state doesn't allow this (e.g. running a completed import, switching an organization that is not imported or whose imported connections are not all active, or a run already in progress)"},
	CodeCompatSupportExpired:                 {Type: "permission", HTTPStatus: 410, Retryable: false, DeveloperMessage: "The `compat/workos` window for this environment ended. Use the SDK, or buy the extended support add-on (D42)"},
	CodeInternalError:                        {Type: "internal", HTTPStatus: 500, Retryable: false, DeveloperMessage: "Unexpected; `request_id` is required in support tickets"},
	CodeNotImplemented:                       {Type: "internal", HTTPStatus: 501, Retryable: false, DeveloperMessage: "The operation is declared in the API but not available yet"},
	CodeDomainPendingDns:                     {Type: "portal_state", HTTPStatus: 0, Retryable: false, DeveloperMessage: "The domain's DNS TXT record hasn't been found yet"},
	CodeFirstSCIMRequestPending:              {Type: "portal_state", HTTPStatus: 0, Retryable: false, DeveloperMessage: "No SCIM request has arrived from the IdP yet"},
	CodeTestSignInPending:                    {Type: "portal_state", HTTPStatus: 0, Retryable: false, DeveloperMessage: "The connection hasn't completed a test sign-in yet"},
	CodeDomainUnverified:                     {Type: "portal_state", HTTPStatus: 0, Retryable: false, DeveloperMessage: "The setup portal can't turn on single sign-on until the organization has a verified domain. Activation through the management API is unchanged"},
	CodePortalLinkExpired:                    {Type: "portal_state", HTTPStatus: 0, Retryable: false, DeveloperMessage: "The portal link expired"},
	CodePortalLinkAlreadyUsed:                {Type: "portal_state", HTTPStatus: 0, Retryable: false, DeveloperMessage: "The single-use portal link was already used"},
	CodePortalLinkRevoked:                    {Type: "portal_state", HTTPStatus: 0, Retryable: false, DeveloperMessage: "The portal link was revoked"},
	CodeSignInConfirmationAbandoned:          {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The unsolicited sign-in confirmation was cancelled or expired; no application ticket was issued"},
	CodeSignInPolicyDenied:                   {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The configured sign-in policy did not admit this verified identity"},
	CodePortalProfileIDPNotAllowed:           {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The selected IdP is not allowed by this portal profile"},
	CodePortalProfileFeatureNotAllowed:       {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The selected feature is not allowed by this portal profile"},
	CodePortalProfileAttributesMissing:       {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "Required attributes are not mapped for every selected feature"},
	CodePortalSetupFinished:                  {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "This portal setup is already finished and cannot be edited"},
	CodePortalProfileInUse:                   {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "This portal profile is referenced by an unexpired setup link"},
	CodeCustomAttributeInUse:                 {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "This custom attribute is required by a portal profile"},
	CodeSPKeyRotationConflict:                {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The signing-key rotation generation or state no longer matches this request"},
	CodeSPKeyRotationNotReady:                {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The required signing-key rotation steps are not complete"},
	CodeSPKeyRotationKeyInUse:                {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The operation would remove this connection's current signing key"},
	CodeSPKeyRotationUnsupported:             {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "SP signing-key rotation is available only for a configured SAML connection"},
	CodeSPKeyRotationRetirementWarning:       {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The retained rollback signing key needs review"},
	CodeMcpKeyTooBroad:                       {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "Use a restricted secret key with mcp:use and only the required tool scopes"},
	CodeMcpOriginRefused:                     {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "Send MCP requests without an Origin header or from an origin on the configured allow-list"},
	CodeBillingProviderUnavailable:           {Type: "unavailable", HTTPStatus: 503, Retryable: true, DeveloperMessage: "No payment provider in the configured chain for this billing country could open a checkout or apply the change. Retry later"},
	CodeBillingWebhookSignatureInvalid:       {Type: "authentication", HTTPStatus: 401, Retryable: false, DeveloperMessage: "The payment webhook's signature, timestamp or event ID did not verify; nothing was recorded"},
	CodeBillingSubscriptionRequired:          {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "This change needs an active paid subscription on the project; start a checkout first"},
	CodeBillingHandoverUnsupported:           {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The item's subscription is cancelled at the period end, but no payment provider in this region can start a new subscription at that date; buy again after the period ends"},
	CodeGstinInvalid:                         {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The GSTIN fails the shape or check-digit test, or its state code differs from `state_code`"},
	CodeStaffApprovalRequired:                {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The change needs a second staff member's approval: nothing changed yet. Held routes answer 202 with the pending StaffApprovalRequest; the code names that outcome and is never sent as an error body"},
	CodeStaffApprovalSelf:                    {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The member who requested the change cannot approve or reject it; a second member must"},
	CodeStaffApprovalExpired:                 {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The approval request expired (24 hours after it was made); request the change again"},
	CodeStaffApprovalNotPending:              {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The approval request was already approved, rejected, cancelled or expired"},
	CodeStaffApprovalDuplicate:               {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "A pending approval request for the same action and target exists; approve, reject or cancel it first"},
	CodeStaffApprovalStale:                   {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The target changed since the request was made, so the action no longer applies (for example the account is already suspended); the request stays pending, reject or cancel it"},
	CodeOIDCTenantMismatch:                   {Type: "idp_error", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The ID token's Google `hd` or Entra `tid` isn't the organization's verified domain or recorded tenant"},
	CodeOIDCTenantClaimMissing:               {Type: "idp_error", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The ID token has no valid Entra `tid` claim, so it can't be bound to a tenant"},
	CodeOIDCPersonalAccountRefused:           {Type: "idp_error", HTTPStatus: 403, Retryable: false, DeveloperMessage: "A personal Google or Microsoft account signed in and the connection doesn't allow personal accounts"},
	CodeManagedOIDCUnavailable:               {Type: "unavailable", HTTPStatus: 503, Retryable: true, DeveloperMessage: "Sign in with Google Workspace or Microsoft Entra isn't configured on this deployment"},
	CodeConnectionShareInvalid:               {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "A connection can be shared only with another organization of the same environment, with the owner organization's explicit consent and at least one domain the target organization verified"},
	CodeOrganizationSlugTaken:                {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "Another organization in this environment already uses this `slug`"},
	CodeAPIKeyOrganizationForbidden:          {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The key is limited to one organization and this request is outside it"},
	CodeAPIKeySecretLimit:                    {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The key already has the most active secrets allowed (5); revoke one first"},
	CodeSAMLSloDisabled:                      {Type: "not_found", HTTPStatus: 404, Retryable: false, DeveloperMessage: "Single logout is not enabled for this SAML connection, or the IdP has no single logout URL"},
	CodeSAMLLogoutSessionNotFound:            {Type: "not_found", HTTPStatus: 404, Retryable: false, DeveloperMessage: "No recorded SAML sign-in matches the NameID and SessionIndex of the logout"},
	CodeTrustedTokenInvalid:                  {Type: "authentication", HTTPStatus: 401, Retryable: false, DeveloperMessage: "The subject token failed verification against the trusted issuer (signature, algorithm, iss, aud or time checks)"},
	CodeTrustedTokenReplayed:                 {Type: "authentication", HTTPStatus: 401, Retryable: false, DeveloperMessage: "The subject token's jti was already exchanged"},
	CodeDirectorySourceMismatch:              {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The directory's provider does not take this source: CSV imports need provider csv, source credentials need a pulled source (HRIS, Google Workspace or SFTP), and only SCIM directories accept SCIM requests or rotate a SCIM token"},
	CodeDirectoryCsvInvalid:                  {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The CSV file is invalid; `message` names up to 20 problems by row and column"},
	CodeDirectorySourceCredentialsInvalid:    {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The directory source refused the credentials, or they are not in the expected shape"},
	CodeDirectorySourceMethodUnavailable:     {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "This provider takes no pasted credentials: bamboohr and google_admin_sdk directories connect only through POST /v1/directories/{dir}/source/oauth/start (Connect BambooHR, Connect Google)"},
	CodeDirectorySourceWorkspaceRequired:     {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "Connect Google needs a Google Workspace account: the signed-in account has no Workspace domain (hd), so it is a personal Gmail account without a directory. Sent as error_code on the redirect to return_url"},
	CodeDirectorySourceAdminRequired:         {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "Connect Google: the Directory API refused the signed-in user (403), so they cannot read the Workspace's users and groups. Sign in as a Workspace admin with directory read access. Sent as error_code on the redirect to return_url"},
	CodeDirectorySourceScopeMissing:          {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The source refused the OAuth access token for a missing scope; the message names the scopes requested and what the source said. Add the scope to the OAuth app, then connect again"},
	CodeDirectorySourceUnavailable:           {Type: "invalid_request", HTTPStatus: 400, Retryable: true, DeveloperMessage: "The directory source couldn't be reached or kept failing; try again later"},
	CodeDirectorySourceMassDeleteRefused:     {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The sync would delete more than half of a directory of 20 or more users at once, so nothing was changed"},
	CodeSignInHookDenied:                     {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The environment's sign-in hook answered deny for this sign-in"},
	CodeSignInHookUnavailable:                {Type: "unavailable", HTTPStatus: 503, Retryable: true, DeveloperMessage: "The fail-closed sign-in hook did not return a valid signed answer in time (timeout, transport error, bad signature, oversized or invalid response, or a reserved claim)"},
	CodeSignInRiskDenied:                     {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The environment's risk policy refused this sign-in because of its risk signals"},
	CodeEncryptionKeyUnavailable:             {Type: "unavailable", HTTPStatus: 503, Retryable: true, DeveloperMessage: "A customer-managed key (BYOK) that protects this data is revoked, disabled, deleted, denied or unreachable, so the request fails closed. Restore the key or its key policy in AWS KMS, then POST /v1/encryption_keys/{key}/validate"},
	CodeEncryptionKeyInvalid:                 {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The key reference is not a full AWS KMS key ARN (aliases are refused), or a test encrypt and decrypt with it failed. Check the ARN and that the key policy grants our service role kms:Encrypt and kms:Decrypt"},
	CodeEncryptionKeyPolicyTooBroad:          {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The key worked under another environment's encryption context: its key policy must pin kms:EncryptionContext:ohauthcat_environment (and ohauthcat_scope for an organization key) as documented"},
	CodeVaultDecryptFailed:                   {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The ciphertext did not decrypt: it is malformed or tampered, belongs to another organization or key, or the context differs from the one used to encrypt. The reason is never given"},
	CodeEncryptionKeyStateConflict:           {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The key version's state doesn't allow this change: the active version cannot be deprecated, and deprecated or revoked versions cannot be activated"},
	CodeAuditLogSchemaNotFound:               {Type: "not_found", HTTPStatus: 404, Retryable: false, DeveloperMessage: "No schema is registered for this audit event's action (or the version it names); register one with POST /v1/audit_log_schemas first"},
	CodeFeatureFlagRevisionStale:             {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The feature flag changed since the `revision` you sent; read it again and reapply your change"},
	CodePipeProviderNotConfigured:            {Type: "not_found", HTTPStatus: 404, Retryable: false, DeveloperMessage: "No Pipes provider is configured under this slug; configure it with PUT /v1/pipes/providers/{provider}"},
	CodePipeConnectionNotConnected:           {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The Pipes connection has no usable token (pending, disconnected, or in error after the provider refused a refresh); send the end user through a new connect flow"},
	CodeDirectorySourceOauthDomainUnverified: {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The Google Workspace domain that approved the connection is not a verified domain of the directory's organization; verify the domain first, or connect with an admin of a verified domain"},
	CodeDirectorySourceOauthConfirmRequired:  {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The connect link was not confirmed on its OhAuthCat confirm page (missing or mismatched confirm cookie, CSRF field or Origin), so nothing was connected; open the authorization_url and press Continue"},
	CodePipeProviderHasConnections:           {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The provider has connected end users, so its authorize, token or revoke endpoint host can't change; disconnect those connections first, or configure the new endpoints under a new provider slug"},
	CodePipeConnectBrowserMismatch:           {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "This Pipes connect flow was opened in another browser, or the browser did not send its binding cookie; send the end user through a new connect flow from their own session"},
	CodePasswordTooShort:                     {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The password is shorter than the environment's minimum length (`param` is `password`)"},
	CodePasswordTooLong:                      {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The password is longer than 256 characters"},
	CodePasswordInvalid:                      {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The password is not valid text (invalid UTF-8 or control characters), or lacks a character type the environment requires (D191: an uppercase letter, a lowercase letter, a number and a symbol; error.details lists each missing type)"},
	CodePasswordContainsAccountDetails:       {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The password contains the user's email address or the app name"},
	CodePasswordBreached:                     {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The password appears in a public breach corpus (checked with the k-anonymity range API; the password never leaves OhAuthCat)"},
	CodePasswordReused:                       {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The password matches one of the user's recent passwords (the environment's history rule)"},
	CodeUserCredentialsInvalid:               {Type: "authentication", HTTPStatus: 401, Retryable: false, DeveloperMessage: "The email address or password is wrong; the answer is the same for unknown addresses and for throttled attempts"},
	CodeUserCodeInvalid:                      {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The email code or link is wrong, expired, already used, out of attempts or opened in another browser; every case answers alike"},
	CodeUserBlocked:                          {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The user is blocked in this environment and cannot sign in"},
	CodeUserFactorInvalid:                    {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The second-factor code (authenticator app or recovery code) is wrong, already used or throttled; every case answers alike"},
	CodeUserSocialSignInFailed:               {Type: "authentication", HTTPStatus: 400, Retryable: true, DeveloperMessage: "The social sign-in could not complete: unknown or disabled provider, a used, expired or other-browser state, a refused code exchange or an ID token that failed its checks"},
	CodeUserSocialDemoUnavailable:            {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "This provider has no OhAuthCat demo app configured, so a sandbox environment must send its own client_id and client_secret to enable it"},
	CodeUserLinkRefused:                      {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "An account with this email exists, and the provider did not assert the email is verified or the environment's linking rule is never; sign in with the existing method and link from there"},
	CodeRefreshTokenReused:                   {Type: "authentication", HTTPStatus: 400, Retryable: false, DeveloperMessage: "A refresh token that was already exchanged was presented again; the whole session is revoked (rotation with reuse detection). Sign the user in again"},
	CodeSignInBlocked:                        {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "A bot-protection rule of the environment blocked this sign-in or sign-up (see GET /v1/user_management/bot_detections)"},
	CodeBotChallengeFailed:                   {Type: "invalid_request", HTTPStatus: 400, Retryable: true, DeveloperMessage: "The CAPTCHA answer was missing, refused, already used or solved for another browser or flow; a new challenge is shown"},
	CodeOauthResourceExists:                  {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "A resource server with this identifier is already registered in the environment"},
	CodeOauthThirdPartyDenied:                {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The user's organization does not allow third-party apps, so this app cannot be authorized for its members"},
	CodePassportCodeInvalid:                  {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The Passport email code is wrong or expired, or too many tries were made; start again for a new code"},
	CodePassportRotationNotPermitted:         {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "This Passport link was not enrolled through a portal link with the `sso` or `edit_connection` intent (or was enrolled before intents were recorded), so it cannot start a connection rotation; enrol again through such a portal link"},
	CodePassportConsentRequired:              {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The Passport holds no active grant for this scope on this vendor (or on the vendor the data would come from); the scope is in `param`"},
	CodePassportDomainProofFailed:            {Type: "conflict", HTTPStatus: 409, Retryable: true, DeveloperMessage: "The reused domain proof no longer holds: the verification record was not found in DNS at the time of use, or the domain is no longer verified at the vendor that proved it"},
	CodePassportHandoffInvalid:               {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The Passport handoff token is unknown, expired or already used; start the Passport enrolment again from the vendor portal"},
	CodeLoginMethodDisabled:                  {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The sign-in method is turned off for this environment (or the organization the flow signs into); see GET /v1/user_management/settings"},
	CodeNativeOriginInvalid:                  {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "A native app origin must be android:apk-key-hash:<base64url SHA-256 of the signing certificate> or ios:<team id>.<bundle id> (served in the sign-in host's webcredentials association); at most 20 per environment"},
	CodeExplanationBudgetExhausted:           {Type: "rate_limited", HTTPStatus: 429, Retryable: false, DeveloperMessage: "The environment's monthly sign-in explanation cost cap is reached; explanations resume when the period resets or the cap is raised"},
	CodeExplanationRateLimited:               {Type: "rate_limited", HTTPStatus: 429, Retryable: true, DeveloperMessage: "Too many sign-in explanation requests for this environment; retry after `Retry-After`"},
	CodeLoginRequired:                        {Type: "authentication", HTTPStatus: 401, Retryable: false, DeveloperMessage: "The hosted account route needs a signed-in browser session of this environment (the session cookie on the login host); sign in through the hosted login flow, then retry"},
	CodeUserLastFactor:                       {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The environment or organization requires MFA and this is the user's last second factor; add another passkey or authenticator app before removing it"},
	CodeDirectorySourceHostKeyMismatch:       {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The SFTP server presented a host key whose SHA256 fingerprint differs from host_key_fingerprint, so the connection was closed"},
	CodeDirectorySourceFileNotFound:          {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "No file on the SFTP server matches the source path"},
	CodeDirectorySourceOauthStateInvalid:     {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The OAuth callback's state is unknown, already used or expired (30 minutes); start the connection again"},
	CodeDirectorySourceOauthExpired:          {Type: "invalid_request", HTTPStatus: 400, Retryable: false, DeveloperMessage: "The OAuth connect flow was not completed within 30 minutes. Sent as error_code on the redirect to return_url, never as a response; start the connection again"},
	CodeDirectorySourceCompanyRequired:       {Type: "invalid_request", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The connected HRIS account reaches several companies; choose one with POST /v1/directories/{dir}/source/oauth/company"},
	CodeEndUserSSORequired:                   {Type: "permission", HTTPStatus: 403, Retryable: false, DeveloperMessage: "The organization that owns this email domain requires SSO; password, magic-code, passkey and social sign-in are refused for its users (organization user_management_settings sso_required)"},
	CodeSigninHostnameLimitReached:           {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "Unused: no longer returned. Custom sign-in hostnames are now their own passkey relying parties (PUT /v1/environment/signin_host), so there is no platform-wide limit on them. Kept in the catalogue so clients that handle it keep working"},
	CodeSignUpNotPending:                     {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "POST /login/{env}/api/sign_up/resend was called on a flow with no pending sign-up (none started, already verified, or replaced by a sign-in code); start the sign-up again"},
	CodeProviderComingSoon:                   {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "The directory source is coming soon in this environment: new directories of it can't be created yet (existing ones keep syncing). Register interest with POST /dashboard/providers/{provider}/interest to be emailed when it goes live"},
	CodeProviderNotComingSoon:                {Type: "conflict", HTTPStatus: 409, Retryable: false, DeveloperMessage: "Interest can be registered only for a directory source that is coming soon in this environment; this one is available now (create the directory instead)"},
}

// CatalogueDefinition returns a copy of the pinned metadata for a code.
func CatalogueDefinition(code ErrorCode) (ErrorDefinition, bool) {
	definition, ok := errorCatalogue[code]
	return definition, ok
}
