package backtrack

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
