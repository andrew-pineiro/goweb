package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"goweb/controllers"
	"goweb/handlers"

	"github.com/gorilla/mux"
)

var (
	AuthToken    string
	ConfigureApi bool
	Port         string
)

func init() {
	flag.BoolVar(&ConfigureApi, "Api", false, "Enables API routes")
	flag.StringVar(&AuthToken, "AuthToken", "", "API auth token for secure acess")
	flag.StringVar(&Port, "Port", "8080", "Port to run application on")
	flag.Parse()
}

func main() {
	initializeApp()
}
func initializeApp() {
	// Set API Token
	if ConfigureApi {
		log.Println("STARTUP: Setting API token")
		err := handlers.SetToken(AuthToken)
		if err != nil {
			log.Printf("ERROR: Unable to set API token; Supply with --AuthToken or create token.secret file in root directory;")
			log.Printf("ERROR: %s", err)
			os.Exit(1)
		}
	}
	// Setup Router
	router := setupRouter()

	// Load Users
	log.Printf("STARTUP: Loading users from file")
	controllers.LoadUsers()

	startServer(router)
}
func startServer(router *mux.Router) {
	server := &http.Server{
		Addr:    ":" + Port,
		Handler: router,
	}

	log.Printf("STARTUP COMPLETE: LISTENING ON PORT %s", Port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
func setupRouter() *mux.Router {
	log.Printf("STARTUP: Setting up router")
	router := mux.NewRouter()

	//Static + JS Files
	log.Println("STARTUP: Configuring static file handler")
	router.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./www/static"))))
	router.HandleFunc("/js/{file}", handlers.LoadJSFile)

	//Redirects
	log.Println("STARTUP: Configuring redirect handler")
	router.HandleFunc("/favicon.ico", handlers.Redirects)
	router.HandleFunc("/", handlers.Redirects)

	//API Endpoints
	if ConfigureApi {
		log.Println("STARTUP: Configuring API endpoint handler")
		router.HandleFunc("/api/{endpoint}", handlers.APIHandler)
	}

	//HTML Pages
	log.Println("STARTUP: Configuring web page handler")
	router.HandleFunc("/{page}", handlers.LoadPage).Methods("GET")

	//NOT FOUND
	log.Println("STARTUP: Configuring 404 handler.")
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s NOT FOUND: %s", r.RemoteAddr, r.RequestURI)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 page not found"))
	})

	log.Println("STARTUP: Router configured!")
	return router
}
