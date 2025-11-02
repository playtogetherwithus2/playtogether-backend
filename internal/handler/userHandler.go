package handler

import (
	"context"
	"play-together/internal/model"
	"play-together/internal/repository"
)

type UserHandler struct {
	repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) HandleGetUsers(ctx context.Context) ([]model.User, error) {
	return h.repo.GetUsers(ctx)
}

func (h *UserHandler) HandleGetUserByID(ctx context.Context, userID string) (model.User, error) {
	return h.repo.GetUserByID(ctx, userID)
}

func (h *UserHandler) HandleGetUsersByIDs(ctx context.Context, userIDs []string) ([]model.User, error) {
	return h.repo.GetUsersByIDs(ctx, userIDs)
}

func (h *UserHandler) HandleUpdateUser(ctx context.Context, id string, req model.UpdateUserRequest) error {
	return h.repo.UpdateUser(ctx, id, req)
}
