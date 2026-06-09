package array

import "strconv"

// https://leetcode.com/problems/find-numbers-with-even-number-of-digits/description/?envType=daily-question&envId=2025-04-30
func findNumbers(nums []int) int {
	var result int

	for _, num := range nums {
		strNum := strconv.Itoa(num)
		if len(strNum)%2 == 0 {
			result++
		}
	}

	return result
}
