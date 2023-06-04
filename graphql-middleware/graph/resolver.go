package graph

import "github.com/Dagosu/BookingApp/graphql-middleware/graph/upstream"

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Server struct {
	testClient *upstream.TestClient
}

type Resolver struct {
	server *Server
}
