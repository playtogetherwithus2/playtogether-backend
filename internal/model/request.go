package model

import "time"

type Request struct {
	ID          string    `json:"id,omitempty" firestore:"id,omitempty"`
	SendersId   string    `json:"senders_id" firestore:"senders_id"`
	ReceiversId string    `json:"receivers_id" firestore:"receivers_id"`
	MatchId     string    `json:"match_id" firestore:"match_id"`
	MatchName   string    `json:"match_name" firestore:"match_name"`
	Status      string    `json:"status" firestore:"status"`
	CreatedAt   time.Time `json:"created_at" firestore:"created_at"`

	Sender   *UserDetails `json:"sender,omitempty" firestore:"-"`
	Receiver *UserDetails `json:"receiver,omitempty" firestore:"-"`
}
