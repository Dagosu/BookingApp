package domain

import (
	"context"

	dt "github.com/Dagosu/BookingApp/datatypes"
)

// UserFlightsMappingUsecase defines the user usecase interface.
type UserFlightsMappingUsecase interface {
	PurchaseFlight(ctx context.Context, userId, flightId string) (*dt.Flight, error)
}

// UserFlightsMappingRepository defines the fluseright repository interface.
type UserFlightsMappingRepository interface {
	PurchaseFlight(ctx context.Context, email, password string) (*dt.User, error)
}
