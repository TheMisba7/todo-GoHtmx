-- name: CreateUser :one
INSERT INTO users (id, username, password)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetUser :one
SELECT * FROM users where username = $1;

-- name: GetUserById :one
SELECT * FROM users where id = $1;

-- name: DeleteUser :exec
delete from users where id = $1;