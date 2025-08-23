package env

import (
	"log"

	"github.com/joho/godotenv"
)

// Load loads environment variables from .env files.
//
// If no paths are provided, it will try to load the default .env file.
//
// Parameters:
//
//	...envFilePath: The paths to the .env files to load.
//
// Returns:
//
//	None.
func Load(envFilePath ...string) {
	paths := []string{".env"}

	paths = append(paths, envFilePath...)

	for _, path := range paths {
		if fileExists(path) {
			err := godotenv.Load(path)
			if err != nil {
				log.Fatal("Error loading " + path + " file")
			}
		}
	}
}
