package binarysearch

func searchMatrix(matrix [][]int, target int) bool {
	top := 0
	bottom := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1
	var middle int
	for top < bottom {
		// if this part wrong, could make this function error out of bounds by 1
		mid := top + ((bottom - top) / 2)
		if target < matrix[mid][left] {
			// if this part wrong, could make this function error out of bounds by 1
			bottom = mid
			middle = bottom
		}
		if target > matrix[mid][right] {
			top = mid + 1
			middle = top
		}
		if target >= matrix[mid][left] && target <= matrix[mid][right] {
			middle = mid
			break
		}
	}

	for left <= right {
		mid := left + ((right - left) / 2)
		if target < matrix[middle][mid] {
			right = mid - 1
		}
		if target > matrix[middle][mid] {
			left = mid + 1
		}
		if target == matrix[middle][mid] {
			return true
		}
	}

	return false
}
