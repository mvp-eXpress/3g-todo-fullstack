//go:generate go run scripts/gqlgen.go -v

package gqlgen_core

import (
	"context"
)

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateCollection(ctx context.Context, input CreateCollectionInput) (*Collection, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetCollections(ctx context.Context) ([]*Collection, error) {
	panic("not implemented")
}
