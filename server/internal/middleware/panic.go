package middleware

import (
	"fmt"
	"net/http"
)

//PanicMiddleware provides defensive from panics
func PanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic:", r)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
