package dynamicprogramming

// https://leetcode.com/problems/minimum-path-sum/
func MinPathSum(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i-1 < 0 && j-1 < 0 {
				continue
			}

			if i-1 < 0 {
				grid[i][j] = grid[i][j] + grid[i][j-1]
				continue
			}

			if j-1 < 0 {
				grid[i][j] = grid[i][j] + grid[i-1][j]
				continue
			}

			grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}
