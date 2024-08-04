package main

import (
	"log"
	"net/http"
	"os"

	"goweb/controllers"
	"goweb/handlers"

	"github.com/gorilla/mux"
)

const (
	Port = "8080"
)

var AuthToken string

func main() {

	err := handlers.SetToken()
	if err != nil {
		log.Printf("ERROR: unable to get token; %s", err)
		os.Exit(1)
	}
	router := mux.NewRouter()

	controllers.LoadUsers()

	//Static Files
	router.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./www/static"))))
	//JS Files
	router.HandleFunc("/js/{file}", handlers.LoadJSFile)
	//Redirects
	router.HandleFunc("/favicon.ico", handlers.Redirects)
	router.HandleFunc("/", handlers.Redirects)
	//API Endpoints
	router.HandleFunc("/api/{endpoint}", handlers.APIHandler)
	//HTML Pages
	router.HandleFunc("/{page}", handlers.LoadPage).Methods("GET")
	//NOT FOUND
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s NOT FOUND: %s", r.RemoteAddr, r.RequestURI)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 page not found"))
	})

	server := &http.Server{
		Addr:    ":" + Port,
		Handler: router,
	}

	log.Printf("START: PORT %s; TOKEN: %s", Port, handlers.Token)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
