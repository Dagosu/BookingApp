package grpc

import (
	"context"

	"github.com/Dagosu/BookingApp/booking-api/app/domain"
	dt "github.com/Dagosu/BookingApp/datatypes"
)

type userServiceServer struct {
	uu domain.UserUsecase
}

var _ dt.UserServiceServer = &userServiceServer{}

func newUserServiceServer(uu domain.UserUsecase) dt.UserServiceServer {
	return &userServiceServer{uu}
}

func (us *userServiceServer) CheckCredentials(ctx context.Context, req *dt.CheckCredentialsRequest) (*dt.CheckCredentialsResponse, error) {
	user, err := us.uu.CheckCredentials(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, err
	}

	authorized := true
	if user == nil {
		authorized = false
	}

	return &dt.CheckCredentialsResponse{
		UserId:     user.GetId(),
		Auhtorized: authorized,
	}, nil
}
