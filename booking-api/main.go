package main

import (
	"fmt"
	"os"

	"github.com/Dagosu/BookingApp/booking-api/config"
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
	test := config.C.MongoURI
	fmt.Println(test)

	// services, err := service.NewServices(cc)
	// if err != nil {
	// 	return err
	// }
	// defer services.Close()

	// s, err := gs.New(
	// 	gs.WithGRPCPort(config.C.GrpcPort),
	// 	gs.WithMetricsPort(config.C.MetricsPort),
	// 	gs.WithReflection(config.C.Development),
	// )
	// if err != nil {
	// 	return fmt.Errorf("Couldn't create grpcserver, %v", err)
	// }

	// r := persistence.NewRepositories(d, dLogs, services.IamClient)
	// u := usecase.NewUsecases(r, services)

	// s.Init(func(s *grpc.Server) {
	// 	grpcDelivery.RegisterServices(s, u, d)
	// })
	// defer s.Close()

	// keep.WaitForSignal()

	return nil
}
