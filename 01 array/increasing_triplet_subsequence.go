package array

import "math"

// https://leetcode.com/problems/increasing-triplet-subsequence/submissions/?envType=study-plan-v2&envId=leetcode-75
func IncreasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	i, j := math.MaxInt, math.MaxInt

	for _, val := range nums {
		if val <= i {
			i = val
		} else if val <= j {
			j = val
		} else {
			return true
		}
	}

	return false
}
