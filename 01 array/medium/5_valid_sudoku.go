package array

// https://leetcode.com/problems/valid-sudoku/description/
func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			val := board[r][c]
			if val == '.' {
				continue
			}

			// Check row
			if rows[r][val] {
				return false
			}
			rows[r][val] = true

			// Check column
			if cols[c][val] {
				return false
			}
			cols[c][val] = true

			// Check 3x3 box
			boxIndex := (r/3)*3 + c/3
			if boxes[boxIndex][val] {
				return false
			}
			boxes[boxIndex][val] = true
		}
	}

	return true
}
