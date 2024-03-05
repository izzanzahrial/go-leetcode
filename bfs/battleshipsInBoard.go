package bfs

// https://leetcode.com/problems/battleships-in-a-board/description/
func CountBattleships(board [][]byte) int {
	var result int
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'X' {
				result++
				changeBattleShip(i, j, &board)
			}
		}
	}

	return result
}

func changeBattleShip(i, j int, board *[][]byte) {
	// Check top
	if i > 0 && (*board)[i-1][j] == 'X' {
		(*board)[i-1][j] = '.'
		changeBattleShip(i-1, j, board)
	}

	// Check right
	if j < len((*board)[0])-1 && (*board)[i][j+1] == 'X' {
		(*board)[i][j+1] = '.'
		changeBattleShip(i, j+1, board)
	}

	// Check bottom
	if i < len(*board)-1 && (*board)[i+1][j] == 'X' {
		(*board)[i+1][j] = '.'
		changeBattleShip(i+1, j, board)
	}

	// Check left
	if j > 0 && (*board)[i][j-1] == 'X' {
		(*board)[i][j-1] = '.'
		changeBattleShip(i, j-1, board)
	}
}
