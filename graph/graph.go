package graph

import (
	context "context"
	"net/http"

	sources "github.com/emil-nasso/daily-digest/sources"
	"github.com/vektah/gqlgen/handler"
)

type User struct {
	Id       string
	Username string
	Password string
	Sessions []*Session
}

type Session struct {
	Id   string
	Key  string
	User *User
}

type App struct {
	users    []*User
	sessions []*Session
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	handler := handler.GraphQL(MakeExecutableSchema(app))
	handler(w, r)
}

func (app *App) Query_sources(context.Context) ([]sources.Source, error) {
	return sources.Get(), nil
}
