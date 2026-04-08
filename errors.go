package env

import (
	"fmt"
	"strings"
)

// MissingEnvError describes an unset required environment variable
// with optional context to aid debugging.
type MissingEnvError struct {
	Key     string
	Context string
}

// Error returns the formatted error describing the missing environment variable.
func (e MissingEnvError) Error() string {
	if strings.TrimSpace(e.Context) == "" {
		return fmt.Sprintf("env: required variable %q is missing", e.Key)
	}
	return fmt.Sprintf("env: required variable %q is missing: %s", e.Key, e.Context)
}
