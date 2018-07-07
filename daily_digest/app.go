package daily_digest

import (
	context "context"
	"errors"
	"net/http"
	"strings"

	"github.com/vektah/gqlgen/handler"
)

type App struct {
	digestService       DigestService
	subscriptionService SubscriptionService
	entriesService      EntryService
	authService         AuthService
	accessDenied        error
}

type contextKey string

func NewApp() *App {
	return &App{
		digestService:       DigestService{},
		subscriptionService: SubscriptionService{},
		entriesService:      EntryService{},
		authService:         AuthService{},
		accessDenied:        errors.New("Access denied"),
	}
}

func (app *App) Seed() {
	app.entriesService.Seed()
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Setup response headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	// Add auth to context
	c := r.Context()
	sessionKey := r.Header.Get("Authorization")
	sessionKey = strings.Replace(sessionKey, "Bearer ", "", 1)
	user := app.authService.UserForSession(sessionKey)
	c = context.WithValue(c, contextKey("user"), user)
	r = r.WithContext(c)

	// Handle request
	handler.GraphQL(MakeExecutableSchema(app))(w, r)
}

func (app *App) Query_sources(ctx context.Context) ([]Source, error) {
	sources := make([]Source, 0)
	for _, source := range Get() {
		sources = append(sources, *source)
	}

	return sources, nil
}

func (app *App) Mutation_newSubscription(ctx context.Context, input *NewSubscriptionInput) (Subscription, error) {
	user := currentUser(ctx)
	if user == nil {
		return Subscription{}, app.accessDenied
	}

	source := GetById(*input.SourceId)
	if source == nil {
		return Subscription{}, errors.New("Invalid sourceId")
	}
	return *app.subscriptionService.Create(user, source), nil
}

func (app *App) Query_subscriptions(ctx context.Context) ([]Subscription, error) {
	user := currentUser(ctx)
	if user == nil {
		return []Subscription{}, app.accessDenied
	}
	return app.subscriptionService.ListAll(user), nil
}

func (app *App) Query_digests(ctx context.Context, date string) ([]Digest, error) {
	user := currentUser(ctx)
	if user == nil {
		return []Digest{}, app.accessDenied
	}
	return app.digestService.Get(user, date, &app.subscriptionService), nil
}

func (app *App) Mutation_register(ctx context.Context, input RegisterInput) (*string, error) {
	app.authService.Register(input.Username, input.Password)
	return app.authService.Login(input.Username, input.Password)
}

func (app *App) Mutation_login(ctx context.Context, input LoginInput) (*string, error) {
	return app.authService.Login(input.Username, input.Password)
}

func currentUser(c context.Context) *User {
	user, ok := c.Value(contextKey("user")).(*User)
	if ok {
		return user
	}
	return nil
}
