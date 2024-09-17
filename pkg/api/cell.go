package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func killCell(c *gin.Context) {
	index_str := c.Param("index")
	index, err := strconv.Atoi(index_str)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid index"})
		return
	}

	game := GetInstance().Game
	game.Board.SetCellByIndex(index, false)

	c.HTML(200, "cell.go.html", gin.H{
		"index": index,
		"life":  false,
	})
}

func birthCell(c *gin.Context) {
	index_str := c.Param("index")
	index, err := strconv.Atoi(index_str)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid index"})
		return
	}

	game := GetInstance().Game
	game.Board.SetCellByIndex(index, true)

	c.HTML(200, "cell.go.html", gin.H{
		"index": index,
		"life":  true,
	})
}
