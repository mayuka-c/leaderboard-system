// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: leaderboards.sql

package db

import (
	"context"
	"database/sql"
)

const getPlayersScoreByGame = `-- name: GetPlayersScoreByGame :many
SELECT games.name, players.username, leaderboards.score
FROM leaderboards
JOIN games ON games.id = leaderboards.game_id
JOIN players ON players.id = leaderboards.player_id
WHERE games.id = $1 
ORDER BY players.username
`

type GetPlayersScoreByGameRow struct {
	Name     string        `json:"name"`
	Username string        `json:"username"`
	Score    sql.NullInt64 `json:"score"`
}

func (q *Queries) GetPlayersScoreByGame(ctx context.Context, id int64) ([]GetPlayersScoreByGameRow, error) {
	rows, err := q.db.QueryContext(ctx, getPlayersScoreByGame, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetPlayersScoreByGameRow{}
	for rows.Next() {
		var i GetPlayersScoreByGameRow
		if err := rows.Scan(&i.Name, &i.Username, &i.Score); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const upsertPlayerScore = `-- name: UpsertPlayerScore :one
INSERT INTO leaderboards (
  game_id, player_id, score
) VALUES (
  $1, $2, $3
)
RETURNING game_id, player_id, score, updated_at
`

type UpsertPlayerScoreParams struct {
	GameID   int64         `json:"game_id"`
	PlayerID int64         `json:"player_id"`
	Score    sql.NullInt64 `json:"score"`
}

func (q *Queries) UpsertPlayerScore(ctx context.Context, arg UpsertPlayerScoreParams) (Leaderboard, error) {
	row := q.db.QueryRowContext(ctx, upsertPlayerScore, arg.GameID, arg.PlayerID, arg.Score)
	var i Leaderboard
	err := row.Scan(
		&i.GameID,
		&i.PlayerID,
		&i.Score,
		&i.UpdatedAt,
	)
	return i, err
}