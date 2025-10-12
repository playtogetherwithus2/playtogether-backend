package service

import (
	"context"
	"errors"
	"play-together/internal/handler"

	"firebase.google.com/go/auth"
)

type LoginService struct {
	handler *handler.LoginHandler
}

func NewLoginService(handler *handler.LoginHandler) *LoginService {
	return &LoginService{handler: handler}
}

func (s *LoginService) Login(ctx context.Context, email, password string) (string, error) {
	if email == "" || password == "" {
		return "", errors.New("email and password required")
	}

	token, err := s.handler.HandleLogin(ctx, email, password)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *LoginService) Signup(ctx context.Context, email, password string) (*auth.UserRecord, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password required")
	}

	user, err := s.handler.HandleSignup(ctx, email, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
