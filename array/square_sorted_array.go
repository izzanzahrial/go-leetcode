package array

import "sort"

// https://leetcode.com/problems/squares-of-a-sorted-array/description/
func sortedSquares(nums []int) []int {
	for i, num := range nums {
		nums[i] = num * num
	}

	sort.Ints(nums)
	return nums
}
