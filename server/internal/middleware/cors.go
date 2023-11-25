package middleware

import (
	"net/http"
)

//const allowedOrigin = "*"
//const allowedOrigin = "*"
const allowedOrigin = "http://84.23.53.216"
const allowedOriginCsat = "http://84.23.53.216:8030"

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Add("Access-Control-Allow-Origin", allowedOriginCsat)
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

func CorsCredentionalsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(w, r)
	})
}
