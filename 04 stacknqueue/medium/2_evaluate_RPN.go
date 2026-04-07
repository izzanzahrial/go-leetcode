package stacknqueue

import "strconv"

// https://leetcode.com/problems/evaluate-reverse-polish-notation/
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

func evalRPN2(tokens []string) int {
	var stack []int

	for _, char := range tokens {
		switch char {
		case "+":
			num1 := stack[len(stack)-2]
			num2 := stack[len(stack)-1]
			addResult := num1 + num2
			stack = append(stack[:len(stack)-2], addResult)
		case "-":
			num1 := stack[len(stack)-2]
			num2 := stack[len(stack)-1]
			subResult := num1 - num2
			stack = append(stack[:len(stack)-2], subResult)
		case "*":
			num1 := stack[len(stack)-2]
			num2 := stack[len(stack)-1]
			timesResult := num1 * num2
			stack = append(stack[:len(stack)-2], timesResult)
		case "/":
			num1 := stack[len(stack)-2]
			num2 := stack[len(stack)-1]
			divResult := num1 / num2
			stack = append(stack[:len(stack)-2], divResult)
		default:
			num, _ := strconv.Atoi(char)
			stack = append(stack, num)
		}
	}

	return stack[0]
}
