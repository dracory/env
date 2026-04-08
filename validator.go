package env

import "strings"

// Validator collects validation errors for required environment variables.
// Use it to accumulate all missing var errors before returning them at once,
// rather than failing on the first missing variable.
type Validator struct {
	errs []error
}

// Add appends err to the validator when it is non-nil.
func (v *Validator) Add(err error) {
	if err != nil {
		v.errs = append(v.errs, err)
	}
}

// MustString retrieves the required string value for key, recording any error.
func (v *Validator) MustString(key, context string) string {
	value, err := RequireString(key, context)
	v.Add(err)
	return value
}

// GetString retrieves an optional string value for key (no validation).
func (v *Validator) GetString(key string) string {
	return GetString(key)
}

// GetBool retrieves an optional bool value for key (no validation).
func (v *Validator) GetBool(key string) bool {
	return GetBool(key)
}

// MustWhen validates that value is present when condition is true, recording any error.
func (v *Validator) MustWhen(condition bool, key, context, value string) {
	v.Add(RequireWhen(condition, key, context, value))
}

// Err returns a ValidationError wrapping all collected issues, or nil if none.
func (v *Validator) Err() error {
	if len(v.errs) == 0 {
		return nil
	}
	return ValidationError{errs: v.errs}
}

// ValidationError aggregates multiple missing/invalid environment variable errors.
type ValidationError struct {
	errs []error
}

// Error returns a formatted string listing all validation failures.
func (e ValidationError) Error() string {
	if len(e.errs) == 0 {
		return "env: validation failed"
	}
	var b strings.Builder
	b.WriteString("env: validation failed:\n")
	for i, err := range e.errs {
		if i > 0 {
			b.WriteByte('\n')
		}
		b.WriteString(" - ")
		b.WriteString(err.Error())
	}
	return b.String()
}

// Errors returns the accumulated error slice.
func (e ValidationError) Errors() []error {
	return append([]error(nil), e.errs...)
}
