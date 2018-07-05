package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	_ "github.com/emil-nasso/daily-digest/plugins"

	"github.com/emil-nasso/daily-digest/server"
	"github.com/vektah/gqlgen/handler"
)

var app *App

var digestService server.DigestService
var subscriptionService server.SubscriptionService
var entriesService server.EntryService

func init() {
	digestService = server.DigestService{}
	subscriptionService = server.SubscriptionService{}
	entriesService = server.EntryService{}

	app = &App{}
}

func main() {
	entriesService.Seed()

	http.Handle("/", handler.Playground("Daily-Digest", "/graphql"))
	http.Handle("/graphql", app)
	fmt.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type App struct {
	users    []*server.User
	sessions []*server.Session
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	handler := handler.GraphQL(server.MakeExecutableSchema(app))
	handler(w, r)
}

func (app *App) Query_sources(context.Context) ([]server.Source, error) {
	sources := make([]server.Source, 0)
	for _, source := range server.Get() {
		sources = append(sources, *source)
	}

	return sources, nil
}

func (app *App) Mutation_newSubscription(ctx context.Context, input *server.NewSubscriptionInput) (server.Subscription, error) {
	source := server.GetById(*input.SourceId)
	if source == nil {
		return server.Subscription{}, errors.New("Invalid sourceId")
	}
	return *subscriptionService.Create(source), nil
}

func (app *App) Query_subscriptions(ctx context.Context) ([]server.Subscription, error) {
	return subscriptionService.ListAll(), nil
}

func (app *App) Query_digests(ctx context.Context, date string) ([]server.Digest, error) {
	return digestService.Get(date, &subscriptionService), nil
}
