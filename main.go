package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"goweb/handlers"

	"github.com/gorilla/mux"
)

const (
	Port = "8080"
)

var AuthToken string

func CaselessMatcher(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.ToLower(r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
func main() {

	err := handlers.SetToken()
	if err != nil {
		log.Printf("ERROR: unable to get token; %s", err)
		os.Exit(1)
	}
	router := mux.NewRouter()

	//Static Files
	router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("./www/images"))))
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./www/js"))))
	router.PathPrefix("/content/").Handler(http.StripPrefix("/content/", http.FileServer(http.Dir("./www/content"))))

	//Redirects
	router.HandleFunc("/favicon.ico", handlers.Redirects)
	router.HandleFunc("/", handlers.Redirects)

	//API Endpoints
	router.HandleFunc("/api/{endpoint}", handlers.APIHandler)

	//HTML Pages
	router.HandleFunc("/{page}", handlers.LoadPage).Methods("GET")

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
