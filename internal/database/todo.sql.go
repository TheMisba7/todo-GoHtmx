// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: todo.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createTodo = `-- name: CreateTodo :one
INSERT INTO todo(id, owner, name, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5) RETURNING id, name, owner, status, created_at, updated_at
`

type CreateTodoParams struct {
	ID        uuid.UUID
	Owner     uuid.UUID
	Name      sql.NullString
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, createTodo,
		arg.ID,
		arg.Owner,
		arg.Name,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Owner,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findByUser = `-- name: findByUser :many
SELECT id, name, owner, status, created_at, updated_at FROM todo where owner = $1
`

func (q *Queries) findByUser(ctx context.Context, owner uuid.UUID) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, findByUser, owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Owner,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
