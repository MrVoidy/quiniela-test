package portsuser

import "context"

// RegisterUserResult is what HTTP clients need after a successful registration.
type RegisterUserResult struct {
	Message string
	Name    string
	UserID  int32
}

// UserService is the application API for user workflows (what handlers call).
type UserService interface {
	RegisterUser(ctx context.Context, name string) (*RegisterUserResult, error)
}
