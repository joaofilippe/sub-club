package database

import (
	"context"
	"fmt"
	"sync"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var (
	connection *Connection
	once       sync.Once
	connErr    error
)

type Connection struct {
	db *sqlx.DB
}

// NewConnection garante que a conexão com o banco seja inicializada
// apenas uma única vez, independente do número de chamadas.
func NewConnection(ctx context.Context, driver, url string) (*Connection, error) {
	once.Do(func() {
		var db *sqlx.DB
		db, connErr = sqlx.ConnectContext(ctx, driver, url)
		if connErr != nil {
			connErr = fmt.Errorf("unable to connect to database: %w", connErr)
			return
		}

		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(2)
		db.SetConnMaxLifetime(time.Hour)

		connection = &Connection{
			db: db,
		}
	})

	if connErr != nil {
		return nil, connErr
	}

	return connection, nil
}

func (c *Connection) GetDB() *sqlx.DB {
	return c.db
}

func (c *Connection) Close() error {
	if c.db != nil {
		return c.db.Close()
	}
	return nil
}
