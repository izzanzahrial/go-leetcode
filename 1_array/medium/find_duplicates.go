package array

import "sort"

// https://leetcode.com/problems/find-the-duplicate-number/
func findDuplicate(nums []int) int {
	numsMap := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		if _, ok := numsMap[num]; ok {
			return num
		}
		numsMap[num] = struct{}{}
	}

	return -1
}

func findDuplicate2(nums []int) int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return nums[i]
		}
	}

	return -1
}
