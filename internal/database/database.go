package database

import (
	"context"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // Standard library driver
	"github.com/jmoiron/sqlx"
)

// NewConnection creates a connection pool to the database
func NewConnection(ctx context.Context, connectionString string) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(ctx, "pgx", connectionString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	// Optimization: Configure pool settings here if needed
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(time.Hour)

	return db, nil
}
