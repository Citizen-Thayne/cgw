package board

import (
	"strings"
)

type Board struct {
	rows [][]bool
}

func NewBoard(numCols int, numRows int) *Board {
	board := Board{
		rows: make([][]bool, numRows),
	}
	for i := 0; i < numRows; i++ {
		row := make([]bool, numCols)
		for j := range row {
			row[j] = false
		}
		board.rows[i] = row
	}

	return &board
}

func (b *Board) GetCell(x, y int) bool {
	return b.rows[y][x]
}

func (b *Board) SetCell(x, y int, value bool) {
	b.rows[y][x] = value
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
			if x+i < 0 || x+i >= len(b.rows[0]) {
				continue
			}
			if y+j < 0 || y+j >= len(b.rows) {
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
	return NewBoard(len(b.rows[0]), len(b.rows))
}

func (b *Board) String() string {
	var sb strings.Builder
	for _, row := range b.rows {
		for _, cell := range row {
			if cell {
				sb.WriteRune('0')
			} else {
				sb.WriteRune('*')
			}
		}
		sb.WriteRune('\n')
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
