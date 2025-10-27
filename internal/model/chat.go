package model

import "time"

type CreateGroupRequest struct {
	MatchId   string    `json:"match_id" firestore:"match_id"`
	GroupName string    `json:"group_name" firestore:"group_name"`
	CreatedBy string    `json:"created_by" firestore:"created_by"`
	Members   []string  `json:"members" firestore:"members"`
	CreatedAt time.Time `json:"created_at" firestore:"created_at"`
}

type SendMessageRequest struct {
	SenderID  string    `json:"sender_id" firestore:"sender_id"`
	Text      string    `json:"text" firestore:"text"`
	Timestamp time.Time `json:"timestamp" firestore:"timestamp"`
}

type ModifyMemberRequest struct {
	UserID string `json:"user_id" firestore:"user_id"`
}

type Message struct {
	SenderID  string    `json:"sender_id" firestore:"sender_id"`
	Text      string    `json:"text" firestore:"text"`
	Timestamp time.Time `json:"timestamp" firestore:"timestamp"`
}

type GroupDetails struct {
	ID        string    `json:"id" firestore:"id"`
	GroupName string    `json:"group_name" firestore:"group_name"`
	CreatedBy string    `json:"created_by" firestore:"created_by"`
	Members   []string  `json:"members" firestore:"members"`
	CreatedAt time.Time `json:"created_at" firestore:"created_at"`
}
