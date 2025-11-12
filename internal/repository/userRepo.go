package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"play-together/config"
	"play-together/internal/model"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type UserRepository interface {
	GetUsers(ctx context.Context) ([]model.UserDetails, error)
	GetUserByID(ctx context.Context, userID string) (model.UserDetails, error)
	GetUsersByIDs(ctx context.Context, userIDs []string) ([]model.UserDetails, error)
	UpdateUser(ctx context.Context, id string, req model.UpdateUserRequest) error
}

type userRepository struct {
	firebaseClient *config.FirebaseClient
}

func NewUserRepository(firebaseClient *config.FirebaseClient) UserRepository {
	return &userRepository{firebaseClient: firebaseClient}
}

func (r *userRepository) GetUsers(ctx context.Context) ([]model.UserDetails, error) {
	fsClient := r.firebaseClient.Firestore
	if fsClient == nil {
		return nil, fmt.Errorf("firestore client is not initialized")
	}

	iter := fsClient.Collection("users").Documents(ctx)
	defer iter.Stop()

	var users []model.UserDetails
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error iterating users: %w", err)
		}

		var user model.UserDetails
		if err := doc.DataTo(&user); err != nil {
			return nil, fmt.Errorf("error parsing user data: %w", err)
		}
		user.UID = doc.Ref.ID
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, userID string) (model.UserDetails, error) {
	fsClient := r.firebaseClient.Firestore
	if fsClient == nil {
		return model.UserDetails{}, fmt.Errorf("firestore client is not initialized")
	}

	doc, err := fsClient.Collection("users").Doc(userID).Get(ctx)
	if err != nil {
		return model.UserDetails{}, fmt.Errorf("failed to get user: %w", err)
	}

	var user model.UserDetails
	if err := doc.DataTo(&user); err != nil {
		return model.UserDetails{}, fmt.Errorf("error parsing user data: %w", err)
	}
	user.UID = doc.Ref.ID

	return user, nil
}

func (r *userRepository) GetUsersByIDs(ctx context.Context, userIDs []string) ([]model.UserDetails, error) {
	fsClient := r.firebaseClient.Firestore
	if fsClient == nil {
		return nil, fmt.Errorf("firestore client is not initialized")
	}

	var users []model.UserDetails

	for _, id := range userIDs {
		doc, err := fsClient.Collection("users").Doc(id).Get(ctx)
		if err != nil {
			continue
		}

		var user model.UserDetails
		if err := doc.DataTo(&user); err != nil {
			continue
		}
		user.UID = doc.Ref.ID
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, id string, req model.UpdateUserRequest) error {
	fsClient := r.firebaseClient.Firestore
	if fsClient == nil {
		return fmt.Errorf("firestore client is not initialized")
	}

	data := map[string]interface{}{}

	if req.UserName != "" {
		iter := fsClient.Collection("users").Where("user_name", "==", req.UserName).Documents(ctx)
		defer iter.Stop()

		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return fmt.Errorf("error checking user_name uniqueness: %w", err)
			}

			if doc.Ref.ID != id {
				return fmt.Errorf("user_name '%s' already exists for another user", req.UserName)
			}
		}

		data["user_name"] = req.UserName
	}

	if req.UserName != "" {
		data["user_name"] = req.UserName
	}
	if req.Name != "" {
		data["display_name"] = req.Name
	}
	if req.Age != 0 {
		data["age"] = req.Age
	}
	if req.Gender != "" {
		data["gender"] = req.Gender
	}
	if req.ProfilePhotoURL != "" {
		imageURL, err := uploadImageToCloudinary(req.ProfilePhotoURL)
		if err != nil {
			return fmt.Errorf("image upload failed: %w", err)
		}
		data["profile_photo_url"] = imageURL
	}
	if req.Bio != "" {
		data["bio"] = req.Bio
	}
	if len(req.SportsInterested) > 0 {
		data["sports_interested"] = req.SportsInterested
	}
	if len(req.AvailabilityDays) > 0 {
		data["availability_days"] = req.AvailabilityDays
	}
	if req.PreferredTime != "" {
		data["preferred_time"] = req.PreferredTime
	}
	if len(req.PreferredLocations) > 0 {
		data["preferred_locations"] = req.PreferredLocations
	}
	if req.City != "" {
		data["city"] = req.City
	}

	data["updated_at"] = time.Now()

	if len(data) == 0 {
		return fmt.Errorf("no fields provided for update")
	}

	_, err := fsClient.Collection("users").Doc(id).Set(ctx, data, firestore.MergeAll)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func uploadImageToCloudinary(filePath string) (string, error) {
	cfg := config.LoadConfig()
	data, err := os.ReadFile(cfg.FirebaseConfigPath)
	if err != nil {
		return "", fmt.Errorf("failed to read service account file: %w", err)
	}

	type ServiceAccount struct {
		CloudinaryURL string `json:"cloudinary_url"`
	}

	var sa ServiceAccount
	if err := json.Unmarshal(data, &sa); err != nil {
		return "", fmt.Errorf("failed to parse service account json: %w", err)
	}

	if sa.CloudinaryURL == "" {
		return "", fmt.Errorf("cloudinary_url not found in service account json")
	}

	cld, err := cloudinary.NewFromURL(sa.CloudinaryURL)
	if err != nil {
		return "", fmt.Errorf("failed to initialize cloudinary: %w", err)
	}

	uploadResult, err := cld.Upload.Upload(context.Background(), filePath, uploader.UploadParams{
		Folder: "user_profiles",
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload to cloudinary: %w", err)
	}

	return uploadResult.SecureURL, nil
}
