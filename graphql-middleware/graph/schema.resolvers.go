package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"

	graph "github.com/Dagosu/BookingApp/graphql-middleware/graph/generated"
	"github.com/Dagosu/BookingApp/graphql-middleware/graph/model"
)

// TestEndpoint is the resolver for the testEndpoint field.
func (r *queryResolver) TestEndpoint(ctx context.Context, in model.TestEndpointInput) (*model.TestEndpointResponse, error) {
	return r.resolveTestEndpoint(ctx, in)
}

// TestList is the resolver for the testList field.
func (r *subscriptionResolver) TestList(ctx context.Context, in model.TestListInput) (<-chan *model.TestListResponse, error) {
	panic(fmt.Errorf("not implemented: TestList - testList"))
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

// Subscription returns graph.SubscriptionResolver implementation.
func (r *Resolver) Subscription() graph.SubscriptionResolver { return &subscriptionResolver{r} }

type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
