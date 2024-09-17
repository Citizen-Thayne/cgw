package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	result := NewBoard(10, 10)

	assert.NotNil(t, result)
}

func TestGetCell(t *testing.T) {
	board := NewBoard(10, 10)

	assert.False(t, board.GetCell(0, 0))
}

func TestSetCell(t *testing.T) {
	board := NewBoard(10, 10)

	board.SetCell(0, 0, true)

	assert.True(t, board.GetCell(0, 0))
}

func TestGetAdjacent(t *testing.T) {
	board := NewBoard(10, 10)

	board.SetCell(0, 0, true)
	board.SetCell(1, 1, true)

	assert.Equal(t, 0, board.GetAdjacent(3, 3))
	assert.Equal(t, 1, board.GetAdjacent(0, 0))
	assert.Equal(t, 2, board.GetAdjacent(0, 1))
}

func TestEmptyClone(t *testing.T) {
	board := NewBoard(10, 10)

	board.SetCell(0, 0, true)
	clone := board.EmptyClone()

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			cellValue := clone.GetCell(i, j)
			assert.False(t, cellValue)
		}
	}
}

func TestString(t *testing.T) {
	board := NewBoard(3, 3)

	board.SetCell(0, 0, true)
	board.SetCell(1, 1, true)
	board.SetCell(2, 2, true)

	expected := "0**\n"
	expected += "*0*\n"
	expected += "**0\n"

	assert.Equal(t, expected, board.String())
}

func TestParseBoardString(t *testing.T) {
	boardString := "0**\n"
	boardString += "*0*\n"
	boardString += "**0\n"
	board := ParseBoardString(boardString)

	assert.Equal(t, 3, board.Width)
	assert.Equal(t, 3, board.Height)
	assert.True(t, board.GetCell(0, 0))
	assert.True(t, board.GetCell(1, 1))
	assert.True(t, board.GetCell(2, 2))
}
