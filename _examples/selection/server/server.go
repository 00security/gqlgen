package main

import (
	"log"
	"net/http"

	"github.com/00security/gqlgen/example/selection"
	"github.com/00security/gqlgen/graphql/handler"
	"github.com/00security/gqlgen/graphql/playground"
)

func main() {
	http.Handle("/", playground.Handler("Selection Demo", "/query"))
	http.Handle("/query", handler.NewDefaultServer(selection.NewExecutableSchema(selection.Config{Resolvers: &selection.Resolver{}})))
	log.Fatal(http.ListenAndServe(":8086", nil))
}
