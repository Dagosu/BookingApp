// inspired from
// https://dev.to/craicoverflow/a-no-nonsense-guide-to-environment-variables-in-go-a2f
package env

import (
	"log"
	"os"
	"runtime/debug"
	"strconv"
)

// TODO use existing package for env processing
// https://github.com/kelseyhightower/envconfig

// GetEnv returns provided env value, or the fallback when it doesn't exist
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

// GetEnvMustExist returns provided env value, or panics when it doesn't exits
func GetEnvMustExist(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	debug.PrintStack()
	log.Fatalln("Please provide env value for", key)

	return ""
}

// GetEnvAsBool returns provided env value, or the fallback when it doesn't exist
func GetEnvAsBool(name string, fallback bool) bool {
	valStr := GetEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return fallback
}

// GetEnvAsIntMustExist returns provided env value converted to Int,
// or panics when it doesn't exits or it has a invalid value
func GetEnvAsIntMustExist(name string) int {
	valueStr := GetEnvMustExist(name)
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Fatalln("Invalid ", name, err)
	}

	return value
}
