package config

import (
	"os"
)

type Config struct {
	DatabaseURL string
}

func Load() (*Config, error) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/subclub?sslmode=disable"
	}

	return &Config{
		DatabaseURL: dbURL,
	}, nil
}
