package models

import "time"

type CreateProfile struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int32  `json:"age"`
	Gender    string `json:"gender"`
	PlayerID  int64  `json:"player_id"`
}

type UpdateProfile struct {
	ID        int64
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int32  `json:"age"`
	Gender    string `json:"gender"`
}

type Profile struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Age       int32     `json:"age"`
	Gender    string    `json:"gender"`
	UpdatedAt time.Time `json:"updated_at"`
}
