package dynamicprogramming

// https://leetcode.com/problems/arithmetic-slices/submissions/

func numberOfArithmeticSlices(nums []int) int {
	var result int
	dp := make([]int, len(nums))

	for i := 2; i < len(nums); i++ {
		if nums[i-1]-nums[i-2] == nums[i]-nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
		result += dp[i]
	}

	return result
}
