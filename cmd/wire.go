//go:build wireinject
// +build wireinject

package main

import (
	"play-together/config"
	"play-together/internal/routes"
	"play-together/internal/server"

	"github.com/google/wire"
)

func InitializeServer() (*server.Server, error) {
	wire.Build(
		config.LoadConfig,
		config.NewFirebaseClient,
		// handler.NewHealthHandler,
		routes.SetupRouter,
		providePortFromConfig,
		server.NewServer,
	)
	return &server.Server{}, nil
}

func providePortFromConfig(cfg *config.Config) string {
	return cfg.Port
}
