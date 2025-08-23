package env

import (
	"os"
	"testing"
)

func TestGetFloat64(t *testing.T) {
	os.Setenv("TEST_FLOAT", "123.45")
	value := GetFloat64("TEST_FLOAT")
	if value != 123.45 {
		t.Errorf("Expected 123.45, got %f", value)
	}
	os.Unsetenv("TEST_FLOAT")

	value = GetFloat64("NON_EXISTENT")
	if value != 0.0 {
		t.Errorf("Expected 0.0, got %f", value)
	}

	os.Setenv("TEST_FLOAT_INVALID", "abc")
	value = GetFloat64("TEST_FLOAT_INVALID")
	if value != 0.0 {
		t.Errorf("Expected 0.0, got %f", value)
	}
	os.Unsetenv("TEST_FLOAT_INVALID")
}

func TestGetFloat64OrDefault(t *testing.T) {
	os.Setenv("TEST_FLOAT", "123.45")
	value := GetFloat64OrDefault("TEST_FLOAT", 678.9)
	if value != 123.45 {
		t.Errorf("Expected 123.45, got %f", value)
	}
	os.Unsetenv("TEST_FLOAT")

	value = GetFloat64OrDefault("NON_EXISTENT", 678.9)
	if value != 678.9 {
		t.Errorf("Expected 678.9, got %f", value)
	}
}

func TestGetFloat64OrError(t *testing.T) {
	os.Setenv("TEST_FLOAT", "123.45")
	value, err := GetFloat64OrError("TEST_FLOAT")
	if err != nil {
		t.Errorf("Expected nil error, got '%s'", err)
	}
	if value != 123.45 {
		t.Errorf("Expected 123.45, got %f", value)
	}
	os.Unsetenv("TEST_FLOAT")

	_, err = GetFloat64OrError("NON_EXISTENT")
	if err == nil {
		t.Error("Expected error, got nil")
	}

	os.Setenv("TEST_FLOAT_INVALID", "abc")
	_, err = GetFloat64OrError("TEST_FLOAT_INVALID")
	if err == nil {
		t.Error("Expected error, got nil")
	}
	os.Unsetenv("TEST_FLOAT_INVALID")
}

func TestGetFloat64OrPanic(t *testing.T) {
	os.Setenv("TEST_FLOAT", "123.45")
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code panicked: %v", r)
		}
	}()
	GetFloat64OrPanic("TEST_FLOAT")
	os.Unsetenv("TEST_FLOAT")

	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic")
		}
	}()
	GetFloat64OrPanic("NON_EXISTENT")
}
