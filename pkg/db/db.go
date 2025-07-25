package db

import (
	"context"
	"os"

	"github.com/Cenzios/pf-backend/pkg/db/dbiface"
	"github.com/Cenzios/pf-backend/pkg/db/mongo"
	"github.com/Cenzios/pf-backend/pkg/db/mysql"
	"github.com/Cenzios/pf-backend/pkg/db/postgres"
	"github.com/Cenzios/pf-backend/pkg/logger"
)

// DB is the global database instance
var DB dbiface.Database

// Init initializes the global DB and CacheStore based on environment variables
func Init() {
	switch os.Getenv("DB_DRIVER") {
	case "postgres":
		DB = postgres.New()
		logger.Info.Println("✅ Connected to Postgres")
	case "mysql":
		DB = mysql.New()
		logger.Info.Println("✅ Connected to MySQL")
	default:
		DB = mongo.New()
		logger.Info.Println("✅ Connected to MongoDB")
	}

	// switch os.Getenv("CACHE_DRIVER") {
	// case "memcache":
	// 	CacheStore = memcache.New()
	// default:
	// 	CacheStore = redis.New()
	// }
}

func FindOne(ctx context.Context, collection string, filter interface{}) (interface{}, error) {
	return DB.FindOne(ctx, collection, filter)
}
