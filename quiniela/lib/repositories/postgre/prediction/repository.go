package predictionrepo

import (
	"context"

	domainprediction "quiniela-app/quiniela/lib/domain/prediction"
	portsprediction "quiniela-app/quiniela/lib/ports/prediction"
	"quiniela-app/quiniela/sqlcrepository"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository implements portsprediction.PredictionRepository using sqlc.
type Repository struct {
	q *sqlcrepository.Queries
}

var _ portsprediction.PredictionRepository = (*Repository)(nil)

// NewRepository wires sqlc queries to a pgx pool.
func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{q: sqlcrepository.New(pool)}
}

// Save inserts a prediction row.
func (r *Repository) Save(ctx context.Context, p *domainprediction.Prediction) error {
	return r.q.CreatePrediction(ctx, sqlcrepository.CreatePredictionParams{
		FixtureID: p.FixtureID,
		UserID:    p.UserID,
		PredA:     p.PredA,
		PredB:     p.PredB,
	})
}

// GetUserScore returns the aggregate score for a user.
func (r *Repository) GetUserScore(ctx context.Context, userID int32) (int64, error) {
	return r.q.GetUserScore(ctx, userID)
}
