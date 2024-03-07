package main

import (
	"log"
	"net/http"

	"goweb/handlers"

	"github.com/gorilla/mux"
)

const (
	Port = "8080"
)

func main() {

	router := mux.NewRouter()

	//Static Files
	router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("./www/images"))))
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./www/js"))))

	//HTML Pages
	router.HandleFunc("/", handlers.LoadPage).Methods("GET")
	router.HandleFunc("/test", handlers.LoadPage).Methods("GET")

	//API Endpoints
	router.HandleFunc("/api/gettasks", handlers.APIHandler).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("NOT FOUND: %s", r.RequestURI)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 page not found"))
	})

	server := &http.Server{
		Addr:    ":" + Port,
		Handler: router,
	}

	log.Printf("START: PORT %s", Port)
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
