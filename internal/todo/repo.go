package todo

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

const createTodoQuery = `
insert into
todos
	(title, description, completed, created_at, updated_at)
values
	(?, 	?,			 ?, 		?,			?)
returning *;
`

func CreateTodo(db *bun.DB, todo Todo) (*Todo, error) {
	newTodo := &todo

	now := time.Now().UTC()
	newTodo.CreatedAt = now
	newTodo.UpdatedAt = now

	_, err := db.NewInsert().
		Model(newTodo).
		Exec(context.Background())
	if err != nil {
		return nil, err
	}

	return newTodo, nil
}

const getAllTodosQuery = `
select * from todos
order by
	completed asc,
	created_at desc;
`

func GetTodos(db *bun.DB) ([]Todo, error) {
	var todos []Todo
	err := db.NewSelect().
		Model(&todos).
		OrderExpr("completed ASC, created_at DESC").
		Scan(context.Background())
	if err != nil {
		return []Todo{}, err
	}
	return todos, nil
}

const updateTodoQuery = `
update todos
set
    completed = ?,
    completed_at = ?,
    updated_at = ?
where
    id = ?
returning *;
`

func UpdateTodo(db *bun.DB, updateTodo Todo) (*Todo, error) {
	var todo Todo

	err := db.NewUpdate().
		Model(&todo).
		Set("completed = ?, completed_at = ?, updated_at = ?",
			updateTodo.Completed,
			updateTodo.CompletedAt,
			updateTodo.UpdatedAt,
		).
		Where("id = ?", updateTodo.ID).
		Returning("*").
		Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

const deleteTodoQuery = `delete from todos where id = ?;`

func DeleteTodo(db *bun.DB, id uint) error {
	_, err := db.NewDelete().Model((*Todo)(nil)).Where("id = ?", id).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
