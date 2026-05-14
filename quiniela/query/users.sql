-- name: CreateUser :one
INSERT INTO users (created_at, updated_at, name, api_key)
VALUES ($1, $2, $3, $4)
RETURNING id, created_at, updated_at, name, api_key;
