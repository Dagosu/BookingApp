package graph

import (
	"context"

	dt "github.com/Dagosu/BookingApp/datatypes"
	"github.com/Dagosu/BookingApp/graphql-middleware/graph/model"
	"github.com/Dagosu/BookingApp/graphql-middleware/graph/parser"
)

func (r *mutationResolver) resolvePurchaseFlight(ctx context.Context, in model.PurchaseFlightInput) (*model.PurchaseFlightResponse, error) {
	res, err := r.server.bookingClient.UserFlightsMappingService.PurchaseFlight(ctx, &dt.PurchaseFlightRequest{
		UserId:   parser.StrDerefer(in.UserID),
		FlightId: parser.StrDerefer(in.FlightID),
	})
	if err != nil {
		return nil, err
	}

	return &model.PurchaseFlightResponse{
		PurchasedFlight: parser.ParseFlight(res.GetPurchasedFlight()),
	}, nil
}
