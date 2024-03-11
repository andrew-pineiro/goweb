package handlers

import (
	"encoding/json"
	"goweb/controllers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var Token string

func SetToken() error {
	tk, err := os.ReadFile("token.secret")
	if err != nil {
		return err
	}
	Token = string(tk)
	return nil
}
func checkToken(token string, w http.ResponseWriter, r *http.Request) bool {
	if !(token == Token) {
		log.Printf("%s UNAUTHORIZED: %s", r.RemoteAddr, r.RequestURI)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 - unauthorized"))
		return false
	} else {
		log.Printf("%s AUTHORIZED TOKEN %s", r.RemoteAddr, token)
		return true
	}
}
func APIHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("token")
	endpoint := mux.Vars(r)["endpoint"]

	w.Header().Set("Content-Type", "application/json")

	log.Printf("%s %s: %s", r.RemoteAddr, r.Method, r.RequestURI)

	//TODO(#1): implement rate limiting
	switch endpoint {
	case "gettasks":
		if checkToken(token, w, r) {
			data, err := controllers.GetAllTasks()
			if err != nil {
				http.Error(w, "500 internal server error", http.StatusInternalServerError)
				break
			}
			json.NewEncoder(w).Encode(data)
		}
	case "contact":
		switch r.Method {
		case http.MethodGet:
			//TODO(#2): implement GET endpoint for contact
			break
		case http.MethodPost:
			var msg controllers.Message
			err := json.NewDecoder(r.Body).Decode(&msg)
			if err != nil {
				log.Printf("ERROR: could not decode json; %s", err)
				http.Error(w, "500 internal server error", http.StatusInternalServerError)
				break
			}

			err = controllers.RecvMessage(msg, r)
			if err != nil {
				log.Printf("ERROR: %s", err)
				http.Error(w, "500 internal server error", http.StatusInternalServerError)
			}

		}
	default:
		http.Error(w, "404 page not found", http.StatusNotFound)
	}

}
