package repository_test

import (
	"database/sql"
	"database/sql/driver"

	sqlite "modernc.org/sqlite"
)

func init() {
	// ent expects the SQLite dialect to be registered as "sqlite3".
	// modernc.org/sqlite registers as "sqlite", so we alias it here.
	sql.Register("sqlite3", &sqliteDriver{new(sqlite.Driver)})
}

type sqliteDriver struct {
	driver.Driver
}
