// Command existingauth is a runnable recipe: it links a profile JSON file
// (the output of GET /v1/sso/profile) to a user in an existing auth system
// and starts a session there. Configuration comes from environment variables;
// see examples/recipes/existing-auth/ for each provider.
//
//	go run ./examples/existingauth <auth0|cognito|firebase|supabase|clerk> profile.json
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	ohauthcat "ohauthcat.com/sdk-go"
	"ohauthcat.com/sdk-go/existingauth"
)

func main() {
	if err := run(context.Background(), os.Args[1:]); err != nil {
		var typed *existingauth.Error
		if errors.As(err, &typed) {
			fmt.Fprintf(os.Stderr, "refused: %s (%s, HTTP %d)\n", typed.Code, typed.Provider, typed.Status)
		} else {
			fmt.Fprintln(os.Stderr, "error:", err)
		}
		os.Exit(1)
	}
}

func env(name string) string { return os.Getenv(name) }

func policy() existingauth.Policy {
	if env("LINK_ON_MISSING") == "create" {
		return existingauth.Policy{OnMissing: existingauth.OnMissingCreate}
	}
	return existingauth.Policy{OnMissing: existingauth.OnMissingRefuse}
}

func run(ctx context.Context, args []string) error {
	if len(args) != 2 {
		return errors.New("usage: existingauth <auth0|cognito|firebase|supabase|clerk> profile.json")
	}
	raw, err := os.ReadFile(args[1]) //nolint:gosec // The operator names the profile file to read.
	if err != nil {
		return err
	}
	var profile ohauthcat.Profile
	if err := json.Unmarshal(raw, &profile); err != nil {
		return fmt.Errorf("profile.json is not a profile: %w", err)
	}
	in := existingauth.FromProfile(profile)
	switch args[0] {
	case "auth0":
		l, err := existingauth.NewAuth0Linker(existingauth.Auth0Options{
			Domain: env("AUTH0_DOMAIN"), Connection: env("AUTH0_CONNECTION"),
			ClientID: env("AUTH0_CLIENT_ID"), ClientSecret: env("AUTH0_CLIENT_SECRET"),
		})
		if err != nil {
			return err
		}
		res, err := existingauth.Link(ctx, in, l, policy())
		if err != nil {
			return err
		}
		fmt.Printf("auth0 user %s (created %t); start your application session for it\n", res.User.UserID, res.Created)
	case "cognito":
		tr, err := existingauth.SigV4Transport(existingauth.SigV4Options{
			Region: env("AWS_REGION"),
			Credentials: existingauth.AWSCredentials{
				AccessKeyID: env("AWS_ACCESS_KEY_ID"), SecretAccessKey: env("AWS_SECRET_ACCESS_KEY"), SessionToken: env("AWS_SESSION_TOKEN"),
			},
		})
		if err != nil {
			return err
		}
		l, err := existingauth.NewCognitoLinker(existingauth.CognitoOptions{
			UserPoolID: env("COGNITO_USER_POOL_ID"), ClientID: env("COGNITO_CLIENT_ID"), ClientSecret: env("COGNITO_CLIENT_SECRET"), Transport: tr,
		})
		if err != nil {
			return err
		}
		res, err := existingauth.Link(ctx, in, l, policy())
		if err != nil {
			return err
		}
		session, err := l.StartSession(ctx, res.User, env("COGNITO_CHALLENGE_ANSWER"))
		if err != nil {
			return err
		}
		fmt.Printf("cognito user %s (created %t); tokens issued, expires in %d s\n", res.User.Username, res.Created, session.ExpiresIn)
	case "supabase":
		l, err := existingauth.NewSupabaseLinker(existingauth.SupabaseOptions{URL: env("SUPABASE_URL"), ServiceRoleKey: env("SUPABASE_SERVICE_ROLE_KEY")})
		if err != nil {
			return err
		}
		res, err := existingauth.Link(ctx, in, l, policy())
		if err != nil {
			return err
		}
		session, err := l.StartSession(ctx, res.User)
		if err != nil {
			return err
		}
		fmt.Printf("supabase user %s (created %t); session expires in %d s\n", res.User.ID, res.Created, session.ExpiresIn)
	case "clerk":
		l, err := existingauth.NewClerkLinker(existingauth.ClerkOptions{SecretKey: env("CLERK_SECRET_KEY")})
		if err != nil {
			return err
		}
		res, err := existingauth.Link(ctx, in, l, policy())
		if err != nil {
			return err
		}
		token, err := l.StartSession(ctx, res.User, 5*time.Minute)
		if err != nil {
			return err
		}
		fmt.Printf("clerk user %s (created %t); redeem the sign-in token with the ticket strategy (url set: %t)\n", res.User.ID, res.Created, token.URL != "")
	case "firebase":
		return errors.New("firebase needs the Firebase Admin SDK as the user store; see examples/recipes/existing-auth/firebase.md")
	default:
		return fmt.Errorf("unknown provider %q", args[0])
	}
	return nil
}
