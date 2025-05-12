package slidingwindow

import "math"

func minSubArrayLen(target int, nums []int) int {
	left, total := 0, 0
	result := math.MaxInt

	for right, num := range nums {
		total += num
		for total >= target {
			if result > right-left+1 {
				result = right - left + 1
			}
			total -= nums[left]
			left++
		}
	}

	if result == math.MaxInt {
		return 0
	}

	return result
}
