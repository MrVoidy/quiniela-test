package prediction

import "github.com/google/uuid"

// Prediction is a user's scoreline guess for a fixture (fixture_predictions row).
type Prediction struct {
	FixtureID int32
	UserID    uuid.UUID
	PredA     int32
	PredB     int32
}
