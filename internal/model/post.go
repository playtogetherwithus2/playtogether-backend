package model

import "time"

type GamePost struct {
	ID              string    `json:"id,omitempty" firestore:"id,omitempty"`
	Name            string    `json:"name" binding:"required" firestore:"name"`
	Venue           string    `json:"venue" binding:"required" firestore:"venue"`
	PlayersRequired int       `json:"players_required" binding:"required" firestore:"players_required"`
	Timing          string    `json:"timing" binding:"required" firestore:"timing"`
	NeedExtraPlayer bool      `json:"need_extra_player" firestore:"need_extra_player"`
	Fee             string    `json:"fee,omitempty" firestore:"fee,omitempty"`
	About           string    `json:"about,omitempty" firestore:"about,omitempty"`
	SkillLevel      string    `json:"skill_level,omitempty" firestore:"skill_level,omitempty"`
	BackendUserId   string    `json:"backend_user_id" binding:"required" firestore:"backend_user_id"`
	CreatedAt       time.Time `json:"createdAt" firestore:"createdAt"`
}
