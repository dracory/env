package env

import (
	"os"
	"reflect"
	"testing"
)

func TestGetArray(t *testing.T) {
	tests := []struct {
		name       string
		envVal     string
		setEnv     bool
		separators []string
		want       []string
	}{
		{
			name:   "comma and semicolon separated default",
			envVal: "  apple,  banana ; orange , grape  ",
			setEnv: true,
			want:   []string{"apple", "banana", "orange", "grape"},
		},
		{
			name:       "custom separator pipe",
			envVal:     "apple|banana|orange",
			setEnv:    true,
			separators: []string{"|"},
			want:       []string{"apple", "banana", "orange"},
		},
		{
			name:       "multiple custom separators",
			envVal:     "apple:banana;orange|grape",
			setEnv:    true,
			separators: []string{":", ";", "|"},
			want:       []string{"apple", "banana", "orange", "grape"},
		},
		{
			name:   "auto detect json array",
			envVal: `["apple", "banana", "orange"]`,
			setEnv: true,
			want:   []string{"apple", "banana", "orange"},
		},
		{
			name:   "auto detect json array mixed types",
			envVal: `["apple", 123, true]`,
			setEnv: true,
			want:   []string{"apple", "123", "true"},
		},
		{
			name:   "missing key returns nil",
			setEnv: false,
			want:   nil,
		},
		{
			name:   "empty elements trimmed out",
			envVal: "apple,,  ; banana, ",
			setEnv: true,
			want:   []string{"apple", "banana"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Unsetenv("TEST_ARRAY")
			if tt.setEnv {
				os.Setenv("TEST_ARRAY", tt.envVal)
				defer os.Unsetenv("TEST_ARRAY")
			}

			got := GetArray("TEST_ARRAY", tt.separators...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetArrayOrDefault(t *testing.T) {
	def := []string{"default1", "default2"}

	os.Setenv("TEST_ARRAY_DEF", "a,b,c")
	got := GetArrayOrDefault("TEST_ARRAY_DEF", def)
	if !reflect.DeepEqual(got, []string{"a", "b", "c"}) {
		t.Errorf("expected [a b c], got %v", got)
	}
	os.Unsetenv("TEST_ARRAY_DEF")

	got = GetArrayOrDefault("NON_EXISTENT", def)
	if !reflect.DeepEqual(got, def) {
		t.Errorf("expected default %v, got %v", def, got)
	}
}

func TestGetArrayOrError(t *testing.T) {
	os.Setenv("TEST_ARRAY_ERR", "x,y")
	got, err := GetArrayOrError("TEST_ARRAY_ERR")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(got, []string{"x", "y"}) {
		t.Errorf("expected [x y], got %v", got)
	}
	os.Unsetenv("TEST_ARRAY_ERR")

	_, err = GetArrayOrError("NON_EXISTENT")
	if err == nil {
		t.Error("expected error for missing key, got nil")
	}
}

func TestGetArrayOrPanic(t *testing.T) {
	os.Setenv("TEST_ARRAY_PANIC", "1,2")
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("unexpected panic: %v", r)
		}
	}()
	got := GetArrayOrPanic("TEST_ARRAY_PANIC")
	if !reflect.DeepEqual(got, []string{"1", "2"}) {
		t.Errorf("expected [1 2], got %v", got)
	}
	os.Unsetenv("TEST_ARRAY_PANIC")

	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for missing key, got none")
		}
	}()
	GetArrayOrPanic("NON_EXISTENT")
}

func TestGetJSONArray(t *testing.T) {
	tests := []struct {
		name    string
		envVal  string
		setEnv  bool
		want    []string
		wantErr bool
	}{
		{
			name:   "valid json array strings",
			envVal: `["foo", "bar"]`,
			setEnv: true,
			want:   []string{"foo", "bar"},
		},
		{
			name:   "valid json array numbers and bools",
			envVal: `[1, 2.5, false]`,
			setEnv: true,
			want:   []string{"1", "2.5", "false"},
		},
		{
			name:    "invalid json string",
			envVal:  `[foo, bar]`,
			setEnv:  true,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "missing env var",
			setEnv:  false,
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Unsetenv("TEST_JSON_ARRAY")
			if tt.setEnv {
				os.Setenv("TEST_JSON_ARRAY", tt.envVal)
				defer os.Unsetenv("TEST_JSON_ARRAY")
			}

			got := GetJSONArray("TEST_JSON_ARRAY")
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetJSONArray() = %v, want %v", got, tt.want)
			}

			gotErr, err := GetJSONArrayOrError("TEST_JSON_ARRAY")
			if tt.wantErr {
				if err == nil {
					t.Errorf("GetJSONArrayOrError() expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("GetJSONArrayOrError() unexpected error: %v", err)
				}
				if !reflect.DeepEqual(gotErr, tt.want) {
					t.Errorf("GetJSONArrayOrError() = %v, want %v", gotErr, tt.want)
				}
			}
		})
	}
}

func TestGetJSONArrayOrDefaultAndPanic(t *testing.T) {
	def := []string{"def"}

	// GetJSONArrayOrDefault
	got := GetJSONArrayOrDefault("NON_EXISTENT", def)
	if !reflect.DeepEqual(got, def) {
		t.Errorf("expected default %v, got %v", def, got)
	}

	os.Setenv("TEST_JSON_DEF", `["a"]`)
	got = GetJSONArrayOrDefault("TEST_JSON_DEF", def)
	if !reflect.DeepEqual(got, []string{"a"}) {
		t.Errorf("expected [a], got %v", got)
	}
	os.Unsetenv("TEST_JSON_DEF")

	// GetJSONArrayOrPanic
	os.Setenv("TEST_JSON_PANIC", `["p1", "p2"]`)
	gotPanic := GetJSONArrayOrPanic("TEST_JSON_PANIC")
	if !reflect.DeepEqual(gotPanic, []string{"p1", "p2"}) {
		t.Errorf("expected [p1 p2], got %v", gotPanic)
	}
	os.Unsetenv("TEST_JSON_PANIC")

	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for invalid json array, got none")
		}
	}()
	os.Setenv("TEST_JSON_PANIC_INVALID", `invalid json`)
	defer os.Unsetenv("TEST_JSON_PANIC_INVALID")
	GetJSONArrayOrPanic("TEST_JSON_PANIC_INVALID")
}

func TestValidator_ArrayMethods(t *testing.T) {
	v := &Validator{}

	os.Setenv("V_ARR", "val1, val2")
	defer os.Unsetenv("V_ARR")

	got := v.GetArray("V_ARR")
	if !reflect.DeepEqual(got, []string{"val1", "val2"}) {
		t.Errorf("Validator.GetArray() = %v, want [val1 val2]", got)
	}

	gotReq := v.GetArrayOrError("V_ARR", "context")
	if !reflect.DeepEqual(gotReq, []string{"val1", "val2"}) {
		t.Errorf("Validator.GetArrayOrError() = %v, want [val1 val2]", gotReq)
	}
	if v.Err() != nil {
		t.Errorf("unexpected validator error: %v", v.Err())
	}

	v.GetArrayOrError("NON_EXISTENT_ARR", "missing array")
	if v.Err() == nil {
		t.Error("expected validator error for missing array, got nil")
	}

	gotDef := v.GetArrayOrDefault("NON_EXISTENT_ARR", []string{"default"})
	if !reflect.DeepEqual(gotDef, []string{"default"}) {
		t.Errorf("Validator.GetArrayOrDefault() = %v, want [default]", gotDef)
	}
}

func TestGetArrayMapped(t *testing.T) {
	os.Setenv("TEST_TRANSFORM", "  Password, WebAuthn, OAUTH ")
	defer os.Unsetenv("TEST_TRANSFORM")

	// Custom mapper
	clean := GetArrayMapped("TEST_TRANSFORM", func(s string) string {
		return "role_" + s
	})
	want := []string{"role_Password", "role_WebAuthn", "role_OAUTH"}
	if !reflect.DeepEqual(clean, want) {
		t.Errorf("GetArrayMapped() = %v, want %v", clean, want)
	}

	// Nil mapper
	raw := GetArrayMapped("TEST_TRANSFORM", nil)
	wantRaw := []string{"Password", "WebAuthn", "OAUTH"}
	if !reflect.DeepEqual(raw, wantRaw) {
		t.Errorf("GetArrayMapped(nil) = %v, want %v", raw, wantRaw)
	}

	// Missing key
	if got := GetArrayMapped("NON_EXISTENT_MAP", func(s string) string { return s }); got != nil {
		t.Errorf("GetArrayMapped(missing) = %v, want nil", got)
	}
}

func TestGetArrayMappedOrDefaultAndOrErrorAndPanic(t *testing.T) {
	def := []string{"default"}

	// Default
	gotDef := GetArrayMappedOrDefault("NON_EXISTENT", def, nil)
	if !reflect.DeepEqual(gotDef, def) {
		t.Errorf("GetArrayMappedOrDefault(missing) = %v, want %v", gotDef, def)
	}

	os.Setenv("TEST_MAP_DEF", "a,b")
	defer os.Unsetenv("TEST_MAP_DEF")
	gotDefPresent := GetArrayMappedOrDefault("TEST_MAP_DEF", def, func(s string) string { return s + "!" })
	if !reflect.DeepEqual(gotDefPresent, []string{"a!", "b!"}) {
		t.Errorf("GetArrayMappedOrDefault(present) = %v, want [a! b!]", gotDefPresent)
	}

	// Error
	_, err := GetArrayMappedOrError("NON_EXISTENT", nil)
	if err == nil {
		t.Error("GetArrayMappedOrError(missing) expected error, got nil")
	}

	// Panic
	defer func() {
		if r := recover(); r == nil {
			t.Error("GetArrayMappedOrPanic(missing) expected panic, got none")
		}
	}()
	GetArrayMappedOrPanic("NON_EXISTENT", nil)
}

func TestGetArrayLower(t *testing.T) {
	os.Setenv("TEST_LOWER", "Password, WebAuthn, OAUTH")
	defer os.Unsetenv("TEST_LOWER")

	got := GetArrayLower("TEST_LOWER")
	want := []string{"password", "webauthn", "oauth"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetArrayLower() = %v, want %v", got, want)
	}

	// JSON array
	os.Setenv("TEST_LOWER_JSON", `["Foo", "BAR"]`)
	defer os.Unsetenv("TEST_LOWER_JSON")
	gotJSON := GetArrayLower("TEST_LOWER_JSON")
	wantJSON := []string{"foo", "bar"}
	if !reflect.DeepEqual(gotJSON, wantJSON) {
		t.Errorf("GetArrayLower(JSON) = %v, want %v", gotJSON, wantJSON)
	}

	// Default
	def := []string{"def"}
	gotDef := GetArrayLowerOrDefault("NON_EXISTENT", def)
	if !reflect.DeepEqual(gotDef, def) {
		t.Errorf("GetArrayLowerOrDefault() = %v, want %v", gotDef, def)
	}

	// Error
	_, err := GetArrayLowerOrError("NON_EXISTENT")
	if err == nil {
		t.Error("GetArrayLowerOrError(missing) expected error, got nil")
	}

	// Panic
	defer func() {
		if r := recover(); r == nil {
			t.Error("GetArrayLowerOrPanic(missing) expected panic, got none")
		}
	}()
	GetArrayLowerOrPanic("NON_EXISTENT")
}

func TestGetArrayUpper(t *testing.T) {
	os.Setenv("TEST_UPPER", "Password, WebAuthn, oauth")
	defer os.Unsetenv("TEST_UPPER")

	got := GetArrayUpper("TEST_UPPER")
	want := []string{"PASSWORD", "WEBAUTHN", "OAUTH"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetArrayUpper() = %v, want %v", got, want)
	}

	// Default
	def := []string{"DEF"}
	gotDef := GetArrayUpperOrDefault("NON_EXISTENT", def)
	if !reflect.DeepEqual(gotDef, def) {
		t.Errorf("GetArrayUpperOrDefault() = %v, want %v", gotDef, def)
	}

	// Error
	_, err := GetArrayUpperOrError("NON_EXISTENT")
	if err == nil {
		t.Error("GetArrayUpperOrError(missing) expected error, got nil")
	}

	// Panic
	defer func() {
		if r := recover(); r == nil {
			t.Error("GetArrayUpperOrPanic(missing) expected panic, got none")
		}
	}()
	GetArrayUpperOrPanic("NON_EXISTENT")
}

func TestValidator_ArrayMappedAndCaseMethods(t *testing.T) {
	v := &Validator{}

	os.Setenv("V_MAP", "Foo, Bar")
	defer os.Unsetenv("V_MAP")

	gotMapped := v.GetArrayMapped("V_MAP", func(s string) string { return s + "1" })
	if !reflect.DeepEqual(gotMapped, []string{"Foo1", "Bar1"}) {
		t.Errorf("Validator.GetArrayMapped() = %v, want [Foo1 Bar1]", gotMapped)
	}

	gotLower := v.GetArrayLower("V_MAP")
	if !reflect.DeepEqual(gotLower, []string{"foo", "bar"}) {
		t.Errorf("Validator.GetArrayLower() = %v, want [foo bar]", gotLower)
	}

	gotUpper := v.GetArrayUpper("V_MAP")
	if !reflect.DeepEqual(gotUpper, []string{"FOO", "BAR"}) {
		t.Errorf("Validator.GetArrayUpper() = %v, want [FOO BAR]", gotUpper)
	}

	// OrError methods
	v.GetArrayMappedOrError("NON_EXISTENT", "ctx1", nil)
	v.GetArrayLowerOrError("NON_EXISTENT", "ctx2")
	v.GetArrayUpperOrError("NON_EXISTENT", "ctx3")

	if v.Err() == nil {
		t.Error("expected validator errors, got nil")
	}

	// OrDefault methods
	def := []string{"def"}
	if got := v.GetArrayMappedOrDefault("NON_EXISTENT", def, nil); !reflect.DeepEqual(got, def) {
		t.Errorf("Validator.GetArrayMappedOrDefault() = %v, want %v", got, def)
	}
	if got := v.GetArrayLowerOrDefault("NON_EXISTENT", def); !reflect.DeepEqual(got, def) {
		t.Errorf("Validator.GetArrayLowerOrDefault() = %v, want %v", got, def)
	}
	if got := v.GetArrayUpperOrDefault("NON_EXISTENT", def); !reflect.DeepEqual(got, def) {
		t.Errorf("Validator.GetArrayUpperOrDefault() = %v, want %v", got, def)
	}
}

func TestGetJSONArrayMappedLowerUpper(t *testing.T) {
	os.Setenv("TEST_JSON_MAP", `["Alpha", "Beta"]`)
	defer os.Unsetenv("TEST_JSON_MAP")

	// GetJSONArrayMapped
	gotMap := GetJSONArrayMapped("TEST_JSON_MAP", func(s string) string { return s + "!" })
	if !reflect.DeepEqual(gotMap, []string{"Alpha!", "Beta!"}) {
		t.Errorf("GetJSONArrayMapped() = %v, want [Alpha! Beta!]", gotMap)
	}

	// GetJSONArrayLower
	gotLower := GetJSONArrayLower("TEST_JSON_MAP")
	if !reflect.DeepEqual(gotLower, []string{"alpha", "beta"}) {
		t.Errorf("GetJSONArrayLower() = %v, want [alpha beta]", gotLower)
	}

	// GetJSONArrayUpper
	gotUpper := GetJSONArrayUpper("TEST_JSON_MAP")
	if !reflect.DeepEqual(gotUpper, []string{"ALPHA", "BETA"}) {
		t.Errorf("GetJSONArrayUpper() = %v, want [ALPHA BETA]", gotUpper)
	}

	// Invalid JSON / Missing key returns nil / default / error / panic
	def := []string{"default"}
	if got := GetJSONArrayMapped("NON_EXISTENT_JSON", nil); got != nil {
		t.Errorf("GetJSONArrayMapped(missing) = %v, want nil", got)
	}
	if got := GetJSONArrayMappedOrDefault("NON_EXISTENT_JSON", def, nil); !reflect.DeepEqual(got, def) {
		t.Errorf("GetJSONArrayMappedOrDefault(missing) = %v, want %v", got, def)
	}
	if _, err := GetJSONArrayMappedOrError("NON_EXISTENT_JSON", nil); err == nil {
		t.Error("GetJSONArrayMappedOrError(missing) expected error, got nil")
	}

	if got := GetJSONArrayLowerOrDefault("NON_EXISTENT_JSON", def); !reflect.DeepEqual(got, def) {
		t.Errorf("GetJSONArrayLowerOrDefault(missing) = %v, want %v", got, def)
	}
	if _, err := GetJSONArrayLowerOrError("NON_EXISTENT_JSON"); err == nil {
		t.Error("GetJSONArrayLowerOrError(missing) expected error, got nil")
	}

	if got := GetJSONArrayUpperOrDefault("NON_EXISTENT_JSON", def); !reflect.DeepEqual(got, def) {
		t.Errorf("GetJSONArrayUpperOrDefault(missing) = %v, want %v", got, def)
	}
	if _, err := GetJSONArrayUpperOrError("NON_EXISTENT_JSON"); err == nil {
		t.Error("GetJSONArrayUpperOrError(missing) expected error, got nil")
	}

	// Panics
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("GetJSONArrayMappedOrPanic expected panic, got none")
			}
		}()
		GetJSONArrayMappedOrPanic("NON_EXISTENT_JSON", nil)
	}()

	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("GetJSONArrayLowerOrPanic expected panic, got none")
			}
		}()
		GetJSONArrayLowerOrPanic("NON_EXISTENT_JSON")
	}()

	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("GetJSONArrayUpperOrPanic expected panic, got none")
			}
		}()
		GetJSONArrayUpperOrPanic("NON_EXISTENT_JSON")
	}()
}

func TestValidator_JSONArrayMappedAndCaseMethods(t *testing.T) {
	v := &Validator{}

	os.Setenv("V_JSON_MAP", `["Foo", "Bar"]`)
	defer os.Unsetenv("V_JSON_MAP")

	gotMapped := v.GetJSONArrayMapped("V_JSON_MAP", func(s string) string { return s + "2" })
	if !reflect.DeepEqual(gotMapped, []string{"Foo2", "Bar2"}) {
		t.Errorf("Validator.GetJSONArrayMapped() = %v, want [Foo2 Bar2]", gotMapped)
	}

	gotLower := v.GetJSONArrayLower("V_JSON_MAP")
	if !reflect.DeepEqual(gotLower, []string{"foo", "bar"}) {
		t.Errorf("Validator.GetJSONArrayLower() = %v, want [foo bar]", gotLower)
	}

	gotUpper := v.GetJSONArrayUpper("V_JSON_MAP")
	if !reflect.DeepEqual(gotUpper, []string{"FOO", "BAR"}) {
		t.Errorf("Validator.GetJSONArrayUpper() = %v, want [FOO BAR]", gotUpper)
	}

	// OrError methods
	v.GetJSONArrayMappedOrError("NON_EXISTENT", "ctx1", nil)
	v.GetJSONArrayLowerOrError("NON_EXISTENT", "ctx2")
	v.GetJSONArrayUpperOrError("NON_EXISTENT", "ctx3")

	if v.Err() == nil {
		t.Error("expected validator errors, got nil")
	}

	// OrDefault methods
	def := []string{"def"}
	if got := v.GetJSONArrayMappedOrDefault("NON_EXISTENT", def, nil); !reflect.DeepEqual(got, def) {
		t.Errorf("Validator.GetJSONArrayMappedOrDefault() = %v, want %v", got, def)
	}
	if got := v.GetJSONArrayLowerOrDefault("NON_EXISTENT", def); !reflect.DeepEqual(got, def) {
		t.Errorf("Validator.GetJSONArrayLowerOrDefault() = %v, want %v", got, def)
	}
	if got := v.GetJSONArrayUpperOrDefault("NON_EXISTENT", def); !reflect.DeepEqual(got, def) {
		t.Errorf("Validator.GetJSONArrayUpperOrDefault() = %v, want %v", got, def)
	}
}
