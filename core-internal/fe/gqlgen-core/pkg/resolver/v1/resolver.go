// go:generate go run scripts/gqlgen.go -v

package graph

import (
	"context"

	gqlgen_core "github.com/mvp-eXpress/3g-todo-fullstack/core-internal/fe/gqlgen-core"
	// graph "github.com/mvp-eXpress/3g-todo-fullstack/core-internal/fe/gqlgen-core/pkg/resolver/v1"
)

type Resolver struct{}

func (r *Resolver) Mutation() gqlgen_core.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() gqlgen_core.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct {
	server *Server
}

func (r *mutationResolver) CreateCollection(ctx context.Context, input CreateCollectionInput) (*Collection, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetCollections(ctx context.Context) ([]*Collection, error) {
	gc := NewGrpcClient("adas")
	// panic("not implemented")
}
