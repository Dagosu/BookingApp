package usecase

import (
	"context"
	"fmt"

	"github.com/Dagosu/BookingApp/booking-api/app/domain"
	dt "github.com/Dagosu/BookingApp/datatypes"
	"go.mongodb.org/mongo-driver/mongo"
)

type userUsecase struct {
	ur domain.UserRepository
}

func newuserUsecase(ur domain.UserRepository) *userUsecase {
	return &userUsecase{ur}
}

func (fu *userUsecase) CheckCredentials(ctx context.Context, email, password string) (*dt.User, error) {
	user, err := fu.ur.CheckCredentials(ctx, email, password)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}

	if user == nil || err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("Invalid credentials!")
	}

	return user, nil
}
