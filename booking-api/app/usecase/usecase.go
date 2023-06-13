package usecase

import (
	"github.com/Dagosu/BookingApp/booking-api/app/persistence"
)

// Usecases contains all usecases used in the delivery layers.
type Usecases struct {
	FlightUsecase             *flightUsecase
	UserUsecase               *userUsecase
	UserFlightsMappingUsecase *userFlightsMappingUsecase
}

// NewUsecases returns usecases instantiated with required repositories
func NewUsecases(r *persistence.Repositories) *Usecases {
	return &Usecases{
		FlightUsecase:             newFlightUsecase(r.FlightRepository),
		UserUsecase:               newuserUsecase(r.UserRepository),
		UserFlightsMappingUsecase: newUserFlightsMappingUsecase(r.UserFlightsMappingRepository, r.UserRepository, r.FlightRepository),
	}
}
