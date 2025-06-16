package env

import (
	"os"
	"testing"

	"github.com/gouniverse/envenc"
)

func TestVaultLoad(t *testing.T) {
	password := "password%%1234567890"

	// Create a temporary .vault file, so that we can use the name later
	tempFile, err := os.CreateTemp("", "test.vault")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	tempFile.Close()

	// Remove the temporary .vault file, so that we can use the same path later
	err = os.Remove(tempFile.Name())

	if err != nil {
		t.Fatalf("Error removing temporary file: %v", err)
	}

	err = envenc.Init(tempFile.Name(), password)

	defer os.Remove(tempFile.Name())

	if err != nil {
		t.Fatal(err.Error())
	}

	// Write some content to the .env file
	err = envenc.KeySet(tempFile.Name(), password, "TEST_VAULT_VAR", "test_vault_value")
	if err != nil {
		t.Fatalf("Error writing to temporary file: %v", err)
	}

	// Call the EnvInitialize function
	VaultLoad(struct {
		Password      string
		VaultFilePath string
		VaultContent  string
	}{
		Password:      password,
		VaultFilePath: tempFile.Name(),
	})

	// Assert that the environment variable was loaded
	if os.Getenv("TEST_VAULT_VAR") != "test_vault_value" {
		t.Errorf("Expected TEST_VAULT_VAR to be 'test_vault_value', but got '%s'", os.Getenv("TEST_VAULT_VAR"))
	}
}
