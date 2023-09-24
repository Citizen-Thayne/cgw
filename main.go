package main

import (
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Game struct {
	numCols int
	numRows int
	cells   [][]bool
}

func render(game *Game, w io.Writer) {
	var builder strings.Builder
	builder.WriteString(`<div class="grid-container">`)
	for _, row := range game.cells {
		for _, cell := range row {
			if cell {
				builder.WriteString(`<div class="cell alive"></div>`)
			} else {
				builder.WriteString(`<div class="cell dead"></div>`)
			}
		}
	}
	builder.WriteString(`</div>`)
	w.Write([]byte(builder.String()))
}

func handleStatic(path string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path)
	}
}

func NewGame(numCols int, numRows int) Game {
	var cells [][]bool
	for i := 0; i < numRows; i++ {
		row := make([]bool, numCols)
		for j := range row {
			row[j] = false
		}
		cells = append(cells, row)
	}

	return Game{
		numCols: numCols,
		numRows: numRows,
		cells:   cells,
	}
}

func main() {
	r := gin.Default()
	r.StaticFile("/", "index.html")
	r.StaticFile("/styles.css", "styles.css")
	r.LoadHTMLGlob("templates/*")

	game := NewGame(3, 3)

	r.GET("/api/game", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "game.go.html", gin.H{"cells": game.cells} )
	})
	r.Run(":3333")
}
