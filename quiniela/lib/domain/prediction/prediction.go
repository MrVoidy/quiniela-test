package prediction

// Prediction is a user's scoreline guess for a fixture (fixture_predictions row).
type Prediction struct {
	FixtureID int32
	UserID    int32
	PredA     int32
	PredB     int32
}
