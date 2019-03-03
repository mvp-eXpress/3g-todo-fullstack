package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/kelseyhightower/envconfig"
	graph "github.com/mvp-eXpress/3g-todo-fullstack/core-internal/fe/gqlgen-core/pkg/resolver/v1"
	// "github.com/tinrab/spidey/graphql/graph"
)

const defaultPort = "8080"

type Config struct {
	GrpcServiceuRL string `envconfig:"SERVICE_URL"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	s, err := graph.NewGraphQLServer(cfg.GrpcServiceuRL)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/playground", handler.Playground("GraphQL playground", "/graphql"))
	http.Handle("/graphql", handler.GraphQL(s.ToExecutableSchema()))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// func main() {
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = defaultPort
// 	}

// 	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
// 	http.Handle("/query", handler.GraphQL(gqlgen_core.NewExecutableSchema(gqlgen_core.Config{Resolvers: &gqlgen_core.Resolver{}})))

// 	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
// 	log.Fatal(http.ListenAndServe(":"+port, nil))
// }
