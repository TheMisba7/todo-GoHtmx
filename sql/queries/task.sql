-- name: CreateTask :one
INSERT INTO task(id, name, created_at, updated_at, todo_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetTasks :many
SELECT * FROM task where todo_id = $1;

-- name: GetTaskById :one
SELECT * FROM task where id = $1;

-- name: DeleteTask :exec
delete from task where id = $1;