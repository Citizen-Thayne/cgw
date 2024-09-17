package board

import (
	"strings"
)

type Board struct {
	Height int
	Width  int
	Cells  []bool
}

func NewBoard(numCols int, numRows int) *Board {
	board := Board{
		Height: numRows,
		Width:  numCols,
		Cells:  make([]bool, numCols*numRows),
	}

	for i := range board.Cells {
		board.Cells[i] = false
	}
	return &board
}

func (b *Board) GetCell(x, y int) bool {
	return b.Cells[y*b.Width+x]
}

func (b *Board) SetCell(x, y int, value bool) {
	b.Cells[y*b.Width+x] = value
}

func (b *Board) SetCellByIndex(index int, value bool) {
	b.Cells[index] = value
}

func (b *Board) GetCellByIndex(index int, value bool) bool {
	return b.Cells[index]
}

func (b *Board) GetAdjacent(x, y int) int {
	adjacent := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			// Skip the cell itself
			if i == 0 && j == 0 {
				continue
			}
			// Skip out of bounds cells
			if x+i < 0 || x+i >= b.Width {
				continue
			}
			if y+j < 0 || y+j >= b.Height {
				continue
			}

			if b.GetCell(x+i, y+j) {
				adjacent++
			}
		}
	}
	return adjacent
}

func (b *Board) EmptyClone() *Board {
	return NewBoard(b.Width, b.Height)
}

func (b *Board) String() string {
	var sb strings.Builder
	for i, cell := range b.Cells {
		if cell {
			sb.WriteRune('0')
		} else {
			sb.WriteRune('*')
		}
		if (i+1)%b.Width == 0 {
			sb.WriteRune('\n')
		}
	}
	return sb.String()
}

func ParseBoardString(input string) *Board {
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
