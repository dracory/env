package env

import (
	"os"
	"testing"
)

func TestGetBool(t *testing.T) {
	os.Setenv("TEST_BOOL_TRUE", "true")
	value := GetBool("TEST_BOOL_TRUE")
	if value != true {
		t.Errorf("Expected true, got %v", value)
	}
	os.Unsetenv("TEST_BOOL_TRUE")

	os.Setenv("TEST_BOOL_FALSE", "false")
	value = GetBool("TEST_BOOL_FALSE")
	if value != false {
		t.Errorf("Expected false, got %v", value)
	}
	os.Unsetenv("TEST_BOOL_FALSE")

	value = GetBool("NON_EXISTENT")
	if value != false {
		t.Errorf("Expected false, got %v", value)
	}

	os.Setenv("TEST_BOOL_INVALID", "abc")
	value = GetBool("TEST_BOOL_INVALID")
	if value != false {
		t.Errorf("Expected false, got %v", value)
	}
	os.Unsetenv("TEST_BOOL_INVALID")
}

func TestGetBoolOrDefault(t *testing.T) {
	os.Setenv("TEST_BOOL_TRUE", "true")
	value := GetBoolOrDefault("TEST_BOOL_TRUE", false)
	if value != true {
		t.Errorf("Expected true, got %v", value)
	}
	os.Unsetenv("TEST_BOOL_TRUE")

	value = GetBoolOrDefault("NON_EXISTENT", true)
	if value != true {
		t.Errorf("Expected true, got %v", value)
	}
}

func TestGetBoolOrError(t *testing.T) {
	os.Setenv("TEST_BOOL_TRUE", "true")
	value, err := GetBoolOrError("TEST_BOOL_TRUE")
	if err != nil {
		t.Errorf("Expected nil error, got '%s'", err)
	}
	if value != true {
		t.Errorf("Expected true, got %v", value)
	}
	os.Unsetenv("TEST_BOOL_TRUE")

	_, err = GetBoolOrError("NON_EXISTENT")
	if err == nil {
		t.Error("Expected error, got nil")
	}

	os.Setenv("TEST_BOOL_INVALID", "abc")
	_, err = GetBoolOrError("TEST_BOOL_INVALID")
	if err == nil {
		t.Error("Expected error, got nil")
	}
	os.Unsetenv("TEST_BOOL_INVALID")
}

func TestGetBoolOrPanic(t *testing.T) {
	os.Setenv("TEST_BOOL_TRUE", "true")
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code panicked: %v", r)
		}
	}()
	GetBoolOrPanic("TEST_BOOL_TRUE")
	os.Unsetenv("TEST_BOOL_TRUE")

	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic")
		}
	}()
	GetBoolOrPanic("NON_EXISTENT")
}

func TestGetBool_EdgeCases(t *testing.T) {
	type tc struct {
		val       string
		expect    bool
		shouldErr bool
	}

	cases := []tc{
		// Truthy tokens
		{"T", true, false},
		{"TRUE", true, false},
		{"Yes", true, false},
		{"1", true, false},
		// Falsy tokens
		{"F", false, false},
		{"FALSE", false, false},
		{"No", false, false},
		{"0", false, false},
		// Whitespace around tokens
		{"  yes  ", true, false},
		{"  no  ", false, false},
		{"  1  ", true, false},
		{"  0  ", false, false},
		// Numeric values
		{"2", true, false},
		{"-2", false, false},
		{"0", false, false},
		{"0.00", false, false},
		{"+0", false, false},
		{"-0", false, false},
		{"+2.5", true, false},
		{"-3.14", false, false},
		{"1e-3", true, false},
		{"-1e-3", false, false},
		// Invalid numerics/tokens should error in GetBoolOrError and be false in GetBool
		{"NaN", false, true},
		{"Inf", false, true},
		{"-Inf", false, true},
		{"maybe", false, true},
	}

	const key = "TEST_BOOL_EDGE"
	for i, c := range cases {
		// GetBoolOrError behavior
		os.Setenv(key, c.val)
		got, err := GetBoolOrError(key)
		if c.shouldErr {
			if err == nil {
				t.Fatalf("case %d (%q): expected error, got nil (value=%v)", i, c.val, got)
			}
		} else {
			if err != nil {
				t.Fatalf("case %d (%q): unexpected error: %v", i, c.val, err)
			}
			if got != c.expect {
				t.Fatalf("case %d (%q): expected %v, got %v", i, c.val, c.expect, got)
			}
		}

		// GetBool should return false on error cases, or the value otherwise
		b := GetBool(key)
		if c.shouldErr {
			if b != false {
				t.Fatalf("case %d (%q): GetBool expected false on error case, got %v", i, c.val, b)
			}
		} else {
			if b != c.expect {
				t.Fatalf("case %d (%q): GetBool expected %v, got %v", i, c.val, c.expect, b)
			}
		}

		os.Unsetenv(key)
	}
}
