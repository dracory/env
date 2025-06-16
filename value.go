package env

import (
	"encoding/base64"
	"os"
	"strings"

	"github.com/gouniverse/envenc"
)

// Env returns the value for an environment key
func Value(key string) string {
	value := os.Getenv(key)

	valueProcessed := envProcess(value)

	return valueProcessed
}

func envProcess(value string) string {
	valueTrimmed := strings.TrimSpace(value)

	if strings.HasPrefix(valueTrimmed, "base64:") {
		valueNoPrefix := strings.TrimPrefix(valueTrimmed, `base64:`)

		valueDecoded, err := base64.URLEncoding.DecodeString(valueNoPrefix)

		if err != nil {
			return err.Error()
		}

		return string(valueDecoded)
	}

	if strings.HasPrefix(valueTrimmed, "obfuscated:") {
		valueNoPrefix := strings.TrimPrefix(valueTrimmed, `obfuscated:`)

		valueDecoded, err := envenc.Deobfuscate(valueNoPrefix)

		if err != nil {
			return err.Error()
		}

		return string(valueDecoded)
	}

	return valueTrimmed
}
