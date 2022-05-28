package dynamicprogramming

func uniquePaths(m int, n int) int {
	paths := make([][]int, m)
	for i := range paths {
		paths[i] = make([]int, n)
	}

	for r := 0; r < len(paths); r++ {
		for c := 0; c < len(paths[0]); c++ {
			if r == 0 && c == 0 {
				paths[r][c] = 1
			} else if r == 0 {
				paths[r][c] = paths[r][c-1]
			} else if c == 0 {
				paths[r][c] = paths[r-1][c]
			} else {
				paths[r][c] = paths[r-1][c] + paths[r][c-1]
			}
		}
	}

	return paths[m-1][n-1]
}
