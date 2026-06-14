package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// EnsureDatabase connects to the postgres system database and creates the
// target database if it does not exist. This is safe to call on every startup.
func EnsureDatabase(ctx context.Context, databaseURL string) error {
	u, err := url.Parse(databaseURL)
	if err != nil {
		return fmt.Errorf("parse database URL: %w", err)
	}

	dbName := u.Path
	if len(dbName) > 0 && dbName[0] == '/' {
		dbName = dbName[1:]
	}

	adminURL := *u
	adminURL.Path = "/postgres"

	db, err := sql.Open("pgx", adminURL.String())
	if err != nil {
		return fmt.Errorf("open admin connection: %w", err)
	}
	defer db.Close()

	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("ping postgres: %w", err)
	}

	var exists bool
	err = db.QueryRowContext(ctx,
		"SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)", dbName,
	).Scan(&exists)
	if err != nil {
		return fmt.Errorf("check database existence: %w", err)
	}

	if exists {
		return nil
	}

	log.Printf("[Database] Database %q not found — creating...", dbName)
	if _, err := db.ExecContext(ctx, fmt.Sprintf(`CREATE DATABASE %q`, dbName)); err != nil {
		return fmt.Errorf("create database %q: %w", dbName, err)
	}
	log.Printf("[Database] Database %q created.", dbName)
	return nil
}
