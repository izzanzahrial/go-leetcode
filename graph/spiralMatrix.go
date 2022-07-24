package graph

func spiralOrder(matrix [][]int) []int {
	var result []int

	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1

	for {
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}

		top++
		if top > bottom {
			break
		}

		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}

		right--
		if right < left {
			break
		}

		for i := right; i >= left; i-- {
			result = append(result, matrix[bottom][i])
		}

		bottom--
		if bottom < top {
			break
		}

		for i := bottom; i >= top; i-- {
			result = append(result, matrix[i][left])
		}

		left++
		if left > right {
			break
		}
	}

	return result
}
