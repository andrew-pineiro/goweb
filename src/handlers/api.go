package handlers

import (
	"encoding/json"
	"goweb/controllers"
	"goweb/models"
	"goweb/utils"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

var Token string

func SetToken(token string) error {
	if len(token) == 0 {
		tk, err := os.ReadFile("token.secret")
		if err != nil {
			return err
		}
		Token = string(tk)
	} else {
		Token = token
	}
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
	//TODO: unhardcode API endpoints
	token := r.Header.Get("token")
	endpoint := mux.Vars(r)["endpoint"]

	w.Header().Set("Content-Type", "application/json")

	if !utils.CheckRateCount(strings.Split(r.RemoteAddr, ":")[0]) {
		log.Printf("%s RATE LIMIT EXCEEDED", r.RemoteAddr)
		http.Error(w, "429 too many request", http.StatusTooManyRequests)
		return
	}

	log.Printf("%s %s: %s", r.RemoteAddr, r.Method, r.RequestURI)

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
	case "login":
		switch r.Method {
		case http.MethodGet:
			break
		case http.MethodPost:
			var tempUser models.User
			err := json.NewDecoder(r.Body).Decode(&tempUser)
			if err != nil {
				log.Printf("ERROR: could not decode json; %s", err)
				http.Error(w, "500 internal server error", http.StatusInternalServerError)
				break
			}

			existingUser := controllers.Login(tempUser.Username, tempUser.Password)
			if existingUser.Id < 0 {
				http.Error(w, "invalid username or password", http.StatusForbidden)
				break
			}

			log.Printf("%s LOGIN: %s", r.RemoteAddr, existingUser.Username)
			w.Header().Add("X-Auth-Token", existingUser.AuthToken)

		}
	default:
		http.Error(w, "404 page not found", http.StatusNotFound)
	}

}
