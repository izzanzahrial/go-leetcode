package dynamicprogramming

import "fmt"

func targetSum(nums []int, target int) int {
	dp := make(map[string]int)
	var dfs func(i, total int) int

	dfs = func(i, total int) int {
		if i >= len(nums) {
			if total == target {
				return 1
			}
			return 0
		}

		key := fmt.Sprintf("%d, %d", i, total)
		if val, ok := dp[key]; ok {
			return val
		}

		dp[key] = dfs(i+1, total+nums[i]) + dfs(i+1, total-nums[i])

		return dp[key]
	}

	return dfs(0, 0)
}
