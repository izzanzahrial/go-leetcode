package dynamicprogramming

import "math"

// https://leetcode.com/problems/edit-distance/description/?envType=study-plan-v2&envId=leetcode-75
func MinDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)

	for i := 0; i < len(word1)+1; i++ {
		dp[i] = make([]int, len(word2)+1)
	}

	for i := 0; i < len(word1)+1; i++ {
		dp[i][0] = i
	}

	for i := 0; i < len(word2)+1; i++ {
		dp[0][i] = i
	}

	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			if word1[i-1] != word2[j-1] {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			} else {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}

	return dp[len(word1)][len(word2)]
}

func min(nums ...int) int {
	result := math.MaxInt

	for _, val := range nums {
		if val < result {
			result = val
		}
	}

	return result
}
