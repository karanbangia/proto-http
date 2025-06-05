package main

import (
	"log"
	"net/http"

	"github.com/proto-http/server"
)

func main() {
	// Create and start server
	userServer := server.NewUserServer()
	http.HandleFunc("/user", userServer.HandleUser)

	// Start server
	log.Printf("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
