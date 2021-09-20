package handler

import (
	"log"
	"net/http"
)

//LoggerMiddleware ...
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request: %v %v %v\n", r.Method, r.Host, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
