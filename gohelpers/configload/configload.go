package configload

import (
	"log"
	"os"

	"github.com/k0kubun/pp"
	"github.com/kelseyhightower/envconfig"
)

var hydratedConfigs = map[interface{}]bool{}

// Init loads config from environment variables
func Hydrate(c interface{}) {
	if hydrated := hydratedConfigs[c]; hydrated {
		// config already hydrated
		return
	}

	hydratedConfigs[c] = true

	err := envconfig.Process("", c)

	pp.SetDefaultOutput(os.Stderr)
	pp.Printf("Config specification: %v\n", c)

	if err != nil {
		log.Fatalln("Environment error:", err.Error())
	}
}
