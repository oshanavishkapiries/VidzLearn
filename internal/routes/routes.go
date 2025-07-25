package routes

import (
	"net/http"

	appHandlers "github.com/Cenzios/pf-backend/internal/_app/handler"
	metaHandlers "github.com/Cenzios/pf-backend/internal/_meta/handler"
	userHandlers "github.com/Cenzios/pf-backend/internal/_user/handler"
)

func RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	// Versioned API
	apiV1 := http.NewServeMux()
	// Health Check
	apiV1.HandleFunc("/healthz", appHandlers.HealthCheckHandler)

	//apiV1.Handle("/healthz", middleware.JWTAuthMiddleware(http.HandlerFunc(appHandlers.HealthCheckHandler)))

	// Meta Data
	apiV1.HandleFunc("/countries/all", metaHandlers.GetAllCountries)

	// User
	apiV1.HandleFunc("/user/register", userHandlers.RegisterUser)
	apiV1.HandleFunc("/user/login", userHandlers.LoginUser)

	// Mount the versioned mux
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", apiV1))

	return mux
}
