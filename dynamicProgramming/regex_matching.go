package dynamicprogramming

// https://leetcode.com/problems/regular-expression-matching/description/?envType=problem-list-v2&envId=dynamic-programming&
func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)

	for i, _ := range dp {
		dp[i] = make([]bool, len(p)+1)
	}

	for i, _ := range dp {
		dp[i][0] = true
	}

	for i, _ := range dp[0] {
		dp[0][i] = true
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if p[j-1] == '*' || p[j-1] == '.' || s[i-1] == p[i-1] {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			} else {
				dp[i][j] = false
			}
		}
	}

	return dp[len(dp)][len(dp[0])]
}
