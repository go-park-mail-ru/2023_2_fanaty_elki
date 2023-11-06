package middleware

import (
	"fmt"
	"net/http"
)

const allowedOrigin = "http://84.23.53.216"

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Println("Cors")
		w.Header().Add("Access-Control-Allow-Origin", allowedOrigin)
		if r.Method == http.MethodOptions{
			fmt.Println("FDKSJFLKDSFJ")
			w.Header().Add("Access-Control-Allow-Credentials", "true")
			w.Header().Set("content-type", "application/json")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func CorsCredentionalsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Println("Cookie")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(w, r)
	})
}