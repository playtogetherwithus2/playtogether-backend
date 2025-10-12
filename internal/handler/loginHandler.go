package handler

import (
	"context"
	"firebase.google.com/go/auth"
	"play-together/internal/repository"
)

type LoginHandler struct {
	repo repository.LoginRepository
}

func NewLoginHandler(repo repository.LoginRepository) *LoginHandler {
	return &LoginHandler{repo: repo}
}

func (h *LoginHandler) HandleLogin(ctx context.Context, email, password string) (string, error) {
	token, err := h.repo.LoginWithEmailAndPassword(ctx, email, password)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (h *LoginHandler) HandleSignup(ctx context.Context, email, password string) (*auth.UserRecord, error) {
	user, err := h.repo.SignupWithEmailAndPassword(ctx, email, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
