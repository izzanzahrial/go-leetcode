package main

/**
 * @param rooms: m x n 2D grid
 * @return: nothing
 */
func WallsAndGates(rooms [][]int) [][]int {
	lenR := len(rooms)
	lenC := len(rooms[0])

	for r := 0; r < lenR; r++ {
		for c := 0; c < lenC; c++ {
			if rooms[r][c] == 0 {
				dfs(r, c, lenR, lenC, 0, &rooms)
			}
		}
	}

	return rooms
}

func dfs(r, c, lenR, lenC, val int, rooms *[][]int) {
	if r < 0 || c < 0 || r >= lenR || c >= lenC || (*rooms)[r][c] == -1 {
		return
	}

	val += 1
	if (*rooms)[r][c] == 2147483647 {
		(*rooms)[r][c] = min((*rooms)[r][c], val)
	}
	dfs(r+1, c, lenR, lenC, val, rooms)
	dfs(r-1, c, lenR, lenC, val, rooms)
	dfs(r, c+1, lenR, lenC, val, rooms)
	dfs(r, c-1, lenR, lenC, val, rooms)
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
