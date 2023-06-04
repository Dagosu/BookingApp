//nolint:gochecknoglobals
package config

type specGraphQLMiddleware struct {
	Port        string `envconfig:"PORT" default:"8080"`
	Development bool   `envconfig:"DEVELOPMENT" default:"false"`

	QueryComplexityLimit int `envconfig:"QUERY_COMPLEXITY_LIMIT" default:"4000"`
}

// C is used across the project to use its values
var C specGraphQLMiddleware
