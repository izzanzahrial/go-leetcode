package backtrack

// https://leetcode.com/problems/word-search/
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

// version 2
func wordSearch2(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				temp := board[i][j]
				board[i][j] = '.'
				if search(board, i, j, word[1:]) {
					return true
				}
				board[i][j] = temp
			}
		}
	}

	return false
}

func search(board [][]byte, i, j int, word string) bool {
	if len(word) == 0 {
		return true
	}

	// check top
	if i-1 >= 0 && board[i-1][j] == word[0] {
		temp := board[i-1][j]
		board[i-1][j] = '.'
		if search(board, i-1, j, word[1:]) {
			return true
		}
		board[i-1][j] = temp
	}

	// check right
	if j+1 < len(board[0]) && board[i][j+1] == word[0] {
		temp := board[i][j+1]
		board[i][j+1] = '.'
		if search(board, i, j+1, word[1:]) {
			return true
		}
		board[i][j+1] = temp
	}

	// check bottom
	if i+1 < len(board) && board[i+1][j] == word[0] {
		temp := board[i+1][j]
		board[i+1][j] = '.'
		if search(board, i+1, j, word[1:]) {
			return true
		}
		board[i+1][j] = temp
	}

	// check left
	if j-1 >= 0 && board[i][j-1] == word[0] {
		temp := board[i][j-1]
		board[i][j-1] = '.'
		if search(board, i, j-1, word[1:]) {
			return true
		}
		board[i][j-1] = temp
	}

	return false
}
