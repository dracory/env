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
