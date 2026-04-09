package binarysearch

// https://neetcode.io/problems/sqrtx/question
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	left, right := 0, x
	var result int
	for left <= right {
		mid := left + ((right - left) / 2)
		currNum := mid * mid

		if currNum == x {
			return mid
		}

		if currNum > x {
			right = mid - 1
		}

		if currNum < x {
			left = mid + 1
			// since it's lower and the result could be rounded down
			// that's why we store the current mid value as a result
			// if it's rounded up, we store the result in right index
			result = mid
		}
	}

	return result
}
