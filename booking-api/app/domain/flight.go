package domain

import (
	"context"

	dt "github.com/Dagosu/BookingApp/datatypes"
)

type FlightUsecase interface {
	FlightList(req *dt.FlightListRequest, stream dt.FlightService_FlightListServer) error
	GetFlight(ctx context.Context, id string) (*dt.Flight, error)
	WriteReview(ctx context.Context, flightId, userId, text string) (*dt.Flight, error)
}

type FlightRepository interface {
	FlightList(req *dt.FlightListRequest, stream dt.FlightService_FlightListServer) error
	Get(ctx context.Context, id string) (*dt.Flight, error)
	GetFutureFlights(ctx context.Context) ([]*dt.Flight, error)
	WriteReview(ctx context.Context, flightId, userName string, review *dt.Review) (*dt.Flight, error)
}
