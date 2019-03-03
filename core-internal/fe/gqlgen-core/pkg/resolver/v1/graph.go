package graph

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/mvp-eXpress/3g-todo-fullstack/core-internal/fe/gqlgen-core/grpc_client"
)

type Server struct {
	grpcClient *grpc_client.Client
}

func NewGraphQLServer(grpcURL string) (*Server, error) {
	grpcClient, err := grpc_client.NewClient(grpcURL)
	if err != nil {
		return nil, err
	}
	return &Server{grpcClient}, nil

}

func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: s,
	})
}
