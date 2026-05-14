package portsprediction

import (
	"context"

	domainprediction "quiniela-app/quiniela/lib/domain/prediction"

	"github.com/google/uuid"
)

// PredictionRepository defines persistence for predictions and score reads.
type PredictionRepository interface {
	Save(ctx context.Context, p *domainprediction.Prediction) error
	GetUserScore(ctx context.Context, userID uuid.UUID) (int64, error)
}
