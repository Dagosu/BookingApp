package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"

	graph "github.com/Dagosu/BookingApp/graphql-middleware/graph/generated"
	"github.com/Dagosu/BookingApp/graphql-middleware/graph/model"
)

// PurchaseFlight is the resolver for the purchaseFlight field.
func (r *mutationResolver) PurchaseFlight(ctx context.Context, in model.PurchaseFlightInput) (*model.PurchaseFlightResponse, error) {
	return r.resolvePurchaseFlight(ctx, in)
}

// FavoriteFlight is the resolver for the favoriteFlight field.
func (r *mutationResolver) FavoriteFlight(ctx context.Context, in model.FavoriteFlightInput) (*model.FavoriteFlightResponse, error) {
	return r.resolveFavoriteFlight(ctx, in)
}

// WriteReview is the resolver for the writeReview field.
func (r *mutationResolver) WriteReview(ctx context.Context, in model.WriteReviewInput) (*model.WriteReviewResponse, error) {
	return r.resolveWriteReview(ctx, in)
}

// CheckCredentials is the resolver for the checkCredentials field.
func (r *queryResolver) CheckCredentials(ctx context.Context, in model.CheckCredentialsInput) (*model.CheckCredentialsResponse, error) {
	return r.resolveCheckCredentials(ctx, in)
}

// GetFlight is the resolver for the getFlight field.
func (r *queryResolver) GetFlight(ctx context.Context, in model.GetFlightInput) (*model.GetFlightResponse, error) {
	return r.resolveGetFlight(ctx, in)
}

// GetPurchasedFlights is the resolver for the getPurchasedFlights field.
func (r *queryResolver) GetPurchasedFlights(ctx context.Context, in model.GetPurchasedFlightsInput) (*model.GetPurchasedFlightsResponse, error) {
	return r.resolveGetPurchasedFlights(ctx, in)
}

// GetFavoritedFlights is the resolver for the getFavoritedFlights field.
func (r *queryResolver) GetFavoritedFlights(ctx context.Context, in model.GetFavoritedFlightsInput) (*model.GetFavoritedFlightsResponse, error) {
	return r.resolveGetFavoritedFlights(ctx, in)
}

// RecommendFlight is the resolver for the recommendFlight field.
func (r *queryResolver) RecommendFlight(ctx context.Context, in model.RecommendFlightInput) (*model.RecommendFlightResponse, error) {
	return r.resolveRecommendFlight(ctx, in)
}

// CheckFlightPurchase is the resolver for the checkFlightPurchase field.
func (r *queryResolver) CheckFlightPurchase(ctx context.Context, in model.CheckFlightPurchaseInput) (*model.CheckFlightPurchaseResponse, error) {
	return r.resolveCheckFlightPurchase(ctx, in)
}

// FlightList is the resolver for the flightList field.
func (r *subscriptionResolver) FlightList(ctx context.Context, in model.FlightListInput) (<-chan *model.FlightListResponse, error) {
	return r.resolveFlightList(ctx, in)
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

// Subscription returns graph.SubscriptionResolver implementation.
func (r *Resolver) Subscription() graph.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
