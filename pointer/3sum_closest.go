package pointer

import (
	"math"
	"sort"
)

// https://leetcode.com/problems/3sum-closest/submissions/
func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var result int
	closest := math.MaxInt

	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1

		for left < right {
			val := nums[i] + nums[left] + nums[right]
			diff := int(math.Abs(float64(val - target)))

			if val == target {
				result = target
				return result
			} else if diff < closest {
				closest = diff
				result = val
			}

			if val > target {
				right--
			}

			if val < target {
				left++
			}
		}
	}

	return result
}
