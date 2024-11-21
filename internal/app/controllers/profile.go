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

func (p *Application) CreateProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var profile models.CreateProfile
		if err := c.BindJSON(&profile); err != nil {
			log.Error(err)
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		profile.PlayerID = int64(id)

		resp, err := p.handler.CreateProfile(ctx, profile)
		if err != nil {
			log.Error(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL SERVER ERROR"})
			return
		}

		c.IndentedJSON(http.StatusCreated, resp)
	}
}

func (p *Application) GetPlayerProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		player_id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := p.handler.GetPlayerProfile(ctx, int64(player_id))
		if err != nil {
			log.Error(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL SERVER ERROR"})
			return
		}

		if resp.ID == 0 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("There exists no profile for player with id: %d, hence cannot be updated", player_id)})
			return
		}

		c.IndentedJSON(http.StatusOK, resp)
	}
}

func (p *Application) UpdateProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var profile models.UpdateProfile
		if err := c.BindJSON(&profile); err != nil {
			log.Error(err)
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		profile.ID = int64(id)

		resp, err := p.handler.UpdateProfile(ctx, profile)
		if err != nil {
			log.Error(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL SERVER ERROR"})
			return
		}

		if resp.ID == 0 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("There exists no profile with id: %d, hence cannot be updated", id)})
			return
		}

		c.IndentedJSON(http.StatusOK, resp)
	}
}
