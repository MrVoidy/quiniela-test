package user

import (
	"time"

	"github.com/google/uuid"
)

// User is an API user persisted in Postgres (table users).
type User struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	ApiKey    string
}

// NewRegistered builds a user ready to insert (id and api_key generated).
func NewRegistered(name string) *User {
	now := time.Now()
	return &User{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
		ApiKey:    uuid.New().String(),
	}
}
