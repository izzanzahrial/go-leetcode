package dynamicprogramming

// https://leetcode.com/problems/house-robber/description/?envType=study-plan-v2&envId=leetcode-75
func Rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = maxInt(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = maxInt(dp[i-1], nums[i]+dp[i-2])
	}

	return maxInt(dp[len(nums)-1], dp[len(nums)-2])
}

func maxInt(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
