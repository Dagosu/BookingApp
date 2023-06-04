package domain

import (
	"context"
)

// TestUsecase defines the test usecase interface.
type TestUsecase interface {
	TestEndpoint(ctx context.Context, request string) (string, error)
}

// TestRepository defines the test repository interface.
type TestRepository interface {
	TestEndpoint(ctx context.Context, request string) (string, error)
}
