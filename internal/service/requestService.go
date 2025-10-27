package service

import (
	"context"
	"errors"
	"play-together/internal/handler"
	"play-together/internal/model"
)

type RequestService struct {
	handler *handler.RequestHandler
}

func NewRequestService(handler *handler.RequestHandler) *RequestService {
	return &RequestService{handler: handler}
}

func (s *RequestService) CreateRequest(ctx context.Context, req model.Request) (string, error) {
	if req.SendersId == "" || req.ReceiversId == "" || req.MatchId == "" {
		return "", errors.New("All details are required")
	}
	return s.handler.CreateRequest(ctx, req)
}

func (s *RequestService) GetAllRequests(ctx context.Context, senderID, receiverID string) ([]*model.Request, error) {
	return s.handler.HandleGetAllRequests(ctx, senderID, receiverID)
}

func (s *RequestService) GetRequestByID(ctx context.Context, id string) (*model.Request, error) {
	if id == "" {
		return nil, errors.New("request ID is required")
	}
	return s.handler.HandleGetRequestByID(ctx, id)
}
