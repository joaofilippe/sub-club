package config

import (
	"os"
)

// Config holds all configuration for our application
type Config struct {
	DatabaseURL string
}

// Load reads configuration from environment variables
func Load() (*Config, error) {
	// For now, we can simple read from env directly.
	// In the future we might use something like viper or godotenv.
	dbURL := os.Getenv("DATABASE_URL")

	// Default to a local connection string if not set (useful for dev)
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/subclub?sslmode=disable"
	}

	return &Config{
		DatabaseURL: dbURL,
	}, nil
}
