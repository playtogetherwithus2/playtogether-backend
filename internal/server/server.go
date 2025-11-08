package server

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	port   string
}

func NewServer(router *gin.Engine, port string) *Server {
	return &Server{
		router: router,
		port:   port,
	}
}

func (s *Server) Start() {
	// ✅ Get port from Render environment
	port := os.Getenv("PORT")
	if port == "" {
		port = s.port
	}

	addr := ":" + port
	log.Printf("🚀 Starting server on %s", addr)

	// ✅ This call blocks, so Render detects it
	if err := s.router.Run(addr); err != nil && err != http.ErrServerClosed {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
