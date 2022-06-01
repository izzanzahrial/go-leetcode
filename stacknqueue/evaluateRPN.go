package stacknqueue

import "strconv"

func evalRPN(tokens []string) int {
	var result int
	var stack []string
	operators := map[string]string{
		"+": "+",
		"-": "-",
		"*": "*",
		"/": "/",
	}

	for i := 0; i < len(tokens); i++ {
		if len(stack) == 0 {
			stack = append(stack, tokens[i])
		} else {
			if _, ok := operators[string(tokens[i])]; ok {
				var val int
				int1, _ := strconv.Atoi(stack[len(stack)-1])
				int2, _ := strconv.Atoi(stack[len(stack)-2])

				switch string(tokens[i]) {
				case "+":
					val = int2 + int1
				case "-":
					val = int2 - int1
				case "*":
					val = int2 * int1
				case "/":
					val = int2 / int1
				}
				stack = append(stack[:len(stack)-2], strconv.Itoa(val))
			} else {
				stack = append(stack, tokens[i])
			}
		}
	}

	result, _ = strconv.Atoi(stack[0])

	return result
}
