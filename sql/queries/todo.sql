-- name: CreateTodo :one
INSERT INTO todo(id, owner, name, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: findByUser :many
SELECT * FROM todo where owner = $1;