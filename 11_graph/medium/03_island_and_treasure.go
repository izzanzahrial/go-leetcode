package graph

// https://neetcode.io/problems/islands-and-treasure/question
func islandsAndTreasure(grid [][]int) {
	rows, cols := len(grid), len(grid[0])
	queue := make([][2]int, 0)

	// push all treasures (0) into queue first
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 0 {
				queue = append(queue, [2]int{r, c})
			}
		}
	}

	directions := [][2]int{
		{-1, 0}, // up
		{1, 0},  // down
		{0, -1}, // left
		{0, 1},  // right
	}

	// BFS from all treasures simultaneously
	for len(queue) > 0 {
		r, c := queue[0][0], queue[0][1]
		queue = queue[1:]

		for _, d := range directions {
			nr, nc := r+d[0], c+d[1]

			// only process INF cells
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == 2147483647 {
				grid[nr][nc] = grid[r][c] + 1
				queue = append(queue, [2]int{nr, nc})
			}
		}
	}
}
