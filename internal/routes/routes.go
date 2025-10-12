package routes

import (
	"play-together/config"
	"play-together/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(firebaseClient *config.FirebaseClient) *gin.Engine {
	router := gin.Default()

	// Initialize handlers
	healthHandler := handler.NewHealthHandler()

	public := router.Group("/api/v1")
	{
		public.GET("/health", healthHandler.HealthCheck)
	}

	return router
}
