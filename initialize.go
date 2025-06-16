package env

import (
	"log"

	"github.com/joho/godotenv"
)

// Initialize initializes the environment variables
func Initialize(envFilePath ...string) {
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
