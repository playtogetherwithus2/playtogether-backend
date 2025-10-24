package service

import (
	"context"
	"errors"
	"play-together/internal/handler"
	"play-together/internal/model"
)

type PostService struct {
	handler *handler.PostHandler
}

func NewPostService(handler *handler.PostHandler) *PostService {
	return &PostService{handler: handler}
}

func (s *PostService) CreatePost(ctx context.Context, post *model.GamePost) (string, error) {
	if post.Name == "" || post.Venue == "" || post.BackendUserId == "" {
		return "", errors.New("name and venue are required")
	}
	return s.handler.HandleCreatePost(ctx, post)
}

func (s *PostService) GetAllPosts(ctx context.Context) ([]*model.GamePost, error) {
	return s.handler.HandleGetAllPosts(ctx)
}

func (s *PostService) GetPostByID(ctx context.Context, id string) (*model.GamePost, error) {
	if id == "" {
		return nil, errors.New("post ID is required")
	}
	return s.handler.HandleGetPostByID(ctx, id)
}
