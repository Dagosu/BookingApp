package upstream

import (
	dt "github.com/Dagosu/BookingApp/datatypes"
	"google.golang.org/grpc"
)

// FlightsClient is the grpc client that calls flights-api
type TestClient struct {
	conn        *grpc.ClientConn
	TestService dt.TestServiceClient
}

func TestClientNew(url string, opts []grpc.DialOption) (*TestClient, error) {
	conn, err := grpc.Dial(url, opts...)
	if err != nil {
		return nil, err
	}

	t := dt.NewTestServiceClient(conn)

	return &TestClient{conn, t}, nil
}

func (c *TestClient) Close() {
	c.conn.Close()
}
