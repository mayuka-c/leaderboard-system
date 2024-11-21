-- name: CreateProfile :one
INSERT INTO profiles (
  first_name, last_name, email, age, gender, player_id
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id;

-- name: GetProfile :one
SELECT * FROM profiles
WHERE id = $1;

-- name: UpdateProfile :one
UPDATE profiles 
  set first_name = $2,
  last_name = $3,
  email = $4,
  age = $5,
  gender = $6,
  player_id = $7
WHERE id = $1
RETURNING *;

-- name: DeleteProfile :exec
DELETE FROM profiles
WHERE id = $1;

-- name: PlayerProfile :one
SELECT profiles.*
FROM players
JOIN profiles ON profiles.player_id = players.id
WHERE players.id = $1;