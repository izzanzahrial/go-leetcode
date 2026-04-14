package pointer

import "math"

// https://leetcode.com/problems/container-with-most-water/submissions/?envType=study-plan-v2&envId=leetcode-75
func MaxArea(height []int) int {
	max := 0
	left, right := 0, len(height)-1

	for left < right {
		var count int
		currLen := right - left
		if height[left] < height[right] {
			count = currLen * height[left]
			left += 1
		} else {
			count = currLen * height[right]
			right -= 1
		}

		if max < count {
			max = count
		}
	}

	return max
}

func maxArea(heights []int) int {
	maxWater := math.MinInt
	left, right := 0, len(heights)-1
	for left < right {
		currContainerHeight := min(heights[left], heights[right])
		currWater := (right - left) * currContainerHeight
		maxWater = max(maxWater, currWater)

		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}

	return maxWater
}
