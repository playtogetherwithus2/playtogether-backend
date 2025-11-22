package repository

import (
	"context"
	"errors"
	"play-together/config"
	"play-together/internal/model"
	"regexp"
	"strings"
	"time"
)

type PostRepository interface {
	CreatePost(ctx context.Context, post *model.GamePost) (string, error)
	GetAllPosts(ctx context.Context, searchKey string, memberIn string, memberNotIn string) ([]*model.GamePost, error)
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
	docRef, _, err := client.Collection("matches").Add(ctx, post)
	if err != nil {
		return "", errors.New("failed to create post: " + err.Error())
	}
	return docRef.ID, nil
}

func (r *postRepository) GetAllPosts(ctx context.Context, searchKey, memberIn, memberNotIn string) ([]*model.GamePost, error) {
	client := r.firebaseClient.Firestore

	iter := client.Collection("matches").Documents(ctx)
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

	if memberIn != "" || memberNotIn != "" {
		groupIter := client.Collection("groups").Documents(ctx)

		allMatchIDs := make(map[string]bool)
		userMatchIDs := make(map[string]bool)

		for {
			doc, err := groupIter.Next()
			if err != nil {
				break
			}

			var group model.GroupDetails
			if err := doc.DataTo(&group); err != nil {
				continue
			}

			allMatchIDs[group.MatchId] = true

			for _, m := range group.Members {
				if m == memberIn || m == memberNotIn {
					userMatchIDs[group.MatchId] = true
					break
				}
			}
		}

		var filtered []*model.GamePost

		if memberIn != "" {
			for _, p := range posts {
				if userMatchIDs[p.ID] {
					filtered = append(filtered, p)
				}
			}
		}

		if memberNotIn != "" {
			for _, p := range posts {
				if allMatchIDs[p.ID] && !userMatchIDs[p.ID] {
					filtered = append(filtered, p)
				}
			}
		}

		posts = filtered
	}

	now := time.Now().Local()
	var futurePosts []*model.GamePost

	cleanSpaces := func(s string) string {
		s = strings.TrimSpace(s)
		s = strings.ReplaceAll(s, "\u202F", " ") 
		s = strings.ReplaceAll(s, "\u00A0", " ") 
		s = strings.ReplaceAll(s, "\t", " ")     
		s = strings.ReplaceAll(s, "  ", " ")   
		return strings.TrimSpace(s)
	}

	for _, p := range posts {

		dateStr := cleanSpaces(p.Date)
		dateParsed, err := time.Parse("02-01-2006", dateStr)
		if err != nil {
			continue
		}

		timing := cleanSpaces(p.Timing)
		timing = strings.ToUpper(timing)
		timeParts := strings.Split(timing, "-")
		if len(timeParts) != 2 {
			continue
		}

		startTimeStr := strings.TrimSpace(timeParts[0]) 

		startTimeParsed, err := time.Parse("3:04 PM", startTimeStr)
		if err != nil {
			continue
		}

		matchStart := time.Date(
			dateParsed.Year(),
			dateParsed.Month(),
			dateParsed.Day(),
			startTimeParsed.Hour(),
			startTimeParsed.Minute(),
			0, 0,
			time.Local,
		)

		if matchStart.After(now) {
			futurePosts = append(futurePosts, p)
		}
	}

	posts = futurePosts

	if searchKey == "" {
		return posts, nil
	}

	searchKey = strings.ToLower(strings.TrimSpace(searchKey))
	pattern := ".*"
	for _, ch := range searchKey {
		pattern += regexp.QuoteMeta(string(ch)) + ".*"
	}
	re, _ := regexp.Compile("(?i)" + pattern)

	var filtered []*model.GamePost
	for _, p := range posts {
		if re.MatchString(strings.ToLower(p.Name)) || re.MatchString(strings.ToLower(p.Venue)) {
			filtered = append(filtered, p)
		}
	}

	return filtered, nil
}

func (r *postRepository) GetPostByID(ctx context.Context, id string) (*model.GamePost, error) {
	client := r.firebaseClient.Firestore

	doc, err := client.Collection("matches").Doc(id).Get(ctx)
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
