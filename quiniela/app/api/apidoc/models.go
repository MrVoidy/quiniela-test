// Package apidoc holds types referenced by swag annotations for OpenAPI schemas.
package apidoc

// HealthResponse is returned by GET /v1/healthz.
type HealthResponse struct {
	Status string `json:"status"`
}

// ErrorResponse is a generic JSON error body.
type ErrorResponse struct {
	Error string `json:"error"`
}

// CreateUserRequest is the body for POST /v1/users.
type CreateUserRequest struct {
	Name string `json:"name"`
}

// CreateUserResponse is returned on successful user creation.
type CreateUserResponse struct {
	Message string `json:"message"`
	Name    string `json:"name"`
	UserID  string `json:"user_id" example:"550e8400-e29b-41d4-a716-446655440000"`
}

// PostPredictionRequest is the body for POST /v1/predictions.
type PostPredictionRequest struct {
	FixtureID int32  `json:"fixture_id"`
	UserID    string `json:"user_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	PredA     int32  `json:"pred_a"`
	PredB     int32  `json:"pred_b"`
}

// UserScoreResponse is returned by GET /v1/users/{userID}/score.
type UserScoreResponse struct {
	TotalPoints int64 `json:"total_points"`
}
