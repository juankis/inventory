package config

import (
	"fmt"
	"os"
)

const (
	PRODUCTION      string = "production"
	TESTING         string = "testing"
	TEST            string = "test"
	DEVELOP         string = "develop"
	SCOPE_METRICS   string = "scope-metrics"
	DB_DSN_TEMPLATE string = "%s:%s@tcp(%s)/inventory"
)

var (
	Environment string
	LoggerLevel string
)

func init() {
	fmt.Println("config.init()")

	var isSet bool
	LoggerLevel = "info"

	// Check (or Sanitize) GO Environment
	if Environment, isSet = os.LookupEnv("GO_ENVIRONMENT"); !isSet {
		Environment = DEVELOP
		os.Setenv("GO_ENVIRONMENT", Environment)
	}
	if Environment == DEVELOP || Environment == TESTING {
		LoggerLevel = "debug"
	}
	fmt.Println("GO-Environment:", Environment, ", was set in env?", isSet)
}
