package stacknqueue

import "strconv"

// https://leetcode.com/problems/baseball-game/
func calPoints(operations []string) int {
	var stack []int

	for _, op := range operations {
		switch op {
		case "C":
			stack = stack[:len(stack)-1]
		case "D":
			stack = append(stack, stack[len(stack)-1]*2)
		case "+":
			prevTwo := stack[len(stack)-1] + stack[len(stack)-2]
			stack = append(stack, prevTwo)
		default:
			numOP, _ := strconv.Atoi(op)
			stack = append(stack, numOP)
		}
	}

	var result int
	for _, num := range stack {
		result += num
	}

	return result
}
