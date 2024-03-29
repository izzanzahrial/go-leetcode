package hash

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/?envType=study-plan-v2&envId=leetcode-75
func LetterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	digitMap := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	var str string
	var result []string
	var dfs func(str string, idx int)

	dfs = func(str string, idx int) {
		if len(str) == len(digits) {
			result = append(result, str)
			return
		}

		for _, val := range digitMap[digits[idx]] {
			dfs(str+val, idx+1)
		}
	}
	dfs(str, 0)

	return result
}
