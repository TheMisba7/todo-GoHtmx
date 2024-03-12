-- name: CreateTodo :one
INSERT INTO todo(id, owner, name, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: FindByUser :many
SELECT * FROM todo where owner = $1;

-- name: FindTodoById :one
SELECT * FROM todo where id = $1;

-- name: DeleteTodo :exec
delete from todo where id = $1;