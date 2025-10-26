package repository

import (
	"context"
	"fmt"
	"time"

	"play-together/config"
	"play-together/internal/model"

	"cloud.google.com/go/firestore"
)

type ChatRepository interface {
	CreateGroup(ctx context.Context, req model.CreateGroupRequest) (string, error)
	GetMessages(ctx context.Context, groupID string) ([]model.Message, error)
	SendMessage(ctx context.Context, groupID string, req model.SendMessageRequest) error
	AddMember(ctx context.Context, groupID string, req model.ModifyMemberRequest) error
	RemoveMember(ctx context.Context, groupID string, req model.ModifyMemberRequest) error
	GetGroupDetails(ctx context.Context, groupID string) (model.GroupDetails, error)
}

type chatRepository struct {
	firebaseClient *config.FirebaseClient
}

func NewChatRepository(firebaseClient *config.FirebaseClient) ChatRepository {
	return &chatRepository{firebaseClient: firebaseClient}
}

func (r *chatRepository) CreateGroup(ctx context.Context, req model.CreateGroupRequest) (string, error) {
	fs := r.firebaseClient.Firestore

	docRef := fs.Collection("groups").NewDoc()
	groupData := model.CreateGroupRequest{
		GroupName: req.GroupName,
		CreatedBy: "lkAIJVjXppZJNB15AEy7sRblgg23",
		Members:   req.Members,
		CreatedAt: time.Now(),
	}

	_, err := docRef.Set(ctx, groupData)
	if err != nil {
		return "", fmt.Errorf("failed to create group: %w", err)
	}
	return docRef.ID, nil
}

func (r *chatRepository) GetMessages(ctx context.Context, groupID string) ([]model.Message, error) {
	fs := r.firebaseClient.Firestore
	msgs := []model.Message{}

	docs, err := fs.Collection("groups").Doc(groupID).Collection("messages").
		OrderBy("Timestamp", firestore.Asc).Documents(ctx).GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch messages: %w", err)
	}

	for _, d := range docs {
		var m model.Message
		d.DataTo(&m)
		msgs = append(msgs, m)
	}
	return msgs, nil
}

func (r *chatRepository) SendMessage(ctx context.Context, groupID string, req model.SendMessageRequest) error {
	fs := r.firebaseClient.Firestore

	message := model.SendMessageRequest{
		SenderID:  "7XovFYxKtrSjh96BYl6s6GYwqXO2",
		Text:      req.Text,
		Timestamp: time.Now(),
	}

	_, _, err := fs.Collection("groups").Doc(groupID).Collection("messages").Add(ctx, message)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	return nil
}

func (r *chatRepository) AddMember(ctx context.Context, groupID string, req model.ModifyMemberRequest) error {
	fs := r.firebaseClient.Firestore
	docRef := fs.Collection("groups").Doc(groupID)

	_, err := docRef.Update(ctx, []firestore.Update{
		{Path: "Members", Value: firestore.ArrayUnion(req.UserID)},
	})
	if err != nil {
		return fmt.Errorf("failed to add member: %w", err)
	}
	return nil
}

func (r *chatRepository) RemoveMember(ctx context.Context, groupID string, req model.ModifyMemberRequest) error {
	fs := r.firebaseClient.Firestore
	docRef := fs.Collection("groups").Doc(groupID)

	_, err := docRef.Update(ctx, []firestore.Update{
		{Path: "Members", Value: firestore.ArrayRemove(req.UserID)},
	})
	if err != nil {
		return fmt.Errorf("failed to remove member: %w", err)
	}
	return nil
}

func (r *chatRepository) GetGroupDetails(ctx context.Context, groupID string) (model.GroupDetails, error) {
	fs := r.firebaseClient.Firestore
	doc, err := fs.Collection("groups").Doc(groupID).Get(ctx)
	if err != nil {
		return model.GroupDetails{}, fmt.Errorf("failed to get group: %w", err)
	}

	var group model.GroupDetails
	doc.DataTo(&group)
	group.ID = doc.Ref.ID
	return group, nil
}
