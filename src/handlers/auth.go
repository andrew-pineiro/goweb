package handlers

import (
	"encoding/json"
	"goweb/middleware"
	"net/http"
)

// User represents a simple user model
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// TODO: Unhardcode this into a db.
var users = map[string]string{
	"admin": "password123",
	"user1": "mypassword",
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse JSON request
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Authenticate user
	if password, ok := users[user.Username]; ok && password == user.Password {
		// Set session if authentication is successful
		middleware.SetSession(w, user.Username)

		redirect := r.URL.Query().Get("redirect")
		if redirect == "" {
			redirect = "/"
		}

		http.Redirect(w, r, redirect, http.StatusFound)
		return
	}

	http.Error(w, "Invalid username or password", http.StatusUnauthorized)
}

// LogoutHandler handles user logout
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	middleware.ClearSession(w, r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logged out successfully"))
}
