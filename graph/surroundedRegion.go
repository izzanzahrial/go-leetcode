package graph

func solve(board [][]byte) [][]byte {
	lenR := len(board)
	lenC := len(board[0])

	for r := 0; r < lenR; r++ {
		for c := 0; c < lenC; c++ {
			if r == 0 || c == 0 || r == lenR-1 || c == lenC-1 && board[r][c] == 'O' {
				regionBorder(r, c, lenR, lenC, &board)
			}
		}
	}

	for r := 0; r < lenR; r++ {
		for c := 0; c < lenC; c++ {
			if board[r][c] == 'O' {
				region(r, c, lenR, lenC, &board)
			}
		}
	}

	for r := 0; r < lenR; r++ {
		for c := 0; c < lenC; c++ {
			if board[r][c] == 'B' {
				board[r][c] = 'O'
			}
		}
	}

	return board
}

func regionBorder(r, c, lenR, lenC int, board *[][]byte) {
	if r < 0 || c < 0 || r >= lenR || c >= lenC || (*board)[r][c] != 'O' {
		return
	}

	(*board)[r][c] = 'B'
	regionBorder(r+1, c, lenR, lenC, board)
	regionBorder(r-1, c, lenR, lenC, board)
	regionBorder(r, c+1, lenR, lenC, board)
	regionBorder(r, c-1, lenR, lenC, board)
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
