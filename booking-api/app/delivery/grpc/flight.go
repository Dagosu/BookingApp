package grpc

import (
	"context"

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

func (fs *flightServiceServer) GetFlight(ctx context.Context, req *dt.GetFlightRequest) (*dt.GetFlightResponse, error) {
	flight, err := fs.fu.GetFlight(ctx, req.GetFlightId())
	if err != nil {
		return nil, err
	}

	return &dt.GetFlightResponse{
		Flight: flight,
	}, nil
}

func (fs *flightServiceServer) WriteReview(ctx context.Context, req *dt.WriteReviewRequest) (*dt.WriteReviewResponse, error) {
	flight, err := fs.fu.WriteReview(ctx, req.GetFlightId(), req.GetUserId(), req.GetText())
	if err != nil {
		return nil, err
	}

	return &dt.WriteReviewResponse{
		Flight: flight,
	}, nil
}
