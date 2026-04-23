package graph

// https://leetcode.com/problems/rotting-oranges
func orangesRotting(grid [][]int) int {
	var fresh int
	orangeQ := make([][]int, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				fresh++
			}

			if grid[i][j] == 2 {
				orangeQ = append(orangeQ, []int{i, j})
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	var result int
	for len(orangeQ) > 0 {
		for _, q := range orangeQ {
			row, col := q[0], q[1]

			// top
			if row-1 >= 0 && grid[row-1][col] == 1 {
				grid[row-1][col] = 2
				orangeQ = append(orangeQ, []int{row - 1, col})
				fresh--
			}

			// right
			if col+1 < len(grid[0]) && grid[row][col+1] == 1 {
				grid[row][col+1] = 2
				orangeQ = append(orangeQ, []int{row, col + 1})
				fresh--
			}

			// bottom
			if row+1 < len(grid) && grid[row+1][col] == 1 {
				grid[row+1][col] = 2
				orangeQ = append(orangeQ, []int{row + 1, col})
				fresh--
			}

			// left
			if col-1 >= 0 && grid[row][col-1] == 1 {
				grid[row][col-1] = 2
				orangeQ = append(orangeQ, []int{row, col - 1})
				fresh--
			}

			orangeQ = orangeQ[1:]
		}

		result++

		if fresh == 0 {
			return result
		}
	}

	return -1
}

func orangesRotting2(grid [][]int) int {
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
