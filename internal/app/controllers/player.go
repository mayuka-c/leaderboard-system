package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/mayuka-c/leaderboard-system-go/internal/app/models"
)

func (p *Application) InsertPlayer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var player models.InsertPlayer
		if err := c.BindJSON(&player); err != nil {
			log.Error(err)
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := p.handler.InsertPlayer(ctx, player)
		if err != nil {
			log.Error(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL SERVER ERROR"})
			return
		}

		c.IndentedJSON(http.StatusCreated, resp)
	}
}

func (p *Application) ListPlayers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		resp, err := p.handler.ListPlayers(ctx)
		if err != nil {
			log.Error(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL SERVER ERROR"})
			return
		}

		c.IndentedJSON(http.StatusOK, resp)
	}
}

func (p *Application) UpdatePlayer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var player models.UpdatePlayer
		if err := c.BindJSON(&player); err != nil {
			log.Error(err)
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		player.ID = int64(id)

		resp, err := p.handler.UpdatePlayer(ctx, player)
		if err != nil {
			log.Error(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL SERVER ERROR"})
			return
		}

		if resp == 0 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("There exists no player with id: %d, hence cannot be updated", id)})
			return
		}

		c.IndentedJSON(http.StatusOK, map[string]string{"msg": fmt.Sprintf("Player with id: %d successfully updated", resp)})
	}
}

func (p *Application) DeletePlayer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var player models.DeletePlayer

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		player.ID = int64(id)

		err = p.handler.DeletePlayer(ctx, player)
		if err != nil {
			log.Error(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL SERVER ERROR"})
			return
		}

		c.Status(http.StatusNoContent)
	}
}
