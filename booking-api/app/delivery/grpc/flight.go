package grpc

import (
	"github.com/Dagosu/BookingApp/booking-api/app/domain"
	dt "github.com/Dagosu/BookingApp/datatypes"
)

type flightServiceServer struct {
	fu domain.FlightUsecase
}

var _ dt.FlightServiceServer = &flightServiceServer{}

func newFlightServiceServer(tu domain.FlightUsecase) dt.FlightServiceServer {
	return &flightServiceServer{tu}
}

func (fs *flightServiceServer) FlightList(req *dt.FlightListRequest, stream dt.FlightService_FlightListServer) error {
	return fs.fu.FlightList(req, stream)
}
