package matrix

// https://leetcode.com/problems/matrix-diagonal-sum/
func diagonalSum(mat [][]int) int {
	var total int
	n, m := len(mat), len(mat[0])

	// diagonal top left into bottom right
	for i, j := 0, 0; i < n && j < m; i, j = i+1, j+1 {
		if mat[i][j] == 0 {
			continue
		}

		total += mat[i][j]
		mat[i][j] = 0
	}

	// diagonal top right into bottom left
	for i, j := 0, m-1; i < n && j >= 0; i, j = i+1, j-1 {
		if mat[i][j] == 0 {
			continue
		}

		total += mat[i][j]
		mat[i][j] = 0
	}

	return total
}
