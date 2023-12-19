package middleware

import (
	"net/http"
)

//const allowedOrigin = "*"
const allowedOrigin = "http://prinesy-poday.ru"

//CorsMiddleware provides answer to OPTIONS request and set CORS and CSRF headers. For GET method sets conent-type application/json
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", allowedOrigin)
		if r.Method == http.MethodOptions {
			if r.URL.Path != "/api/users" {
				w.Header().Add("Access-Control-Allow-Credentials", "true")
			}
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PATCH, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
		}
		next.ServeHTTP(w, r)
	})
}

//CorsCredentionalsMiddleware provides header allow credentials
func CorsCredentionalsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(w, r)
	})
}
