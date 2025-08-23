package env

import (
	"fmt"
	"os"
)

// GetString retrieves the string value of an environment variable.
// It returns an empty string if the key is not found.
func GetString(key string) string {
	return envProcess(os.Getenv(key))
}

// GetStringOrDefault retrieves the string value of an environment variable with a default.
func GetStringOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return envProcess(value)
}

// GetStringOrError retrieves the string value of an environment variable,
// returning an error if the key is not found.
func GetStringOrError(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("environment variable '%s' not found", key)
	}
	return envProcess(value), nil
}

// GetStringOrPanic retrieves the string value of an environment variable,
// panicking if not set.
func GetStringOrPanic(key string) string {
	value, err := GetStringOrError(key)
	if err != nil {
		panic(fmt.Sprintf("Environment variable '%s' is required, but not set.", key))
	}
	return value
}
