package handlers

import (
	"net/http"
	"time"

	"github.com/Cenzios/pf-backend/pkg/response"
)

// HealthCheckHandler returns a 200 OK status and basic info
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response.Success(w, map[string]interface{}{
		"status":     "OK",
		"timestamp":  time.Now().UTC(),
		"service":    "PFBackend",
		"version":    "v1.0.0", // You can replace this dynamically later
		"request_id": r.Header.Get("X-Request-ID"),
	}, "Health check passed")
}
