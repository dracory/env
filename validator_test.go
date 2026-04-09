package env

import (
	"os"
	"testing"
)

// helper to set and later unset an env var within a test
func setEnv(t *testing.T, key, value string) {
	t.Helper()
	os.Setenv(key, value)
	t.Cleanup(func() { os.Unsetenv(key) })
}

// ============================================================================
// GetStringOrError
// ============================================================================

func TestValidator_GetStringOrError(t *testing.T) {
	tests := []struct {
		name      string
		setup     func()
		key       string
		wantValue string
		wantErr   bool
	}{
		{
			name:      "present value is returned",
			setup:     func() { os.Setenv("V_STR", "hello") },
			key:       "V_STR",
			wantValue: "hello",
		},
		{
			name:    "missing key records error",
			setup:   func() {},
			key:     "V_STR_MISSING",
			wantErr: true,
		},
		{
			name:    "whitespace-only value records error",
			setup:   func() { os.Setenv("V_STR_WS", "   ") },
			key:     "V_STR_WS",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Unsetenv(tt.key)
			tt.setup()
			t.Cleanup(func() { os.Unsetenv(tt.key) })

			v := &Validator{}
			got := v.GetStringOrError(tt.key, "test context")

			if tt.wantErr {
				if v.Err() == nil {
					t.Error("expected error, got nil")
				}
				if got != "" {
					t.Errorf("expected empty string on error, got %q", got)
				}
			} else {
				if v.Err() != nil {
					t.Errorf("unexpected error: %v", v.Err())
				}
				if got != tt.wantValue {
					t.Errorf("expected %q, got %q", tt.wantValue, got)
				}
			}
		})
	}
}

// ============================================================================
// GetStringOrDefault
// ============================================================================

func TestValidator_GetStringOrDefault(t *testing.T) {
	tests := []struct {
		name         string
		envValue     string
		setEnv       bool
		defaultValue string
		want         string
	}{
		{"returns env value when set", "custom", true, "default", "custom"},
		{"returns default when missing", "", false, "default", "default"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Unsetenv("V_STR_DEF")
			if tt.setEnv {
				setEnv(t, "V_STR_DEF", tt.envValue)
			}
			v := &Validator{}
			got := v.GetStringOrDefault("V_STR_DEF", tt.defaultValue)
			if got != tt.want {
				t.Errorf("expected %q, got %q", tt.want, got)
			}
			if v.Err() != nil {
				t.Errorf("unexpected error: %v", v.Err())
			}
		})
	}
}

// ============================================================================
// GetBoolOrError
// ============================================================================

func TestValidator_GetBoolOrError(t *testing.T) {
	tests := []struct {
		name    string
		envVal  string
		setEnv  bool
		want    bool
		wantErr bool
	}{
		{"true value", "true", true, true, false},
		{"false value", "false", true, false, false},
		{"missing key records error", "", false, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Unsetenv("V_BOOL")
			if tt.setEnv {
				setEnv(t, "V_BOOL", tt.envVal)
			}
			v := &Validator{}
			got := v.GetBoolOrError("V_BOOL", "test context")
			if tt.wantErr {
				if v.Err() == nil {
					t.Error("expected error, got nil")
				}
			} else {
				if v.Err() != nil {
					t.Errorf("unexpected error: %v", v.Err())
				}
				if got != tt.want {
					t.Errorf("expected %v, got %v", tt.want, got)
				}
			}
		})
	}
}

// ============================================================================
// GetBoolOrDefault
// ============================================================================

func TestValidator_GetBoolOrDefault(t *testing.T) {
	tests := []struct {
		name   string
		envVal string
		setEnv bool
		defVal bool
		want   bool
	}{
		{"returns env value", "true", true, false, true},
		{"returns default when missing", "", false, true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Unsetenv("V_BOOL_DEF")
			if tt.setEnv {
				setEnv(t, "V_BOOL_DEF", tt.envVal)
			}
			v := &Validator{}
			got := v.GetBoolOrDefault("V_BOOL_DEF", tt.defVal)
			if got != tt.want {
				t.Errorf("expected %v, got %v", tt.want, got)
			}
		})
	}
}

// ============================================================================
// GetIntOrError
// ============================================================================

func TestValidator_GetIntOrError(t *testing.T) {
	tests := []struct {
		name    string
		envVal  string
		setEnv  bool
		want    int
		wantErr bool
	}{
		{"valid int", "42", true, 42, false},
		{"missing key records error", "", false, 0, true},
		{"invalid int records error", "abc", true, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Unsetenv("V_INT")
			if tt.setEnv {
				setEnv(t, "V_INT", tt.envVal)
			}
			v := &Validator{}
			got := v.GetIntOrError("V_INT", "test context")
			if tt.wantErr {
				if v.Err() == nil {
					t.Error("expected error, got nil")
				}
			} else {
				if v.Err() != nil {
					t.Errorf("unexpected error: %v", v.Err())
				}
				if got != tt.want {
					t.Errorf("expected %d, got %d", tt.want, got)
				}
			}
		})
	}
}

// ============================================================================
// GetIntOrDefault
// ============================================================================

func TestValidator_GetIntOrDefault(t *testing.T) {
	tests := []struct {
		name   string
		envVal string
		setEnv bool
		defVal int
		want   int
	}{
		{"returns env value", "10", true, 99, 10},
		{"returns default when missing", "", false, 99, 99},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Unsetenv("V_INT_DEF")
			if tt.setEnv {
				setEnv(t, "V_INT_DEF", tt.envVal)
			}
			v := &Validator{}
			got := v.GetIntOrDefault("V_INT_DEF", tt.defVal)
			if got != tt.want {
				t.Errorf("expected %d, got %d", tt.want, got)
			}
		})
	}
}

// ============================================================================
// GetFloat64OrError
// ============================================================================

func TestValidator_GetFloat64OrError(t *testing.T) {
	tests := []struct {
		name    string
		envVal  string
		setEnv  bool
		want    float64
		wantErr bool
	}{
		{"valid float", "3.14", true, 3.14, false},
		{"missing key records error", "", false, 0, true},
		{"invalid float records error", "abc", true, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Unsetenv("V_FLOAT")
			if tt.setEnv {
				setEnv(t, "V_FLOAT", tt.envVal)
			}
			v := &Validator{}
			got := v.GetFloat64OrError("V_FLOAT", "test context")
			if tt.wantErr {
				if v.Err() == nil {
					t.Error("expected error, got nil")
				}
			} else {
				if v.Err() != nil {
					t.Errorf("unexpected error: %v", v.Err())
				}
				if got != tt.want {
					t.Errorf("expected %v, got %v", tt.want, got)
				}
			}
		})
	}
}

// ============================================================================
// GetFloat64OrDefault
// ============================================================================

func TestValidator_GetFloat64OrDefault(t *testing.T) {
	tests := []struct {
		name   string
		envVal string
		setEnv bool
		defVal float64
		want   float64
	}{
		{"returns env value", "2.71", true, 0, 2.71},
		{"returns default when missing", "", false, 9.99, 9.99},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Unsetenv("V_FLOAT_DEF")
			if tt.setEnv {
				setEnv(t, "V_FLOAT_DEF", tt.envVal)
			}
			v := &Validator{}
			got := v.GetFloat64OrDefault("V_FLOAT_DEF", tt.defVal)
			if got != tt.want {
				t.Errorf("expected %v, got %v", tt.want, got)
			}
		})
	}
}

// ============================================================================
// RequireWhen
// ============================================================================

func TestValidator_RequireWhen(t *testing.T) {
	tests := []struct {
		name      string
		condition bool
		value     string
		wantErr   bool
	}{
		{"condition false, empty value - no error", false, "", false},
		{"condition true, value present - no error", true, "something", false},
		{"condition true, empty value - records error", true, "", true},
		{"condition true, whitespace value - records error", true, "   ", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Validator{}
			v.RequireWhen(tt.condition, "SOME_KEY", "test context", tt.value)
			if tt.wantErr && v.Err() == nil {
				t.Error("expected error, got nil")
			}
			if !tt.wantErr && v.Err() != nil {
				t.Errorf("unexpected error: %v", v.Err())
			}
		})
	}
}

// ============================================================================
// Err / ValidationError accumulation
// ============================================================================

func TestValidator_AccumulatesMultipleErrors(t *testing.T) {
	os.Unsetenv("V_MISSING_1")
	os.Unsetenv("V_MISSING_2")

	v := &Validator{}
	v.GetStringOrError("V_MISSING_1", "first missing")
	v.GetStringOrError("V_MISSING_2", "second missing")

	err := v.Err()
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	verr, ok := err.(ValidationError)
	if !ok {
		t.Fatalf("expected ValidationError, got %T", err)
	}

	if len(verr.Errors()) != 2 {
		t.Errorf("expected 2 errors, got %d", len(verr.Errors()))
	}
}

func TestValidator_ErrNilWhenNoErrors(t *testing.T) {
	v := &Validator{}
	if v.Err() != nil {
		t.Error("expected nil error on empty validator")
	}
}

func TestMissingEnvError_Format(t *testing.T) {
	tests := []struct {
		name    string
		err     MissingEnvError
		wantMsg string
	}{
		{
			name:    "with context",
			err:     MissingEnvError{Key: "MY_KEY", Context: "set the API key"},
			wantMsg: `env: required variable "MY_KEY" is missing: set the API key`,
		},
		{
			name:    "without context",
			err:     MissingEnvError{Key: "MY_KEY"},
			wantMsg: `env: required variable "MY_KEY" is missing`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.wantMsg {
				t.Errorf("expected %q, got %q", tt.wantMsg, got)
			}
		})
	}
}
