package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Cenzios/pf-backend/config"
	"github.com/Cenzios/pf-backend/internal/routes"
	"github.com/Cenzios/pf-backend/middleware"
	"github.com/Cenzios/pf-backend/pkg/db"
	"github.com/Cenzios/pf-backend/pkg/firebase"
	"github.com/Cenzios/pf-backend/pkg/logger"
	"github.com/Cenzios/pf-backend/pkg/smtp"
	//"github.com/Cenzios/pf-backend/seed"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// Initialize logger
	logger.Init()

	// Initialize database and cache (auto-selects backend based on env)
	db.Init()

	// Initialize firebase
	firebase.Init()

	// Initialize smtp
	smtp.Init()

	// Seed Data
	//seed.Init()

	// Register router
	router := routes.RegisterRoutes()

	// Apply middleware
	handler := middleware.Init(router)

	// Setup server
	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Start server
	go func() {
		logger.Info.Printf("🚀 Server running on port %s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error.Fatalf("Listen error: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	logger.Info.Println("🛑 Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error.Fatalf("Server Shutdown Failed: %v", err)
	}

	logger.Info.Println("✅ Server exited cleanly")
}
