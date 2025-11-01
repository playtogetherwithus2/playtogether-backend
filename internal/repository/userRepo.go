package repository

import (
	"context"
	"fmt"

	"play-together/config"
	"play-together/internal/model"

	"google.golang.org/api/iterator"
)

type UserRepository interface {
	GetUsers(ctx context.Context) ([]model.User, error)
	GetUserByID(ctx context.Context, userID string) (model.User, error)
	GetUsersByIDs(ctx context.Context, userIDs []string) ([]model.User, error)
}

type userRepository struct {
	firebaseClient *config.FirebaseClient
}

func NewUserRepository(firebaseClient *config.FirebaseClient) UserRepository {
	return &userRepository{firebaseClient: firebaseClient}
}

func (r *userRepository) GetUsers(ctx context.Context) ([]model.User, error) {
	fsClient := r.firebaseClient.Firestore
	if fsClient == nil {
		return nil, fmt.Errorf("firestore client is not initialized")
	}

	iter := fsClient.Collection("users").Documents(ctx)
	defer iter.Stop()

	var users []model.User
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error iterating users: %w", err)
		}

		var user model.User
		if err := doc.DataTo(&user); err != nil {
			return nil, fmt.Errorf("error parsing user data: %w", err)
		}
		user.UID = doc.Ref.ID
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, userID string) (model.User, error) {
	fsClient := r.firebaseClient.Firestore
	if fsClient == nil {
		return model.User{}, fmt.Errorf("firestore client is not initialized")
	}

	doc, err := fsClient.Collection("users").Doc(userID).Get(ctx)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to get user: %w", err)
	}

	var user model.User
	if err := doc.DataTo(&user); err != nil {
		return model.User{}, fmt.Errorf("error parsing user data: %w", err)
	}
	user.UID = doc.Ref.ID

	return user, nil
}

func (r *userRepository) GetUsersByIDs(ctx context.Context, userIDs []string) ([]model.User, error) {
	fsClient := r.firebaseClient.Firestore
	if fsClient == nil {
		return nil, fmt.Errorf("firestore client is not initialized")
	}

	var users []model.User

	for _, id := range userIDs {
		doc, err := fsClient.Collection("users").Doc(id).Get(ctx)
		if err != nil {
			continue
		}

		var user model.User
		if err := doc.DataTo(&user); err != nil {
			continue
		}
		user.UID = doc.Ref.ID
		users = append(users, user)
	}

	return users, nil
}
