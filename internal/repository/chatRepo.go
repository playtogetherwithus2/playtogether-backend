package repository

import (
	"context"
	"fmt"
	"time"

	"play-together/config"
	"play-together/internal/model"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type ChatRepository interface {
	CreateGroup(ctx context.Context, req model.CreateGroupRequest) (string, error)
	GetMessages(ctx context.Context, groupID string) ([]model.Message, error)
	SendMessage(ctx context.Context, groupID string, req model.SendMessageRequest) error
	AddMember(ctx context.Context, groupID string, req model.ModifyMemberRequest) error
	AddMemberByMatchID(ctx context.Context, matchID string, req model.ModifyMemberRequest) error
	RemoveMember(ctx context.Context, groupID string, req model.ModifyMemberRequest) error
	GetGroupDetails(ctx context.Context, groupID string) (model.GroupDetails, error)
	GetAllGroups(ctx context.Context, memberID string) ([]*model.GroupDetails, error)
}

type chatRepository struct {
	firebaseClient *config.FirebaseClient
}

func NewChatRepository(firebaseClient *config.FirebaseClient) ChatRepository {
	return &chatRepository{firebaseClient: firebaseClient}
}

func (r *chatRepository) CreateGroup(ctx context.Context, req model.CreateGroupRequest) (string, error) {
	fs := r.firebaseClient.Firestore

	// Ensure Members is initialized (avoid nil Firestore field)
	if req.Members == nil {
		req.Members = []string{}
	}

	// Always set CreatedAt from server-side for consistency
	req.CreatedAt = time.Now()

	docRef := fs.Collection("groups").NewDoc()

	_, err := docRef.Set(ctx, map[string]interface{}{
		"group_name": req.GroupName,
		"created_by": req.CreatedBy,
		"match_id":   req.MatchId,
		"members":    req.Members,
		"created_at": req.CreatedAt,
	})
	if err != nil {
		return "", fmt.Errorf("failed to create group: %w", err)
	}

	return docRef.ID, nil
}

func (r *chatRepository) GetAllGroups(ctx context.Context, memberID string) ([]*model.GroupDetails, error) {
	client := r.firebaseClient.Firestore
	var iter *firestore.DocumentIterator

	if memberID != "" {
		iter = client.Collection("groups").Where("members", "array-contains", memberID).Documents(ctx)
	} else {
		iter = client.Collection("groups").Documents(ctx)
	}

	groups := make([]*model.GroupDetails, 0)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var group model.GroupDetails
		if err := doc.DataTo(&group); err == nil {
			group.ID = doc.Ref.ID
			groups = append(groups, &group)
		}
	}

	return groups, nil
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
		{Path: "members", Value: firestore.ArrayUnion(req.UserID)},
	})
	if err != nil {
		return fmt.Errorf("failed to add member: %w", err)
	}
	return nil
}

func (r *chatRepository) AddMemberByMatchID(ctx context.Context, matchID string, req model.ModifyMemberRequest) error {
	fs := r.firebaseClient.Firestore
	iter := fs.Collection("groups").Where("match_id", "==", matchID).Documents(ctx)
	batch := fs.Batch()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("error fetching groups: %w", err)
		}

		batch.Update(doc.Ref, []firestore.Update{
			{Path: "members", Value: firestore.ArrayUnion(req.UserID)},
		})
	}

	_, err := batch.Commit(ctx)
	if err != nil {
		return fmt.Errorf("failed to add member to groups: %w", err)
	}
	return nil
}

func (r *chatRepository) RemoveMember(ctx context.Context, groupID string, req model.ModifyMemberRequest) error {
	fs := r.firebaseClient.Firestore
	docRef := fs.Collection("groups").Doc(groupID)

	_, err := docRef.Update(ctx, []firestore.Update{
		{Path: "members", Value: firestore.ArrayRemove(req.UserID)},
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
