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

// Save inserts the user and sets u.ID from the database.
func (r *Repository) Save(ctx context.Context, u *domainuser.User) error {
	row, err := r.q.CreateUser(ctx, sqlcrepository.CreateUserParams{
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Name:      u.Name,
		ApiKey:    u.ApiKey,
	})
	if err != nil {
		return err
	}
	u.ID = row.ID
	return nil
}
