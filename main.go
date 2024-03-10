package main

import (
	"log"
	"net/http"
	"os"

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
		log.Printf("ERROR: unable to get token, %s", err)
		os.Exit(1)
	}
	router := mux.NewRouter()

	//Static Files
	router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("./www/images"))))
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./www/js"))))
	router.PathPrefix("/content/").Handler(http.StripPrefix("/content/", http.FileServer(http.Dir("./www/content"))))

	//HTML Pages
	router.HandleFunc("/index.html", handlers.LoadPage).Methods("GET")
	router.HandleFunc("/contact.html", handlers.LoadPage).Methods("GET")

	//Redirects
	router.HandleFunc("/favicon.ico", handlers.Redirects)
	router.HandleFunc("/", handlers.Redirects)

	//API Endpoints
	router.HandleFunc("/api/gettasks", handlers.APIHandler).Methods("GET")
	router.HandleFunc("/api/contact", handlers.APIHandler)

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
