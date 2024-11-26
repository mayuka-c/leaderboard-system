package handlers

import (
	"context"
	"strings"

	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"

	"github.com/mayuka-c/leaderboard-system-go/internal/app/models"
	db "github.com/mayuka-c/leaderboard-system-go/internal/pkg/db/sqlc"
)

func (h *service) InsertPlayer(ctx context.Context, input models.InsertPlayer) (map[string]int64, error) {
	id, err := h.dbQueries.CreatePlayer(ctx, db.CreatePlayerParams{
		Username: input.Username,
		Password: input.Password,
	})
	if err != nil {
		log.Fatalf("Failed to insert player record into DB. Err: %s", err.Error())
		return map[string]int64{}, err
	}

	return map[string]int64{"id": id}, err
}

func (h *service) ListPlayers(ctx context.Context) (models.ListPlayer, error) {
	var result models.ListPlayer

	dbPlayers, err := h.dbQueries.ListPlayers(ctx)
	if err != nil {
		log.Fatalf("Failed to get player records from DB. Err: %s", err.Error())
		return result, err
	}

	players := []models.GetPlayer{}
	if err := copier.Copy(&players, dbPlayers); err != nil {
		log.Fatalf("Error while doing the copy to players model. Err: %s", err.Error())
		return result, err
	}

	result.Players = players
	result.Total = int64(len(players))

	return result, err
}

func (h *service) UpdatePlayer(ctx context.Context, input models.UpdatePlayer) (int64, error) {
	id, err := h.dbQueries.UpdatePlayer(ctx, db.UpdatePlayerParams{
		ID:       input.ID,
		Password: input.Password,
	})
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return 0, nil
		}
		log.Fatalf("Failed to update player record into DB. Err: %s", err.Error())
		return 0, err
	}

	return id, err
}

func (h *service) DeletePlayer(ctx context.Context, input models.DeletePlayer) error {
	err := h.dbQueries.DeletePlayer(ctx, input.ID)
	if err != nil {
		log.Fatalf("Failed to delete player record from DB. Err: %s", err.Error())
		return err
	}

	return err
}
