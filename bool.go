package env

import (
	"strconv"
	"strings"
)

// Bool returns the value for an environment key
//
// Any of the following values are considered true: "true", "True", "TRUE", "T", "t", "1"
// Any of the following values are considered false: "false", "False", "FALSE", "F", "f", "0"
// Any other value returns the default false.
//
// Parameters:
//   - key: The environment key
//
// Returns:
//   - The value for the environment key, false if not set
func Bool(key string) bool {
	return getEnvBool(key, false)
}

// BoolDefault returns the value for an environment key with a default value
//
// Any of the following values are considered true: "true", "True", "TRUE", "T", "t", "1"
// Any of the following values are considered false: "false", "False", "FALSE", "F", "f", "0"
// Any other value returns the default value.
//
// Parameters:
//   - key: The environment key
//   - defaultValue: The default value
//
// Returns:
//   - The value for the environment key
func BoolDefault(key string, defaultValue bool) bool {
	return getEnvBool(key, defaultValue)
}

// getEnvBool returns the value for an environment key as a boolean, or the default value if not set
//
// Any of the following values are considered true: "true", "True", "TRUE", "T", "t", "1", "yes", "Yes", "YES"
// Any of the following values are considered false: "false", "False", "FALSE", "F", "f", "0", "no", "No", "NO"
// Any other value returns the default value.
//
// This function handles base64 and obfuscated prefixes.
//
// Parameters:
//   - key: The environment key
//   - defaultValue: The default value
//
// Returns:
//   - The value for the environment key
func getEnvBool(key string, defaultValue bool) bool {
	valueStr := Value(key)

	if valueStr == "" {
		return defaultValue
	}

	if strings.ToLower(valueStr) == "yes" {
		return true
	}

	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		// Return the default value if the value can't be parsed as a boolean
		return defaultValue
	}

	return value
}
