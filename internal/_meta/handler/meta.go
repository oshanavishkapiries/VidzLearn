package handlers

import (
	"context"
	"net/http"
	"time"

	metaService "github.com/Cenzios/pf-backend/internal/_meta/service"
	"github.com/Cenzios/pf-backend/pkg/response"
)

func GetAllCountries(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	countries, err := metaService.GetAllCountries(ctx)
	if err != nil {
		response.InternalServerError(w, "Failed to fetch countries")
		return
	}

	response.Success(w, countries, "Countries fetched successfully")
}
