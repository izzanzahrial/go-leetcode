package array

// https://leetcode.com/problems/build-array-from-permutation/
func buildArray(nums []int) []int {
	var result []int

	for _, num := range nums {
		result = append(result, nums[num])
	}

	return result
}
