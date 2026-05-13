package health

import (
	"net/http"

	_ "quiniela-app/quiniela/app/api/apidoc" // referenced by swag @Success types

	"quiniela-app/quiniela/app/api/httpjson"
)

// Readiness returns a simple JSON status for load balancers.
//
//	@Summary		Health check
//	@Description	Liveness probe; returns status ok when the process is up.
//	@Tags			health
//	@Produce		json
//	@Success		200	{object}	apidoc.HealthResponse
//	@Router			/v1/healthz [get]
func Readiness(w http.ResponseWriter, r *http.Request) {
	_ = httpjson.Write(w, http.StatusOK, struct {
		Status string `json:"status"`
	}{Status: "ok"})
}

// Err always responds with HTTP 500 (for testing error paths).
//
//	@Summary		Force error response
//	@Description	Returns a JSON error with HTTP 500.
//	@Tags			health
//	@Produce		json
//	@Success		500	{object}	apidoc.ErrorResponse
//	@Router			/v1/err [get]
func Err(w http.ResponseWriter, r *http.Request) {
	_ = httpjson.Write(w, http.StatusInternalServerError, struct {
		Error string `json:"error"`
	}{Error: "Internal Server Error"})
}
