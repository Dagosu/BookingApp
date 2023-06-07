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

var _ dt.TestServiceServer = &testServiceServer{}

func newTestServiceServer(tu domain.TestUsecase) dt.TestServiceServer {
	return &testServiceServer{tu}
}

func (ts *testServiceServer) TestEndpoint(ctx context.Context, req *dt.TestEndpointRequest) (*dt.TestEndpointResponse, error) {
	response, err := ts.tu.TestEndpoint(ctx, req.GetRequest())
	if err != nil {
		return nil, fmt.Errorf("error")
	}

	return &dt.TestEndpointResponse{
		Response: response,
	}, nil
}

func (ts *testServiceServer) TestList(req *dt.TestListRequest, stream dt.TestService_TestListServer) error {
	return ts.tu.ListMyObjects(req, stream)
}
