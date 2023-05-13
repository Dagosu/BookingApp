package config

import "github.com/kelseyhightower/envconfig"

type specBooking struct {
	Development bool `envconfig:"DEVELOPMENT"`

	MongoURI    string `envconfig:"MONGO_URI" required:"true"`
	GrpcPort    int    `envconfig:"GRPC_PORT" default:"50052"`
	MetricsPort int    `envconfig:"METRICS_PORT" default:"9091"`
}

// C is used across the project to use its values
var C specBooking

func EnvSetup() error {
	return envconfig.Process("", &C)
}
