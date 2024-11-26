package handlers

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	"github.com/mayuka-c/leaderboard-system-go/internal/app/config"
	"github.com/mayuka-c/leaderboard-system-go/internal/app/models"
	db "github.com/mayuka-c/leaderboard-system-go/internal/pkg/db/sqlc"
)

type IServiceHandler interface {
	// players
	InsertPlayer(ctx context.Context, input models.InsertPlayer) (map[string]int64, error)
	ListPlayers(ctx context.Context) (models.ListPlayer, error)
	UpdatePlayer(ctx context.Context, input models.UpdatePlayer) (int64, error)
	DeletePlayer(ctx context.Context, input models.DeletePlayer) error
	// profile
	CreateProfile(ctx context.Context, input models.CreateProfile) (map[string]int64, error)
	GetPlayerProfile(ctx context.Context, player_id int64) (models.Profile, error)
	UpdateProfile(ctx context.Context, input models.UpdateProfile) (models.Profile, error)
	// games
	InsertGame(ctx context.Context, input models.InsertGame) (map[string]int64, error)
	ListGames(ctx context.Context) (models.ListGames, error)
	DeleteGame(ctx context.Context, input models.DeleteGame) error
	//leaderboard
	UpdateLeaderboard(ctx context.Context, input models.UpdateLeaderboard) (models.LeaderBoard, error)
	GetPlayersScorebyGame(ctx context.Context, gameID int64) ([]models.LeaderBoard, error)
}

type service struct {
	dbQueries *db.Queries
}

func NewServiceHandler() IServiceHandler {
	dbConfig := config.GetDBConfig()
	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", dbConfig.Username, dbConfig.Password, dbConfig.DB_URL, dbConfig.Database)

	// Connect to database
	dbClient, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return &service{
		dbQueries: db.New(dbClient),
	}
}
