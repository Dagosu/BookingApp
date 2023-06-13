package graph

import (
	"context"

	dt "github.com/Dagosu/BookingApp/datatypes"
	"github.com/Dagosu/BookingApp/graphql-middleware/graph/model"
	"github.com/Dagosu/BookingApp/graphql-middleware/graph/parser"
)

func (r *queryResolver) resolveCheckCredentials(ctx context.Context, in model.CheckCredentialsInput) (*model.CheckCredentialsResponse, error) {
	res, err := r.server.bookingClient.UserService.CheckCredentials(ctx, &dt.CheckCredentialsRequest{
		Email:    parser.StrDerefer(in.Email),
		Password: parser.StrDerefer(in.Password),
	})
	if err != nil {
		return nil, err
	}

	return &model.CheckCredentialsResponse{
		UserID:     parser.StrRefer(res.GetUserId()),
		Authorized: parser.BoolRefer(res.GetAuhtorized()),
	}, nil
}
