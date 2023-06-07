package domain

import (
	"context"

	dt "github.com/Dagosu/BookingApp/datatypes"
)

// TestUsecase defines the test usecase interface.
type TestUsecase interface {
	TestEndpoint(ctx context.Context, request string) (string, error)
	ListMyObjects(req *dt.TestListRequest, stream dt.TestService_TestListServer) error
}

// TestRepository defines the test repository interface.
type TestRepository interface {
	TestEndpoint(ctx context.Context, request string) (string, error)
	ListMyObjects(req *dt.TestListRequest, stream dt.TestService_TestListServer) error
}
