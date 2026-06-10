package array

// https://leetcode.com/problems/is-subsequence/description/
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	node := 0
	for _, char := range t {
		if string(char) == string(s[node]) {
			node++
		}

		if node == len(s) {
			return true
		}
	}

	return false
}
