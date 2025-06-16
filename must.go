package env

import (
	"log"
	"os"
)

// Must retrieves the value of an environment variable, panicking if not set.
func Must(key string) string {
	value := os.Getenv(key)

	if value == "" {
		log.Panicf(`Environment variable "%s" is required but not set`, key)
	}

	valueProcessed := envProcess(value)

	return valueProcessed
}
