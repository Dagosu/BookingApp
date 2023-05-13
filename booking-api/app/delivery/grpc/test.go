package grpc

import (
	"context"
	"fmt"

	"github.com/Dagosu/BookingApp/booking-api/app/domain"
	dt "github.com/Dagosu/BookingApp/datatypes"
)

type testServiceServer struct {
	tu domain.TestUsecase
}

// var _ dt.FlightServiceServer = &flightServiceServer{}

func newTestServiceServer(tu domain.TestUsecase) dt.TestServiceServer {
	return &testServiceServer{tu}
}

func (ts *testServiceServer) TestEndpoint(ctx context.Context, req *dt.TestEndpointRequest) (*dt.TestEndpointResponse, error) {
	err := ts.tu.TestEndpoint(ctx)
	if err != nil {
		return nil, fmt.Errorf("error")
	}

	return nil, fmt.Errorf("error")
}
