package routes

import (
	"play-together/config"
	"play-together/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	firebaseClient *config.FirebaseClient, loginService *service.LoginService, healthService *service.HealthService,
) *gin.Engine {

	router := gin.Default()

	public := router.Group("/api/v1")
	{
		public.GET("/health", healthService.HealthCheck)
		AddLoginRoutes(public, loginService)
	}

	return router
}
