package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        uuid.UUID
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
