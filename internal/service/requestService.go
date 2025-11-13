package service

import (
	"context"
	"errors"
	"fmt"
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

func (s *RequestService) UpdateRequest(ctx context.Context, id string, updateData map[string]interface{}) error {
	if id == "" {
		return errors.New("request ID is required")
	}

	// Optionally validate allowed fields to update
	allowedFields := map[string]bool{
		"status": true,
	}

	for key := range updateData {
		if !allowedFields[key] {
			return fmt.Errorf("updating field '%s' is not allowed", key)
		}
	}

	return s.handler.UpdateRequest(ctx, id, updateData)
}

func (s *RequestService) GetAllRequests(ctx context.Context, senderID, receiverID string, includeUserData bool) ([]*model.Request, error) {
	return s.handler.HandleGetAllRequests(ctx, senderID, receiverID, includeUserData)
}

func (s *RequestService) GetRequestByID(ctx context.Context, id string, includeUserData bool) (*model.Request, error) {
	if id == "" {
		return nil, errors.New("request ID is required")
	}
	return s.handler.HandleGetRequestByID(ctx, id, includeUserData)
}

func (s *RequestService) DeleteRequestByID(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("request ID is required")
	}
	return s.handler.HandleDeleteRequestByID(ctx, id)
}
