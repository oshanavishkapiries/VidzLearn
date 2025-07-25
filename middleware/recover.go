package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/Cenzios/pf-backend/pkg/logger"
	"github.com/Cenzios/pf-backend/pkg/response"
)

// RecoverMiddleware catches panics and prevents server crash
func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				// Log stack trace
				logger.Error.Printf("PANIC recovered: %v\n%s", rec, debug.Stack())

				// Optional: log structured error or send to alerting tool

				// Respond with 500
				response.InternalServerError(w, fmt.Sprintf("Unexpected server error: %v", rec))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
