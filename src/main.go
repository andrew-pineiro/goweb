package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"goweb/handlers"
	"goweb/middleware"

	"github.com/gorilla/mux"
)

var (
	AuthToken    string
	ConfigureApi bool
	UseDB        bool
	DBPath       string
	Port         string
)

func init() {
	flag.BoolVar(&ConfigureApi, "Api", false, "Enables API routes")
	flag.BoolVar(&UseDB, "Db", true, "Enables use of database for logins")
	flag.StringVar(&DBPath, "DbPath", "./data", "Path to store local DB")
	flag.StringVar(&AuthToken, "AuthToken", "", "API auth token for secure acess")
	flag.StringVar(&Port, "Port", "8080", "Port to run application on")
	flag.Parse()
}

func main() {
	//TODO: look into CGO_ENABLED depenedencies for cross-platform support
	initializeApp()
}
func initializeApp() {
	// Set API Token
	if ConfigureApi {
		log.Println("STARTUP: Setting API token")
		err := handlers.SetToken(AuthToken)
		if err != nil {
			log.Printf("ERROR: Unable to set API token; Supply with --AuthToken or create token.secret file in root directory;")
			log.Fatalf("ERROR: %s", err)
		}
	}
	// Setup Router
	router := setupRouter()

	// Load DB for Users
	if UseDB {
		setupDb()
	}

	startServer(router)
}
func setupDb() {
	log.Println("STARTUP: Initializing Database")
	if _, err := os.Stat(DBPath); os.IsNotExist(err) {
		os.Mkdir(DBPath, 0755)
	}
	var filePath = fmt.Sprintf("%s/%s", DBPath, "goweb.db")
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.Create(filePath)
	}
	log.Printf("STARTUP: Database setup: %s", filePath)
}
func startServer(router *mux.Router) {
	server := &http.Server{
		Addr:    ":" + Port,
		Handler: router,
	}

	log.Printf("STARTUP COMPLETE: LISTENING ON PORT %s", Port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
func setupRouter() *mux.Router {
	log.Printf("STARTUP: Setting up router")
	router := mux.NewRouter()

	// Public routes
	// Static + JS Files
	log.Println("STARTUP: Configuring static file handler")
	router.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./www/static"))))
	router.HandleFunc("/js/{file}", handlers.LoadJSFile)
	// Redirects
	log.Println("STARTUP: Configuring redirect handler")
	router.HandleFunc("/favicon.ico", handlers.Redirects)
	router.HandleFunc("/", handlers.Redirects)
	// API Endpoints
	if ConfigureApi {
		log.Println("STARTUP: Configuring API endpoint handler")
		router.HandleFunc("/api/{endpoint}", handlers.APIHandler)
	}
	// HTML Pages
	log.Println("STARTUP: Configuring web page handler")
	router.HandleFunc("/{page}", handlers.LoadPage).Methods("GET")

	// Authentication handlers
	log.Println("STARTUP: Configuring authentication handlers")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/logout", handlers.LogoutHandler).Methods("GET")

	// Protected routes (apply middleware)
	protectedRoutes := router.PathPrefix("/s").Subrouter()
	protectedRoutes.Use(middleware.AuthMiddleware)
	//TODO: fix this
	protectedRoutes.HandleFunc("/s/{page}", handlers.LoadPage).Methods("GET")

	//NOT FOUND
	log.Println("STARTUP: Configuring 404 handler")
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s NOT FOUND: %s", r.RemoteAddr, r.RequestURI)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 page not found"))
	})
	log.Println("STARTUP: Router configured!")
	return router
}
