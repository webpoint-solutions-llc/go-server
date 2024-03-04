package config

import (
	"fmt"
	"os"
)

type envVars struct {
	PGURL string
	PORT  string
}

var Env envVars

// retrieve the value of the environment variable or panics if not set
func getEnv(varName string) string {
	value, exists := os.LookupEnv(varName)
	if !exists {
		panic(fmt.Sprintf("%s must be set", varName))
	}
	return value
}

// retrieve the value of the environment variable or returns a default value if not set
func getOptEnv(varName, defaultValue string) string {
	value, exists := os.LookupEnv(varName)
	if !exists {
		return defaultValue
	}
	return value
}

func LoadEnv() {
	Env = envVars{
		PORT:  getOptEnv("PORT", "8080"),
		PGURL: getEnv("POSTGRESQL_URL"),
	}
}
