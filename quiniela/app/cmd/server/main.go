// Package main runs the Quiniela HTTP API (Postgres, gorilla/mux).
//
//	@title						Quiniela API
//	@version					1.0
//	@description				HTTP API for users, predictions, and health checks. Set DB_URL (Postgres). Optional SWAGGER_HOST for Swagger UI "Try it out" (defaults to localhost:PORT).
//
//	@host						localhost:8080
//	@BasePath					/
//	@schemes					http https
//
//	@tag.name					health
//	@tag.description			Liveness and error test routes
//
//	@tag.name					users
//	@tag.description			User registration
//
//	@tag.name					predictions
//	@tag.description			Quiniela predictions and scores
package main

//go:generate go run github.com/swaggo/swag/cmd/swag@v1.16.4 init -g cmd/server/main.go -o ../../docs -d ../..

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"quiniela-app/quiniela/app/api/health"
	"quiniela-app/quiniela/app/api/predictions"
	"quiniela-app/quiniela/app/api/users"
	"quiniela-app/quiniela/docs"
	"quiniela-app/quiniela/migrations"
	predictionrepo "quiniela-app/quiniela/lib/repositories/postgre/prediction"
	userrepo "quiniela-app/quiniela/lib/repositories/postgre/user"
	predictionservice "quiniela-app/quiniela/lib/service/prediction"
	userservice "quiniela-app/quiniela/lib/service/user"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	fmt.Println("--- STARTING SERVER ---")

	if err := godotenv.Load(); err != nil {
		log.Println("warning: .env not loaded from cwd:", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("CRITICAL: DB_URL not found in .env")
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalf("DB pool: %v", err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("DB connection failed! Error: %v", err)
	}
	fmt.Println("Successfully connected to quiniela_db!")

	if err := migrations.Up(ctx, pool); err != nil {
		log.Fatalf("migrations: %v", err)
	}
	fmt.Println("Database migrations applied.")

	userRepository := userrepo.NewRepository(pool)
	userSvc := userservice.New(userRepository)
	predictionRepository := predictionrepo.NewRepository(pool)
	predictionSvc := predictionservice.New(predictionRepository)

	r := mux.NewRouter()
	r.HandleFunc("/v1/healthz", health.Readiness).Methods(http.MethodGet)
	r.HandleFunc("/v1/err", health.Err).Methods(http.MethodGet)

	uh := users.NewHandler(userSvc)
	r.HandleFunc("/v1/users", uh.CreateUser).Methods(http.MethodPost)

	ph := predictions.NewHandler(predictionSvc)
	r.HandleFunc("/v1/predictions", ph.PostPrediction).Methods(http.MethodPost)
	r.HandleFunc("/v1/users/{userID}/score", ph.GetUserScore).Methods(http.MethodGet)

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	addr := listenAddr(port)
	if h := swaggerHost(os.Getenv("SWAGGER_HOST"), addr); h != "" {
		docs.SwaggerInfo.Host = h
	}

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	fmt.Printf("Server starting at http://localhost:%s\n", strings.TrimPrefix(addr, ":"))
	fmt.Printf("Swagger UI: http://localhost:%s/swagger/index.html\n", strings.TrimPrefix(addr, ":"))
	log.Fatal(srv.ListenAndServe())
}

func listenAddr(port string) string {
	p := strings.TrimSpace(port)
	if p == "" {
		return ":8080"
	}
	if strings.HasPrefix(p, ":") {
		return p
	}
	return ":" + p
}

// swaggerHost prefers SWAGGER_HOST, else falls back to localhost:<port> for Swagger "Try it out".
func swaggerHost(envHost, addr string) string {
	if strings.TrimSpace(envHost) != "" {
		return strings.TrimSpace(envHost)
	}
	p := strings.TrimPrefix(addr, ":")
	if p == "" {
		return ""
	}
	return "localhost:" + p
}
