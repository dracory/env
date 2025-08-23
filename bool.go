package env

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Precompute truthy/falsy token sets from constants in constants.go
var (
	trueSet  map[string]struct{}
	falseSet map[string]struct{}
	// Strict numeric pattern: optional sign, digits with optional decimal, optional exponent
	numericRe = regexp.MustCompile(`^[+-]?((\d+\.?\d*)|(\.\d+))([eE][+-]?\d+)?$`)
)

func init() {
	trueSet = make(map[string]struct{})
	for _, v := range strings.Split(TrueValues, ",") {
		v = strings.TrimSpace(v)
		if v != "" {
			trueSet[v] = struct{}{}
		}
	}
	falseSet = make(map[string]struct{})
	for _, v := range strings.Split(FalseValues, ",") {
		v = strings.TrimSpace(v)
		if v != "" {
			falseSet[v] = struct{}{}
		}
	}
}

// GetBool retrieves the boolean value of an environment variable.
// It returns false if the key is not found or the value is not a valid boolean.
func GetBool(key string) bool {
	value, err := GetBoolOrError(key)
	if err != nil {
		return false
	}
	return value
}

// GetBoolOrDefault retrieves the boolean value of an environment variable with a default.
func GetBoolOrDefault(key string, defaultValue bool) bool {
	value, err := GetBoolOrError(key)
	if err != nil {
		return defaultValue
	}
	return value
}

// GetBoolOrError retrieves the boolean value of an environment variable,
// returning an error if the key is not found or the value is not a valid boolean.
func GetBoolOrError(key string) (bool, error) {
	valueStr := strings.TrimSpace(GetString(key))
	if valueStr == "" {
		return false, fmt.Errorf("environment variable '%s' not found", key)
	}

	// First, honor the explicit truthy/falsy token lists from constants.go
	if _, ok := trueSet[valueStr]; ok {
		return true, nil
	}
	if _, ok := falseSet[valueStr]; ok {
		return false, nil
	}

	// Next, handle numeric values according to the documented rules in constants.go:
	// any positive number => true; zero or any negative number => false
	if numericRe.MatchString(valueStr) {
		if n, errNum := strconv.ParseFloat(valueStr, 64); errNum == nil {
			if n > 0 {
				return true, nil
			}
			return false, nil
		}
	}

	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		return false, fmt.Errorf("environment variable '%s' with value '%s' cannot be parsed as a boolean", key, valueStr)
	}
	return value, nil
}

// GetBoolOrPanic retrieves the boolean value of an environment variable,
// panicking if not set or on parsing error.
func GetBoolOrPanic(key string) bool {
	value, err := GetBoolOrError(key)
	if err != nil {
		panic(err)
	}
	return value
}
