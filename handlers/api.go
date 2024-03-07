package handlers

import (
	"encoding/json"
	"goweb/controllers"
	"log"
	"net/http"
)

const (
	Token = "abc123"
)

func APIHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("token")

	if !validToken(token) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 - unauthorized"))
		return
	}
	log.Printf("AUTHORIZED: %s", r.RemoteAddr)

	w.Header().Set("Content-Type", "application/json")

	log.Printf("%s: %s", r.Method, r.RequestURI)
	switch r.RequestURI[4:] {
	case "/gettasks":
		data, err := controllers.GetAllTasks()
		if err != nil {
			http.Error(w, "500 internal server error", http.StatusInternalServerError)
			break
		}
		json.NewEncoder(w).Encode(data)
	case "/getweather":
		//json.NewEncoder(w).Encode(getWeather())
	default:
		http.Error(w, "404 page not found", http.StatusNotFound)
	}

}
func validToken(token string) bool {
	return token == Token
}
