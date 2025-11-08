package main

import (
	"log"
	"os"

	"project/config"
	"project/interval/server"
)

func main() {
	cfg := config.LoadConfig()

	srv, err := server.InitializeServer(cfg)
	if err != nil {
		log.Fatalf("❌ Failed to initialize server: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.Port
	}

	log.Printf("🚀 Starting server on port %s...", port)
	if err := srv.Start(port); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}
