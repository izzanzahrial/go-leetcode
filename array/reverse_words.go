package main

import (
	"strings"
)

// https://leetcode.com/problems/reverse-words-in-a-string/?envType=study-plan-v2&envId=leetcode-75
func ReverseWords(s string) string {
	// Split the input string into words
	words := strings.Fields(s)

	low, high := 0, len(words)-1
	for low < high {
		words[low], words[high] = words[high], words[low]
		low++
		high--
	}

	return strings.TrimSpace(strings.Join(words, " "))
}
