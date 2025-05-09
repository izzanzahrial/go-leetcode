package array

// https://neetcode.io/problems/concatenation-of-array
// https://leetcode.com/problems/concatenation-of-array/
func GetConcatenationArray(nums []int) []int {
	return append(nums, nums...)
}

func GetConcatenationArray2(nums []int) []int {
	for _, num := range nums {
		nums = append(nums, num)
	}

	return nums
}
