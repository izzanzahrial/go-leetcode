package graph

func maxAareaOfIsland(grid [][]int) int {
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
