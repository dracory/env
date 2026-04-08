package env

import "strings"

// RequireString trims and retrieves the environment value for the provided key,
// returning a MissingEnvError when the value is absent.
func RequireString(key, context string) (string, error) {
	value := strings.TrimSpace(GetString(key))
	if err := EnsureRequired(value, key, context); err != nil {
		return "", err
	}
	return value, nil
}

// RequireWhen validates that the given value is present when the condition is true.
func RequireWhen(condition bool, key, context, value string) error {
	if !condition {
		return nil
	}
	return EnsureRequired(value, key, context)
}

// EnsureRequired returns a MissingEnvError when the supplied value is blank.
func EnsureRequired(value, key, context string) error {
	if strings.TrimSpace(value) != "" {
		return nil
	}
	return MissingEnvError{Key: key, Context: context}
}
