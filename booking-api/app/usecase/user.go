package usecase

import (
	"context"
	"fmt"

	"github.com/Dagosu/BookingApp/booking-api/app/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

// userUsecase is the struct that implements the domain flight usecase
type userUsecase struct {
	ur domain.UserRepository
}

func newuserUsecase(ur domain.UserRepository) *userUsecase {
	return &userUsecase{ur}
}

func (fu *userUsecase) CheckCredentials(ctx context.Context, email, password string) (bool, error) {
	user, err := fu.ur.CheckCredentials(ctx, email, password)
	if err != nil && err != mongo.ErrNoDocuments {
		return false, err
	}

	if user == nil || err == mongo.ErrNoDocuments {
		return false, fmt.Errorf("Invalid credentials!")
	}

	return true, nil
}
