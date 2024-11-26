package models

import "time"

type LeaderBoard struct {
	GameID     int64     `json:"game_id,omitempty"`
	Gamename   string    `json:"game_name,omitempty"`
	PlayerID   int64     `json:"player_id,omitempty"`
	Playername string    `json:"player_name,omitempty"`
	Score      int64     `json:"score"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UpdateLeaderboard struct {
	LeaderBoard
}
