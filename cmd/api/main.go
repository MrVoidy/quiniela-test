package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"quiniela-app/internal/database"
	"time"

	_ "github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	fmt.Println("--- STARTING SERVER ---")

	// 1. Load .env
	err := godotenv.Load("../../.env")
	if err != nil {
		godotenv.Load(".env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("CRITICAL: DB_URL not found in .env")
	}

	// 2. Connect to Database
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("DB connection failed! Error: %v", err)
	}

	fmt.Println("Successfully connected to quiniela_db!")

	// 3. Initialize sqlc queries
	apiCfg := apiConfig{
		DB: database.New(db),
	}

	// 4. Routing
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/healthz", handlerReadiness)
	mux.HandleFunc("GET /v1/err", handlerErr)

	// New endpoint to test creating a user
	mux.HandleFunc("POST /v1/users", apiCfg.handlerCreateUser)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Printf("Server starting at http://localhost:%s\n", port)
	log.Fatal(srv.ListenAndServe())
}

// --- HANDLERS ---

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Invalid JSON body")
		return
	}

	// sqlc generated CreateUser with :execresult returns (sql.Result, error)
	_, err = apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        []byte(uuid.New().String()[0:16]), // Convert UUID to 16-byte slice
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
		ApiKey:    uuid.New().String(),
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	respondWithJSON(w, 201, struct {
		Message string `json:"message"`
		Name    string `json:"name"`
	}{
		Message: "User created successfully!",
		Name:    params.Name,
	})
}

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct {
		Status string `json:"status"`
	}{Status: "ok"})
}

func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 500, "Internal Server Error")
}

// --- HELPERS ---

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, struct {
		Error string `json:"error"`
	}{Error: msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, _ := json.Marshal(payload)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
