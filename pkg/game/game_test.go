package game

import (
	"cgw/pkg/board"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewGame(t *testing.T) {
	b := board.NewBoard(10, 10)
	game := NewGame(b, NewEventDispatcher())

	assert.ObjectsAreEqual(b, game.Board)
}

func assertGameEqual(t *testing.T, expected string, actual *Game) {
	actualStr := actual.Board.String()

	assert.Equal(t, expected, actualStr, "Expected %v, got %v", expected, actualStr)
}

type MockEventDispatcher struct {
	mock.Mock
}

func TestGameNextGeneration(t *testing.T) {
	b := board.NewBoard(4, 4)
	var dispatcher = createMockEventDispatcher()
	game := NewGame(b, dispatcher)

	emptyStr := "****\n"
	emptyStr += "****\n"
	emptyStr += "****\n"
	emptyStr += "****\n"

	// Test empty board
	game.NextGeneration()

	// No change
	assert.Equal(t, emptyStr, game.Board.String())
	assert.Equal(t, 1, len(dispatcher.dispatchEvents))
	dispatcher.reset()

	// Test 1 cell
	game.Board.SetCell(1, 1, true)
	game.NextGeneration()
	assert.Equal(t, emptyStr, game.Board.String())
	assert.Equal(t, 1, len(dispatcher.dispatchEvents))
	dispatcher.reset()

	// Test 2 adjancent cells
	game.Reset()
	game.Board.SetCell(1, 1, true)
	game.Board.SetCell(1, 2, true)
	game.NextGeneration()
	assert.Equal(t, emptyStr, game.Board.String())
	assert.Equal(t, 1, len(dispatcher.dispatchEvents))
	dispatcher.reset()

	// Test corner shape
	start := "****\n"
	start += "*00*\n"
	start += "*0**\n"
	start += "****\n"

	expected := "****\n"
	expected += "*00*\n"
	expected += "*00*\n"
	expected += "****\n"

	game = NewGame(board.ParseBoardString(start), dispatcher)

	game.NextGeneration()
	assertGameEqual(t, expected, game)
	assert.Equal(t, 1, len(dispatcher.dispatchEvents))
	dispatcher.reset()

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

	game = NewGame(board.ParseBoardString(start), dispatcher)

	game.NextGeneration()
	assertGameEqual(t, expected, game)
	assert.Equal(t, 1, len(dispatcher.dispatchEvents))
	dispatcher.reset()
}

func TestSubscribr(t *testing.T) {
	b := board.NewBoard(4, 4)
	var dispatcher = createMockEventDispatcher()
	c := make(chan int)
	game := NewGame(b, dispatcher)

	game.Subscribe(c)

	assert.Equal(t, 1, len(dispatcher.subscribers))
	assert.ElementsMatch(t, []chan int{c}, dispatcher.subscribers)
}
