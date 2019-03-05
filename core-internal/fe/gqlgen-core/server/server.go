package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/kelseyhightower/envconfig"
	graph "github.com/mvp-eXpress/3g-todo-fullstack/core-internal/fe/gqlgen-core/pkg/resolver/v1"
	"github.com/rs/cors"
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

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	router.Handle("/playground", handler.Playground("GraphQL playground", "/graphql"))
	router.Handle("/graphql", handler.GraphQL(s.ToExecutableSchema()))

	log.Fatal(http.ListenAndServe(":8080", router))
}
