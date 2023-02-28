package main

import (
	"log"
	"net/http"

	"github.com/tkwang0530/music-streaming/internal/config"
	"github.com/tkwang0530/music-streaming/internal/utils"
	"github.com/tkwang0530/music-streaming/pkg/server"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database connection
	db, err := utils.NewDatabase(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Start server
	svr, err := server.New(db, cfg.Server.Port)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}
	if err := svr.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}
