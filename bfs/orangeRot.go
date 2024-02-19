package bfs

// https://leetcode.com/problems/rotting-oranges

func OrangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var rotten [][2]int

	fresh := 0
	// check all fresh and rotten oranges
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// check fresh
			if grid[i][j] == 1 {
				fresh++
			}

			// check rotten
			if grid[i][j] == 2 {
				rotten = append(rotten, [2]int{i, j})
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	minutes := 0
	for len(rotten) > 0 {
		for _, rot := range rotten {
			// check top
			x := rot[0] - 1
			y := rot[1]
			if x >= 0 && grid[x][y] == 1 {
				grid[x][y] = 2
				rotten = append(rotten, [2]int{x, y})
				fresh--
			}

			// check right
			x = rot[0]
			y = rot[1] + 1
			if y < n && grid[x][y] == 1 {
				grid[x][y] = 2
				rotten = append(rotten, [2]int{x, y})
				fresh--
			}

			// check bottom
			x = rot[0] + 1
			y = rot[1]
			if x < m && grid[x][y] == 1 {
				grid[x][y] = 2
				rotten = append(rotten, [2]int{x, y})
				fresh--
			}

			// check left
			x = rot[0]
			y = rot[1] - 1
			if y >= 0 && grid[x][y] == 1 {
				grid[x][y] = 2
				rotten = append(rotten, [2]int{x, y})
				fresh--
			}

			rotten = rotten[1:]
		}

		minutes++
		if fresh == 0 {
			return minutes
		}
	}

	return -1
}
