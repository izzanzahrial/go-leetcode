package array

// https://leetcode.com/problems/special-array-i/description/
func isArraySpecial(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	if len(nums) == 2 {
		if nums[0]%2 != nums[1]%2 {
			return true
		}

		return false
	}

	for i := 1; i < len(nums)-1; i++ {
		if nums[i-1]%2 == nums[i]%2 || nums[i+1]%2 == nums[i]%2 {
			return false
		}
	}

	return true
}
