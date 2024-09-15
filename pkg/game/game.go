package game

import (
	"cgw/pkg/board"
)

type Game struct {
	board *board.Board
}

func NewGame(board *board.Board) *Game {
	return &Game{
		board: board,
	}
}

func (g *Game) NextGeneration() {
	nextBoard := g.board.EmptyClone()
	currCursor := board.NewCursor(g.board)
	nextCursor := board.NewCursor(nextBoard)

	for {
		adjacent := currCursor.Adjacent()
		value, ok := currCursor.Next()
		if !ok {
			break
		}
		if value {
			if adjacent == 2 || adjacent == 3 {
				nextCursor.Set(true)
			}
		} else {
			if adjacent == 3 {
				nextCursor.Set(true)
			}
		}
		nextCursor.Next()
	}
	g.board = nextBoard
}

func (g *Game) Reset() {
	g.board = g.board.EmptyClone()
}
