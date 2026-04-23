package graph

/**
 * @param rooms: m x n 2D grid
 * @return: nothing
 */

// https://www.lintcode.com/problem/663/
// TODO: still error
func WallsAndGates(rooms [][]int) [][]int {
	lenR := len(rooms)
	lenC := len(rooms[0])

	for r := 0; r < lenR; r++ {
		for c := 0; c < lenC; c++ {
			if rooms[r][c] == 0 {
				dfsWalls(r, c, lenR, lenC, 0, &rooms)
			}
		}
	}

	return rooms
}

func dfsWalls(r, c, lenR, lenC, val int, rooms *[][]int) {
	if r < 0 || c < 0 || r >= lenR || c >= lenC || (*rooms)[r][c] == -1 {
		return
	}

	val += 1
	if (*rooms)[r][c] != 0 {
		(*rooms)[r][c] = min((*rooms)[r][c], val)
	}
	dfsWalls(r+1, c, lenR, lenC, val, rooms)
	dfsWalls(r-1, c, lenR, lenC, val, rooms)
	dfsWalls(r, c+1, lenR, lenC, val, rooms)
	dfsWalls(r, c-1, lenR, lenC, val, rooms)
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func WallsAndGates2(rooms [][]int) [][]int {
	var queue [][]int
	lenR := len(rooms)
	lenC := len(rooms[0])

	for r := 0; r < lenR; r++ {
		for c := 0; c < lenC; c++ {
			if rooms[r][c] == 0 {
				queue = append(queue, []int{r, c})
			}
		}
	}

	val := -1
	for queue != nil {
		r, c := queue[0][0], queue[0][1]
		queue = queue[1:]
		checkNeighbor(r+1, c, lenR, lenC, &queue, &rooms)
		checkNeighbor(r-1, c, lenR, lenC, &queue, &rooms)
		checkNeighbor(r, c+1, lenR, lenC, &queue, &rooms)
		checkNeighbor(r, c-1, lenR, lenC, &queue, &rooms)
		val += 1
		rooms[r][c] = val
	}

	return rooms
}

func checkNeighbor(r, c, lenR, lenC int, queue *[][]int, rooms *[][]int) {
	if r < 0 || c < 0 || r >= lenR || c >= lenR || (*rooms)[r][c] == -1 {
		return
	}

	if (*rooms)[r][c] == 2147483647 {
		(*queue) = append((*queue), []int{r, c})
	}

	return
}
