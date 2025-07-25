package middleware

import (
	"log"
	"net/http"
	"time"
)

// CustomLoggerMiddleware logs detailed HTTP request lifecycle info
func CustomLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf("Started %s %s from %s", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
		log.Printf("Completed %s in %v", r.RequestURI, time.Since(start))
	})
}
