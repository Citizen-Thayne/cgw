package board

import (
	"strings"
)

type Board interface {
	Height() int
	Width() int
	Cells() []bool
	GetCell(x, y int) bool
	GetCellByIndex(index int, value bool) bool
	SetCell(x, y int, value bool)
	SetCellByIndex(index int, value bool)
	GetAdjacent(x, y int) int
	String() string
	EmptyClone() Board
	LayoutElements() []int
}

type board struct {
	height int
	width  int
	cells  []bool
}

func NewBoard(numCols int, numRows int) Board {
	board := board{
		height: numRows,
		width:  numCols,
		cells:  make([]bool, numCols*numRows),
	}

	for i := range board.cells {
		board.cells[i] = false
	}
	return &board
}

func (b *board) Height() int {
	return b.height
}

func (b *board) Width() int {
	return b.width
}

func (b *board) Cells() []bool {
	return b.cells
}

func (b *board) GetCell(x, y int) bool {
	return b.cells[y*b.width+x]
}

func (b *board) SetCell(x, y int, value bool) {
	b.cells[y*b.width+x] = value
}

func (b *board) SetCellByIndex(index int, value bool) {
	b.cells[index] = value
}

func (b *board) GetCellByIndex(index int, value bool) bool {
	return b.cells[index]
}

func (b *board) GetAdjacent(x, y int) int {
	adjacent := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			// Skip the cell itself
			if i == 0 && j == 0 {
				continue
			}
			// Skip out of bounds cells
			if x+i < 0 || x+i >= b.width {
				continue
			}
			if y+j < 0 || y+j >= b.height {
				continue
			}

			if b.GetCell(x+i, y+j) {
				adjacent++
			}
		}
	}
	return adjacent
}

func (b *board) EmptyClone() Board {
	return NewBoard(b.width, b.height)
}

func (b *board) String() string {
	var sb strings.Builder
	for i, cell := range b.cells {
		if cell {
			sb.WriteRune('0')
		} else {
			sb.WriteRune('*')
		}
		if (i+1)%b.width == 0 {
			sb.WriteRune('\n')
		}
	}
	return sb.String()
}

func ParseBoardString(input string) Board {
	rows := strings.Split(input, "\n")
	rowCount := len(rows)

	if rows[rowCount-1] == "" {
		rowCount--
	}

	b := NewBoard(len(rows[0]), rowCount)
	for i := 0; i < rowCount; i++ {
		row := rows[i]
		if row == "" {
			continue
		}
		for j, cell := range row {
			if cell == '0' {
				b.SetCell(j, i, true)
			}
		}
	}
	return b
}
