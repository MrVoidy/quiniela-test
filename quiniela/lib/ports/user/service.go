package portsuser

import (
	"context"

	"github.com/google/uuid"
)

// RegisterUserResult is what HTTP clients need after a successful registration.
type RegisterUserResult struct {
	Message string
	Name    string
	UserID  uuid.UUID
}

// UserService is the application API for user workflows (what handlers call).
type UserService interface {
	RegisterUser(ctx context.Context, name string) (*RegisterUserResult, error)
}
