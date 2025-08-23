package env

import (
	"fmt"
	"strconv"
)

// GetInt retrieves the integer value of an environment variable.
// It returns 0 if the key is not found or the value is not a valid integer.
func GetInt(key string) int {
	value, err := GetIntOrError(key)
	if err != nil {
		return 0
	}
	return value
}

// GetIntOrDefault retrieves the integer value of an environment variable with a default.
func GetIntOrDefault(key string, defaultValue int) int {
	value, err := GetIntOrError(key)
	if err != nil {
		return defaultValue
	}
	return value
}

// GetIntOrError retrieves the integer value of an environment variable,
// returning an error if the key is not found or the value is not a valid integer.
func GetIntOrError(key string) (int, error) {
	valueStr := GetString(key)
	if valueStr == "" {
		return 0, fmt.Errorf("environment variable '%s' not found", key)
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, fmt.Errorf("environment variable '%s' with value '%s' cannot be parsed as an integer", key, valueStr)
	}
	return value, nil
}

// GetIntOrPanic retrieves the integer value of an environment variable,
// panicking if not set or on parsing error.
func GetIntOrPanic(key string) int {
	value, err := GetIntOrError(key)
	if err != nil {
		panic(err)
	}
	return value
}
