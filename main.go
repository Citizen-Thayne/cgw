package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Game struct {
	numCols int
	numRows int
	cells [][]bool
}

func render(game *Game, w io.StringWriter) {

}



func handleStatic(path string) func (w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path)
	}
}

func NewGame(numCols int, numRows int) {
	var cells [][]int
	for i := 0; i < numRows; i++ {
		row := make([]int, numCols)
		cells = append(cells, row)
	}
	return Game {
		numCols: numCols,
		numRows: numRows,
		cells: cells
	}
}

func main() {
	game := NewGame(3, 3)
	http.HandleFunc("/", handleStatic("index.html"))
	http.HandleFunc("/styles.css", handleStatic("styles.css"))
	http.HandleFunc("/api/game", func(w http.ResponseWriter) {
		render
	})
	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
