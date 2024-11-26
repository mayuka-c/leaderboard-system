package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"

	"github.com/mayuka-c/leaderboard-system-go/internal/app/models"
	db "github.com/mayuka-c/leaderboard-system-go/internal/pkg/db/sqlc"
)

func (h *service) UpdateLeaderboard(ctx context.Context, input models.UpdateLeaderboard) (models.LeaderBoard, error) {
	var result models.LeaderBoard
	dbScore, err := h.dbQueries.UpsertPlayerScore(ctx, db.UpsertPlayerScoreParams{
		GameID:   input.GameID,
		PlayerID: input.PlayerID,
		Score:    sql.NullInt64{Int64: input.Score, Valid: true},
	})
	if err != nil {
		if strings.Contains(err.Error(), `violates foreign key constraint "leaderboards_player_id_fkey"`) {
			return result, fmt.Errorf("invalid query for player with id: %d does not exist", input.PlayerID)
		}
		if strings.Contains(err.Error(), `violates foreign key constraint "leaderboards_game_id_fkey"`) {
			return result, fmt.Errorf("invalid query for game with id: %d does not exist", input.GameID)
		}
		log.Fatalf("Failed to upsert player score with id: %d with gameID: %d. Err: %s", input.PlayerID, input.GameID, err.Error())
		return result, err
	}

	if err := copier.Copy(&result, dbScore); err != nil {
		log.Fatalf("Error while doing the copy to leaderboard model. Err: %s", err.Error())
		return result, err
	}

	return result, err
}

func (h *service) GetPlayersScorebyGame(ctx context.Context, gameID int64) ([]models.LeaderBoard, error) {
	var result []models.LeaderBoard
	dbPlayers, err := h.dbQueries.GetPlayersScoreByGame(ctx, gameID)
	if err != nil {
		log.Fatalf("Failed to get players score for game with id: %d. Err: %s", gameID, err.Error())
		return result, err
	}

	if err := copier.Copy(&result, dbPlayers); err != nil {
		log.Fatalf("Error while doing the copy to leaderboard model. Err: %s", err.Error())
		return result, err
	}

	return result, err
}
