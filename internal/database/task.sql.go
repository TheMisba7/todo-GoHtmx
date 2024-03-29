// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: task.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createTask = `-- name: CreateTask :one
INSERT INTO task(id, name, created_at, updated_at, todo_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, name, start_date, end_date, status, created_at, updated_at, todo_id
`

type CreateTaskParams struct {
	ID        uuid.UUID
	Name      sql.NullString
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	TodoID    uuid.UUID
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, createTask,
		arg.ID,
		arg.Name,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.TodoID,
	)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.StartDate,
		&i.EndDate,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.TodoID,
	)
	return i, err
}

const deleteTask = `-- name: DeleteTask :exec
delete from task where id = $1
`

func (q *Queries) DeleteTask(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTask, id)
	return err
}

const getTaskById = `-- name: GetTaskById :one
SELECT id, name, start_date, end_date, status, created_at, updated_at, todo_id FROM task where id = $1
`

func (q *Queries) GetTaskById(ctx context.Context, id uuid.UUID) (Task, error) {
	row := q.db.QueryRowContext(ctx, getTaskById, id)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.StartDate,
		&i.EndDate,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.TodoID,
	)
	return i, err
}

const getTasks = `-- name: GetTasks :many
SELECT id, name, start_date, end_date, status, created_at, updated_at, todo_id FROM task where todo_id = $1
`

func (q *Queries) GetTasks(ctx context.Context, todoID uuid.UUID) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, getTasks, todoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.StartDate,
			&i.EndDate,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.TodoID,
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

const updateTaskStatus = `-- name: UpdateTaskStatus :exec
update task set status = $1 where id = $2
`

type UpdateTaskStatusParams struct {
	Status int32
	ID     uuid.UUID
}

func (q *Queries) UpdateTaskStatus(ctx context.Context, arg UpdateTaskStatusParams) error {
	_, err := q.db.ExecContext(ctx, updateTaskStatus, arg.Status, arg.ID)
	return err
}
