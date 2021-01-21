package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Markogoodman/gqltest/graph"
	"github.com/Markogoodman/gqltest/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	c := generated.Config{Resolvers: &graph.Resolver{}}
	countComplexity := func(childComplexity, count int) int {
		fmt.Println("childComplexity: ", childComplexity, ", count: ", count)
		return count * childComplexity
	}
	c.Complexity.Todo.Related = countComplexity

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))
	srv.Use(extension.FixedComplexityLimit(5))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
