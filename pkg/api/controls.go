package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func playButton(c *gin.Context) {

	game := GetInstance().Game
	fmt.Printf("Game is Playing: %v\n", game.IsPlaying)
	c.HTML(200, "controls_play.go.html", gin.H{
		"isPlaying": game.IsPlaying,
	})
}

func playGame(c *gin.Context) {
	game := GetInstance().Game
	game.Play()
	fmt.Printf("Game is Playing: %v\n", game.IsPlaying)
	c.HTML(200, "controls_play.go.html", gin.H{
		"isPlaying": game.IsPlaying,
	})
}

func pauseGame(c *gin.Context) {
	game := GetInstance().Game
	game.Pause()
	fmt.Printf("Game is Playing: %v\n", game.IsPlaying)
	c.HTML(200, "controls_play.go.html", gin.H{
		"isPlaying": game.IsPlaying,
	})
}
