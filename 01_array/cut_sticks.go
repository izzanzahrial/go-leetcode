package array

import "math"

// https://www.hackerrank.com/challenges/cut-the-sticks/problem?isFullScreen=true
func CutTheSticks(sticks []int32) []int32 {
	// Write your code here
	var result []int32
	for len(sticks) > 0 {
		result = append(result, int32(len(sticks)))

		minStick := int32(math.MaxInt32)
		// search the min stick
		for _, length := range sticks {
			smaller := min(minStick, length)
			minStick = smaller
		}

		// cut the stick for the current min stick length
		// if the stick is smaller than 0 dont add the stick into the newSticks
		// because if we remove directly in sticks while sticks being used in the for loop
		// it will generate out of range index
		// e.g. original length of the sticks is 6 while the loop is starting
		// because we remove one of the value from the sticks, the length become 5
		// but the loop still accessing the until length 6 (index 5) meaning out of range
		// because the one of the value already removed and the new length is 5 (max index 4)
		var newSticks []int32
		for _, length := range sticks {
			newLength := length - minStick
			if newLength > 0 {
				newSticks = append(newSticks, newLength)
			}
		}

		// copy the value of the newSticks into sticks
		// because we will be using it for the next loop
		sticks = newSticks
	}

	return result
}
