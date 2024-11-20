-- name: CreateGame :one
INSERT INTO games (
  name, created_at
) VALUES (
  $1, $2
)
RETURNING id;

-- name: ListGames :many
SELECT * FROM games
ORDER BY name;

-- name: DeleteGame :exec
DELETE FROM games
WHERE id = $1;