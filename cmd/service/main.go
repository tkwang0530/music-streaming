package main

import (
	"log"

	"github.com/tkwang0530/music-streaming-project/internal/config"
	"github.com/tkwang0530/music-streaming-project/internal/routers"
	"github.com/tkwang0530/music-streaming-project/pkg/server"
)

func main() {
	// Load the configuration values
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration values: %s", err)
	}

	// Create a new router
	r := routers.NewRouter()

	// Create a new server instance
	s := server.NewServer(cfg.Server.Port, r)

	// Start the server
	if err := s.Start(); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
