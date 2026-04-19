package backtrack

func totalNQueens(n int) int {
	var result int
	board := makeBoard2(n)

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			result++
			return
		}

		for col := 0; col < n; col++ {
			if isValid2(board, row, col) {
				board[row][col] = "Q"
				backtrack(row + 1)
				board[row][col] = "."
			}
		}
	}

	backtrack(0)
	return result
}

func makeBoard2(n int) [][]string {
	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}
	return board
}

func isValid2(board [][]string, row, col int) bool {
	// column
	for i := 0; i < row; i++ {
		if board[i][col] == "Q" {
			return false
		}
	}

	// diag left
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
			return false
		}
	}

	// diag right
	for i, j := row-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}

	return true
}
