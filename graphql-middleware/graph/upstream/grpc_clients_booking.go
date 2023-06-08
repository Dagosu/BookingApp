package upstream

import (
	dt "github.com/Dagosu/BookingApp/datatypes"
	"google.golang.org/grpc"
)

// FlightsClient is the grpc client that calls booking-api
type BookingClient struct {
	conn          *grpc.ClientConn
	FlightService dt.FlightServiceClient
	UserService   dt.UserServiceClient
}

func BookingClientNew(url string, opts []grpc.DialOption) (*BookingClient, error) {
	conn, err := grpc.Dial(url, opts...)
	if err != nil {
		return nil, err
	}

	fsc := dt.NewFlightServiceClient(conn)
	usc := dt.NewUserServiceClient(conn)

	return &BookingClient{conn, fsc, usc}, nil
}

func (c *BookingClient) Close() {
	c.conn.Close()
}
