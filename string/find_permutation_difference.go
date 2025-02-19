package string

import "math"

// https://leetcode.com/problems/permutation-difference-between-two-strings/
func findPermutationDifference(s string, t string) int {
	var result int
	valueMap := make(map[rune]int, len(s))
	for idx, val := range s {
		valueMap[val] = idx
	}

	for idx, val := range t {
		result += int(math.Abs(float64(idx - valueMap[val])))
	}

	return result
}
