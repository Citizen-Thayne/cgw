package api

import (
	"github.com/gin-gonic/gin"
)

func playButton(c *gin.Context) {

	game := GetInstance().Game
	c.HTML(200, "controls_play.go.html", gin.H{
		"isPlaying": game.IsPlaying,
	})
}

func playGame(c *gin.Context) {
	game := GetInstance().Game
	game.Play()
	c.HTML(200, "controls_play.go.html", gin.H{
		"isPlaying": game.IsPlaying,
	})
}

func pauseGame(c *gin.Context) {
	game := GetInstance().Game
	game.Pause()
	c.HTML(200, "controls_play.go.html", gin.H{
		"isPlaying": game.IsPlaying,
	})
}
