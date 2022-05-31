package backtrack

func wordSearch(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				res := dfs(board, i, j, 0, word)

				if res {
					return true
				}
			}
		}
	}

	return false
}

func dfs(board [][]byte, i, j, pos int, word string) bool {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] != word[pos] || board[i][j] == '*' {
		return false
	}

	if pos == len(word)-1 {
		return true
	}

	pos++
	temp := board[i][j]
	board[i][j] = '*'

	res := dfs(board, i+1, j, pos, word) || dfs(board, i, j+1, pos, word) || dfs(board, i-1, j, pos, word) || dfs(board, i, j-1, pos, word)
	board[i][j] = temp

	return res
}
