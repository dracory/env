package env

import (
	"encoding/json"
	"fmt"
	"strings"
)

// GetArray retrieves a slice of strings from an environment variable.
// It splits the string by the provided separators (defaulting to ',' and ';').
// If the value is formatted as a JSON array, it will automatically parse it as JSON.
// It returns nil if the key is not found.
func GetArray(key string, separators ...string) []string {
	value, err := GetArrayOrError(key, separators...)
	if err != nil {
		return nil
	}
	return value
}

// GetArrayOrDefault retrieves a slice of strings from an environment variable with a default fallback.
func GetArrayOrDefault(key string, defaultValue []string, separators ...string) []string {
	value, err := GetArrayOrError(key, separators...)
	if err != nil {
		return defaultValue
	}
	return value
}

// GetArrayOrError retrieves a slice of strings from an environment variable,
// returning an error if the key is not found.
func GetArrayOrError(key string, separators ...string) ([]string, error) {
	valueStr := GetString(key)
	if valueStr == "" {
		return nil, fmt.Errorf("environment variable '%s' not found", key)
	}

	trimmed := strings.TrimSpace(valueStr)
	if strings.HasPrefix(trimmed, "[") && strings.HasSuffix(trimmed, "]") {
		if res, err := parseJSONArray(trimmed); err == nil {
			return res, nil
		}
	}

	return parseDelimitedArray(trimmed, separators...), nil
}

// GetArrayOrPanic retrieves a slice of strings from an environment variable,
// panicking if not set.
func GetArrayOrPanic(key string, separators ...string) []string {
	value, err := GetArrayOrError(key, separators...)
	if err != nil {
		panic(err)
	}
	return value
}

// GetJSONArray retrieves a slice of strings from a JSON array environment variable.
// It returns nil if the key is not found or cannot be parsed as JSON.
func GetJSONArray(key string) []string {
	value, err := GetJSONArrayOrError(key)
	if err != nil {
		return nil
	}
	return value
}

// GetJSONArrayOrDefault retrieves a slice of strings from a JSON array environment variable with a default fallback.
func GetJSONArrayOrDefault(key string, defaultValue []string) []string {
	value, err := GetJSONArrayOrError(key)
	if err != nil {
		return defaultValue
	}
	return value
}

// GetJSONArrayOrError retrieves a slice of strings from a JSON array environment variable,
// returning an error if the key is not found or cannot be parsed as JSON.
func GetJSONArrayOrError(key string) ([]string, error) {
	valueStr := GetString(key)
	if valueStr == "" {
		return nil, fmt.Errorf("environment variable '%s' not found", key)
	}

	res, err := parseJSONArray(strings.TrimSpace(valueStr))
	if err != nil {
		return nil, fmt.Errorf("environment variable '%s' with value '%s' cannot be parsed as a JSON array: %w", key, valueStr, err)
	}

	return res, nil
}

// GetJSONArrayOrPanic retrieves a slice of strings from a JSON array environment variable,
// panicking if not set or cannot be parsed as JSON.
func GetJSONArrayOrPanic(key string) []string {
	value, err := GetJSONArrayOrError(key)
	if err != nil {
		panic(err)
	}
	return value
}

func parseDelimitedArray(val string, separators ...string) []string {
	if len(separators) == 0 {
		separators = []string{",", ";"}
	}

	for _, sep := range separators {
		val = strings.ReplaceAll(val, sep, "\x00")
	}

	parts := strings.Split(val, "\x00")
	result := make([]string, 0, len(parts))
	for _, p := range parts {
		trimmed := strings.TrimSpace(p)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

func parseJSONArray(val string) ([]string, error) {
	var strSlice []string
	if err := json.Unmarshal([]byte(val), &strSlice); err == nil {
		result := make([]string, 0, len(strSlice))
		for _, s := range strSlice {
			trimmed := strings.TrimSpace(s)
			if trimmed != "" {
				result = append(result, trimmed)
			}
		}
		return result, nil
	}

	var anySlice []any
	if err := json.Unmarshal([]byte(val), &anySlice); err == nil {
		result := make([]string, 0, len(anySlice))
		for _, v := range anySlice {
			trimmed := strings.TrimSpace(fmt.Sprintf("%v", v))
			if trimmed != "" {
				result = append(result, trimmed)
			}
		}
		return result, nil
	}

	return nil, fmt.Errorf("invalid JSON array syntax")
}
