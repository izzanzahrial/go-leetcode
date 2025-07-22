package array

// https://leetcode.com/problems/first-missing-positive/
// https://www.youtube.com/watch?v=8g78yfzMlao
// the idea is to check from 1 to the max "missing" positive number within nums
// how to get the max "missing" positive number within nums
// you can get it by getting the length of the nums
// e.g. if you have nums with the length of 9, meaning the max "missing" positive number is 9
func firstMissingPositive(nums []int) int {
	maxPositive := len(nums)

	numsMap := make(map[int]struct{})
	for _, num := range nums {
		numsMap[num] = struct{}{}
	}

	// result start with 1 since we have to find the "minimum missing" positive number
	result := 1
	for i := 1; i <= maxPositive; i++ {
		// check if the current "minimum missing" positive number is exists
		// if is, go the next number else break
		// e.g. 1 exists in numsMap, check if 2 exists
		if _, exists := numsMap[i]; exists {
			result++
		} else {
			break
		}
	}

	return result
}
