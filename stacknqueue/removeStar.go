package stacknqueue

import "strings"

// https://leetcode.com/problems/removing-stars-from-a-string/description/?envType=study-plan-v2&envId=leetcode-75
func RemoveStars(s string) string {
	var stack []string

	for _, val := range s {
		if string(val) == "*" {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, string(val))
		}
	}

	return strings.Join(stack, "")
}
