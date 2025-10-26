package handler

import (
	"context"
	"play-together/internal/model"
	"play-together/internal/repository"
)

type RequestHandler struct {
	repo repository.RequestRepository
}

func NewRequestHandler(repo repository.RequestRepository) *RequestHandler {
	return &RequestHandler{repo: repo}
}

func (h *RequestHandler) CreateRequest(ctx context.Context, req model.Request) (string, error) {
	return h.repo.CreateRequest(ctx, req)
}

func (h *RequestHandler) HandleGetAllRequests(ctx context.Context) ([]*model.Request, error) {
	return h.repo.GetAllRequests(ctx)
}

func (h *RequestHandler) HandleGetRequestByID(ctx context.Context, id string) (*model.Request, error) {
	return h.repo.GetRequestByID(ctx, id)
}
