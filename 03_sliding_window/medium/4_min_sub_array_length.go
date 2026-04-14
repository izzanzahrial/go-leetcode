package slidingwindow

import "math"

// https://leetcode.com/problems/minimum-size-subarray-sum/
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

func minSubArrayLen2(target int, nums []int) int {
	if target == 0 {
		return 0
	}

	minLength, left, currTotal := math.MaxInt, 0, 0
	for right, num := range nums {
		currTotal += num

		for currTotal >= target {
			// right - left + 1 = calculate current length
			minLength = min(minLength, right-left+1)
			currTotal -= nums[left]
			left++
		}
	}

	if minLength == math.MaxInt {
		return 0
	}

	return minLength
}
