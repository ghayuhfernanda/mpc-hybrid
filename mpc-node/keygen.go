package main

import (
	"encoding/json"
	"net/http"
)

// KeygenHandler handles key generation
func KeygenHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{
		"status": "keygen_complete",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
