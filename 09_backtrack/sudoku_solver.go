package backtrack

func solveSudoku(board [][]byte) {
	solve(board)
}

func solve(board [][]byte) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				// check 1-9
				for num := byte('1'); num <= '9'; num++ {
					if isValid(board, row, col, num) {
						board[row][col] = num
						// dfs into the next number
						if solve(board) {
							return true
						}

						// backtrack if it can solve the current number
						board[row][col] = '.'
					}
				}

				// backtrack here by returning false
				return false
			}
		}
	}

	// done checking every node in board
	return true
}

func isValid(board [][]byte, row, col int, num byte) bool {
	// search diagonaly and horizontaly
	for i := 0; i < 9; i++ {
		if board[row][i] == num || board[i][col] == num {
			return false
		}
	}

	// search within it's box
	startRow, startCol := (row/3)*3, (col/3)*3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == num {
				return false
			}
		}
	}

	return true
}
