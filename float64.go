package env

import (
	"fmt"
	"strconv"
)

// GetFloat64 retrieves the float64 value of an environment variable.
// It returns 0.0 if the key is not found or the value is not a valid float64.
func GetFloat64(key string) float64 {
	value, err := GetFloat64OrError(key)
	if err != nil {
		return 0.0
	}
	return value
}

// GetFloat64OrDefault retrieves the float64 value of an environment variable with a default.
func GetFloat64OrDefault(key string, defaultValue float64) float64 {
	value, err := GetFloat64OrError(key)
	if err != nil {
		return defaultValue
	}
	return value
}

// GetFloat64OrError retrieves the float64 value of an environment variable,
// returning an error if the key is not found or the value is not a valid float64.
func GetFloat64OrError(key string) (float64, error) {
	valueStr := GetString(key)
	if valueStr == "" {
		return 0.0, fmt.Errorf("environment variable '%s' not found", key)
	}

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return 0.0, fmt.Errorf("environment variable '%s' with value '%s' cannot be parsed as a float64", key, valueStr)
	}
	return value, nil
}

// GetFloat64OrPanic retrieves the float64 value of an environment variable,
// panicking if not set or on parsing error.
func GetFloat64OrPanic(key string) float64 {
	value, err := GetFloat64OrError(key)
	if err != nil {
		panic(err)
	}
	return value
}

// The following Float helpers are aliases for the Float64 variants to provide
// naming consistency with String/Bool/Int while maintaining backward compatibility.

// GetFloat retrieves the float64 value of an environment variable.
// It returns 0.0 if the key is not found or the value is not a valid float.
func GetFloat(key string) float64 {
	return GetFloat64(key)
}

// GetFloatOrDefault retrieves the float64 value of an environment variable with a default.
func GetFloatOrDefault(key string, defaultValue float64) float64 {
	return GetFloat64OrDefault(key, defaultValue)
}

// GetFloatOrError retrieves the float64 value of an environment variable,
// returning an error if the key is not found or the value is not a valid float.
func GetFloatOrError(key string) (float64, error) {
	return GetFloat64OrError(key)
}

// GetFloatOrPanic retrieves the float64 value of an environment variable,
// panicking if not set or on parsing error.
func GetFloatOrPanic(key string) float64 {
	return GetFloat64OrPanic(key)
}
