package middleware

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

        // Do stuff here
		log.Println("Consuming API, " + r.Method + ": "+ r.RemoteAddr + r.RequestURI)
        // Call the next handler, which can be another middleware in the chain, or the final handler.

        next.ServeHTTP(w, r)
    })
}