package portsprediction

import (
	"context"

	domainprediction "quiniela-app/quiniela/lib/domain/prediction"

	"github.com/google/uuid"
)

// PredictionService is the application API for prediction workflows.
type PredictionService interface {
	SubmitPrediction(ctx context.Context, p *domainprediction.Prediction) error
	ScoreForUser(ctx context.Context, userID uuid.UUID) (int64, error)
}
