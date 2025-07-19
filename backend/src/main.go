package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// A simple handler for an API endpoint
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from the Go Backend API!")
	})

	log.Println("Starting Go backend server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
