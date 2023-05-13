package main

import (
	"fmt"
	"os"

	grpcDelivery "github.com/Dagosu/BookingApp/booking-api/app/delivery/grpc"
	"github.com/Dagosu/BookingApp/booking-api/app/persistence"
	"github.com/Dagosu/BookingApp/booking-api/app/usecase"
	"github.com/Dagosu/BookingApp/booking-api/config"
	"github.com/Dagosu/BookingApp/gohelpers/database"
	gs "github.com/Dagosu/BookingApp/gohelpers/grpcserver"
	"github.com/Dagosu/BookingApp/gohelpers/keep"
	"google.golang.org/grpc"
)

const (
	// exitFail is the exit code if the program fails.
	exitFail = 1
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run() error {
	err := config.EnvSetup()
	if err != nil {
		panic(err)
	}

	d, err := database.Init(
		config.C.MongoURI,
	)
	if err != nil {
		return fmt.Errorf("Cannot connect to Mongo %v", err)
	}
	defer d.Close()

	// services, err := service.NewServices(cc)
	// if err != nil {
	// 	return err
	// }
	// defer services.Close()

	s, err := gs.New(
		gs.WithGRPCPort(config.C.GrpcPort),
		gs.WithMetricsPort(config.C.MetricsPort),
		gs.WithReflection(config.C.Development),
	)
	if err != nil {
		return fmt.Errorf("Couldn't create grpcserver, %v", err)
	}

	r := persistence.NewRepositories(d)
	u := usecase.NewUsecases(r)

	s.Init(func(s *grpc.Server) {
		grpcDelivery.RegisterServices(s, u, d)
	})
	defer s.Close()

	keep.WaitForSignal()

	return nil
}
