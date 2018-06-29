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
var dailiesService server.DailiesService
var entriesService server.EntryService

func init() {
	digestService = server.DigestService{}
	dailiesService = server.DailiesService{}
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

func (app *App) Mutation_newDigest(ctx context.Context, input *server.NewDigestInput) (server.Digest, error) {
	source := server.GetById(*input.SourceId)
	if source == nil {
		return server.Digest{}, errors.New("Invalid sourceId")
	}
	return *digestService.Create(source), nil
}

func (app *App) Query_digests(ctx context.Context) ([]server.Digest, error) {
	digests := digestService.ListAll()
	return digests, nil
}

func (app *App) Query_daily(ctx context.Context, date *string) (server.Daily, error) {
	return dailiesService.Get(*date, &digestService), nil
}
