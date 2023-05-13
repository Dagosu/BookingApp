package config

type specGoHelpers struct {
	ServiceName string `envconfig:"SERVICE_NAME" default:"gohelpers"`
}

// C is used across the project to use its values
var C specGoHelpers
