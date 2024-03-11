package handlers

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gorilla/mux"
)

const (
	PageRoot = "www"
	BaseFile = "base.html"
)

func Redirects(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s REDIRECT: %s", r.RemoteAddr, r.RequestURI)

	switch r.RequestURI {
	case "/favicon.ico":
		http.Redirect(w, r, "/static/images/favicon.ico", http.StatusMovedPermanently)
	case "/":
		http.Redirect(w, r, "index.html", http.StatusMovedPermanently)
	}
}
func LoadJSFile(w http.ResponseWriter, r *http.Request) {
	file := mux.Vars(r)["file"]
	http.ServeFile(w, r, "www/js/"+file)
}
func LoadPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	page := mux.Vars(r)["page"]
	if !strings.ContainsAny(page, ".") {
		page += ".html"
	}

	file := path.Join(PageRoot, strings.ToLower(page))
	baseFile := path.Join(PageRoot, BaseFile)

	//check if file exists
	if _, err := os.Stat(file); err != nil {
		log.Printf("%s NOT FOUND: %s", r.RemoteAddr, file)
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}
	log.Printf("%s LOAD: %s", r.RemoteAddr, file)

	tmpl, err := template.ParseFiles(file, baseFile)
	if err != nil {
		log.Printf("%s ERROR: %s", r.RemoteAddr, err.Error())
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "base", ""); err != nil {
		log.Printf("%s ERROR: %s", r.RemoteAddr, err.Error())
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}
}
