package handlers

import (
	"goweb/controllers"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/gorilla/mux"
)

const (
	PageRoot = "www"
	BaseFile = "base.html"
)

// Running list of restricted pages that should return a 403 forbidden
var RestrictedPages []string = []string{
	"base.html",
}

// TODO(#7): hard code pages that require auth into the page itself rather than an array
var AuthorizedPages []string = []string{
	"admin-panel.html",
}

func checkRestrictedPages(page string) bool {
	restPages := RestrictedPages
	for i := 0; i < len(restPages); i++ {
		if restPages[i] == page {
			return true
		}
	}
	return false
}
func Redirects(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s REDIRECT: %s", r.RemoteAddr, r.RequestURI)

	switch r.RequestURI {
	case "/favicon.ico":
		http.Redirect(w, r, "/static/images/favicon.ico", http.StatusMovedPermanently)
	case "/":
		http.Redirect(w, r, "index.html", http.StatusMovedPermanently)
	}
}

func checkAuthorizedPages(page string, rawCookie string) bool {
	authPages := AuthorizedPages
	for i := 0; i < len(authPages); i++ {
		if authPages[i] == page &&
			!controllers.CheckAuthToken(rawCookie) {
			return true
		}
	}
	return false
}
func LoadJSFile(w http.ResponseWriter, r *http.Request) {
	file := mux.Vars(r)["file"]
	http.ServeFile(w, r, PageRoot+"/js/"+file)
}
func LoadPage(w http.ResponseWriter, r *http.Request) {
	var data string
	var rawCookie string
	var baseExists = true
	//TODO(#4): implement data injection to pages

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if !CheckRateCount(strings.Split(r.RemoteAddr, ":")[0]) {
		log.Printf("%s RATE LIMIT EXCEEDED", r.RemoteAddr)
		http.Error(w, "429 too many request", http.StatusTooManyRequests)
		return
	}

	page := mux.Vars(r)["page"]
	cookie, err := r.Cookie("X-Auth-Token")
	if err == nil {
		rawCookie, _ = url.PathUnescape(cookie.Value)
	}

	if !strings.ContainsAny(page, ".") {
		page += ".html"
	}

	if checkRestrictedPages(page) {
		http.Error(w, "403 forbidden", http.StatusForbidden)
		log.Printf("%s RESTRICTED: %s", r.RemoteAddr, r.RequestURI)
		return
	}

	if checkAuthorizedPages(page, rawCookie) {
		http.Error(w, "401 unauthorized", http.StatusUnauthorized)
		log.Printf("%s UNAUTHORIZED: %s", r.RemoteAddr, r.RequestURI)
		return
	}

	file := path.Join(PageRoot, strings.ToLower(page))

	baseFile := path.Join(PageRoot, BaseFile)
	if _, err := os.Stat(baseFile); err != nil {
		baseExists = false
	}
	//check if file exists
	if _, err := os.Stat(file); err != nil {
		log.Printf("%s NOT FOUND: %s", r.RemoteAddr, file)
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}

	log.Printf("%s LOAD: %s", r.RemoteAddr, file)

	if strings.Contains(file, ".html") && baseExists {
		tmpl, err := template.ParseFiles(file, baseFile)

		if err != nil {
			log.Printf("%s ERROR: %s", r.RemoteAddr, err.Error())
			http.Error(w, "500 internal server error", http.StatusInternalServerError)
			return
		}

		if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
			log.Printf("%s ERROR: %s", r.RemoteAddr, err.Error())
			http.Error(w, "500 internal server error", http.StatusInternalServerError)
			return
		}
	} else {
		tmpl, err := template.ParseFiles(file)

		if err != nil {
			log.Printf("%s ERROR: %s", r.RemoteAddr, err.Error())
			http.Error(w, "500 internal server error", http.StatusInternalServerError)
			return
		}

		if err := tmpl.Execute(w, data); err != nil {
			log.Printf("%s ERROR: %s", r.RemoteAddr, err.Error())
			http.Error(w, "500 internal server error", http.StatusInternalServerError)
			return
		}
	}
}
