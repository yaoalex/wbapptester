package testapi

import (
	"encoding/json"
	"net/http"
)

// Post does nothing
func Post(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("OK")
}
