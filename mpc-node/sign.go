package main

import (
	"encoding/json"
	"net/http"
)

// SignHandler handles threshold signing
func SignHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{
		"signature": "dummy_signature_123456",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
