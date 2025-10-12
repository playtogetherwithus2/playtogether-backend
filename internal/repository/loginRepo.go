package repository

import (
	"context"
	"errors"
	"play-together/config"

	"firebase.google.com/go/auth"
)

type LoginRepository interface {
	LoginWithEmailAndPassword(ctx context.Context, email, password string) (string, error)
	SignupWithEmailAndPassword(ctx context.Context, email, password string) (*auth.UserRecord, error)
}

type loginRepository struct {
	firebaseClient *config.FirebaseClient
}

func NewLoginRepository(firebaseClient *config.FirebaseClient) LoginRepository {
	return &loginRepository{firebaseClient: firebaseClient}
}

func (r *loginRepository) LoginWithEmailAndPassword(ctx context.Context, email, password string) (string, error) {
	user, err := r.firebaseClient.Auth.GetUserByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	customToken, err := r.firebaseClient.Auth.CustomToken(ctx, user.UID)
	if err != nil {
		return "", errors.New("failed to create authentication token")
	}

	return customToken, nil
}

func (r *loginRepository) SignupWithEmailAndPassword(ctx context.Context, email, password string) (*auth.UserRecord, error) {
	params := (&auth.UserToCreate{}).Email(email).Password(password)
	user, err := r.firebaseClient.Auth.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}
	return user, nil
}
