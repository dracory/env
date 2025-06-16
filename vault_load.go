package env

import (
	"errors"
	"os"

	"github.com/gouniverse/envenc"
)

// EncInitialize initializes environment variables from a vault file or content using the provided password.
// It takes an options struct containing the password and either a vault file path or vault content.
// It returns an error if the initialization fails.
func VaultLoad(options struct {
	Password      string
	VaultFilePath string
	VaultContent  string
}) error {
	if options.Password == "" {
		return errors.New("password is required")
	}

	if options.VaultFilePath == "" && options.VaultContent == "" {
		return errors.New("vault file path or vault content is required")
	}

	if options.VaultFilePath != "" && options.VaultContent != "" {
		return errors.New("vault file path and vault content are mutually exclusive")
	}

	var err error
	keys := map[string]string{}

	if options.VaultFilePath != "" {
		if !fileExists(options.VaultFilePath) {
			return errors.New("Vault file not found: " + options.VaultFilePath)
		}

		keys, err = envenc.KeyListFromFile(options.VaultFilePath, options.Password)

		if err != nil {
			return err
		}
	}

	if options.VaultContent != "" {
		keys, err = envenc.KeyListFromString(options.VaultContent, options.Password)

		if err != nil {
			return err
		}
	}

	for k, v := range keys {
		err := os.Setenv(k, v)

		if err != nil {
			return err
		}
	}

	return nil
}
