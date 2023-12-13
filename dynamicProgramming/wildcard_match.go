package dynamicprogramming

// https://leetcode.com/problems/wildcard-matching/description/
func IsMatch(s string, p string) bool {
	m := len(s)
	n := len(p)

	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true

	for j := 0; j < n; j++ {
		if p[j] == '*' {
			dp[0][j+1] = dp[0][j]
		} else {
			break
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if s[i] == p[j] || p[j] == '?' {
				dp[i+1][j+1] = dp[i][j]
			} else if p[j] == '*' {
				dp[i+1][j+1] = dp[i][j] || dp[i][j+1] || dp[i+1][j]
			}
		}
	}

	return dp[m][n]
}
