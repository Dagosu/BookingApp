package graph

import (
	"time"

	"github.com/99designs/gqlgen/graphql"
	gdomain "github.com/Dagosu/BookingApp/gohelpers/domain"
	generated "github.com/Dagosu/BookingApp/graphql-middleware/graph/generated"
	"github.com/Dagosu/BookingApp/graphql-middleware/graph/upstream"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

//go:generate ../../graphql-schema/build.sh

var kacp = keepalive.ClientParameters{
	Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
	Timeout:             1 * time.Second,  // wait 1 seconds for ping ack before considering the connection dead
	PermitWithoutStream: true,             // send pings even without active streams
}

func NewGraphQLServer() (*Server, error) {
	s := &Server{}
	var err error

	opts := append([]grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(kacp),
	})

	s.bookingClient, err = upstream.BookingClientNew(gdomain.GetServiceURL(gdomain.BookingService), opts)
	if err != nil {
		s.Close()
		return nil, err
	}

	return s, nil
}

// Close terminates any open dependencies.
func (s *Server) Close() {
	if s.bookingClient != nil {
		s.bookingClient.Close()
	}
}

func (s *Server) BookingClient() *upstream.BookingClient {
	return s.bookingClient
}

func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: s,
	})
}

func (s *Server) Mutation() generated.MutationResolver {
	return &mutationResolver{
		&Resolver{
			server: s,
		},
	}
}

func (s *Server) Query() generated.QueryResolver {
	return &queryResolver{
		&Resolver{
			server: s,
		},
	}
}

func (s *Server) Subscription() generated.SubscriptionResolver {
	return &subscriptionResolver{
		&Resolver{
			server: s,
		},
	}
}
