package domain

import (
	dt "github.com/Dagosu/BookingApp/datatypes"
)

// FlightUsecase defines the flight usecase interface.
type FlightUsecase interface {
	FlightList(req *dt.FlightListRequest, stream dt.FlightService_FlightListServer) error
}

// FlightRepository defines the flight repository interface.
type FlightRepository interface {
	FlightList(req *dt.FlightListRequest, stream dt.FlightService_FlightListServer) error
}
