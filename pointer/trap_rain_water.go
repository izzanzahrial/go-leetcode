package pointer

// https://leetcode.com/problems/trapping-rain-water/description/
func Trap(height []int) int {
	var output int
	left, right := 0, len(height)-1
	maxL, maxR := height[left], height[right]

	for left < right {
		if maxL < maxR {
			left++

			if height[left] > maxL {
				maxL = height[left]
			}

			output += maxL - height[left]
		} else {
			right--

			if height[right] > maxR {
				maxR = height[right]
			}

			output += maxR - height[right]
		}
	}

	return output
}
