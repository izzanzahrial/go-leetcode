package stacknqueue

// https://app.codility.com/programmers/lessons/7-stacks_and_queues/brackets/
// TODO: index out of range for false result, and for true, there is one combination bracket left

func brackets(s string) bool {
	var stack []string
	oBracket := map[string]string{
		"}": "{",
		")": "(",
		"]": "[",
	}

	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, string(s[i]))
		} else {
			if oBracket[string(s[i])] == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, string(s[i]))
			}
		}
	}

	return len(stack) == 0
}
