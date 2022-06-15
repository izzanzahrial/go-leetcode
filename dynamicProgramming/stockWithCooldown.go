package dynamicprogramming

import (
	"fmt"
)

// still miss by one idk
func maxProfit(prices []int) int {
	dp := make(map[string]int)
	var dfs func(i int, state bool) int

	dfs = func(i int, state bool) int {
		if i >= len(prices) {
			return 0
		}

		key := fmt.Sprintf("%d, %t", i, state)
		if val, ok := dp[key]; ok {
			return val
		}

		if state {
			buy := dfs(i+1, false) - prices[i]
			cooldown := dfs(i+1, true)
			dp[key] = max(buy, cooldown)
		} else {
			sell := dfs(i+2, false) + prices[i]
			cooldown := dfs(i+1, true)
			dp[key] = max(sell, cooldown)
		}

		return dp[key]
	}

	return dfs(0, true)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
