package api

import (
	"cgw/pkg/board"
	"cgw/pkg/game"
	"sync"
)

type GameSingleton struct {
	Game *game.Game
}

var instance *GameSingleton
var once sync.Once

func GetInstance() *GameSingleton {
	once.Do(func() {
		instance = &GameSingleton{
			Game: game.NewGame(
				board.NewBoard(10, 10),
				game.NewEventDispatcher(),
			),
		}
	})
	return instance
}
