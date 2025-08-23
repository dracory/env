package env

import (
	"os"
	"testing"
)

func TestGetInt(t *testing.T) {
	os.Setenv("TEST_INT", "123")
	value := GetInt("TEST_INT")
	if value != 123 {
		t.Errorf("Expected 123, got %d", value)
	}
	os.Unsetenv("TEST_INT")

	value = GetInt("NON_EXISTENT")
	if value != 0 {
		t.Errorf("Expected 0, got %d", value)
	}

	os.Setenv("TEST_INT_INVALID", "abc")
	value = GetInt("TEST_INT_INVALID")
	if value != 0 {
		t.Errorf("Expected 0, got %d", value)
	}
	os.Unsetenv("TEST_INT_INVALID")
}

func TestGetIntOrDefault(t *testing.T) {
	os.Setenv("TEST_INT", "123")
	value := GetIntOrDefault("TEST_INT", 456)
	if value != 123 {
		t.Errorf("Expected 123, got %d", value)
	}
	os.Unsetenv("TEST_INT")

	value = GetIntOrDefault("NON_EXISTENT", 456)
	if value != 456 {
		t.Errorf("Expected 456, got %d", value)
	}
}

func TestGetIntOrError(t *testing.T) {
	os.Setenv("TEST_INT", "123")
	value, err := GetIntOrError("TEST_INT")
	if err != nil {
		t.Errorf("Expected nil error, got '%s'", err)
	}
	if value != 123 {
		t.Errorf("Expected 123, got %d", value)
	}
	os.Unsetenv("TEST_INT")

	_, err = GetIntOrError("NON_EXISTENT")
	if err == nil {
		t.Error("Expected error, got nil")
	}

	os.Setenv("TEST_INT_INVALID", "abc")
	_, err = GetIntOrError("TEST_INT_INVALID")
	if err == nil {
		t.Error("Expected error, got nil")
	}
	os.Unsetenv("TEST_INT_INVALID")
}

func TestGetIntOrPanic(t *testing.T) {
	os.Setenv("TEST_INT", "123")
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code panicked: %v", r)
		}
	}()
	GetIntOrPanic("TEST_INT")
	os.Unsetenv("TEST_INT")

	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic")
		}
	}()
	GetIntOrPanic("NON_EXISTENT")
}
