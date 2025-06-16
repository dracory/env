package env

import (
	"os"
	"strings"
	"testing"
)

func TestValue_Plain(t *testing.T) {
	// Test with a plain environment variable
	os.Setenv("TEST_KEY", "test_value")
	if value := Value("TEST_KEY"); value != "test_value" {
		t.Errorf("Expected 'test_value', got '%s'", value)
	}
}

func TestValue_Base64(t *testing.T) {
	// Test with a base64 encoded environment variable
	os.Setenv("TEST_KEY_BASE64", "base64:dGVzdF92YWx1ZQ==") // "test_value" in base64
	if value := Value("TEST_KEY_BASE64"); value != "test_value" {
		t.Errorf("Expected 'test_value', got '%s'", value)
	}
}

func TestValue_InvalidBase64(t *testing.T) {
	// Test with an invalid base64 encoded environment variable
	os.Setenv("TEST_KEY_INVALID_BASE64", "base64:invalid_value==")
	if value := Value("TEST_KEY_INVALID_BASE64"); !strings.Contains(value, "illegal base64 data") {
		t.Errorf("Expected an error message containing 'illegal base64 data', got '%s'", value)
	}
}

func TestValue_Empty(t *testing.T) {
	// Test with an empty environment variable
	os.Setenv("TEST_KEY_EMPTY", "")
	if value := Value("TEST_KEY_EMPTY"); value != "" {
		t.Errorf("Expected an empty string, got '%s'", value)
	}
}
