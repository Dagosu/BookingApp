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
	authorized, err := us.uu.CheckCredentials(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, err
	}

	return &dt.CheckCredentialsResponse{
		Auhtorized: authorized,
	}, nil
}
