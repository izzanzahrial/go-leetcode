package array

import "math"

// https://leetcode.com/problems/contains-duplicate-ii/description/
func containsNearbyDuplicate(nums []int, k int) bool {
	duplicateMap := make(map[int]int)

	for idx, num := range nums {
		duplicateIdx, exists := duplicateMap[num]
		if !exists {
			duplicateMap[num] = idx
		} else {
			if int(math.Abs(float64(idx-duplicateIdx))) <= k {
				return true
			} else {
				// already exists in the map
				// and it's out of k, update the index
				// we don't have to compare because the older values are out of k
				duplicateMap[num] = idx
			}
		}
	}

	return false
}
