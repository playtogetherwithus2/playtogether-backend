package handler

import (
	"context"
	"play-together/internal/model"
	"play-together/internal/repository"
)

type ChatHandler struct {
	repo repository.ChatRepository
}

func NewChatHandler(repo repository.ChatRepository) *ChatHandler {
	return &ChatHandler{repo: repo}
}

func (h *ChatHandler) CreateGroup(ctx context.Context, req model.CreateGroupRequest) (string, error) {
	return h.repo.CreateGroup(ctx, req)
}

func (h *ChatHandler) GetAllGroups(ctx context.Context, memberID, groupName string) ([]*model.GroupDetails, error) {
	return h.repo.GetAllGroups(ctx, memberID, groupName)
}

func (h *ChatHandler) GetMessages(ctx context.Context, groupID string) ([]model.Message, error) {
	return h.repo.GetMessages(ctx, groupID)
}

func (h *ChatHandler) SendMessage(ctx context.Context, groupID string, req model.SendMessageRequest) error {
	return h.repo.SendMessage(ctx, groupID, req)
}

func (h *ChatHandler) AddMember(ctx context.Context, groupID string, req model.ModifyMemberRequest) error {
	return h.repo.AddMember(ctx, groupID, req)
}

func (h *ChatHandler) AddMemberByMatchID(ctx context.Context, matchID string, req model.ModifyMemberRequest) error {
	return h.repo.AddMemberByMatchID(ctx, matchID, req)
}

func (h *ChatHandler) RemoveMember(ctx context.Context, groupID string, req model.ModifyMemberRequest) error {
	return h.repo.RemoveMember(ctx, groupID, req)
}

func (h *ChatHandler) GetGroupDetails(ctx context.Context, groupID string) (model.GroupDetails, error) {
	return h.repo.GetGroupDetails(ctx, groupID)
}
