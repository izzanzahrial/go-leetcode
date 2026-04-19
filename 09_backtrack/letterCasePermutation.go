package backtrack

import "strings"

var numeric = map[string]bool{
	"0": true,
	"1": true,
	"2": true,
	"3": true,
	"4": true,
	"5": true,
	"6": true,
	"7": true,
	"8": true,
	"9": true,
}

// https://leetcode.com/problems/letter-case-permutation/submissions/1209826959/
func LetterCasePermutation(s string) []string {
	news := strings.ToLower(s)
	var result []string
	backtrack3(0, news, &result)
	return result
}

func backtrack3(i int, s string, result *[]string) {
	*result = append(*result, s)

	for j := i; j < len(s); j++ {
		if ok := numeric[string(s[j])]; !ok {
			char := strings.ToUpper(string(s[j]))
			news := s[:j] + char + s[j+1:]
			backtrack3(j+1, news, result)
			char = strings.ToLower(string(s[j]))
			s = s[:j] + char + s[j+1:]
		}
	}
}
