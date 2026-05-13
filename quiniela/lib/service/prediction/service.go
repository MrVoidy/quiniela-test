package predictionservice

import (
	"context"

	domainprediction "quiniela-app/quiniela/lib/domain/prediction"
	portsprediction "quiniela-app/quiniela/lib/ports/prediction"
)

// Service implements portsprediction.PredictionService.
type Service struct {
	repo portsprediction.PredictionRepository
}

var _ portsprediction.PredictionService = (*Service)(nil)

// New constructs a prediction service with its repository port.
func New(repo portsprediction.PredictionRepository) *Service {
	return &Service{repo: repo}
}

// SubmitPrediction persists a prediction.
func (s *Service) SubmitPrediction(ctx context.Context, p *domainprediction.Prediction) error {
	return s.repo.Save(ctx, p)
}

// ScoreForUsuario returns the aggregate score for the given usuario id.
func (s *Service) ScoreForUsuario(ctx context.Context, usuarioID int32) (int64, error) {
	return s.repo.GetUsuarioScore(ctx, usuarioID)
}
