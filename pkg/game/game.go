package game

import (
	"cgw/pkg/board"
	"time"
)

type Game struct {
	Board           board.Board
	IsPlaying       bool
	playChan        chan bool
	eventDispatcher EventDispatcher
}

func NewGame(board board.Board, e EventDispatcher) *Game {
	ticker := time.NewTicker(time.Second)

	game := &Game{
		Board:           board,
		playChan:        make(chan bool),
		IsPlaying:       false,
		eventDispatcher: e,
	}

	go func() {
		for {
			select {
			case <-ticker.C:
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
	g.eventDispatcher.Dispatch(NewGeneration)
}

func (g *Game) Reset() {
	g.Board = g.Board.EmptyClone()
}

func (g *Game) Play() {
	g.playChan <- true
	g.IsPlaying = true
	g.eventDispatcher.Dispatch(GameResumed)
}

func (g *Game) Pause() {
	g.playChan <- false
	g.IsPlaying = false
	g.eventDispatcher.Dispatch(GamePaused)
}

func (g *Game) Subscribe(c chan int) {
	g.eventDispatcher.Subscribe(c)
}
