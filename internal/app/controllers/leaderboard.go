package controllers

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/mayuka-c/leaderboard-system-go/internal/app/models"
)

func (l *Application) UpdateLeaderboard() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var req models.UpdateLeaderboard
		if err := c.BindJSON(&req); err != nil {
			log.Error(err)
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		gameID, err := strconv.Atoi(c.Param("game_id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		playerID, err := strconv.Atoi(c.Param("player_id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		req.GameID = int64(gameID)
		req.PlayerID = int64(playerID)

		resp, err := l.handler.UpdateLeaderboard(ctx, req)
		if err != nil {
			log.Error(err)
			if strings.Contains(err.Error(), "invalid query") {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL SERVER ERROR"})
			return
		}

		c.IndentedJSON(http.StatusOK, resp)
	}
}

func (l *Application) GetPlayersScorebyGame() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		gameID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := l.handler.GetPlayersScorebyGame(ctx, int64(gameID))
		if err != nil {
			log.Error(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL SERVER ERROR"})
			return
		}

		c.IndentedJSON(http.StatusOK, resp)
	}
}
