package array

// https://leetcode.com/problems/first-missing-positive/
// https://www.youtube.com/watch?v=8g78yfzMlao
func firstMissingPositive(nums []int) int {
	// Since we only searching for the first missing positive number
	// the maximum positive number is length of the nums+1
	maxPositive := len(nums) + 1

	// Create hash map for nums
	numsMap := make(map[int]struct{})
	for _, num := range nums {
		numsMap[num] = struct{}{}
	}

	result := 1
	// We only have to check from 1 to maxPositive
	for i := 1; i <= maxPositive; i++ {
		// Since we looking through hash map, the time complexity is O(1)
		if _, exists := numsMap[i]; exists {
			result++
		} else {
			return i
		}
	}

	return result
}
