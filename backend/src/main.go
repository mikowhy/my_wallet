package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"

	"my_wallet/backend/src/handlers"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic recovered: %+v", err)

				// Log the stack trace for easier debugging
				log.Print(string(debug.Stack()))
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Could not connect to database: %s\n", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Could not ping database: %s\n", err)
	}

	log.Println("Successfully connected to the database")

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello in My Wallet!")
	})
	mux.HandleFunc("/auth", handlers.AuthHandler(db))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            true,
	})

	var finalHandler http.Handler = mux
	finalHandler = c.Handler(finalHandler)
	finalHandler = loggingMiddleware(finalHandler)
	finalHandler = recoverMiddleware(finalHandler)

	log.Println("Starting Go backend server")
	if err := http.ListenAndServe(":8080", finalHandler); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
