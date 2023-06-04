package graph

import (
	"context"
	"time"

	dt "github.com/Dagosu/BookingApp/datatypes"
	"github.com/Dagosu/BookingApp/graphql-middleware/graph/model"
	"github.com/Dagosu/BookingApp/graphql-middleware/graph/parser"
	"github.com/k0kubun/pp"
)

func (r *queryResolver) resolveTestEndpoint(ctx context.Context, in model.TestEndpointInput) (*model.TestEndpointResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if r.server == nil {
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
