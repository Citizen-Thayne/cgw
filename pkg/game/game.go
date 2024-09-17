package game

import (
	"cgw/pkg/board"
	"fmt"
	"time"
)

type Game struct {
	Board     *board.Board
	IsPlaying bool
	playChan  chan bool
}

func NewGame(board *board.Board) *Game {
	ticker := time.NewTicker(2 * time.Second)
	game := &Game{
		Board:     board,
		playChan:  make(chan bool),
		IsPlaying: false,
	}

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("Play:", game.IsPlaying)
				if game.IsPlaying {
					game.NextGeneration()
				}

			case play := <-game.playChan:
				game.IsPlaying = play
			}
		}
	}()

	return game
}

func (g *Game) NextGeneration() {
	nextBoard := g.Board.EmptyClone()
	currCursor := board.NewCursor(g.Board)
	nextCursor := board.NewCursor(nextBoard)

	for {
		adjacent := currCursor.Adjacent()
		value, final := currCursor.Value()
		if value {
			if adjacent == 2 || adjacent == 3 {
				nextCursor.Set(true)
			}
		} else {
			if adjacent == 3 {
				nextCursor.Set(true)
			}
		}
		if final {
			break
		}
		nextCursor.Next()
		currCursor.Next()
	}
	g.Board = nextBoard
}

func (g *Game) Reset() {
	g.Board = g.Board.EmptyClone()
}

func (g *Game) Play() {
	g.playChan <- true
	g.IsPlaying = true
}

func (g *Game) Pause() {
	g.playChan <- false
	g.IsPlaying = false
}
