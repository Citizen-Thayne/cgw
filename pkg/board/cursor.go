package board

type Cursor struct {
	board *Board
	x, y  int
}

func NewCursor(board *Board) *Cursor {
	return &Cursor{
		board: board,
		x:     0,
		y:     0,
	}
}

func (c *Cursor) Next() (bool, bool) {
	if c.y >= len(c.board.rows) {
		return false, false
	}
	value := c.board.rows[c.y][c.x]
	c.x++
	if c.x >= len(c.board.rows[0]) {
		c.x = 0
		c.y++
	}
	return value, true
}

func (c *Cursor) Set(value bool) {
	c.board.rows[c.y][c.x] = value
}

func (c *Cursor) Adjacent() int {
	return c.board.GetAdjacent(c.x, c.y)
}
