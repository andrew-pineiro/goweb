package middleware

import (
	"net/http"
)

// AuthMiddleware ensures user is authenticated
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, authenticated := GetSession(r); !authenticated {
			http.Redirect(w, r, "/login?redirect="+r.URL.Path, http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}
