package models

import "time"

type InsertGame struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type GetGame struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type ListGames struct {
	Games []GetGame `json:"games"`
	Total int64     `json:"total"`
}

type DeleteGame struct {
	ID int64 `json:"id"`
}
