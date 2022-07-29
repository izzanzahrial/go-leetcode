package sorting

import (
	"sort"
)

// https://leetcode.com/problems/3sum/submissions/

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)

	for i, num := range nums {
		if i > 0 && num == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1
		for left < right {
			threeSum := num + nums[left] + nums[right]
			if threeSum > 0 {
				right--
			} else if threeSum < 0 {
				left++
			} else {
				result = append(result, []int{num, nums[left], nums[right]})
				left++
				for nums[left] == nums[left-1] && left < right {
					left++
				}
			}

		}
	}

	return result
}
