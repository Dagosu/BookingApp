package grpc

import (
	"github.com/Dagosu/BookingApp/booking-api/app/usecase"
	dt "github.com/Dagosu/BookingApp/datatypes"
	"github.com/Dagosu/BookingApp/gohelpers/database"
	"google.golang.org/grpc"
)

// RegisterServices instantiates and registers all grpc delivery services
func RegisterServices(s *grpc.Server, u *usecase.Usecases, d *database.Db) {
	dt.RegisterFlightServiceServer(s, newFlightServiceServer(u.FlightUsecase))
	dt.RegisterUserServiceServer(s, newUserServiceServer(u.UserUsecase))
}
