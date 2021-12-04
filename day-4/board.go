package main

type number struct {
	value  int64
	marked bool
}

type board struct {
	rows [][]*number
}

func (b *board) markNumber(n int64) {
	for _, row := range b.rows {
		for _, col := range row {
			if col.value == n {
				col.marked = true
			}
		}
	}
}

func (b *board) hasWon() bool {
	return b.hasCompleteRow() || b.hasCompleteCol()
}

func (b *board) hasCompleteRow() bool {
	for _, row := range b.rows {
		allMarked := true
		for _, col := range row {
			allMarked = allMarked && col.marked
		}
		if allMarked {
			return true
		}
	}
	return false
}

func (b *board) hasCompleteCol() bool {
	rows := len(b.rows)
	cols := len(b.rows[0])

	for col := 0; col < cols; col++ {
		allMarked := true
		for row := 0; row < rows; row++ {
			allMarked = allMarked && b.rows[row][col].marked
		}
		if allMarked {
			return true
		}
	}
	return false
}

func (b *board) sumOfUnmarked() int64 {
	var sum int64 = 0
	for _, row := range b.rows {
		for _, col := range row {
			if !col.marked {
				sum += col.value
			}
		}
	}
	return sum
}

func (b *board) reset() {
	for _, row := range b.rows {
		for _, col := range row {
			col.marked = false
		}
	}
}
