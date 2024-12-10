package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

// Load loads environment variables from a .env file at the given path.
func Load(path string) error {
	// Load the .env file
	err := godotenv.Load(path)
	if err != nil {
		// If loading the .env file fails, return a wrapped error
		return fmt.Errorf("failed to load config from path %s: %w", path, err)
	}

	return nil
}
