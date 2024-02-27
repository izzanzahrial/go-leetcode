package mod

import (
	"strconv"
	"strings"
)

// https://leetcode.com/problems/add-digits/description/
func AddDigits(num int) int {
	result := cal(num)
	for result > 9 {
		result = cal(result)
	}

	return result
}

func AddDigits2(num int) int {
	if num == 0 {
		return 0
	}

	if num%9 == 0 {
		return 9
	}

	return num % 9
}

func cal(num int) int {
	strNum := strconv.Itoa(num)
	sliceStr := strings.Split(strNum, "")
	sliceNum := parseSlice(sliceStr)
	return add(sliceNum)
}

func parseSlice(slice []string) []int {
	var result []int

	for _, v := range slice {
		num, _ := strconv.Atoi(string(v))
		result = append(result, num)
	}

	return result
}

func add(nums []int) int {
	var result int
	for _, num := range nums {
		result += num
	}

	return result
}
