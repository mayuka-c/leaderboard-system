-- name: CreatePlayer :one
INSERT INTO players (
  username, password
) VALUES (
  $1, $2
)
RETURNING id;

-- name: ListPlayers :many
SELECT * FROM players
ORDER BY username;

-- name: UpdatePlayer :one
UPDATE players
  set password = $2
WHERE id = $1
RETURNING id;

-- name: DeletePlayer :exec
DELETE FROM players
WHERE id = $1;