package pointer

import "math"

// https://leetcode.com/problems/sliding-window-maximum/
// TIME LIMIT EXCEEDED
func maxSlidingWindow(nums []int, k int) []int {
	var result []int

	for i := 0; i < len(nums)-k+1; i++ {
		currMax := math.MinInt
		for j := i; j < i+k; j++ {
			if nums[j] > currMax {
				currMax = nums[j]
			}
		}

		result = append(result, currMax)
	}

	return result
}
