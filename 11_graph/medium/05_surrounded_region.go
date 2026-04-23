package graph

func solve(board [][]byte) [][]byte {
	lenR := len(board)
	lenC := len(board[0])

	// mark all the regions that connected to the border
	// since all the regions that connected to the border cannot be surrounded
	for r := 0; r < lenR; r++ {
		for c := 0; c < lenC; c++ {
			if (r == 0 || c == 0 || r == lenR-1 || c == lenC-1) && board[r][c] == 'O' {
				connectedToBorder(r, c, lenR, lenC, board)
			}
		}
	}

	// turn all the regions that not within the border to 'X'
	for r := 0; r < lenR; r++ {
		for c := 0; c < lenC; c++ {
			if board[r][c] == 'O' {
				region(r, c, lenR, lenC, &board)
			}
		}
	}

	// turn the all border regions back to original value
	for r := 0; r < lenR; r++ {
		for c := 0; c < lenC; c++ {
			if board[r][c] == 'B' {
				board[r][c] = 'O'
			}
		}
	}

	return board
}

func connectedToBorder(r, c, lenR, lenC int, board [][]byte) {
	if r < 0 || c < 0 || r >= lenR || c >= lenC || board[r][c] != 'O' {
		return
	}

	board[r][c] = 'B'
	connectedToBorder(r+1, c, lenR, lenC, board)
	connectedToBorder(r-1, c, lenR, lenC, board)
	connectedToBorder(r, c+1, lenR, lenC, board)
	connectedToBorder(r, c-1, lenR, lenC, board)
}

func region(r, c, lenR, lenC int, board *[][]byte) {
	if r < 0 || c < 0 || r >= lenR || c >= lenC || (*board)[r][c] != 'O' {
		return
	}

	(*board)[r][c] = 'X'
	region(r+1, c, lenR, lenC, board)
	region(r-1, c, lenR, lenC, board)
	region(r, c+1, lenR, lenC, board)
	region(r, c-1, lenR, lenC, board)
}
