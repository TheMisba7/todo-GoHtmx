package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id       uuid.UUID
	Username string
	Password string
}

type Task struct {
	Id        uuid.UUID
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Status    int8 // 0 pending, 1 progress, 2 done
	CreatedAt time.Time
	UpdatedAt time.Time
	TodoId    uuid.UUID // reference to a Todo that this task belongs to
}

type Todo struct {
	Id        uuid.UUID
	Owner     uuid.UUID // reference to user.sql id
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
