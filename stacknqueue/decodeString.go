package stacknqueue

import (
	"strconv"
	"strings"
)

func DecodeString(s string) string {
	var stack []string

	for _, val := range s {
		if val != ']' {
			stack = append(stack, string(val))
		} else {
			var curr string

			for stack[len(stack)-1] != "[" {
				curr = stack[len(stack)-1] + curr
				stack = stack[:len(stack)-1]
			}

			// pop the '[' from the stack
			stack = stack[:len(stack)-1]

			var numStr string
			for len(stack) > 0 && stack[len(stack)-1] >= "0" && stack[len(stack)-1] <= "9" {
				numStr = stack[len(stack)-1] + numStr
				stack = stack[:len(stack)-1]
			}

			num, _ := strconv.Atoi(numStr)

			// repeat the current result to the number of num
			currResult := strings.Repeat(curr, num)
			stack = append(stack, currResult)
		}
	}

	var result string
	for _, s := range stack {
		result = result + string(s)
	}

	return result
}
