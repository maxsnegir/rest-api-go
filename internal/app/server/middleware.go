package server

import (
	"log"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s :: %s ", r.RequestURI, r.Method)
		next.ServeHTTP(w, r)
	})
}
