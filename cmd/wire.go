//go:build wireinject
// +build wireinject

package main

import (
	"play-together/config"
	"play-together/internal/handler"
	"play-together/internal/repository"
	"play-together/internal/routes"
	"play-together/internal/server"
	"play-together/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitializeServer() (*server.Server, error) {
	wire.Build(
		config.LoadConfig,
		config.NewFirebaseClient,

		repository.NewLoginRepository,
		handler.NewLoginHandler,
		service.NewLoginService,

		service.NewHealthService,

		repository.NewPostRepository,
		handler.NewPostHandler,
		service.NewPostService,

		repository.NewChatRepository,
		handler.NewChatHandler,
		service.NewChatService,

		provideRouter,
		providePortFromConfig,
		server.NewServer,
	)
	return &server.Server{}, nil
}

func providePortFromConfig(cfg *config.Config) string {
	return cfg.Port
}

func provideRouter(
	firebaseClient *config.FirebaseClient,
	loginService *service.LoginService,
	healthService *service.HealthService,
	postService *service.PostService,
	chatService *service.ChatService,
) *gin.Engine {
	return routes.SetupRouter(firebaseClient, loginService, healthService, postService, chatService)
}
