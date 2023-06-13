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
	updateableFm *field_mask.FieldMask
}

func newFlightUsecase(fr domain.FlightRepository) *flightUsecase {
	// updateableFm, err := fieldmaskpb.New(&dt.Flight{})
	// if err != nil {
	// 	log.Fatalf("flightUpdateablePaths error: %v", err)
	// }

	return &flightUsecase{fr, nil}
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
