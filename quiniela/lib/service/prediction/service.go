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

// ScoreForUser returns the aggregate score for the given users.id (integer).
func (s *Service) ScoreForUser(ctx context.Context, userID int32) (int64, error) {
	return s.repo.GetUserScore(ctx, userID)
}
