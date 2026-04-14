package stacknqueue

// https://neetcode.io/problems/validate-parentheses/question
func IsValid(s string) bool {
	var stack []string

	// closing bracket partner
	cBracket := map[string]string{
		"}": "{",
		")": "(",
		"]": "[",
	}

	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, string(s[i]))
		} else {
			if cBracket[string(s[i])] == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, string(s[i]))
			}
		}
	}

	return len(stack) == 0
}

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	var parenthesesStack []rune

	for _, char := range s {
		stackLen := len(parenthesesStack)
		if stackLen == 0 {
			parenthesesStack = append(parenthesesStack, char)
			continue
		}

		if char == '}' && parenthesesStack[stackLen-1] == '{' {
			parenthesesStack = parenthesesStack[:stackLen-1]
		} else if char == ']' && parenthesesStack[stackLen-1] == '[' {
			parenthesesStack = parenthesesStack[:stackLen-1]
		} else if char == ')' && parenthesesStack[stackLen-1] == '(' {
			parenthesesStack = parenthesesStack[:stackLen-1]
		} else {
			parenthesesStack = append(parenthesesStack, char)
		}
	}

	return len(parenthesesStack) == 0
}
