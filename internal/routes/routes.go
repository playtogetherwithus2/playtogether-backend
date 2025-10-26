package routes

import (
	"play-together/config"
	"play-together/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	firebaseClient *config.FirebaseClient,
	loginService *service.LoginService,
	healthService *service.HealthService,
	postService *service.PostService,
	chatService *service.ChatService,
	requestService *service.RequestService,

) *gin.Engine {

	router := gin.Default()

	public := router.Group("/api/v1")
	{
		public.GET("/health", healthService.HealthCheck)
		AddLoginRoutes(public, loginService)
		AddPostRoutes(public, postService)
		AddChatRoutes(public, chatService)
		AddRequestRoutes(public, requestService)
	}

	return router
}
