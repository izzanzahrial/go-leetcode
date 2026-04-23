package graph

func numIslands(grid [][]byte) int {
	var islands int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				islands++
				unmarkIsland(i, j, grid)
			}
		}
	}

	return islands
}

func unmarkIsland(row, col int, grid [][]byte) {
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) {
		return
	}

	if grid[row][col] == '1' {
		grid[row][col] = '0'

		// dfs, unmark all direction island teritory
		// top
		unmarkIsland(row-1, col, grid)
		// right
		unmarkIsland(row, col+1, grid)
		// bottom
		unmarkIsland(row+1, col, grid)
		// left
		unmarkIsland(row, col-1, grid)
	}
}

func numIslands2(grid [][]byte) int {
	var result int

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == '1' {
				dfs(&grid, r, c)
				result += 1
			}
		}
	}

	return result
}

func dfs(grid *[][]byte, r, c int) {
	lenR := len(*grid)
	lenC := len((*grid)[0]) // if you use pointer you need to use bracket, i just know this

	if r < 0 || c < 0 || r >= lenR || c >= lenC || (*grid)[r][c] != '1' {
		return
	}

	(*grid)[r][c] = '*'

	dfs(grid, r+1, c)
	dfs(grid, r-1, c)
	dfs(grid, r, c+1)
	dfs(grid, r, c-1)
}
