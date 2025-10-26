package model

import "time"

type Request struct {
	ID          string    `json:"id,omitempty"`
	SendersId   string    `json:"senders_id"`
	ReceiversId string    `json:"receivers_id"`
	MatchId     string    `json:"match_id"`
	CreatedAt   time.Time `json:"created_at"`
}
