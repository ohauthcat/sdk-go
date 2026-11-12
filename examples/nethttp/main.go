// Command nethttp demonstrates process-local sessions and verified webhook handoff.
package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	sdk "ohauthcat.com/sdk-go"
	"ohauthcat.com/sdk-go/frameworkhttp"
)

func run(ctx context.Context, getenv func(string) string, serve func(context.Context, *http.Server) error) error {
	if ctx == nil || getenv == nil || serve == nil {
		return errors.New("example configuration invalid")
	}
	client, err := sdk.NewClient(sdk.ClientConfig{APIKey: getenv("OHAUTHCAT_API_KEY"), ClientID: getenv("OHAUTHCAT_CLIENT_ID"), EnvironmentID: getenv("OHAUTHCAT_ENVIRONMENT_ID")})
	if err != nil {
		return errors.New("SDK configuration invalid")
	}
	defer func() { _ = client.Close() }()
	events := make(chan sdk.Event, 64)
	helper, helperErr := frameworkhttp.New(frameworkhttp.Options{SSO: client.SSO, Environment: getenv("OHAUTHCAT_ENVIRONMENT_ID"), Organization: getenv("OHAUTHCAT_ORGANIZATION_ID"), Connection: getenv("OHAUTHCAT_CONNECTION_ID"), RedirectURI: getenv("OHAUTHCAT_REDIRECT_URI"), WebhookSecrets: []string{getenv("OHAUTHCAT_WEBHOOK_SECRET")}, Queue: func(ctx context.Context, event sdk.Event) error {
		select {
		case events <- event:
			return nil
		case <-ctx.Done():
			return ctx.Err()
		default:
			return errors.New("queue full")
		}
	}})
	if helperErr != nil {
		return errors.New("framework configuration invalid")
	}
	defer func() { _ = helper.Close() }()
	// Replace this demonstration drain with an application-owned durable queue consumer.
	workersDone := make(chan struct{})
	workerCtx, cancelWorker := context.WithCancel(ctx)
	defer func() { cancelWorker(); <-workersDone }()
	go func() {
		defer close(workersDone)
		for {
			select {
			case <-workerCtx.Done():
				return
			case <-events:
			}
		}
	}()
	mux := routes(helper)
	server := &http.Server{Addr: "127.0.0.1:8080", Handler: mux, ReadHeaderTimeout: 5 * time.Second, ReadTimeout: 30 * time.Second, WriteTimeout: 30 * time.Second, IdleTimeout: time.Minute}
	return serve(ctx, server)
}
func routes(helper *frameworkhttp.Helper) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/start", helper.Start)
	mux.HandleFunc("/callback", helper.Callback)
	mux.HandleFunc("/profile", helper.Profile)
	mux.HandleFunc("/webhook", helper.Webhook)
	return mux
}
func serveHTTP(ctx context.Context, server *http.Server) error {
	finished := make(chan error, 1)
	go func() { finished <- server.ListenAndServe() }()
	select {
	case err := <-finished:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return errors.New("HTTP server unavailable")
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		shutdownErr := server.Shutdown(shutdownCtx)
		if shutdownErr != nil {
			_ = server.Close()
		}
		<-finished
		return shutdownErr
	}
}
func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	err := run(ctx, os.Getenv, serveHTTP)
	cancel()
	if err != nil {
		os.Exit(1)
	}
}
