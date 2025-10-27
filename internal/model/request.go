package model

import "time"

type Request struct {
	ID          string    `json:"id,omitempty" firestore:"id,omitempty"`
	SendersId   string    `json:"senders_id" firestore:"senders_id"`
	ReceiversId string    `json:"receivers_id" firestore:"receivers_id"`
	MatchId     string    `json:"match_id" firestore:"match_id"`
	CreatedAt   time.Time `json:"created_at" firestore:"created_at"`
}
