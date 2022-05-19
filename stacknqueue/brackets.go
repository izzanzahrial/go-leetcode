package stacknqueue

// https://app.codility.com/programmers/lessons/7-stacks_and_queues/brackets/
// TODO: index out of range for false result, and for true, there is one combination bracket left

func brackets(str string) bool {
	var stack []string
	hash := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}

	stack = append(stack, string(str[0]))

	for i := 1; i < len(str); i++ {
		if string(stack[len(stack)-1]) != hash[string(str[i])] {
			stack = append(stack, string(str[i]))
		} else {
			stack = append([]string{}, stack[:len(stack)-1]...)
		}
	}

	return stack == nil
}
