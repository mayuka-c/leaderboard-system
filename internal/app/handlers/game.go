package handlers

import (
	"context"

	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"

	"github.com/mayuka-c/leaderboard-system-go/internal/app/models"
	db "github.com/mayuka-c/leaderboard-system-go/internal/pkg/db/sqlc"
)

func (h *service) InsertGame(ctx context.Context, input models.InsertGame) (map[string]int64, error) {
	id, err := h.dbQueries.CreateGame(ctx, db.CreateGameParams{
		Name:      input.Name,
		CreatedAt: input.CreatedAt,
	})
	if err != nil {
		log.Fatalf("Failed to insert game into DB. Err: %s", err.Error())
		return map[string]int64{}, err
	}

	return map[string]int64{"id": id}, err
}

func (h *service) ListGames(ctx context.Context) (models.ListGames, error) {
	var result models.ListGames

	dbGames, err := h.dbQueries.ListGames(ctx)
	if err != nil {
		log.Fatalf("Failed to get game records from DB. Err: %s", err.Error())
		return result, err
	}

	games := []models.GetGame{}
	if err := copier.Copy(&games, dbGames); err != nil {
		log.Fatalf("Error while doing the copy to games model. Err: %s", err.Error())
		return result, err
	}

	result.Games = games
	result.Total = int64(len(games))

	return result, err
}

func (h *service) DeleteGame(ctx context.Context, input models.DeleteGame) error {
	err := h.dbQueries.DeleteGame(ctx, input.ID)
	if err != nil {
		log.Fatalf("Failed to delete game record from DB. Err: %s", err.Error())
		return err
	}

	return err
}
