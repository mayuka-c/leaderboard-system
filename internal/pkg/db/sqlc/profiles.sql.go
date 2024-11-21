// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: profiles.sql

package db

import (
	"context"
)

const createProfile = `-- name: CreateProfile :one
INSERT INTO profiles (
  first_name, last_name, email, age, gender, player_id
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id
`

type CreateProfileParams struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	Age       int32   `json:"age"`
	Gender    GenderT `json:"gender"`
	PlayerID  int64   `json:"player_id"`
}

func (q *Queries) CreateProfile(ctx context.Context, arg CreateProfileParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createProfile,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Age,
		arg.Gender,
		arg.PlayerID,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const deleteProfile = `-- name: DeleteProfile :exec
DELETE FROM profiles
WHERE id = $1
`

func (q *Queries) DeleteProfile(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteProfile, id)
	return err
}

const getProfile = `-- name: GetProfile :one
SELECT id, first_name, last_name, email, age, gender, player_id, updated_at FROM profiles
WHERE id = $1
`

func (q *Queries) GetProfile(ctx context.Context, id int64) (Profile, error) {
	row := q.db.QueryRowContext(ctx, getProfile, id)
	var i Profile
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Age,
		&i.Gender,
		&i.PlayerID,
		&i.UpdatedAt,
	)
	return i, err
}

const playerProfile = `-- name: PlayerProfile :one
SELECT profiles.id, profiles.first_name, profiles.last_name, profiles.email, profiles.age, profiles.gender, profiles.player_id, profiles.updated_at
FROM players
JOIN profiles ON profiles.player_id = players.id
WHERE players.id = $1
`

func (q *Queries) PlayerProfile(ctx context.Context, id int64) (Profile, error) {
	row := q.db.QueryRowContext(ctx, playerProfile, id)
	var i Profile
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Age,
		&i.Gender,
		&i.PlayerID,
		&i.UpdatedAt,
	)
	return i, err
}

const updateProfile = `-- name: UpdateProfile :one
UPDATE profiles 
  set first_name = $2,
  last_name = $3,
  email = $4,
  age = $5,
  gender = $6,
  player_id = $7
WHERE id = $1
RETURNING id, first_name, last_name, email, age, gender, player_id, updated_at
`

type UpdateProfileParams struct {
	ID        int64   `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	Age       int32   `json:"age"`
	Gender    GenderT `json:"gender"`
	PlayerID  int64   `json:"player_id"`
}

func (q *Queries) UpdateProfile(ctx context.Context, arg UpdateProfileParams) (Profile, error) {
	row := q.db.QueryRowContext(ctx, updateProfile,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Age,
		arg.Gender,
		arg.PlayerID,
	)
	var i Profile
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Age,
		&i.Gender,
		&i.PlayerID,
		&i.UpdatedAt,
	)
	return i, err
}
