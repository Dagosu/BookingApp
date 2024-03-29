package graph

import (
	"context"
	"errors"
	"fmt"
	"io"

	dt "github.com/Dagosu/BookingApp/datatypes"
	"github.com/Dagosu/BookingApp/graphql-middleware/graph/model"
	"github.com/Dagosu/BookingApp/graphql-middleware/graph/parser"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *subscriptionResolver) resolveFlightList(ctx context.Context, in model.FlightListInput) (<-chan *model.FlightListResponse, error) {
	stream, err := r.server.bookingClient.FlightService.FlightList(ctx, &dt.FlightListRequest{
		Limit:  parser.IntDerefer(in.Limit),
		Offset: parser.IntDerefer(in.Offset),
		Query:  parser.StrDerefer(in.Query),
		Sorts:  parser.ParseSort(in.Sorts),
		Filter: parser.ParseFilter(in.Filter),
	})
	if err != nil {
		return nil, fmt.Errorf("Flight List error from grpc call to booking-api: %v", err)
	}

	firstMsg, err := stream.Recv()
	if err != nil {
		return nil, fmt.Errorf("error receiving first message from stream: %v", err)
	}

	events := make(chan *model.FlightListResponse)

	flights := []*model.Flight{}
	ready := false

	go func() {
		parser.FlightStreamToChan(events, firstMsg, &flights, &ready)
		defer close(events)

		for {
			msg, err := stream.Recv()

			if errors.Is(err, io.EOF) || status.Code(err) == codes.Canceled || err != nil {
				fmt.Printf("error receiving stream message: %v\n", err)

				break
			}
			parser.FlightStreamToChan(events, msg, &flights, &ready)
		}
	}()

	return events, nil
}

func (r *queryResolver) resolveGetFlight(ctx context.Context, in model.GetFlightInput) (*model.GetFlightResponse, error) {
	res, err := r.server.bookingClient.FlightService.GetFlight(ctx, &dt.GetFlightRequest{
		FlightId: parser.StrDerefer(in.FlightID),
	})
	if err != nil {
		return nil, err
	}

	return &model.GetFlightResponse{
		Flight: parser.ParseFlight(res.GetFlight()),
	}, nil
}

func (r *mutationResolver) resolveWriteReview(ctx context.Context, in model.WriteReviewInput) (*model.WriteReviewResponse, error) {
	res, err := r.server.bookingClient.FlightService.WriteReview(ctx, &dt.WriteReviewRequest{
		FlightId: parser.StrDerefer(in.FlightID),
		UserId:   parser.StrDerefer(in.UserID),
		Text:     parser.StrDerefer(in.Text),
	})
	if err != nil {
		return nil, err
	}

	return &model.WriteReviewResponse{
		Flight: parser.ParseFlight(res.GetFlight()),
	}, nil
}
