package handler

import (
	"context"
	"play-together/internal/model"
	"play-together/internal/repository"
)

type PostHandler struct {
	repo repository.PostRepository
}

func NewPostHandler(repo repository.PostRepository) *PostHandler {
	return &PostHandler{repo: repo}
}

func (h *PostHandler) HandleCreatePost(ctx context.Context, post *model.GamePost) (string, error) {
	return h.repo.CreatePost(ctx, post)
}

func (h *PostHandler) HandleGetAllPosts(ctx context.Context, searchKey, memberIn, memberNotIn string) ([]*model.GamePost, error) {
	return h.repo.GetAllPosts(ctx, searchKey, memberIn, memberNotIn)
}
func (h *PostHandler) HandleGetPostByID(ctx context.Context, id string) (*model.GamePost, error) {
	return h.repo.GetPostByID(ctx, id)
}
