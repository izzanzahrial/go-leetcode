package stacknqueque

// https://app.codility.com/programmers/lessons/7-stacks_and_queues/brackets/

func brackets(str string) bool {
	var stack []string
	hash := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}

	for i := 1; i < len(str); i++ {
		if string(stack[len(stack)-1]) == hash[string(str[i])] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, string(str[i]))
		}
	}

	return stack == nil
}
