package model

import "time"

type User struct {
	UID       string    `json:"id" firestore:"id"`
	Email     string    `json:"email" firestore:"email"`
	Password  string    `json:"password,omitempty" firestore:"password"`
	CreatedAt time.Time `json:"createdAt" firestore:"createdAt"`
}

type UserIDsRequest struct {
	UserIDs []string `json:"id"`
}
