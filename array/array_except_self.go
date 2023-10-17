package main

// https://leetcode.com/problems/product-of-array-except-self/description/?envType=study-plan-v2&envId=leetcode-75
// https://www.youtube.com/watch?v=bNvIQI2wAjk&ab_channel=NeetCode
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	// count for prefix
	prefix := 1
	for idx := range nums {
		result[idx] = prefix
		prefix *= nums[idx]
	}

	// count for postfix
	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= postfix
		postfix *= nums[i]
	}

	return result
}
