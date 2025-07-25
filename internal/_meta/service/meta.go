package service

import (
	"context"

	"github.com/Cenzios/pf-backend/internal/models/mongo"
	"github.com/Cenzios/pf-backend/pkg/db"
)

func GetAllCountries(ctx context.Context) ([]mongo.Country, error) {
	results, err := db.DB.FindMany(ctx, "countries", map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	countries := make([]mongo.Country, 0, len(results))
	for _, r := range results {
		if c, ok := r.(mongo.Country); ok {
			countries = append(countries, c)
		}
	}
	return countries, nil
}
