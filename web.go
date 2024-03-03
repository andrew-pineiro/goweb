package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
)

const (
	PageRoot = "www"
)

func loadPage(w http.ResponseWriter, page string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	file := path.Join(PageRoot, page+".html")

	if _, err := os.Stat(file); err != nil {
		log.Printf("could not find file %s", file)
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles(file)
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	if err := tmpl.Execute(w, ""); err != nil {
		log.Printf("ERROR: %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("new handled request from %s", r.RemoteAddr)
	log.Printf("attempting to open %s", r.RequestURI)
	switch r.RequestURI {
	case "/":
		loadPage(w, "root")
		return
	case "/api/gettasks":
		token := r.Header.Get("token")

		if !validToken(token) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("401 - unauthorized"))
			return
		}

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(getAllTasks())
		return
	default:
		loadPage(w, r.RequestURI[1:])
		return
	}
}
