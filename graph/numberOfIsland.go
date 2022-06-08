package graph

func numIslands(grid [][]byte) int {
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
