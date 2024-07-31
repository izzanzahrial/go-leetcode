package array

import (
	"math"
	"strings"
)

// https://leetcode.com/problems/longest-common-prefix/description/
func longestCommonPrefix(strs []string) string {
	var result strings.Builder
	minLen := math.MaxInt

	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
		}
	}

	for i := 0; i < minLen; i++ {
		char := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != char {
				return result.String()
			}
		}
		result.WriteByte(char)
	}

	return result.String()
}
