package main

import (
	"log"
	"play-together/config"
	"play-together/internal/server"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	// Initialize router
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Server is running 🚀"})
	})

	// Initialize and start server
	srv := server.NewServer(router, cfg.Port)

	log.Println("Starting server...")
	srv.Start()
}
