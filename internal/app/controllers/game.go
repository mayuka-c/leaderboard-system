package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/mayuka-c/leaderboard-system-go/internal/app/models"
)

func (p *Application) InsertGame() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var game models.InsertGame
		if err := c.BindJSON(&game); err != nil {
			log.Error(err)
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := p.handler.InsertGame(ctx, game)
		if err != nil {
			log.Error(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL SERVER ERROR"})
			return
		}

		c.IndentedJSON(http.StatusCreated, resp)
	}
}

func (p *Application) ListGames() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		resp, err := p.handler.ListGames(ctx)
		if err != nil {
			log.Error(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL SERVER ERROR"})
			return
		}

		c.IndentedJSON(http.StatusOK, resp)
	}
}

func (p *Application) DeleteGame() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var game models.DeleteGame

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		game.ID = int64(id)

		err = p.handler.DeleteGame(ctx, game)
		if err != nil {
			log.Error(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL SERVER ERROR"})
			return
		}

		c.Status(http.StatusNoContent)
	}
}
