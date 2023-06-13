package domain

import (
	"context"

	dt "github.com/Dagosu/BookingApp/datatypes"
)

type UserUsecase interface {
	CheckCredentials(ctx context.Context, email, password string) (*dt.User, error)
}

type UserRepository interface {
	CheckCredentials(ctx context.Context, email, password string) (*dt.User, error)
	Get(ctx context.Context, id string) (*dt.User, error)
}
