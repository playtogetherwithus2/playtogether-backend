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
func (h *RequestHandler) UpdateRequest(ctx context.Context, id string, updateData map[string]interface{}) error {
	return h.repo.UpdateRequest(ctx, id, updateData)
}

func (h *RequestHandler) HandleGetAllRequests(ctx context.Context, senderID, receiverID string, includeUserData bool) ([]*model.Request, error) {
	return h.repo.GetAllRequests(ctx, senderID, receiverID, includeUserData)
}
func (h *RequestHandler) HandleGetRequestByID(ctx context.Context, id string) (*model.Request, error) {
	return h.repo.GetRequestByID(ctx, id)
}

func (h *RequestHandler) HandleDeleteRequestByID(ctx context.Context, id string) error {
	return h.repo.DeleteRequestByID(ctx, id)
}
