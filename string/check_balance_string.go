package string

import "strconv"

// https://leetcode.com/problems/check-balanced-string/
func isBalanced(num string) bool {
	var evenIndices, oddIndices int
	for i, val := range num {
		valInt, _ := strconv.Atoi(string(val))
		if i%2 == 0 {
			evenIndices += valInt
		} else {
			oddIndices += valInt
		}
	}

	if evenIndices == oddIndices {
		return true
	}

	return false
}
