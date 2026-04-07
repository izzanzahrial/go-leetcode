package stacknqueue

import "strings"

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
		stack = stack[:len(stack)-1]
	}

	if closing < opening {
		stack = append(stack, ")")
		generate(n, opening, closing+1, stack, result)
		stack = stack[:len(stack)-1]
	}
}
