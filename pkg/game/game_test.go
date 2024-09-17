package game

import (
	"cgw/pkg/board"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	b := board.NewBoard(10, 10)
	game := NewGame(b)

	assert.ObjectsAreEqual(b, game.Board)
}

func assertGameEqual(t *testing.T, expected string, actual *Game) {
	actualStr := actual.Board.String()

	assert.Equal(t, expected, actualStr, "Expected %v, got %v", expected, actualStr)
}

func TestGameNextGeneration(t *testing.T) {
	b := board.NewBoard(4, 4)
	game := NewGame(b)

	emptyStr := "****\n"
	emptyStr += "****\n"
	emptyStr += "****\n"
	emptyStr += "****\n"

	// Test empty board
	game.NextGeneration()

	// No change
	assert.Equal(t, emptyStr, game.Board.String())

	// Test 1 cell
	game.Reset()
	game.Board.SetCell(1, 1, true)
	game.NextGeneration()
	assert.Equal(t, emptyStr, game.Board.String())

	// Test 2 adjancent cells
	game.Reset()
	game.Board.SetCell(1, 1, true)
	game.Board.SetCell(1, 2, true)
	game.NextGeneration()
	assert.Equal(t, emptyStr, game.Board.String())

	// Test corner shape
	start := "****\n"
	start += "*00*\n"
	start += "*0**\n"
	start += "****\n"

	expected := "****\n"
	expected += "*00*\n"
	expected += "*00*\n"
	expected += "****\n"

	game = NewGame(board.ParseBoardString(start))

	game.NextGeneration()
	assertGameEqual(t, expected, game)

	// Test blinker
	blinker := "*****\n"
	blinker += "*****\n"
	blinker += "*000*\n"
	blinker += "*****\n"
	blinker += "*****\n"

	blinker_next := "*****\n"
	blinker_next += "**0**\n"
	blinker_next += "**0**\n"
	blinker_next += "**0**\n"
	blinker_next += "*****\n"

	game = NewGame(board.ParseBoardString(start))

	game.NextGeneration()
	assertGameEqual(t, expected, game)
}
