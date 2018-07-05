package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/emil-nasso/daily-digest/daily_digest"
	_ "github.com/emil-nasso/daily-digest/plugins"
	"github.com/vektah/gqlgen/handler"
)

func main() {
	app := daily_digest.NewApp()
	app.Seed()

	http.Handle("/", handler.Playground("Daily-Digest", "/graphql"))
	http.Handle("/graphql", app)
	fmt.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
