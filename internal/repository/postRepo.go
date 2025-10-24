package repository

import (
	"context"
	"time"
	"errors"
	"play-together/config"
	"play-together/internal/model"
)

type PostRepository interface {
	CreatePost(ctx context.Context, post *model.GamePost) (string, error)
	GetAllPosts(ctx context.Context) ([]*model.GamePost, error)
	GetPostByID(ctx context.Context, id string) (*model.GamePost, error)
}

type postRepository struct {
	firebaseClient *config.FirebaseClient
}

func NewPostRepository(firebaseClient *config.FirebaseClient) PostRepository {
	return &postRepository{firebaseClient: firebaseClient}
}

func (r *postRepository) CreatePost(ctx context.Context, post *model.GamePost) (string, error) {
	client := r.firebaseClient.Firestore
	post.CreatedAt = time.Now()
	docRef, _, err := client.Collection("game_posts").Add(ctx, post)
	if err != nil {
		return "", errors.New("failed to create post: " + err.Error())
	}
	return docRef.ID, nil
}

func (r *postRepository) GetAllPosts(ctx context.Context) ([]*model.GamePost, error) {
	client := r.firebaseClient.Firestore
	iter := client.Collection("game_posts").Documents(ctx)

	var posts []*model.GamePost
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}
		var post model.GamePost
		if err := doc.DataTo(&post); err == nil {
			post.ID = doc.Ref.ID
			posts = append(posts, &post)
		}
	}
	return posts, nil
}

func (r *postRepository) GetPostByID(ctx context.Context, id string) (*model.GamePost, error) {
	client := r.firebaseClient.Firestore

	doc, err := client.Collection("game_posts").Doc(id).Get(ctx)
	if err != nil {
		return nil, errors.New("post not found: " + err.Error())
	}

	var post model.GamePost
	if err := doc.DataTo(&post); err != nil {
		return nil, errors.New("failed to parse post data: " + err.Error())
	}

	post.ID = doc.Ref.ID
	return &post, nil
}
