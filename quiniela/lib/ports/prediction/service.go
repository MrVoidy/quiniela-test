package portsprediction

import (
	"context"

	domainprediction "quiniela-app/quiniela/lib/domain/prediction"
)

// PredictionService is the application API for prediction workflows.
type PredictionService interface {
	SubmitPrediction(ctx context.Context, p *domainprediction.Prediction) error
	ScoreForUser(ctx context.Context, userID int32) (int64, error)
}
