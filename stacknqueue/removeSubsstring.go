package stacknqueue

import "strings"

// https://leetcode.com/problems/remove-all-occurrences-of-a-substring/

func removeOccurrences(s string, part string) string {
	if len(s) < len(part) {
		return s
	}

	stack := make([]string, len(s))

	for i := 0; i < len(s); i++ {
		stack = append(stack, string(s[i]))
		if strings.Join(stack[len(stack)-len(part):], "") == part {
			stack = stack[:len(stack)-len(part)]
		}
	}

	return strings.Join(stack, "")
}
