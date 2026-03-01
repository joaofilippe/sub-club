package config

import (
	"os"
)

type Config struct {
	AppEnv         string
	Port           string
	DatabaseURL    string
	DatabaseDriver string
}

func Load() (*Config, error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbDriver := os.Getenv("DATABASE_DRIVER")
	if dbDriver == "" {
		dbDriver = "pgx"
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/subclub?sslmode=disable"
	}

	return &Config{
		AppEnv:         env,
		Port:           port,
		DatabaseURL:    dbURL,
		DatabaseDriver: dbDriver,
	}, nil
}
