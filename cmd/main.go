package main

import (
	"os"

	"github.com/Cenzios/pf-backend/config"
	"github.com/Cenzios/pf-backend/internal/user"
	"github.com/Cenzios/pf-backend/middleware"
	"github.com/Cenzios/pf-backend/pkg/db"
	"github.com/Cenzios/pf-backend/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	logger.Init()

	// Set Gin mode based on environment variable
	if ginMode := os.Getenv("GIN_MODE"); ginMode != "" {
		gin.SetMode(ginMode)
	} else {
		gin.SetMode(gin.ReleaseMode) // Default to release mode
	}
	r := gin.New()
	// Initialize database connection
	db.ConnectMongoDB()

	// Set up middleware
	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.CustomLogger())

	r.GET("/api/v1/user", user.GetProfile)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Log.Infof("Server running on port %s", port)
	r.Run(":" + port)
}
