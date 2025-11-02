package model

import "time"

type User struct {
	UID                string    `json:"id,omitempty"`
	Email              string    `json:"email" firestore:"email"`
	Password           string    `json:"password,omitempty" firestore:"password"`
	CreatedAt          time.Time `json:"created_at" firestore:"created_at"`
	UserName           string    `json:"user_name,omitempty" firestore:"user_name"`
	Name               string    `json:"display_name,omitempty" firestore:"display_name"`
	Age                int       `json:"age,omitempty" firestore:"age"`
	Gender             string    `json:"gender,omitempty" firestore:"gender"`
	ProfilePhotoURL    string    `json:"profile_photo_url,omitempty" firestore:"profile_photo_url"`
	Bio                string    `json:"bio,omitempty" firestore:"bio"`
	SportsInterested   []string  `json:"sports_interested,omitempty" firestore:"sports_interested"`
	AvailabilityDays   []string  `json:"availability_days,omitempty" firestore:"availability_days"`
	PreferredTime      string    `json:"preferred_time,omitempty" firestore:"preferred_time"`
	PreferredLocations []string  `json:"preferred_locations,omitempty" firestore:"preferred_locations"`
	City               string    `json:"city,omitempty" firestore:"city"`
	UpdatedAt          time.Time `json:"updated_at,omitempty" firestore:"updated_at"`
}

type UserIDsRequest struct {
	UserIDs []string `json:"id"`
}

type UpdateUserRequest struct {
	UserName           string    `json:"user_name" binding:"required"`
	Name               string    `json:"name,omitempty"`
	Age                int       `json:"age,omitempty"`
	Gender             string    `json:"gender,omitempty"`
	ProfilePhotoURL    string    `json:"profile_photo_url,omitempty"`
	Bio                string    `json:"bio,omitempty"`
	SportsInterested   []string  `json:"sports_interested,omitempty"`
	AvailabilityDays   []string  `json:"availability_days,omitempty"`
	PreferredTime      string    `json:"preferred_time,omitempty"`
	PreferredLocations []string  `json:"preferred_locations,omitempty"`
	City               string    `json:"city,omitempty"`
	UpdatedAt          time.Time `json:"updated_at,omitempty"`
}
