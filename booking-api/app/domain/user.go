package domain

import (
	"context"

	dt "github.com/Dagosu/BookingApp/datatypes"
)

// UserUsecase defines the user usecase interface.
type UserUsecase interface {
	CheckCredentials(ctx context.Context, email, password string) (bool, error)
}

// UserRepository defines the fluseright repository interface.
type UserRepository interface {
	CheckCredentials(ctx context.Context, email, password string) (*dt.User, error)
}
