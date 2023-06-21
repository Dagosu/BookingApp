package graph

import (
	"context"

	dt "github.com/Dagosu/BookingApp/datatypes"
	"github.com/Dagosu/BookingApp/graphql-middleware/graph/model"
	"github.com/Dagosu/BookingApp/graphql-middleware/graph/parser"
)

func (r *mutationResolver) resolvePurchaseFlight(ctx context.Context, in model.PurchaseFlightInput) (*model.PurchaseFlightResponse, error) {
	res, err := r.server.bookingClient.UserFlightsMappingService.PurchaseFlight(ctx, &dt.PurchaseFlightRequest{
		UserId:   parser.StrDerefer(in.UserID),
		FlightId: parser.StrDerefer(in.FlightID),
	})
	if err != nil {
		return nil, err
	}

	return &model.PurchaseFlightResponse{
		PurchasedFlight: parser.ParseFlight(res.GetPurchasedFlight()),
	}, nil
}

func (r *mutationResolver) resolveFavoriteFlight(ctx context.Context, in model.FavoriteFlightInput) (*model.FavoriteFlightResponse, error) {
	res, err := r.server.bookingClient.UserFlightsMappingService.FavoriteFlight(ctx, &dt.FavoriteFlightRequest{
		UserId:   parser.StrDerefer(in.UserID),
		FlightId: parser.StrDerefer(in.FlightID),
	})
	if err != nil {
		return nil, err
	}

	return &model.FavoriteFlightResponse{
		FavoritedFlight: parser.ParseFlight(res.GetFavoritedFlight()),
	}, nil
}

func (r *queryResolver) resolveGetPurchasedFlights(ctx context.Context, in model.GetPurchasedFlightsInput) (*model.GetPurchasedFlightsResponse, error) {
	res, err := r.server.bookingClient.UserFlightsMappingService.GetPurchasedFlights(ctx, &dt.GetPurchasedFlightsRequest{
		UserId: parser.StrDerefer(in.UserID),
	})
	if err != nil {
		return nil, err
	}

	flights := []*model.Flight{}
	for _, f := range res.Flights {
		flights = append(flights, parser.ParseFlight(f))
	}

	return &model.GetPurchasedFlightsResponse{
		Flights: flights,
	}, nil
}

func (r *queryResolver) resolveGetFavoritedFlights(ctx context.Context, in model.GetFavoritedFlightsInput) (*model.GetFavoritedFlightsResponse, error) {
	res, err := r.server.bookingClient.UserFlightsMappingService.GetFavoritedFlights(ctx, &dt.GetFavoritedFlightsRequest{
		UserId: parser.StrDerefer(in.UserID),
	})
	if err != nil {
		return nil, err
	}

	flights := []*model.Flight{}
	for _, f := range res.Flights {
		flights = append(flights, parser.ParseFlight(f))
	}

	return &model.GetFavoritedFlightsResponse{
		Flights: flights,
	}, nil
}

func (r *queryResolver) resolveRecommendFlight(ctx context.Context, in model.RecommendFlightInput) (*model.RecommendFlightResponse, error) {
	res, err := r.server.bookingClient.UserFlightsMappingService.RecommendFlight(ctx, &dt.RecommendFlightRequest{
		UserId: parser.StrDerefer(in.UserID),
	})
	if err != nil {
		return nil, err
	}

	flights := []*model.Flight{}
	for _, f := range res.Flights {
		flights = append(flights, parser.ParseFlight(f))
	}

	return &model.RecommendFlightResponse{
		Flights: flights,
	}, nil
}

func (r *queryResolver) resolveCheckFlightPurchase(ctx context.Context, in model.CheckFlightPurchaseInput) (*model.CheckFlightPurchaseResponse, error) {
	res, err := r.server.bookingClient.UserFlightsMappingService.CheckFlightPurchase(ctx, &dt.CheckFlightPurchaseRequest{
		FlightId: parser.StrDerefer(in.FlightID),
		UserId:   parser.StrDerefer(in.UserID),
	})
	if err != nil {
		return nil, err
	}

	return &model.CheckFlightPurchaseResponse{
		Flight: parser.ParseFlight(res.Flight),
	}, nil
}
