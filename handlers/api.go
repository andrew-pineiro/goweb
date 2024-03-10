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
	log.Printf("%s AUTHORIZED TOKEN %s", r.RemoteAddr, token)

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
	case "/contact":
		switch r.Method {
		case http.MethodGet:
			//TODO: not implemenented GET contact
			break
		case http.MethodPost:
			var msg controllers.Message
			err := json.NewDecoder(r.Body).Decode(&msg)
			if err != nil {
				log.Printf("ERROR: could not decode json %s", err)
				http.Error(w, "500 internal server error", http.StatusInternalServerError)
			}

			err = controllers.RecvMessage(msg)
			if err != nil {
				log.Printf("ERROR: %s", err)
				http.Error(w, "500 internal server error", http.StatusInternalServerError)
			}

		}
	default:
		http.Error(w, "404 page not found", http.StatusNotFound)
	}

}
func validToken(token string) bool {
	return token == Token
}
