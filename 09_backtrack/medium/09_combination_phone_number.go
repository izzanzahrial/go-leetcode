package backtrack

import "strings"

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
// dfs
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	digitMap := map[byte][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}

	var result []string
	var curr string
	var dfs func(curr string, i int)

	dfs = func(curr string, i int) {
		if len(curr) == len(digits) {
			result = append(result, curr)
			return
		}

		for _, str := range digitMap[digits[i]] {
			dfs(curr+str, i+1)
		}
	}

	dfs(curr, 0)
	return result
}

// version 2 backtrack
func letterCombinations2(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	var result []string

	digitsMap := map[byte][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}

	var recursive func(idx int, curr []string)
	recursive = func(idx int, curr []string) {
		if len(curr) == len(digits) {
			result = append(result, strings.Join(curr, ""))
			return
		}

		for i := 0; i < len(digitsMap[digits[idx]]); i++ {
			curr = append(curr, digitsMap[digits[idx]][i])
			recursive(idx+1, curr)
			// backtrack
			curr = curr[:len(curr)-1]
		}
	}

	recursive(0, []string{})
	return result
}
