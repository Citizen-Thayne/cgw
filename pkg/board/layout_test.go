package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLayoutElements(t *testing.T) {
	board := NewBoard(2, 2)

	board.SetCell(0, 0, true)
	board.SetCell(0, 1, true)

	elements := board.LayoutElements()

	// Two events, 4 for cells, 1 for end line, 1 for end board
	assert.Equal(t, 6, len(elements))
	assert.Equal(t, Alive, elements[0])
	assert.Equal(t, Dead, elements[1])
	assert.Equal(t, RowEnd, elements[2])
	assert.Equal(t, Alive, elements[3])
	assert.Equal(t, Dead, elements[4])
	assert.Equal(t, BoardEnd, elements[5])
}
