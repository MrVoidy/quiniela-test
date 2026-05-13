package userservice

import (
	"context"

	domainuser "quiniela-app/quiniela/lib/domain/user"
	portsuser "quiniela-app/quiniela/lib/ports/user"
)

// Service implements portsuser.UserService.
type Service struct {
	repo portsuser.UserRepository
}

var _ portsuser.UserService = (*Service)(nil)

// New constructs a user service with its repository port.
func New(repo portsuser.UserRepository) *Service {
	return &Service{repo: repo}
}

// RegisterUser creates a new persisted API user from a display name.
func (s *Service) RegisterUser(ctx context.Context, name string) (*portsuser.RegisterUserResult, error) {
	u := domainuser.NewRegistered(name)
	if err := s.repo.Save(ctx, u); err != nil {
		return nil, err
	}
	return &portsuser.RegisterUserResult{
		Message: "User created successfully!",
		Name:    name,
	}, nil
}
