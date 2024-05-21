package data

import (
	"context"

	"github.com/ewertonls/gotodo/internal/todo"
	"github.com/uptrace/bun"
)

func CreateTables(db *bun.DB) error {
	_, err := db.NewCreateTable().
		Model((*todo.Todo)(nil)).
		Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
