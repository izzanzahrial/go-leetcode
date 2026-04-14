package array

import (
	"strconv"
	"strings"
)

// https://leetcode.com/problems/maximum-swap/submissions/1548643291/
func maximumSwap(num int) int {
	digits := strings.Split(strconv.Itoa(num), "")
	maxNum := num

	for i, digit := range digits {
		if digit == "9" {
			continue
		}

		currentDigit, _ := strconv.Atoi(digit)
		for j := i + 1; j < len(digits); j++ {
			nextDigit, _ := strconv.Atoi(digits[j])
			if nextDigit > currentDigit {
				digits[i], digits[j] = digits[j], digits[i]
				swappedNum, _ := strconv.Atoi(strings.Join(digits, ""))
				if swappedNum > maxNum {
					maxNum = swappedNum
				}
				digits[i], digits[j] = digits[j], digits[i]
			}
		}
	}

	return maxNum
}
