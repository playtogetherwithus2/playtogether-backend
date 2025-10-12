package main

import (
	"log"

	"play-together/config"
	"play-together/internal/routes"
	"play-together/internal/server"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize Firebase client
	firebaseClient, err := config.NewFirebaseClient(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase client: %v", err)
	}

	// Setup routes
	router := routes.SetupRouter(firebaseClient)

	// Create and start server
	srv := server.NewServer(router, cfg.Port)
	srv.Start()
}
