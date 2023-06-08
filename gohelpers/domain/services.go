// Package domain provides a single point of truth regarding the services a system interacts with.
//
// Usage example:
//
// iam.NewAuthIAMConnection(gdomain.GetServiceURL(gdomain.IamService))
package domain

import (
	"fmt"

	"github.com/Dagosu/BookingApp/gohelpers/env"
)

type serviceKey string

const (
	// FlightService represents the unique identifier for the RuleEngine service.
	BookingService serviceKey = "BOOKING"
)

type service struct {
	name string
	url  string
	repo string
}

var services = map[serviceKey]service{
	BookingService: {
		name: "booking",
		url:  "localhost:50052",
		repo: "booking-api",
	},
}

// GetSelfServiceName returns the Service Name of the caller.
//
// it uses the "SERVICE_NAME" value from the service .env file.
func GetSelfServiceName() string {
	key := "SERVICE_NAME"

	return env.GetEnv(key, services[serviceKey(key)].name)
}

// GetServiceName returns service name identified by the constants defined in this package: eg `domain.IamService`
// This name is overridden if the service has the *SERVICE_IDENTIFIER*_API_NAME defined in the .env file.
func GetServiceName(k serviceKey) string {
	key := fmt.Sprintf("%s_API_NAME", k)

	return env.GetEnv(key, services[k].name)
}

// GetServiceURL returns service URL identified by the constants defined in this package: eg `domain.IamService`
// This URL is overridden if the service has the *SERVICE_IDENTIFIER*_API_URL defined in the .env file.
func GetServiceURL(k serviceKey) string {
	key := fmt.Sprintf("%s_API_URL", k)

	return env.GetEnv(key, services[k].url)
}

func GetServiceRepoPath(k string) string {
	return services[serviceKey(k)].repo
}
