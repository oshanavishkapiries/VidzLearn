package seed

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/Cenzios/pf-backend/internal/models/mongo"
	"github.com/Cenzios/pf-backend/pkg/db"
	"github.com/Cenzios/pf-backend/pkg/logger"
)

func Init() {
	if os.Getenv("SEED_DATA") != "true" {
		return
	}

	dbType := os.Getenv("DB_DRIVER")
	switch dbType {
	case "mongo":
		seedMongo()
	case "mysql":
		logger.Info.Println("MySQL is not implemeted")
		//seedMySQL()
	case "postgres":
		logger.Info.Println("Postgres is not implemeted")
		//seedPostgres()
	default:
		logger.Info.Println("No supported DB for seeding")
	}

	// Optionally seed cache
	// seedCache()
}

func seedMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	items, err := db.DB.FindMany(ctx, "countries", map[string]interface{}{})
	if err == nil && len(items) > 0 {
		logger.Info.Println("✅ Countries already seeded in MongoDB, skipping.")
		return
	}

	file, err := os.Open("seed/countries.json")
	if err != nil {
		logger.Error.Fatalf("❌ Failed to open countries.json for MongoDB: %v", err)
	}
	defer file.Close()

	var countries []mongo.Country
	if err := json.NewDecoder(file).Decode(&countries); err != nil {
		logger.Error.Fatalf("❌ Failed to decode countries.json for MongoDB: %v", err)
	}

	seeded := 0
	for _, country := range countries {
		if err := db.DB.InsertOne(ctx, "countries", country); err != nil {
			logger.Error.Printf("❌ Failed to insert %s in MongoDB: %v\n", country.CountryName, err)
		} else {
			seeded++
		}
	}
	if seeded > 0 {
		logger.Info.Printf("✅ Seeded %d countries in MongoDB", seeded)
	} else {
		logger.Info.Println("ℹ️  No new countries seeded in MongoDB")
	}
}

// func seedMySQL() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	items, err := db.DB.FindMany(ctx, "countries", map[string]interface{}{})
// 	if err == nil && len(items) > 0 {
// 		logger.Info.Println("✅ Countries already seeded in MySQL, skipping.")
// 		return
// 	}

// 	file, err := os.Open("seed/countries.json")
// 	if err != nil {
// 		logger.Error.Fatalf("❌ Failed to open countries.json for MySQL: %v", err)
// 	}
// 	defer file.Close()

// 	var countries []mysql.Country
// 	if err := json.NewDecoder(file).Decode(&countries); err != nil {
// 		logger.Error.Fatalf("❌ Failed to decode countries.json for MySQL: %v", err)
// 	}

// 	seeded := 0
// 	for _, country := range countries {
// 		if err := db.DB.InsertOne(ctx, "countries", country); err != nil {
// 			logger.Error.Printf("❌ Failed to insert %s in MySQL: %v\n", country.CountryName, err)
// 		} else {
// 			seeded++
// 		}
// 	}
// 	if seeded > 0 {
// 		logger.Info.Printf("✅ Seeded %d countries in MySQL", seeded)
// 	} else {
// 		logger.Info.Println("ℹ️  No new countries seeded in MySQL")
// 	}
// }

// func seedPostgres() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	items, err := db.DB.FindMany(ctx, "countries", map[string]interface{}{})
// 	if err == nil && len(items) > 0 {
// 		logger.Info.Println("✅ Countries already seeded in Postgres, skipping.")
// 		return
// 	}

// 	file, err := os.Open("seed/countries.json")
// 	if err != nil {
// 		logger.Error.Fatalf("❌ Failed to open countries.json for Postgres: %v", err)
// 	}
// 	defer file.Close()

// 	var countries []postgres.Country
// 	if err := json.NewDecoder(file).Decode(&countries); err != nil {
// 		logger.Error.Fatalf("❌ Failed to decode countries.json for Postgres: %v", err)
// 	}

// 	seeded := 0
// 	for _, country := range countries {
// 		if err := db.DB.InsertOne(ctx, "countries", country); err != nil {
// 			logger.Error.Printf("❌ Failed to insert %s in Postgres: %v\n", country.CountryName, err)
// 		} else {
// 			seeded++
// 		}
// 	}
// 	if seeded > 0 {
// 		logger.Info.Printf("✅ Seeded %d countries in Postgres", seeded)
// 	} else {
// 		logger.Info.Println("ℹ️  No new countries seeded in Postgres")
// 	}
// }

// func seedCache() {
//  // Example: seed static cache data if needed
// }
