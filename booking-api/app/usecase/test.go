package usecase

import (
	"context"

	"github.com/Dagosu/BookingApp/booking-api/app/domain"
	"google.golang.org/genproto/protobuf/field_mask"
)

// testUsecase is the struct that implements the domain test usecase
type testUsecase struct {
	tr           domain.TestRepository
	updateableFm *field_mask.FieldMask
}

func newTestUsecase(fr domain.TestRepository) *testUsecase {
	// updateableFm, err := fieldmaskpb.New(&dt.Flight{})
	// if err != nil {
	// 	log.Fatalf("testUpdateablePaths error: %v", err)
	// }

	return &testUsecase{fr, nil}
}

func (tu *testUsecase) TestEndpoint(ctx context.Context) error {
	tu.tr.TestEndpoint(ctx)

	return nil
}
