package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

const (
	PageRoot = "www"
)

func loadPage(w http.ResponseWriter, page string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if page == "" {
		page = "root"
	}

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
		return
	}

	if err := tmpl.Execute(w, ""); err != nil {
		log.Printf("ERROR: %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("new handled request from %s for %s", r.RemoteAddr, r.RequestURI)

	if strings.HasPrefix(r.RequestURI, "/api/") {
		log.Printf("Handling API Request to %s", r.RequestURI)
		handleRequest(w, r)
		return
	} else {
		log.Printf("attempting to open %s", r.RequestURI)
		loadPage(w, r.RequestURI[1:])
		return
	}
}
