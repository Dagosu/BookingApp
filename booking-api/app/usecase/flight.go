package usecase

import (
	"context"

	"github.com/Dagosu/BookingApp/booking-api/app/domain"
	dt "github.com/Dagosu/BookingApp/datatypes"
	"google.golang.org/genproto/protobuf/field_mask"
)

// flightUsecase is the struct that implements the domain flight usecase
type flightUsecase struct {
	fr           domain.FlightRepository
	ur           domain.UserRepository
	updateableFm *field_mask.FieldMask
}

func newFlightUsecase(fr domain.FlightRepository, ur domain.UserRepository) *flightUsecase {
	// updateableFm, err := fieldmaskpb.New(&dt.Flight{})
	// if err != nil {
	// 	log.Fatalf("flightUpdateablePaths error: %v", err)
	// }

	return &flightUsecase{fr, ur, nil}
}

func (fu *flightUsecase) FlightList(req *dt.FlightListRequest, stream dt.FlightService_FlightListServer) error {
	return fu.fr.FlightList(req, stream)
}

func (fu *flightUsecase) GetFlight(ctx context.Context, id string) (*dt.Flight, error) {
	flight, err := fu.fr.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return flight, nil
}

func (fu *flightUsecase) WriteReview(ctx context.Context, flightId, userId, text string) (*dt.Flight, error) {
	_, err := fu.fr.Get(ctx, flightId)
	if err != nil {
		return nil, err
	}

	user, err := fu.ur.Get(ctx, userId)
	if err != nil {
		return nil, err
	}

	review := &dt.Review{
		UserName: user.GetName(),
		Text:     text,
	}
	resFlight, err := fu.fr.WriteReview(ctx, flightId, user.GetName(), review)
	if err != nil {
		return nil, err
	}

	return resFlight, nil
}
