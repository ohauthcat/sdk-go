# Go SDK

This proprietary SDK requires Go 1.26 and selects toolchain 1.26.11. Its runtime
imports only the standard library; repository-local goimports and govulncheck
tools have separate development dependencies. See LICENSE and THIRD_PARTY_LICENSES.

## Install

```sh
go get ohauthcat.com/sdk-go@v0.0.2
```


Documentation: https://docs.ohauthcat.com. Licensed under the OhAuthCat SDK Licence (see `LICENSE`); third-party notices are in `THIRD_PARTY_LICENSES`.

```go
package example

import (
    "context"
    ohauthcat "ohauthcat.com/sdk-go"
)

func listEvents(ctx context.Context) error {
    client, err := ohauthcat.NewClient(ohauthcat.ClientConfig{APIKey: "sk_example"})
    if err != nil { return err }
    defer client.Close()
    for _, err := range client.Events.List(ctx, ohauthcat.ListEventsArgs{}) {
        if err != nil { return err }
    }
    return nil
}
```

Generated resource methods mirror public management and broker operations in
the pinned OpenAPI file. Every list response has a typed `<Operation>All`
iterator; `Events.List` and `Portal.CreateLink` provide convenience names.
Iterators preserve path/query inputs, paginate forwards, stop on consumer break,
and refuse missing/repeated cursors, empty pages with more data, and page-limit
exhaustion. Duplicate items may be returned by the service.

Constructor configuration is copied and validated. OhAuthCat-Version is omitted
unless an optional `1.x` API version is configured. Requests have a total context
deadline and bounded response bodies. GET/HEAD and writes whose pinned contract
honors Idempotency-Key may retry bounded 429/503 responses with the same key.
Code exchange never retries: reusing a consumed code can revoke the earlier grant.
Custom HTTP transports and response bodies must honor contexts in Do/Read/Close.
The SDK does not create cancellation goroutines to hide a blocking custom body.
Close cancels in-flight work; caller-owned HTTP clients/transports are not closed.

SSO helpers are available on `client.SSO`: `GetAuthorizationURL`,
`GetProfileAndToken`, and `VerifyAccessToken`. Configure `ClientID` (`pk_`) for
authorization/exchange and `EnvironmentID` for local access-token verification.
Verification requires the exact configured issuer (default `https://auth.ohauthcat.com`)
while retrieving `/.well-known/jwks.json` from the active trusted broker origin,
with HTTPS except explicit loopback fixtures. Token-supplied key locations are ignored. No credentials, cookies or
redirects are forwarded to JWKS. ES256/P-256, issuer/audience/environment and
identity/time claims are verified with 60 seconds of clock skew. JWKS cache lifetime
is at most 300 seconds, respects no-store/no-cache, coalesces concurrent refreshes
and throttles unknown keys/failed refreshes for 30 seconds. Individual waiter
cancellation does not cancel a shared refresh; client Close does.

Sign-in: generate the authorization URL with a random `state` (left empty, the
SDK generates one and returns it), keep the state and the PKCE verifier in the
user's session, and on the callback compare the returned state in constant time
before exchanging the code. Refuse a missing or different state: it is a login
CSRF attempt or a stale tab.

```go
package example

import (
    "crypto/subtle"
    "net/http"

    ohauthcat "ohauthcat.com/sdk-go"
)

// Session is your server-side session store (signed cookie, Redis, ...).
type Session interface {
    Set(r *http.Request, w http.ResponseWriter, key, value string)
    Pop(r *http.Request, w http.ResponseWriter, key string) string
}

const redirectURI = "https://app.example.com/callback"

func signIn(client *ohauthcat.Client, sessions Session) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        auth, err := client.SSO.GetAuthorizationURL(ohauthcat.AuthorizationOptions{
            RedirectURI: redirectURI,
            Domain:      "customer.com",
        })
        if err != nil {
            http.Error(w, "sign-in unavailable", http.StatusInternalServerError)
            return
        }
        sessions.Set(r, w, "sso_state", auth.State)
        sessions.Set(r, w, "sso_verifier", auth.CodeVerifier)
        http.Redirect(w, r, auth.URL, http.StatusFound)
    }
}

func callback(client *ohauthcat.Client, sessions Session) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        expected := sessions.Pop(r, w, "sso_state")
        verifier := sessions.Pop(r, w, "sso_verifier")
        got := r.URL.Query().Get("state")
        if expected == "" || verifier == "" || subtle.ConstantTimeCompare([]byte(expected), []byte(got)) != 1 {
            http.Error(w, "sign-in expired, start again", http.StatusBadRequest)
            return
        }
        token, err := client.SSO.GetProfileAndToken(r.Context(), ohauthcat.ExchangeOptions{
            Code:         r.URL.Query().Get("code"),
            CodeVerifier: verifier,
            RedirectURI:  redirectURI,
        })
        if err != nil {
            http.Error(w, "sign-in failed, start again", http.StatusBadRequest)
            return
        }
        _ = token.Profile.Email // create your session for this user
        http.Redirect(w, r, "/", http.StatusFound)
    }
}
```

Guests: a sign-in admitted under the organization's guest policy (an email domain outside its verified domains) has `Profile.IsGuest()` true (`AccessPolicy` is `AccessPolicyGuest`); for every other sign-in the field is absent. `email_domain_verified` stays true for guests, so check this field when guests must get less access.

Portal links:

A setup link lasts `expires_in` seconds, 300 (5 minutes) to 2,592,000 (30 days); leave it out for the default of 604,800 (7 days). A link created with a secret key opens once; revoke it to end it early, and `expires_at` on the result says when it ends. `ttl` (60 to 300 seconds) is only for sealed single-use embed links; never send both.

```go
package example

import (
    "context"

    ohauthcat "ohauthcat.com/sdk-go"
)

func setupLink(ctx context.Context, client *ohauthcat.Client) (string, error) {
    day := int64(24 * 60 * 60)
    link, err := client.Portal.CreateLink(ctx, ohauthcat.CreatePortalLinkArgs{Body: ohauthcat.PortalLinkCreate{
        Organization: "org_01J9Z3K5V7W8X9Y0Z1A2B3C4D5",
        Intents:      []ohauthcat.Intent{"sso"},
        ExpiresIn:    &day,
    }})
    if err != nil {
        return "", err
    }
    return link.URL, nil
}
```

`RestartSignIn` retains API error context while asking for a fresh sign-in flow.
API errors preserve typed catalogue codes/status/context; their Error strings
omit response messages, URLs and credentials.

Webhooks: verify the exact raw request body before parsing it.

```go
package example

import (
    "errors"
    "io"
    "log"
    "net/http"
    "os"

    ohauthcat "ohauthcat.com/sdk-go"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
    body, err := io.ReadAll(io.LimitReader(r.Body, 1<<20)) // the exact raw bytes
    if err != nil {
        http.Error(w, "read failed", http.StatusBadRequest)
        return
    }
    secrets := []string{os.Getenv("OHAUTHCAT_WEBHOOK_SECRET")}
    event, err := ohauthcat.VerifyWebhook(body, r.Header.Get("OhAuthCat-Signature"), secrets, ohauthcat.VerifyOptions{})
    var verr *ohauthcat.WebhookVerificationError
    if errors.As(err, &verr) {
        http.Error(w, verr.Code, http.StatusBadRequest)
        return
    }
    if err != nil {
        http.Error(w, "verification failed", http.StatusInternalServerError)
        return
    }
    log.Println(event.Type, event.ID)
    w.WriteHeader(http.StatusNoContent)
}
```

`VerifySignature` authenticates raw webhook bytes without decoding them.
`VerifyWebhook` (also available on Client) additionally validates the generated
Event envelope. Supply all held rotation secrets exactly, including their prefix.
Default freshness tolerance is 300 seconds; zero is rejected unless
DisableTolerance is explicitly true. The current pinned Event has no per-event
payload union, so strict envelope verification does not validate per-type data.

Profile and DirectoryUser expose `Role *string`: nil means no role. Their codecs
preserve an explicit decoded null separately from an absent role.
`CustomAttributes` holds typed string, string-array, bool and int64 values;
unsupported entries are omitted without failing sign-in. Signed 64-bit bounds
round-trip exactly. For example:

```go
func applicationRole(profile ohauthcat.Profile) string {
    if profile.Role == nil { return "" }
    return *profile.Role
}

func mappedCount(profile ohauthcat.Profile) (int64, bool) {
    if profile.CustomAttributes == nil { return 0, false }
    value, exists := (*profile.CustomAttributes)["count"]
    if !exists || value.Int == nil { return 0, false }
    return *value.Int, true
}
```

Directory user created, updated and deleted events apply the same custom-value
bounds only when their data object identifies itself as `directory_user`; other
event data and WorkOS custom maps remain opaque. Webhook signatures authenticate
the original bytes before this normalization.

Other generated nullable fields use concrete wrappers with Set/Null/Value: Set=false
means absent, Set=true/Null=true means explicit null, and Set=true/Null=false
holds Value. Optional nonnullable fields use pointers. AdditionalFields preserves
future JSON properties; known-name collisions and duplicate JSON keys are refused.
JSON is bounded to 1 MiB and 64 nesting levels. Request schemas remain strict;
responses tolerate new multi-value enum members while retaining singleton
discriminators and structural constraints.

Automatic trusted standby failover follows the shared policy in the repository README. Actual live Go/core sign-in and
hosted test-IdP acceptance have not been executed. Unit/contract proofs use real
standard-library cryptography and independently parsed pinned fixtures. CI
execution is unverified; cross-language tests require Node 24, Python with venv,
and the declared Go toolchain. The complete Go step is a custom/manual pipeline:
its placeholder official Go image lacks Node 24. Supply a verified combined image
before enabling automated PR checks; the step refuses missing prerequisites.

Migration helpers bind to one import through `client.Migration(MigrationConfig{ImportID: ..., LegacyProvider: "workos", ResolveLegacyProfile: resolver})`. The helper's `Routing` points to its routing methods: `ForOrganization(ctx, org)` returns `legacy` or `ohauthcat`; `Invalidate()` immediately discards cached routing after switching or rolling back. `CacheTTL` defaults to 30 seconds and caps at 60; use `DisableCache` to bypass reuse. `MigrateSession(ctx, token, idempotencyKey)` fetches a fresh import, invokes your context-aware UserInfo resolver and sends only the generated import/profile-ID body. The legacy token stays in the app. Resolvers must honor the supplied context. Generated `Profile.LegacyIds` carries identity continuity; callers with an already authenticated legacy ID can use `client.Sessions.MigrateSession` directly.

Automatic trusted edge failover and request-driven dual health probes follow the shared repository README policy. `Standby` pointing to `false` disables it; custom primary origins require both explicit trusted fallback origins to opt in.

`Client.UploadLogo` and `Client.UploadLogoFile` upload a branding logo from bytes
or a file path. Before any request they refuse a file that is empty, larger than
256 KiB (`MaxLogoBytes`), or neither a PNG (by its signature) nor an SVG (by its
root element), returning an `*APIError` with `CodeLogoInvalid`, status 0 and
param `data`; `CheckLogo` runs the same check alone. Accepted bytes are
base64-encoded and sent to `Branding.UploadBrandingLogo`. The API still validates
dimensions and SVG content; its `logo_invalid` refusal arrives as the same typed
error with the HTTP status.

```go
package example

import (
    "context"
    ohauthcat "ohauthcat.com/sdk-go"
)

func uploadLogo(ctx context.Context, client *ohauthcat.Client) (string, error) {
    logo, err := client.UploadLogoFile(ctx, "logo.svg", ohauthcat.UploadLogoOptions{IdempotencyKey: "logo-2026-10"})
    if err != nil { return "", err }
    return logo.LogoURL, nil
}
```

