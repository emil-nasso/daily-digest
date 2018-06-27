package graph

import (
	context "context"
	"net/http"

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

func (app *App) Mutation_signIn(ctx context.Context, username string) (*Session, error) {
	return nil, nil
}

func (app *App) Mutation_register(ctx context.Context, username string, password string) (*Session, error) {
	return nil, nil
}
func (app *App) Query_users(ctx context.Context, input *UsersInput) ([]User, error) {
	return nil, nil
}

func (app *App) User_secretField(ctx context.Context, obj *User) ([]string, error) {
	return nil, nil
}
