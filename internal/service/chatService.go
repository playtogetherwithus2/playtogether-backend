package service

import (
	"context"
	"errors"
	"play-together/internal/handler"
	"play-together/internal/model"
)

type ChatService struct {
	handler *handler.ChatHandler
}

func NewChatService(handler *handler.ChatHandler) *ChatService {
	return &ChatService{handler: handler}
}

func (s *ChatService) CreateGroup(ctx context.Context, req model.CreateGroupRequest) (string, error) {
	if req.GroupName == "" || len(req.Members) == 0 || req.MatchId == "" {
		return "", errors.New("all fields are required")
	}
	return s.handler.CreateGroup(ctx, req)
}

func (s *ChatService) GetAllGroups(ctx context.Context, memberID string) ([]*model.GroupDetails, error) {
	return s.handler.GetAllGroups(ctx, memberID)
}

func (s *ChatService) GetMessages(ctx context.Context, groupID string) ([]model.Message, error) {
	return s.handler.GetMessages(ctx, groupID)
}

func (s *ChatService) SendMessage(ctx context.Context, groupID string, req model.SendMessageRequest) error {
	if req.Text == "" {
		return errors.New("text are required")
	}
	return s.handler.SendMessage(ctx, groupID, req)
}

func (s *ChatService) AddMember(ctx context.Context, groupID string, req model.ModifyMemberRequest) error {
	if req.UserID == "" {
		return errors.New("user_id is required")
	}
	return s.handler.AddMember(ctx, groupID, req)
}

func (s *ChatService) AddMemberByMatchID(ctx context.Context, matchID string, req model.ModifyMemberRequest) error {
	if req.UserID == "" {
		return errors.New("user_id is required")
	}
	return s.handler.AddMemberByMatchID(ctx, matchID, req)
}

func (s *ChatService) RemoveMember(ctx context.Context, groupID string, req model.ModifyMemberRequest) error {
	if req.UserID == "" {
		return errors.New("user_id is required")
	}
	return s.handler.RemoveMember(ctx, groupID, req)
}

func (s *ChatService) GetGroupDetails(ctx context.Context, groupID string) (model.GroupDetails, error) {
	return s.handler.GetGroupDetails(ctx, groupID)
}
