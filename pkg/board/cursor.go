package board

type Cursor struct {
	board Board
	i     int
}

func NewCursor(board Board) *Cursor {
	return &Cursor{
		board: board,
		i:     0,
	}
}

func (c *Cursor) Reset() {
	c.i = 0
}

func (c *Cursor) Value() (bool, bool) {
	cells := c.board.Cells()
	isFinal := c.i == len(cells)-1
	return cells[c.i], isFinal
}

func (c *Cursor) Next() {
	c.i++
	if c.i >= len(c.board.Cells()) {
		c.i = len(c.board.Cells()) - 1
	}
}

func (c *Cursor) Set(value bool) {
	c.board.Cells()[c.i] = value
}

func (c *Cursor) Adjacent() int {
	return c.board.GetAdjacent(
		c.X(),
		c.Y(),
	)
}

func (c *Cursor) X() int {
	return c.i % c.board.Width()
}

func (c *Cursor) Y() int {
	return c.i / c.board.Width()
}
