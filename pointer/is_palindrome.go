package pointer

import (
	"regexp"
	"strings"
)

// https://leetcode.com/problems/valid-palindrome/description/
func isPalindrome(s string) bool {
	rgx, _ := regexp.Compile("[^a-zA-Z0-9]+")
	s = rgx.ReplaceAllLiteralString(s, "")
	s = strings.ToLower(s)

	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}
