package middleware

import (
	"net/http"
	"time"
)

// SessionStore simulates a session storage (replace this with a real DB or cache)
var SessionStore = make(map[string]string)

// SetSession creates a new session for a user
func SetSession(w http.ResponseWriter, username string) {
	// Generate a session token (you can use JWT or a more secure method)
	sessionToken := username + "-session"

	// Store session in memory (use Redis or DB in production)
	SessionStore[sessionToken] = username

	// Create a secure cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Path:     "/",
		HttpOnly: true,  // Prevent JavaScript access
		Secure:   false, // Set to true in HTTPS production
		Expires:  time.Now().Add(24 * time.Hour),
	})
}

// GetSession checks if the session is valid
func GetSession(r *http.Request) (string, bool) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return "", false
	}

	// Check if session exists in storage
	username, exists := SessionStore[cookie.Value]
	return username, exists
}

// ClearSession deletes the session (logout)
func ClearSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err == nil {
		delete(SessionStore, cookie.Value)
	}

	// Expire the cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		Expires:  time.Now().Add(-time.Hour),
	})
}
