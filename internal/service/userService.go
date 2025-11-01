package service

import (
	"context"
	"play-together/internal/handler"
	"play-together/internal/model"
)

type UserService struct {
	handler *handler.UserHandler
}

func NewUserService(handler *handler.UserHandler) *UserService {
	return &UserService{handler: handler}
}

func (s *UserService) GetUsers(ctx context.Context) ([]model.User, error) {
	return s.handler.HandleGetUsers(ctx)
}

func (s *UserService) GetUserByID(ctx context.Context, userID string) (model.User, error) {
	return s.handler.HandleGetUserByID(ctx, userID)
}

func (s *UserService) GetUsersByIDs(ctx context.Context, userIDs []string) ([]model.User, error) {
	return s.handler.HandleGetUsersByIDs(ctx, userIDs)
}
