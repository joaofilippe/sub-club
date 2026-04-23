package repository_test

import (
	"database/sql"
	"database/sql/driver"

	sqlite "modernc.org/sqlite"
)

func init() {
	sql.Register("sqlite3", &sqliteDriver{new(sqlite.Driver)})
}

type sqliteDriver struct {
	driver.Driver
}
