package handlers

import (
	"encoding/json"
	"goweb/middleware"
	"goweb/models"
	"goweb/utils"
	"net/http"
)

func ValidateLogin(username string, password string) bool {
	user := utils.CheckUserByName(username)
	if user.Guid != "" && utils.CheckPasswordHash(password, user.Password) {
		return true
	}
	return false

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse JSON request
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Authenticate user
	if ValidateLogin(user.Username, user.Password) {
		// Set session if authentication is successful
		middleware.SetSession(w, user)

		redirect := r.URL.Query().Get("redirect")
		if redirect == "" {
			redirect = "/"
		}
		//TODO: not working
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
