package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/mayuka-c/leaderboard-system-go/internal/app/config"
	"github.com/mayuka-c/leaderboard-system-go/internal/app/controllers"
	"github.com/mayuka-c/leaderboard-system-go/routes"
)

func loadConfigs() {
	config.LoadServiceConfig()
	config.LoadDBConfig()
}

func main() {
	loadConfigs()

	router := gin.New()
	router.Use(gin.Logger())

	routes.PlayersRoutes(router, controllers.NewApplication())
	routes.GamesRoutes(router, controllers.NewApplication())
	routes.LeaderBoardsRoutes(router, controllers.NewApplication())

	log.Println("LeaderBoard System is running at port: ", config.GetServiceConfig().APIPort)
	log.Fatal(router.Run(":" + strconv.Itoa(config.GetServiceConfig().APIPort)))
}
