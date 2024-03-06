package main

import (
	"encoding/json"
	"log"
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
	log.Printf("Authorized client %s", r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")
	log.Printf("Running switch statement on %s", r.RemoteAddr[4:])
	switch r.RequestURI[4:] {
	case "gettasks":
		log.Printf("Attempting to get all tasks")
		json.NewEncoder(w).Encode(getAllTasks())
	case "getweather":
		//json.NewEncoder(w).Encode(getWeather())
	default:
		http.Error(w, "404 page not found", http.StatusNotFound)
	}

}
func validToken(token string) bool {
	return token == Token
}
