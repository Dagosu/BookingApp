package usecase

import (
	"github.com/Dagosu/BookingApp/booking-api/app/persistence"
)

// Usecases contains all usecases used in the delivery layers.
type Usecases struct {
	TestUsecase *testUsecase
}

// NewUsecases returns usecases instantiated with required repositories
func NewUsecases(r *persistence.Repositories) *Usecases {
	return &Usecases{
		TestUsecase: newTestUsecase(r.TestRepository),
	}
}
