package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCursor(t *testing.T) {
	b := NewBoard(2, 2)
	b.SetCell(1, 0, true)
	c := NewCursor(b)

	value, final := c.Value()
	assert.Equal(t, value, false)
	assert.Equal(t, final, false)

	c.Next()
	value, final = c.Value()
	assert.Equal(t, value, true)
	assert.Equal(t, final, false)

	c.Next()
	value, final = c.Value()
	assert.Equal(t, value, false)
	assert.Equal(t, final, false)

	c.Next()
	value, final = c.Value()
	assert.Equal(t, value, false)
	assert.Equal(t, final, true)
}

func TestCursorSet(t *testing.T) {
	b := NewBoard(2, 2)
	c := NewCursor(b)

	c.Set(true)
	assert.True(t, b.GetCell(0, 0))

	c.Set(false)
	assert.False(t, b.GetCell(0, 0))

	c.Next()
	c.Set(true)
	assert.True(t, b.GetCell(1, 0))
}
