package database

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

func Open(path string) (*bun.DB, error) {
	sqldb, err := sql.Open(sqliteshim.ShimName, path)
	if err != nil {
		return nil, err
	}
	// SQLite allows a single writer; more connections cause "database is locked" errors.
	sqldb.SetMaxOpenConns(1)
	return bun.NewDB(sqldb, sqlitedialect.New()), nil
}
