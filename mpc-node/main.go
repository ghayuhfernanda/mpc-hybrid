package main

import (
	"net/http"
)

func main() {
	// HTTP Routes
	http.HandleFunc("/health", health)
	http.HandleFunc("/keygen", KeygenHandler)
	http.HandleFunc("/sign", SignHandler)

	// Change port per node: 8001, 8002, 8003
	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}
	http.ListenAndServe(":"+port, nil)

}
