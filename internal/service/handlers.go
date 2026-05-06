package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"quiniela-app/internal/database" // Change 'quiniela-app' to your module name

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

type Handler struct {
	DB *database.Queries
}

// PostPrediction saves a user's guess to the database
func (h *Handler) PostPrediction(w http.ResponseWriter, r *http.Request) {
	var req struct {
		FixtureID int32 `json:"fixture_id"`
		UserID    int32 `json:"user_id"`
		PredA     int32 `json:"pred_a"`
		PredB     int32 `json:"pred_b"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := h.DB.CreatePrediction(r.Context(), database.CreatePredictionParams{
		FixtureID:   req.FixtureID,
		UsuarioID:   req.UserID,
		PrediccionA: req.PredA,
		PrediccionB: req.PredB,
	})

	if err != nil {
		http.Error(w, "Could not save prediction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetUserScore fetches the total 1-point wins for a user
func (h *Handler) GetUserScore(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "userID")
	userID, _ := strconv.Atoi(idStr)

	score, err := h.DB.GetUserScore(r.Context(), int32(userID))
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]int64{"total_points": score})
}
