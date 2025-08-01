package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	AppName  string
	AppEnv   string
	SeedData string

	JWTSecret string

	DBDriver    string
	MongoURI    string
	MongoDB     string
	MySQLDSN    string
	PostgresDSN string

	CacheDriver   string
	RedisAddr     string
	RedisPassword string
	RedisDB       string
	MemcacheAddr  string

	FIREBASE_SERVICE_ACCOUNT string
	FIREBASE_STORAGE_BUCKET  string

	SMTP_HOST         string
	SMTP_PORT         string
	SMTP_USERNAME     string
	SMTP_PASSWORD     string
	SMTP_FROM_ADDRESS string
	SMTP_FROM_NAME    string
}

// LoadConfig loads environment variables into a Config struct
func LoadConfig() *Config {
	// Load from .env if exists
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, falling back to system environment")
	}

	cfg := &Config{
		Port:     getEnv("PORT", "8080"),
		AppName:  getEnv("APP_NAME", "PFBackend"),
		AppEnv:   getEnv("APP_ENV", "development"),
		SeedData: getEnv("SEED_DATA", "true"),

		JWTSecret: getEnv("JWT_SECRET", ""),

		DBDriver:    getEnv("DB_DRIVER", "mongo"),
		MongoURI:    getEnv("MONGO_URI", ""),
		MongoDB:     getEnv("MONGO_DB", ""),
		MySQLDSN:    getEnv("MYSQL_DSN", ""),
		PostgresDSN: getEnv("POSTGRES_DSN", ""),

		CacheDriver:   getEnv("CACHE_DRIVER", "redis"),
		RedisAddr:     getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       getEnv("REDIS_DB", "0"),
		MemcacheAddr:  getEnv("MEMCACHE_ADDR", "localhost:11211"),

		FIREBASE_SERVICE_ACCOUNT: getEnv("FIREBASE_SERVICE_ACCOUNT", ""),
		FIREBASE_STORAGE_BUCKET:  getEnv("FIREBASE_STORAGE_BUCKET", ""),

		SMTP_HOST:         getEnv("SMTP_HOST", ""),
		SMTP_PORT:         getEnv("SMTP_PORT", ""),
		SMTP_USERNAME:     getEnv("SMTP_USERNAME", ""),
		SMTP_PASSWORD:     getEnv("SMTP_PASSWORD", ""),
		SMTP_FROM_ADDRESS: getEnv("SMTP_FROM_ADDRESS", ""),
		SMTP_FROM_NAME:    getEnv("SMTP_FROM_NAME", ""),
	}

	// Only enforce Mongo as required for backward compatibility
	if cfg.DBDriver == "mongo" && (cfg.MongoURI == "" || cfg.MongoDB == "") {
		log.Fatal("❌ Required environment variables (MONGO_URI or MONGO_DB) missing for MongoDB.")
	}

	validateEnvVars(cfg)
	return cfg
}

// getEnv returns env value or fallback
func getEnv(key, fallback string) string {

	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

func validateEnvVars(cfg *Config) {
	missing := []string{}

	// Always required
	if cfg.Port == "" {
		missing = append(missing, "PORT")
	}
	if cfg.DBDriver == "" {
		missing = append(missing, "DB_DRIVER")
	}
	if cfg.CacheDriver == "" {
		missing = append(missing, "CACHE_DRIVER")
	}

	// DB-specific
	switch cfg.DBDriver {
	case "mongo":
		if cfg.MongoURI == "" {
			missing = append(missing, "MONGO_URI")
		}
		if cfg.MongoDB == "" {
			missing = append(missing, "MONGO_DB")
		}
	case "mysql":
		if cfg.MySQLDSN == "" {
			missing = append(missing, "MYSQL_DSN")
		}
	case "postgres":
		if cfg.PostgresDSN == "" {
			missing = append(missing, "POSTGRES_DSN")
		}
	}

	// Cache-specific
	switch cfg.CacheDriver {
	case "redis":
		if cfg.RedisAddr == "" {
			missing = append(missing, "REDIS_ADDR")
		}
	case "memcache":
		if cfg.MemcacheAddr == "" {
			missing = append(missing, "MEMCACHE_ADDR")
		}
	}

	// Firebase
	if cfg.FIREBASE_SERVICE_ACCOUNT == "" {
		missing = append(missing, "FIREBASE_SERVICE_ACCOUNT")
	}
	if cfg.FIREBASE_STORAGE_BUCKET == "" {
		missing = append(missing, "FIREBASE_STORAGE_BUCKET")
	}

	// SMTP
	if cfg.SMTP_HOST == "" {
		missing = append(missing, "SMTP_HOST")
	}
	if cfg.SMTP_PORT == "" {
		missing = append(missing, "SMTP_PORT")
	}
	if cfg.SMTP_USERNAME == "" {
		missing = append(missing, "SMTP_USERNAME")
	}
	if cfg.SMTP_PASSWORD == "" {
		missing = append(missing, "SMTP_PASSWORD")
	}
	if cfg.SMTP_FROM_ADDRESS == "" {
		missing = append(missing, "SMTP_FROM_ADDRESS")
	}
	if cfg.SMTP_FROM_NAME == "" {
		missing = append(missing, "SMTP_FROM_NAME")
	}

	if len(missing) > 0 {
		log.Fatalf("❌ Missing required environment variables: %v", missing)
	}
}
