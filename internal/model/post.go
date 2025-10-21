package model

type GamePost struct {
	ID              string `json:"id,omitempty"`
	Name            string `json:"name" binding:"required"`
	Venue           string `json:"venue" binding:"required"`
	PlayersRequired int    `json:"players_required" binding:"required"`
	Timing          string `json:"timing" binding:"required"`
	NeedExtraPlayer bool   `json:"need_extra_player"`
	Fee             string `json:"fee,omitempty"`
	About           string `json:"about,omitempty"`
	SkillLevel      string `json:"skill_level,omitempty"`
	BackendUserName string `json:"backend_user_name" binding:"required"`
}
