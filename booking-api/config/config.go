package config

import "github.com/kelseyhightower/envconfig"

type specBooking struct {
	MongoURI string `envconfig:"MONGO_URI" required:"true"`
}

// C is used across the project to use its values
var C specBooking

func EnvSetup() error {
	return envconfig.Process("", &C)
}
