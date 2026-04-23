package graph

// https://leetcode.com/problems/island-perimeter/
func islandPerimeter(grid [][]int) int {
	var perimeter int

	island := make([][]int, 0)
outerloop:
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				grid[i][j] = 2
				island = append(island, []int{i, j})
				break outerloop
			}
		}
	}

	for len(island) > 0 {
		row, col := island[0][0], island[0][1]
		island = island[1:]

		// check top
		if row == 0 {
			perimeter++
		}

		if row-1 >= 0 {
			if grid[row-1][col] == 0 {
				perimeter++
			}

			if grid[row-1][col] == 1 {
				grid[row-1][col] = 2
				island = append(island, []int{row - 1, col})
			}
		}

		// check right
		if col == len(grid[0])-1 {
			perimeter++
		}
		if col+1 < len(grid[0]) {
			if grid[row][col+1] == 0 {
				perimeter++
			}

			if grid[row][col+1] == 1 {
				grid[row][col+1] = 2
				island = append(island, []int{row, col + 1})
			}
		}

		//check bottom
		if row == len(grid)-1 {
			perimeter++
		}
		if row+1 < len(grid) {
			if grid[row+1][col] == 0 {
				perimeter++
			}

			if grid[row+1][col] == 1 {
				grid[row+1][col] = 2
				island = append(island, []int{row + 1, col})
			}
		}

		// check left
		if col == 0 {
			perimeter++
		}
		if col-1 >= 0 {
			if grid[row][col-1] == 0 {
				perimeter++
			}

			if grid[row][col-1] == 1 {
				grid[row][col-1] = 2
				island = append(island, []int{row, col - 1})
			}
		}
	}

	return perimeter
}
