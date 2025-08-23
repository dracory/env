package env

import (
	"os"
	"testing"
)

func TestGetString(t *testing.T) {
	os.Setenv("TEST_STRING", "hello")
	value := GetString("TEST_STRING")
	if value != "hello" {
		t.Errorf("Expected 'hello', got '%s'", value)
	}
	os.Unsetenv("TEST_STRING")

	value = GetString("NON_EXISTENT")
	if value != "" {
		t.Errorf("Expected '', got '%s'", value)
	}
}

func TestGetStringOrDefault(t *testing.T) {
	os.Setenv("TEST_STRING", "hello")
	value := GetStringOrDefault("TEST_STRING", "default")
	if value != "hello" {
		t.Errorf("Expected 'hello', got '%s'", value)
	}
	os.Unsetenv("TEST_STRING")

	value = GetStringOrDefault("NON_EXISTENT", "default")
	if value != "default" {
		t.Errorf("Expected 'default', got '%s'", value)
	}
}

func TestGetStringOrError(t *testing.T) {
	os.Setenv("TEST_STRING", "hello")
	value, err := GetStringOrError("TEST_STRING")
	if err != nil {
		t.Errorf("Expected nil error, got '%s'", err)
	}
	if value != "hello" {
		t.Errorf("Expected 'hello', got '%s'", value)
	}
	os.Unsetenv("TEST_STRING")

	_, err = GetStringOrError("NON_EXISTENT")
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestGetStringOrPanic(t *testing.T) {
	os.Setenv("TEST_STRING", "hello")
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code panicked: %v", r)
		}
	}()
	GetStringOrPanic("TEST_STRING")
	os.Unsetenv("TEST_STRING")

	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic")
		}
	}()
	GetStringOrPanic("NON_EXISTENT")
}
