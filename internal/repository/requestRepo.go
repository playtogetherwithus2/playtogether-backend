package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"play-together/config"
	"play-together/internal/model"

	"cloud.google.com/go/firestore"
)

type RequestRepository interface {
	CreateRequest(ctx context.Context, req model.Request) (string, error)
	GetAllRequests(ctx context.Context, sendersID, receiversID string) ([]*model.Request, error)
	GetRequestByID(ctx context.Context, id string) (*model.Request, error)
	DeleteRequestByID(ctx context.Context, id string) error
}

type requestRepository struct {
	firebaseClient *config.FirebaseClient
}

func NewRequestRepository(firebaseClient *config.FirebaseClient) RequestRepository {
	return &requestRepository{firebaseClient: firebaseClient}
}

func (r *requestRepository) CreateRequest(ctx context.Context, req model.Request) (string, error) {
	fs := r.firebaseClient.Firestore

	docRef := fs.Collection("requests").NewDoc()
	groupData := model.Request{
		SendersId:   req.SendersId,
		ReceiversId: req.ReceiversId,
		MatchId:     req.MatchId,
		CreatedAt:   time.Now(),
	}

	_, err := docRef.Set(ctx, groupData)
	if err != nil {
		return "", fmt.Errorf("failed to create group: %w", err)
	}
	return docRef.ID, nil
}

func (r *requestRepository) GetAllRequests(ctx context.Context, senderID, receiverID string) ([]*model.Request, error) {
	client := r.firebaseClient.Firestore

	var iter *firestore.DocumentIterator

	if senderID != "" {
		iter = client.Collection("requests").Where("senders_id", "==", senderID).Documents(ctx)
	} else if receiverID != "" {
		iter = client.Collection("requests").Where("receivers_id", "==", receiverID).Documents(ctx)
	} else {
		iter = client.Collection("requests").Documents(ctx)
	}

	var requests []*model.Request

	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}
		var request model.Request
		if err := doc.DataTo(&request); err == nil {
			request.ID = doc.Ref.ID
			requests = append(requests, &request)
		}
	}

	return requests, nil
}

func (r *requestRepository) GetRequestByID(ctx context.Context, id string) (*model.Request, error) {
	client := r.firebaseClient.Firestore

	doc, err := client.Collection("requests").Doc(id).Get(ctx)
	if err != nil {
		return nil, errors.New("post not found: " + err.Error())
	}

	var request model.Request
	if err := doc.DataTo(&request); err != nil {
		return nil, errors.New("failed to parse request data: " + err.Error())
	}

	request.ID = doc.Ref.ID
	return &request, nil
}

func (r *requestRepository) DeleteRequestByID(ctx context.Context, id string) error {
	client := r.firebaseClient.Firestore

	_, err := client.Collection("requests").Doc(id).Get(ctx)
	if err != nil {
		return errors.New("request not found: " + err.Error())
	}

	_, err = client.Collection("requests").Doc(id).Delete(ctx)
	if err != nil {
		return errors.New("failed to delete request: " + err.Error())
	}

	return nil
}
