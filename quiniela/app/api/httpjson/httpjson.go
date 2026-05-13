package httpjson

import (
	"encoding/json"
	"net/http"
)

// Write sets Content-Type and writes JSON with the given status code.
func Write(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}
