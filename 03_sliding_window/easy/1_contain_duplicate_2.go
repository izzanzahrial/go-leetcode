package slidingwindow

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

func containsNearbyDuplicate2(nums []int, k int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				if int(math.Abs(float64(i-j))) <= k {
					return true
				}
			}
		}
	}

	return false
}
