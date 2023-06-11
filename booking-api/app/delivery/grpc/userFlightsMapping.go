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
	// purchasedFlight, err := ufms.uu.PurchaseFlight(ctx, req.GetUserId(), req.GetFlightId())
	// if err != nil {
	// 	return nil, err
	// }

	return &dt.PurchaseFlightResponse{
		PurchasedFlight: nil,
	}, nil
}

func (ufms *userFlightsMappingServiceServer) FavoriteFlight(ctx context.Context, req *dt.FavoriteFlightRequest) (*dt.FavoriteFlightResponse, error) {
	// purchasedFlight, err := ufms.uu.PurchaseFlight(ctx, req.GetUserId(), req.GetFlightId())
	// if err != nil {
	// 	return nil, err
	// }

	return &dt.FavoriteFlightResponse{
		FavoritedFlight: nil,
	}, nil
}
