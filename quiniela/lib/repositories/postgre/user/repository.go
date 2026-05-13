package userrepo

import (
	"context"

	domainuser "quiniela-app/quiniela/lib/domain/user"
	portsuser "quiniela-app/quiniela/lib/ports/user"
	"quiniela-app/quiniela/sqlcrepository"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository implements portsuser.UserRepository using sqlc.
type Repository struct {
	q *sqlcrepository.Queries
}

var _ portsuser.UserRepository = (*Repository)(nil)

// NewRepository wires sqlc queries to a pgx pool.
func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{q: sqlcrepository.New(pool)}
}

// Save maps domain user to sqlc and inserts.
func (r *Repository) Save(ctx context.Context, u *domainuser.User) error {
	return r.q.CreateUser(ctx, sqlcrepository.CreateUserParams{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Name:      u.Name,
		ApiKey:    u.ApiKey,
	})
}
