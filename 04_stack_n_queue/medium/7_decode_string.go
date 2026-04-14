package stacknqueue

import (
	"strconv"
	"strings"
)

// https://neetcode.io/problems/decode-string/question
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

func decodeString2(s string) string {
	var stack []rune
	for _, char := range s {
		if char != ']' {
			stack = append(stack, char)
		} else {
			var currChar []rune
			for stack[len(stack)-1] != '[' {
				// revese the current char stack
				currChar = append([]rune{stack[len(stack)-1]}, currChar...)
				// pop the value
				stack = stack[:len(stack)-1]
			}

			// pop the '['
			stack = stack[:len(stack)-1]

			var strRepetition string
			for len(stack) > 0 && stack[len(stack)-1] >= '0' && stack[len(stack)-1] <= '9' {
				// reverse the num of strRepetition
				strRepetition = string(stack[len(stack)-1]) + strRepetition
				// pop the value
				stack = stack[:len(stack)-1]
			}

			intRepetition, _ := strconv.Atoi(strRepetition)
			for i := 0; i < intRepetition; i++ {
				for _, c := range currChar {
					stack = append(stack, c)
				}
			}
		}
	}

	return string(stack)
}
