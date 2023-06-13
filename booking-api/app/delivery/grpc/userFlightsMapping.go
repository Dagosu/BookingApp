package grpc

import (
	"context"

	"github.com/Dagosu/BookingApp/booking-api/app/domain"
	dt "github.com/Dagosu/BookingApp/datatypes"
)

type userFlightsMappingServiceServer struct {
	uu domain.UserFlightsMappingUsecase
}

var _ dt.UserFlightsMappingServiceServer = &userFlightsMappingServiceServer{}

func newUserFlightsMappingServiceServer(ufmu domain.UserFlightsMappingUsecase) dt.UserFlightsMappingServiceServer {
	return &userFlightsMappingServiceServer{ufmu}
}

func (ufms *userFlightsMappingServiceServer) PurchaseFlight(ctx context.Context, req *dt.PurchaseFlightRequest) (*dt.PurchaseFlightResponse, error) {
	purchasedFlight, err := ufms.uu.PurchaseFlight(ctx, req.GetUserId(), req.GetFlightId())
	if err != nil {
		return nil, err
	}

	return &dt.PurchaseFlightResponse{
		PurchasedFlight: purchasedFlight,
	}, nil
}

func (ufms *userFlightsMappingServiceServer) FavoriteFlight(ctx context.Context, req *dt.FavoriteFlightRequest) (*dt.FavoriteFlightResponse, error) {
	favoritedFlight, err := ufms.uu.FavoriteFlight(ctx, req.GetUserId(), req.GetFlightId())
	if err != nil {
		return nil, err
	}

	return &dt.FavoriteFlightResponse{
		FavoritedFlight: favoritedFlight,
	}, nil
}

func (ufms *userFlightsMappingServiceServer) GetPurchasedFlights(ctx context.Context, req *dt.GetPurchasedFlightsRequest) (*dt.GetPurchasedFlightsResponse, error) {
	flights, err := ufms.uu.GetPurchasedFlights(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}

	return &dt.GetPurchasedFlightsResponse{
		Flights: flights,
	}, nil
}

func (ufms *userFlightsMappingServiceServer) GetFavoritedFlights(ctx context.Context, req *dt.GetFavoritedFlightsRequest) (*dt.GetFavoritedFlightsResponse, error) {
	flights, err := ufms.uu.GetFavoritedFlights(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}

	return &dt.GetFavoritedFlightsResponse{
		Flights: flights,
	}, nil
}

func (ufms *userFlightsMappingServiceServer) RecommendFlight(ctx context.Context, req *dt.RecommendFlightRequest) (*dt.RecommendFlightResponse, error) {
	flights, err := ufms.uu.RecommendFlight(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}

	return &dt.RecommendFlightResponse{
		Flights: flights,
	}, nil
}
