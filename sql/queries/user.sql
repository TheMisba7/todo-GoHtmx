-- name: CreateUser :one
INSERT INTO users (id, username, password)
VALUES ($1, $2, $2) RETURNING *;

-- name: GetUser :one
SELECT * FROM users where username = $1;

-- name: GetUserById :one
SELECT * FROM users where id = $1;