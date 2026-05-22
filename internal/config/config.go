package config

import (
	"os"
)

type Environment string

const (
	Development Environment = "development"
	Production  Environment = "production"
)

type Config struct {
	AppEnv         Environment
	Port           string
	DatabaseURL    string
	DatabaseDriver string
	JWTSecret      string
}

func Load() (*Config, error) {
	env := Environment(os.Getenv("APP_ENV"))
	if env == "" {
		env = Development
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

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "change-me-in-production"
	}

	return &Config{
		AppEnv:         env,
		Port:           port,
		DatabaseURL:    dbURL,
		DatabaseDriver: dbDriver,
		JWTSecret:      jwtSecret,
	}, nil
}

func (c *Config) IsDevelopment() bool {
	return c.AppEnv == Development
}

func (c *Config) IsProduction() bool {
	return c.AppEnv == Production
}
