package array

import "math"

// https://leetcode.com/problems/jump-game-ii/description/
func Jump(nums []int) int {
	min := math.MaxInt64

	dfs(0, 0, len(nums)-1, &min, nums)

	return min
}

func dfs(idx, curr, length int, min *int, nums []int) {
	if idx == length {
		if curr < *min {
			*min = curr
		}
		return
	}

	val := nums[idx]

	for i := idx + 1; i <= idx+val && i <= length; i++ {
		dfs(i, curr+1, length, min, nums)
	}
}
