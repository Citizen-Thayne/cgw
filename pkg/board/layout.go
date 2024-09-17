package board

const (
	Alive = iota
	Dead
	RowEnd
	BoardEnd
)

func (b *Board) LayoutElements() []int {
	element_count := b.Height*b.Width + b.Height
	elements := make([]int, element_count)

	board_index := 0
	element_index := 0
	for {
		board_value := b.Cells[board_index]
		if board_value {
			elements[element_index] = Alive
		} else {
			elements[element_index] = Dead
		}

		board_index++
		element_index++

		if element_index == element_count-1 {
			elements[element_index] = BoardEnd
			break
		} else if board_index%b.Width == 0 {
			elements[element_index] = RowEnd
			element_index++
		}
	}

	return elements
}
