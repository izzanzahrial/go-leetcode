package dynamicprogramming

// https://leetcode.com/problems/min-cost-climbing-stairs/description/?envType=study-plan-v2&envId=leetcode-75
func MinCostClimbingStairs(cost []int) int {
	var dp []int
	dp = append(dp, cost[0])
	dp = append(dp, cost[1])

	for i := 2; i < len(cost); i++ {
		cost1 := cost[i] + dp[i-1]
		cost2 := cost[i] + dp[i-2]
		if cost1 < cost2 {
			dp = append(dp, cost1)
		} else {
			dp = append(dp, cost2)
		}
	}

	if dp[len(cost)-1] < dp[len(cost)-2] {
		return dp[len(cost)-1]
	}

	return dp[len(cost)-2]
}
