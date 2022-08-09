package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/xadamy/coffeeshops/graphql-server/graph/generated"
	"github.com/xadamy/coffeeshops/graphql-server/graph/model"
)

// Coffeeshops is the resolver for the coffeeshops field.
func (r *queryResolver) Coffeeshops(ctx context.Context) ([]*model.Coffeeshop, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
