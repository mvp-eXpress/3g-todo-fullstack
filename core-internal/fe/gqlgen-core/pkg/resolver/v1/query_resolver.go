package graph

import (
	"context"
	"log"
	"time"

	gqlgen_core "github.com/mvp-eXpress/3g-todo-fullstack/core-internal/fe/gqlgen-core"
)

type queryResolver struct {
	server *Server
}

func (r *queryResolver) GetCollections(ctx context.Context) ([]*gqlgen_core.Collection, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	res, err := r.server.grpcClient.GetCollections(ctx)
	if err != nil {
		log.Printf("graph::GetCollections::err: %v", err)
		return nil, err
	}
	return res, nil
}
