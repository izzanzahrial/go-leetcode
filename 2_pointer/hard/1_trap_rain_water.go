package pointer

// https://leetcode.com/problems/trapping-rain-water/description/
func trap(height []int) int {
	var result int
	left, right := 0, len(height)-1
	leftWall, rightWall := height[left], height[right]

	for left < right {
		// search for the smaller wall
		if leftWall < rightWall {
			left++

			// if the new wall is heigher than the previous wall
			if height[left] > leftWall {
				leftWall = height[left]
			} else {
				// add to the result, the left wall subtract by the current wall
				// why we don't need to add the right wall to calculation because the right wall is taller
				// meaning, it will always contain the water
				// now if the leftWall is heigher than the current wall, it will contain some water
				result += leftWall - height[left]
			}
		} else {
			right--

			if height[right] > rightWall {
				rightWall = height[right]
			} else {
				result += rightWall - height[right]
			}
		}
	}

	return result
}

func Trap2(height []int) int {
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
