package backtrack

import "strings"

// https://leetcode.com/problems/generate-parentheses/
func generateParenthesis(n int) []string {
	var stack []string
	var result []string
	generate(n, 0, 0, stack, &result)
	return result
}

func generate(n, opening, closing int, stack []string, result *[]string) {
	if opening == n && closing == n {
		*result = append(*result, strings.Join(stack, ""))
		return
	}

	if opening < n {
		stack = append(stack, "(")
		generate(n, opening+1, closing, stack, result)
		// backtrack
		stack = stack[:len(stack)-1]
	}

	if closing < opening {
		stack = append(stack, ")")
		generate(n, opening, closing+1, stack, result)
		// backtrack
		stack = stack[:len(stack)-1]
	}
}

func generateParenthesis2(n int) []string {
	var result []string

	var generate func(opening, closing int, currStr string)
	generate = func(opening, closing int, currStr string) {
		if len(currStr)/2 == n {
			result = append(result, currStr)
			return
		}

		if opening < n {
			generate(opening+1, closing, currStr+"(")
		}

		if closing < opening {
			generate(opening, closing+1, currStr+")")
		}
	}

	generate(0, 0, "")
	return result
}
