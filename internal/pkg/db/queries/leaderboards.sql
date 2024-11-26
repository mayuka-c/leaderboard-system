-- name: UpsertPlayerScore :one
INSERT INTO leaderboards (
  game_id, player_id, score
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetPlayersScoreByGame :many
SELECT games.name AS gameName, players.username as playerName, leaderboards.score, leaderboards.updated_at
FROM leaderboards
JOIN games ON games.id = leaderboards.game_id
JOIN players ON players.id = leaderboards.player_id
WHERE games.id = $1 
ORDER BY leaderboards.score DESC;