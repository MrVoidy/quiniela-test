package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "quiniela-app/quiniela/app/api/apidoc" // referenced by swag annotations

	"quiniela-app/quiniela/app/api/httpjson"
	portsuser "quiniela-app/quiniela/lib/ports/user"
)

// Handler serves HTTP for user workflows using a UserService port.
type Handler struct {
	Svc portsuser.UserService
}

// NewHandler wires the user service port (typically implemented in lib/service/user).
func NewHandler(svc portsuser.UserService) *Handler {
	return &Handler{Svc: svc}
}

// CreateUser persists a new API user with a generated id and api_key.
//
//	@Summary		Register user
//	@Description	Creates a row in the users table.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			body	body		apidoc.CreateUserRequest	true	"User payload"
//	@Success		201		{object}	apidoc.CreateUserResponse
//	@Failure		400		{object}	apidoc.ErrorResponse
//	@Router			/v1/users [post]
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	var params parameters
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		_ = httpjson.Write(w, http.StatusBadRequest, struct {
			Error string `json:"error"`
		}{Error: "Invalid JSON body"})
		return
	}

	res, err := h.Svc.RegisterUser(r.Context(), params.Name)
	if err != nil {
		_ = httpjson.Write(w, http.StatusBadRequest, struct {
			Error string `json:"error"`
		}{Error: fmt.Sprintf("Couldn't create user: %v", err)})
		return
	}

	_ = httpjson.Write(w, http.StatusCreated, struct {
		Message string `json:"message"`
		Name    string `json:"name"`
	}{
		Message: res.Message,
		Name:    res.Name,
	})
}
