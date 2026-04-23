package graph

// https://leetcode.com/problems/max-area-of-island/
// bfs version
func maxAreaOfIsland(grid [][]int) int {
	var maxArea int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				maxArea = max(maxArea, islandArea(i, j, grid))
			}
		}
	}

	return maxArea
}

func islandArea(i, j int, grid [][]int) int {
	var area int

	islandQueue := make([][]int, 0)
	islandQueue = append(islandQueue, []int{i, j})
	grid[i][j] = 0

	for len(islandQueue) > 0 {
		area++
		row, col := islandQueue[0][0], islandQueue[0][1]
		islandQueue = islandQueue[1:]

		// top
		if row-1 >= 0 && grid[row-1][col] == 1 {
			grid[row-1][col] = 0
			islandQueue = append(islandQueue, []int{row - 1, col})
		}

		// right
		if col+1 < len(grid[0]) && grid[row][col+1] == 1 {
			grid[row][col+1] = 0
			islandQueue = append(islandQueue, []int{row, col + 1})
		}

		// bottom
		if row+1 < len(grid) && grid[row+1][col] == 1 {
			grid[row+1][col] = 0
			islandQueue = append(islandQueue, []int{row + 1, col})
		}

		// left
		if col-1 >= 0 && grid[row][col-1] == 1 {
			grid[row][col-1] = 0
			islandQueue = append(islandQueue, []int{row, col - 1})
		}
	}

	return area
}

// dfs version
func maxAreaOfIsland2(grid [][]int) int {
	var result int
	lenR := len(grid)
	lenC := len(grid[0])

	for r := 0; r < lenR; r++ {
		for c := 0; c < lenC; c++ {
			if grid[r][c] == 1 {
				curr := dfsMax(r, c, lenR, lenC, new(int), &grid)

				if curr > result {
					result = curr
				}
			}
		}
	}

	return result
}

func dfsMax(r, c, lenR, lenC int, curr *int, grid *[][]int) int {
	if r < 0 || c < 0 || r >= lenR || c >= lenC || (*grid)[r][c] != 1 {
		return *curr
	}

	*curr += 1
	(*grid)[r][c] = 0
	dfsMax(r+1, c, lenR, lenC, curr, grid)
	dfsMax(r-1, c, lenR, lenC, curr, grid)
	dfsMax(r, c+1, lenR, lenC, curr, grid)
	dfsMax(r, c-1, lenR, lenC, curr, grid)

	return *curr
}
