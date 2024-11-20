package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/mayuka-c/leaderboard-system-go/internal/app/controllers"
)

func PlayersRoutes(incomingRoutes *gin.Engine, handler *controllers.Application) {
	incomingRoutes.POST("/api/v1/player", handler.InsertPlayer())
	incomingRoutes.GET("/api/v1/players", handler.ListPlayers())
	incomingRoutes.PUT("/api/v1/player/:id", handler.UpdatePlayer())
	incomingRoutes.DELETE("/api/v1/player/:id", handler.DeletePlayer())
}

func GamesRoutes(incomingRoutes *gin.Engine, handler *controllers.Application) {
	incomingRoutes.POST("/api/v1/game", handler.InsertGame())
	incomingRoutes.GET("/api/v1/games", handler.ListGames())
	incomingRoutes.GET("/api/v1/game/:id", handler.DeleteGame())
}

func LeaderBoardsRoutes(incomingRoutes *gin.Engine, handler *controllers.Application) {
	incomingRoutes.PUT("/api/v1/game/:game_id/player/:player_id/stats", handler.UpdateLeaderboard())
	incomingRoutes.GET("/api/v1/game/:id/players/stats", handler.GetPlayersScorebyGame())
}
