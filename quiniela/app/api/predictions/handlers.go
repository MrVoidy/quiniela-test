package predictions

import (
	"encoding/json"
	"net/http"
	"strings"

	_ "quiniela-app/quiniela/app/api/apidoc" // referenced by swag annotations

	domainprediction "quiniela-app/quiniela/lib/domain/prediction"
	portsprediction "quiniela-app/quiniela/lib/ports/prediction"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Handler serves HTTP for prediction workflows using a PredictionService port.
type Handler struct {
	Svc portsprediction.PredictionService
}

// NewHandler wires the prediction service port.
func NewHandler(svc portsprediction.PredictionService) *Handler {
	return &Handler{Svc: svc}
}

// PostPrediction stores a user's score guess for a fixture.
//
//	@Summary		Submit prediction
//	@Description	Inserts into fixture_predictions; user_id must be a users.id UUID from POST /v1/users.
//	@Tags			predictions
//	@Accept			json
//	@Produce		json
//	@Param			body	body		apidoc.PostPredictionRequest	true	"Prediction payload"
//	@Success		201
//	@Failure		400		{string}	string	"plain text error"
//	@Failure		500		{string}	string	"plain text error"
//	@Router			/v1/predictions [post]
func (h *Handler) PostPrediction(w http.ResponseWriter, r *http.Request) {
	var req struct {
		FixtureID int32  `json:"fixture_id"`
		UserID    string `json:"user_id"`
		PredA     int32  `json:"pred_a"`
		PredB     int32  `json:"pred_b"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	uid, err := uuid.Parse(strings.TrimSpace(req.UserID))
	if err != nil {
		http.Error(w, "Invalid user_id (expected UUID)", http.StatusBadRequest)
		return
	}

	err = h.Svc.SubmitPrediction(r.Context(), &domainprediction.Prediction{
		FixtureID: req.FixtureID,
		UserID:    uid,
		PredA:     req.PredA,
		PredB:     req.PredB,
	})
	if err != nil {
		http.Error(w, "Could not save prediction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetUserScore returns the count of correctly guessed outcomes for a user.
//
//	@Summary		User prediction score
//	@Description	Aggregates matching predictions against stored fixture results.
//	@Tags			predictions
//	@Produce		json
//	@Param			userID	path		string	true	"users.id (UUID)"
//	@Success		200		{object}	apidoc.UserScoreResponse
//	@Failure		400		{string}	string	"plain text error"
//	@Failure		404		{string}	string	"plain text error"
//	@Router			/v1/users/{userID}/score [get]
func (h *Handler) GetUserScore(w http.ResponseWriter, r *http.Request) {
	idStr, ok := mux.Vars(r)["userID"]
	if !ok || strings.TrimSpace(idStr) == "" {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	userID, err := uuid.Parse(strings.TrimSpace(idStr))
	if err != nil {
		http.Error(w, "Invalid user id (expected UUID)", http.StatusBadRequest)
		return
	}

	score, err := h.Svc.ScoreForUser(r.Context(), userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]int64{"total_points": score})
}
