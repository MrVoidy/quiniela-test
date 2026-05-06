package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"quiniela-app/internal/database"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	DB *database.Queries
}

func main() {
	dbConn, err := sql.Open("mysql", "user:mypassword@tcp(localhost:3306)/quiniela_db")
	if err != nil {
		log.Fatal(err)
	}

	srv := &Server{DB: database.New(dbConn)}
	r := chi.NewRouter()

	r.Post("/predict", srv.handlePredict)
	r.Get("/score/{userID}", srv.handleGetScore)

	log.Println("Quiniela API starting on :8080")
	http.ListenAndServe(":8080", r)
}

func (s *Server) handlePredict(w http.ResponseWriter, r *http.Request) {
	var req struct {
		FixtureID int32 `json:"fixture_id"`
		UserID    int32 `json:"user_id"`
		PredA     int32 `json:"pred_a"`
		PredB     int32 `json:"pred_b"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	err := s.DB.CreatePrediction(r.Context(), database.CreatePredictionParams{
		FixtureID:   req.FixtureID,
		UsuarioID:   req.UserID,
		PrediccionA: req.PredA,
		PrediccionB: req.PredB,
	})

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) handleGetScore(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "userID"))
	score, err := s.DB.GetUserScore(r.Context(), int32(userID))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(map[string]int64{"total_points": score})
}
