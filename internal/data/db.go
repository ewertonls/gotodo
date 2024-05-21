package data

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

func NewDB() (*bun.DB, error) {
	sqliteDB, err := sql.Open(sqliteshim.ShimName, "file:todos.db?cache=shared")
	if err != nil {
		return nil, err
	}

	db := bun.NewDB(sqliteDB, sqlitedialect.New())
	db.AddQueryHook(bundebug.NewQueryHook())
	return db, nil
}
