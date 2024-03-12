// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Task struct {
	ID        uuid.UUID
	Name      sql.NullString
	StartDate sql.NullTime
	EndDate   sql.NullTime
	Status    int32
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	TodoID    uuid.UUID
}

type Todo struct {
	ID        uuid.UUID
	Name      sql.NullString
	Owner     uuid.UUID
	Status    int32
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

type User struct {
	ID       uuid.UUID
	Username string
	Password string
}
