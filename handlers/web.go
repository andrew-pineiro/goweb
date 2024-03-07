package handlers

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
)

const (
	PageRoot = "www"
)

func Redirects(w http.ResponseWriter, r *http.Request) {
	switch r.RequestURI {
	case "/favicon.ico":
		http.Redirect(w, r, "images/favicon.ico", http.StatusMovedPermanently)
	}
}

func LoadPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	page := r.RequestURI[1:]
	if page == "" {
		page = "root"
	}

	file := path.Join(PageRoot, page+".html")

	//check is file exists
	if _, err := os.Stat(file); err != nil {
		log.Printf("NOT FOUND: %s", file)
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}
	log.Printf("LOAD: %s", file)

	tmpl, err := template.ParseFiles(file)
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, ""); err != nil {
		log.Printf("ERROR: %s", err.Error())
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}
}
