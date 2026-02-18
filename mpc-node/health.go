package main

import (
	"encoding/json"
	"net/http"
)

// health endpoint
func health(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{
		"status": "online",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
