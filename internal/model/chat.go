package model

import "time"

type CreateGroupRequest struct {
	GroupName string    `json:"group_name"`
	CreatedBy string    `json:"created_by"`
	Members   []string  `json:"members"`
	CreatedAt time.Time `json:"created_at"`
}

type SendMessageRequest struct {
	SenderID  string    `json:"sender_id"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

type ModifyMemberRequest struct {
	UserID string `json:"user_id"`
}

type Message struct {
	SenderID  string    `json:"sender_id"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

type GroupDetails struct {
	ID        string    `json:"id"`
	GroupName string    `json:"group_name"`
	CreatedBy string    `json:"created_by"`
	Members   []string  `json:"members"`
	CreatedAt time.Time `json:"created_at"`
}
