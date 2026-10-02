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

// GetJSONArrayMapped retrieves a slice of strings from a JSON array environment variable and applies mapper to each element.
func GetJSONArrayMapped(key string, mapper func(string) string) []string {
	value, err := GetJSONArrayMappedOrError(key, mapper)
	if err != nil {
		return nil
	}
	return value
}

// GetJSONArrayMappedOrDefault retrieves a mapped slice of strings from a JSON array with a default fallback.
func GetJSONArrayMappedOrDefault(key string, defaultValue []string, mapper func(string) string) []string {
	value, err := GetJSONArrayMappedOrError(key, mapper)
	if err != nil {
		return defaultValue
	}
	return value
}

// GetJSONArrayMappedOrError retrieves a slice of strings from a JSON array and applies mapper to each element, returning an error if parsing fails or key is not set.
func GetJSONArrayMappedOrError(key string, mapper func(string) string) ([]string, error) {
	elems, err := GetJSONArrayOrError(key)
	if err != nil {
		return nil, err
	}
	if mapper == nil {
		return elems, nil
	}
	res := make([]string, len(elems))
	for i, e := range elems {
		res[i] = mapper(e)
	}
	return res, nil
}

// GetJSONArrayMappedOrPanic retrieves a mapped slice of strings from a JSON array, panicking if not set or invalid JSON.
func GetJSONArrayMappedOrPanic(key string, mapper func(string) string) []string {
	value, err := GetJSONArrayMappedOrError(key, mapper)
	if err != nil {
		panic(err)
	}
	return value
}

// GetJSONArrayLower retrieves a slice of strings from a JSON array converted to lower case.
func GetJSONArrayLower(key string) []string {
	return GetJSONArrayMapped(key, strings.ToLower)
}

// GetJSONArrayLowerOrDefault retrieves a lowercased slice of strings from a JSON array with a default fallback.
func GetJSONArrayLowerOrDefault(key string, defaultValue []string) []string {
	return GetJSONArrayMappedOrDefault(key, defaultValue, strings.ToLower)
}

// GetJSONArrayLowerOrError retrieves a lowercased slice of strings from a JSON array, returning an error if invalid/missing.
func GetJSONArrayLowerOrError(key string) ([]string, error) {
	return GetJSONArrayMappedOrError(key, strings.ToLower)
}

// GetJSONArrayLowerOrPanic retrieves a lowercased slice of strings from a JSON array, panicking if invalid/missing.
func GetJSONArrayLowerOrPanic(key string) []string {
	return GetJSONArrayMappedOrPanic(key, strings.ToLower)
}

// GetJSONArrayUpper retrieves a slice of strings from a JSON array converted to upper case.
func GetJSONArrayUpper(key string) []string {
	return GetJSONArrayMapped(key, strings.ToUpper)
}

// GetJSONArrayUpperOrDefault retrieves an uppercased slice of strings from a JSON array with a default fallback.
func GetJSONArrayUpperOrDefault(key string, defaultValue []string) []string {
	return GetJSONArrayMappedOrDefault(key, defaultValue, strings.ToUpper)
}

// GetJSONArrayUpperOrError retrieves an uppercased slice of strings from a JSON array, returning an error if invalid/missing.
func GetJSONArrayUpperOrError(key string) ([]string, error) {
	return GetJSONArrayMappedOrError(key, strings.ToUpper)
}

// GetJSONArrayUpperOrPanic retrieves an uppercased slice of strings from a JSON array, panicking if invalid/missing.
func GetJSONArrayUpperOrPanic(key string) []string {
	return GetJSONArrayMappedOrPanic(key, strings.ToUpper)
}

// GetArrayMapped applies custom mapper transformation to each element.
// It returns nil if the key is not found or if the value cannot be retrieved.
func GetArrayMapped(key string, mapper func(string) string, separators ...string) []string {
	value, err := GetArrayMappedOrError(key, mapper, separators...)
	if err != nil {
		return nil
	}
	return value
}

// GetArrayMappedOrDefault retrieves a mapped slice of strings with a default fallback if not found.
func GetArrayMappedOrDefault(key string, defaultValue []string, mapper func(string) string, separators ...string) []string {
	value, err := GetArrayMappedOrError(key, mapper, separators...)
	if err != nil {
		return defaultValue
	}
	return value
}

// GetArrayMappedOrError retrieves a slice of strings and applies the mapper function to each element.
// Returns an error if the key is not found.
func GetArrayMappedOrError(key string, mapper func(string) string, separators ...string) ([]string, error) {
	elems, err := GetArrayOrError(key, separators...)
	if err != nil {
		return nil, err
	}
	if mapper == nil {
		return elems, nil
	}
	res := make([]string, len(elems))
	for i, e := range elems {
		res[i] = mapper(e)
	}
	return res, nil
}

// GetArrayMappedOrPanic retrieves a mapped slice of strings, panicking if not set.
func GetArrayMappedOrPanic(key string, mapper func(string) string, separators ...string) []string {
	value, err := GetArrayMappedOrError(key, mapper, separators...)
	if err != nil {
		panic(err)
	}
	return value
}

// GetArrayLower retrieves a slice of strings converted to lower case.
func GetArrayLower(key string, separators ...string) []string {
	return GetArrayMapped(key, strings.ToLower, separators...)
}

// GetArrayLowerOrDefault retrieves a slice of strings converted to lower case with a default fallback.
func GetArrayLowerOrDefault(key string, defaultValue []string, separators ...string) []string {
	return GetArrayMappedOrDefault(key, defaultValue, strings.ToLower, separators...)
}

// GetArrayLowerOrError retrieves a slice of strings converted to lower case, returning an error if not found.
func GetArrayLowerOrError(key string, separators ...string) ([]string, error) {
	return GetArrayMappedOrError(key, strings.ToLower, separators...)
}

// GetArrayLowerOrPanic retrieves a slice of strings converted to lower case, panicking if not set.
func GetArrayLowerOrPanic(key string, separators ...string) []string {
	return GetArrayMappedOrPanic(key, strings.ToLower, separators...)
}

// GetArrayUpper retrieves a slice of strings converted to upper case.
func GetArrayUpper(key string, separators ...string) []string {
	return GetArrayMapped(key, strings.ToUpper, separators...)
}

// GetArrayUpperOrDefault retrieves a slice of strings converted to upper case with a default fallback.
func GetArrayUpperOrDefault(key string, defaultValue []string, separators ...string) []string {
	return GetArrayMappedOrDefault(key, defaultValue, strings.ToUpper, separators...)
}

// GetArrayUpperOrError retrieves a slice of strings converted to upper case, returning an error if not found.
func GetArrayUpperOrError(key string, separators ...string) ([]string, error) {
	return GetArrayMappedOrError(key, strings.ToUpper, separators...)
}

// GetArrayUpperOrPanic retrieves a slice of strings converted to upper case, panicking if not set.
func GetArrayUpperOrPanic(key string, separators ...string) []string {
	return GetArrayMappedOrPanic(key, strings.ToUpper, separators...)
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
