package persistence

import (
	"github.com/Dagosu/BookingApp/gohelpers/database"
)

// Repositories contains a reference to each repository.
type Repositories struct {
	FlightRepository *flightRepository
	UserRepository   *userRepository
}

// NewRepositories returns repositories instantiated with a database connection
func NewRepositories(d *database.Db) *Repositories {
	// clock := clock.New()

	return &Repositories{
		FlightRepository: newFlightRepository(d),
		UserRepository:   newUserRepository(d),
	}
}
