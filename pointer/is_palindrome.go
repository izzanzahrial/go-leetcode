package pointer

import (
	"regexp"
	"strings"
)

// https://leetcode.com/problems/valid-palindrome/description/
func isPalindrome(s string) bool {
	regex, _ := regexp.Compile("[^a-zA-Z0-9]+")
	s = regex.ReplaceAllString(s, "")
	s = strings.ToLower(s)

	for i, j := 0, len(s)-1; i < j; i++ {
		if s[i] != s[j] {
			return false
		}
		j--
	}

	return true
}
