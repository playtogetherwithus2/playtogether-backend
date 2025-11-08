package main

import (
	"log"
	"play-together/config"
	"play-together/internal/server"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("🟢 Starting main.go")

	cfg := config.LoadConfig()
	log.Printf("✅ Config loaded: PORT=%s", cfg.Port)

	// Initialize router
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Server is running 🚀"})
	})

	log.Println("⚙️  Creating server instance...")
	srv := server.NewServer(router, cfg.Port)

	log.Println("🚀 Calling srv.Start() ...")
	srv.Start()
}
