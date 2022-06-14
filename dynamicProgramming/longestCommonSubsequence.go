package dynamicprogramming

import "math"

// https://leetcode.com/problems/longest-common-subsequence/
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := len(dp) - 2; i >= 0; i-- {
		for j := len(dp[0]) - 2; j >= 0; j-- {
			if text1[i] == text2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				val1 := float64(dp[i][j+1])
				val2 := float64(dp[i+1][j])
				dp[i][j] = int(math.Max(val1, val2))
			}
		}
	}

	return dp[0][0]
}
