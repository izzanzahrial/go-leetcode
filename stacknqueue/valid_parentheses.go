package stacknqueue

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
