package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	password "my_wallet/backend/src/app/users"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AuthHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		println(r)
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		var passwordHash string
		err := db.QueryRow("SELECT password_hash FROM users WHERE email = $1", user.Email).Scan(&passwordHash)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Invalid credentials", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			log.Printf("Database error: %s\n", err)
			return
		}

		match, err := password.ComparePasswordAndHash(user.Password, passwordHash)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			log.Printf("Password comparison error: %s\n", err)
			return
		}

		if match {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "Authentication successful"})
		} else {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		}
	}
}
