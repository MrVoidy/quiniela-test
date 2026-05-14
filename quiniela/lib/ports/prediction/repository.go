package portsprediction

import (
	"context"

	domainprediction "quiniela-app/quiniela/lib/domain/prediction"
)

// PredictionRepository defines persistence for predictions and score reads.
type PredictionRepository interface {
	Save(ctx context.Context, p *domainprediction.Prediction) error
	GetUserScore(ctx context.Context, userID int32) (int64, error)
}
