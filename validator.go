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

// GetStringOrError retrieves the required string value for key.
// Records a MissingEnvError if the value is absent.
func (v *Validator) GetStringOrError(key, context string) string {
	value := strings.TrimSpace(GetString(key))
	if value == "" {
		v.Add(MissingEnvError{Key: key, Context: context})
		return ""
	}
	return value
}

// GetStringOrDefault retrieves a string value for key, returning defaultValue if not set.
func (v *Validator) GetStringOrDefault(key, defaultValue string) string {
	return GetStringOrDefault(key, defaultValue)
}

// GetString retrieves an optional string value for key (no validation).
func (v *Validator) GetString(key string) string {
	return GetString(key)
}

// GetBoolOrError retrieves the required bool value for key.
// Records a MissingEnvError if the value is absent.
func (v *Validator) GetBoolOrError(key, context string) bool {
	value, err := GetBoolOrError(key)
	if err != nil {
		v.Add(MissingEnvError{Key: key, Context: context})
		return false
	}
	return value
}

// GetBoolOrDefault retrieves a bool value for key, returning defaultValue if not set.
func (v *Validator) GetBoolOrDefault(key string, defaultValue bool) bool {
	return GetBoolOrDefault(key, defaultValue)
}

// GetBool retrieves an optional bool value for key (no validation).
func (v *Validator) GetBool(key string) bool {
	return GetBool(key)
}

// GetInt retrieves an optional int value for key (no validation).
func (v *Validator) GetInt(key string) int {
	return GetInt(key)
}

// GetIntOrError retrieves the required int value for key.
// Records a MissingEnvError if the value is absent or unparseable.
func (v *Validator) GetIntOrError(key, context string) int {
	value, err := GetIntOrError(key)
	if err != nil {
		v.Add(MissingEnvError{Key: key, Context: context})
		return 0
	}
	return value
}

// GetIntOrDefault retrieves an int value for key, returning defaultValue if not set.
func (v *Validator) GetIntOrDefault(key string, defaultValue int) int {
	return GetIntOrDefault(key, defaultValue)
}

// GetFloat64 retrieves an optional float64 value for key (no validation).
func (v *Validator) GetFloat64(key string) float64 {
	return GetFloat64(key)
}

// GetFloat64OrError retrieves the required float64 value for key.
// Records a MissingEnvError if the value is absent or unparseable.
func (v *Validator) GetFloat64OrError(key, context string) float64 {
	value, err := GetFloat64OrError(key)
	if err != nil {
		v.Add(MissingEnvError{Key: key, Context: context})
		return 0
	}
	return value
}

// GetFloat64OrDefault retrieves a float64 value for key, returning defaultValue if not set.
func (v *Validator) GetFloat64OrDefault(key string, defaultValue float64) float64 {
	return GetFloat64OrDefault(key, defaultValue)
}

// GetArray retrieves an optional slice of strings for key (no validation).
func (v *Validator) GetArray(key string, separators ...string) []string {
	return GetArray(key, separators...)
}

// GetArrayOrError retrieves the required slice of strings for key.
// Records a MissingEnvError if the value is absent.
func (v *Validator) GetArrayOrError(key, context string, separators ...string) []string {
	value, err := GetArrayOrError(key, separators...)
	if err != nil {
		v.Add(MissingEnvError{Key: key, Context: context})
		return nil
	}
	return value
}

// GetArrayOrDefault retrieves a slice of strings for key, returning defaultValue if not set.
func (v *Validator) GetArrayOrDefault(key string, defaultValue []string, separators ...string) []string {
	return GetArrayOrDefault(key, defaultValue, separators...)
}

// RequireWhen validates that value is present when condition is true, recording any error.
func (v *Validator) RequireWhen(condition bool, key, context, value string) {
	if condition && strings.TrimSpace(value) == "" {
		v.Add(MissingEnvError{Key: key, Context: context})
	}
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
