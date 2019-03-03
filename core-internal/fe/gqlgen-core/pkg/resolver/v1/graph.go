package graph

import (
	"log"

	"github.com/99designs/gqlgen/graphql"
	gqlgen_core "github.com/mvp-eXpress/3g-todo-fullstack/core-internal/fe/gqlgen-core"
	"github.com/mvp-eXpress/3g-todo-fullstack/core-internal/fe/gqlgen-core/grpc_client"
	"google.golang.org/grpc"
)

type gc struct {
	conn *grpc.ClientConn
}
type Server struct {
	grpcClient *grpc_client.Client
}

func NewGraphQLServer(grpcURL string) (*Server, error) {
	log.Printf("NewGraphQLServer::grpcURL: %v", grpcURL)
	grpcClient, err := grpc_client.NewClient(grpcURL)
	if err != nil {
		return nil, err
	}
	return &Server{grpcClient}, nil

}

func (s *Server) Mutation() gqlgen_core.MutationResolver {
	return &mutationResolver{
		server: s,
	}
}

func (s *Server) Query() gqlgen_core.QueryResolver {
	return &queryResolver{
		server: s,
	}
}

func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
	return gqlgen_core.NewExecutableSchema(gqlgen_core.Config{
		Resolvers: s,
	})
}
