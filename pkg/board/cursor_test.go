package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCursor(t *testing.T) {
	b := NewBoard(2, 2)
	b.SetCell(1, 0, true)
	c := NewCursor(b)

	var value, ok bool

	value, ok = c.Next()
	assert.Equal(t, value, false)
	assert.Equal(t, ok, true)

	value, ok = c.Next()
	assert.Equal(t, value, true)
	assert.Equal(t, ok, true)

	value, ok = c.Next()
	assert.Equal(t, value, false)
	assert.Equal(t, ok, true)

	value, ok = c.Next()
	assert.Equal(t, value, false)
	assert.Equal(t, ok, true)

	value, ok = c.Next()
	assert.Equal(t, value, false)
	assert.Equal(t, ok, false)
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
