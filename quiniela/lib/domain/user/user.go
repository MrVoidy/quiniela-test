package user

import (
	"time"

	"github.com/google/uuid"
)

// User is an API user persisted in Postgres (table users).
type User struct {
	ID        int32
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	ApiKey    string
}

// NewRegistered builds a user ready to insert (id assigned by database on save).
func NewRegistered(name string) *User {
	now := time.Now()
	key := uuid.New().String()
	if len(key) > 64 {
		key = key[:64]
	}
	return &User{
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
		ApiKey:    key,
	}
}
