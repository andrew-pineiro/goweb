package main

import (
	"encoding/json"
	"net/http"
)

const (
	Token        = "abc123"
	TaskFilePath = "data/tasks.csv"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("token")

	if !validToken(token) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 - unauthorized"))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	switch r.RequestURI[4:] {
	case "gettasks":
		json.NewEncoder(w).Encode(getAllTasks())
	default:
		http.Error(w, "404 page not found", http.StatusNotFound)
	}

}
func validToken(token string) bool {
	return token == Token
}
