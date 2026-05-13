package portsuser

import (
	"context"

	domainuser "quiniela-app/quiniela/lib/domain/user"
)

// UserRepository defines persistence operations the domain needs for users.
type UserRepository interface {
	Save(ctx context.Context, u *domainuser.User) error
}
