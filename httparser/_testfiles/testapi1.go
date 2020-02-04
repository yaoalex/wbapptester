package testapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Get does nothing
func Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	status := params["return_status"]
	if status == "200" {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("OK")
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// NotHandler1 does nothing
func NotHandler1(w http.ResponseWriter) {
	fmt.Println("do nothing")
}

// NotHandler2 does nothing
func NotHandler2(w http.ResponseWriter) {
	fmt.Println("do nothing again")
}
