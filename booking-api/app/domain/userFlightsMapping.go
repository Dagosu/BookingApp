package domain

import (
	"context"

	dt "github.com/Dagosu/BookingApp/datatypes"
)

type UserFlightsMappingUsecase interface {
	PurchaseFlight(ctx context.Context, userId, flightId string) (*dt.Flight, error)
	FavoriteFlight(ctx context.Context, userId, flightId string) (*dt.Flight, error)
	GetPurchasedFlights(ctx context.Context, userId string) ([]*dt.Flight, error)
	GetFavoritedFlights(ctx context.Context, userId string) ([]*dt.Flight, error)
	RecommendFlight(ctx context.Context, userId string) ([]*dt.Flight, error)
}

type UserFlightsMappingRepository interface {
	Create(ctx context.Context, user *dt.User, flight *dt.Flight) (*dt.UserFlightsMapping, error)
	GetByUser(ctx context.Context, userId string) (*dt.UserFlightsMapping, error)
	AddPurchasedFlight(ctx context.Context, userFlightsId string, flight *dt.Flight) error
	AddFavoritedFlight(ctx context.Context, userFlightsId string, flight *dt.Flight) error
}
