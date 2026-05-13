package portsprediction

import (
	"context"

	domainprediction "quiniela-app/quiniela/lib/domain/prediction"
)

// PredictionService is the application API for prediction workflows.
type PredictionService interface {
	SubmitPrediction(ctx context.Context, p *domainprediction.Prediction) error
	ScoreForUsuario(ctx context.Context, usuarioID int32) (int64, error)
}
