package pointer

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
