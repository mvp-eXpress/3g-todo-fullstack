package graph

import (
	"context"
	"log"
	"time"

	gqlgen_core "github.com/mvp-eXpress/3g-todo-fullstack/core-internal/fe/gqlgen-core"
)

type mutationResolver struct {
	server *Server
}

func (r *mutationResolver) CreateCollection(ctx context.Context, in gqlgen_core.CreateCollectionInput) (*gqlgen_core.Collection, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	res, err := r.server.grpcClient.CreateCollection(ctx, in)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

// type Mutation interface {
// 	Close()
// 	getCollections(ctx context.Context) ([]*gqlgen_core.Collection, error)
// }

// func NewGrpcClient(url string) (Mutation, error) {
// 	conn, err := grpc.Dial(url, grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}
// 	return &gc{conn}, nil
// }

// func (g *gc) Close() {
// 	g.Close()

// }
