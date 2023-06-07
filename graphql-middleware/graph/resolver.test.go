package graph

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	dt "github.com/Dagosu/BookingApp/datatypes"
	"github.com/Dagosu/BookingApp/graphql-middleware/graph/model"
	"github.com/Dagosu/BookingApp/graphql-middleware/graph/parser"
	"github.com/k0kubun/pp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *queryResolver) resolveTestEndpoint(ctx context.Context, in model.TestEndpointInput) (*model.TestEndpointResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if r.server == nil || r.server.testClient == nil || r.server.testClient.TestService == nil {
		return nil, pp.Errorf("server or testClient or TestService is not initialized")
	}

	res, err := r.server.testClient.TestService.TestEndpoint(ctx, &dt.TestEndpointRequest{
		Request: parser.StrDerefer(in.Request),
	})
	if err != nil {
		return nil, err
	}

	return &model.TestEndpointResponse{
		Response: parser.StrRefer(res.GetResponse()),
	}, nil
}

func (r *subscriptionResolver) resolveTestList(ctx context.Context, in model.TestListInput) (<-chan *model.TestListResponse, error) {
	stream, err := r.server.testClient.TestService.TestList(ctx, &dt.TestListRequest{
		Limit:  parser.IntDerefer(in.Limit),
		Offset: parser.IntDerefer(in.Offset),
		Query:  parser.StrDerefer(in.Query),
		Sorts:  parser.ParseSort(in.Sorts),
		Filter: parser.ParseFilter(in.Filter),
	})
	if err != nil {
		return nil, fmt.Errorf("Test List error from grpc call to booking-api: %v", err)
	}

	firstMsg, err := stream.Recv()
	if err != nil {
		return nil, fmt.Errorf("error receiving first message from stream: %v", err)
	}

	events := make(chan *model.TestListResponse)

	myObjects := []*model.MyObject{}
	ready := false

	go func() {
		parser.ObjectsStreamToChan(events, firstMsg, &myObjects, &ready)
		defer close(events)

		for {
			msg, err := stream.Recv()

			if errors.Is(err, io.EOF) || status.Code(err) == codes.Canceled || err != nil {
				fmt.Printf("error receiving stream message: %v\n", err)

				break
			}
			parser.ObjectsStreamToChan(events, msg, &myObjects, &ready)
		}
	}()

	return events, nil
}
