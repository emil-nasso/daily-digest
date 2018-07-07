package server

import (
	context "context"
	"errors"
	"net/http"
	"strings"

	"github.com/emil-nasso/daily-digest/daily_digest"
	"github.com/vektah/gqlgen/handler"
)

type App struct {
	accessDenied error
}

type contextKey string

func NewApp() *App {
	return &App{
		accessDenied: errors.New("Access denied"),
	}
}

func (app *App) Seed() {
	daily_digest.SeedEntries()
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Setup response headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	// Add auth to context
	c := r.Context()
	sessionKey := r.Header.Get("Authorization")
	sessionKey = strings.Replace(sessionKey, "Bearer ", "", 1)
	user := daily_digest.GetUserForSession(sessionKey)
	c = context.WithValue(c, contextKey("user"), user)
	r = r.WithContext(c)

	// Handle request
	handler.GraphQL(daily_digest.MakeExecutableSchema(app))(w, r)
	daily_digest.DebugUsers()
}

func (app *App) Query_sources(ctx context.Context) ([]daily_digest.Source, error) {
	sources := make([]daily_digest.Source, 0)
	for _, source := range daily_digest.GetSources() {
		sources = append(sources, *source)
	}

	return sources, nil
}

func (app *App) Mutation_newSubscription(ctx context.Context, input *daily_digest.NewSubscriptionInput) (daily_digest.Subscription, error) {
	user := currentUser(ctx)
	if user == nil {
		return daily_digest.Subscription{}, app.accessDenied
	}

	source := daily_digest.GetSourceById(*input.SourceId)
	if source == nil {
		return daily_digest.Subscription{}, errors.New("Invalid sourceId")
	}
	return *daily_digest.CreateSubscription(user, source), nil
}

func (app *App) Query_subscriptions(ctx context.Context) ([]daily_digest.Subscription, error) {
	user := currentUser(ctx)
	if user == nil {
		return []daily_digest.Subscription{}, app.accessDenied
	}
	return daily_digest.ListAllSubscriptions(user), nil
}

func (app *App) Query_digests(ctx context.Context, date string) ([]daily_digest.Digest, error) {
	user := currentUser(ctx)
	if user == nil {
		return []daily_digest.Digest{}, app.accessDenied
	}
	return daily_digest.GetDigest(user, date), nil
}

func (app *App) Mutation_register(ctx context.Context, input daily_digest.RegisterInput) (*string, error) {
	daily_digest.RegisterUser(input.Username, input.Password)
	return daily_digest.Login(input.Username, input.Password)
}

func (app *App) Mutation_login(ctx context.Context, input daily_digest.LoginInput) (*string, error) {
	return daily_digest.Login(input.Username, input.Password)
}

func currentUser(c context.Context) *daily_digest.User {
	user, ok := c.Value(contextKey("user")).(*daily_digest.User)
	if ok {
		return user
	}
	return nil
}
