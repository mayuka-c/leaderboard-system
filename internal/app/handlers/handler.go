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
