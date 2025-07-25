package middleware

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const requestIDHeader = "X-Request-ID"

func generateRequestID() string {
	return strconv.FormatInt(time.Now().UnixNano(), 36) + "-" + strconv.Itoa(rand.Intn(100000))
}

// HTTPLoggerMiddleware logs each HTTP request with a request ID, status, and duration, using icons for clarity.
func HTTPLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Get or generate request ID
		reqID := r.Header.Get(requestIDHeader)
		if reqID == "" {
			reqID = generateRequestID()
		}

		// Add request ID to response header
		w.Header().Set(requestIDHeader, reqID)

		// Log request start
		log.Printf("➡️  [REQ %s] %s %s", reqID, r.Method, r.URL.Path)

		// Wrap the ResponseWriter to capture status code
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(rw, r)

		duration := time.Since(start)
		icon := statusIcon(rw.statusCode)
		log.Printf("%s [REQ %s] %s %s %d %s", icon, reqID, r.Method, r.URL.Path, rw.statusCode, duration)
	})
}

func statusIcon(status int) string {
	switch {
	case status >= 200 && status < 300:
		return "✅"
	case status >= 400 && status < 500:
		return "⚠️"
	case status >= 500:
		return "❌"
	default:
		return "ℹ️"
	}
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
