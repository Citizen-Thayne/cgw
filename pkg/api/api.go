package api

import (
	"github.com/gin-gonic/gin"
)

func BindRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.GET("/game", getGame)

	api.POST("/cell/kill/:index", killCell)
	api.POST("/cell/birth/:index", birthCell)

	api.GET("/controls/play", playButton)
	api.POST("/controls/play", playGame)
	api.POST("/controls/pause", pauseGame)

	api.GET("/events", eventStream)
}

func getGame(c *gin.Context) {
	game := GetInstance().Game
	c.HTML(200, "game.go.html", gin.H{
		"game": game,
	})
}
