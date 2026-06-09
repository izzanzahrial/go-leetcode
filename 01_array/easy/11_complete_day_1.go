package array

import "math"

// https://leetcode.com/problems/count-pairs-that-form-a-complete-day-i/description/
func countCompleteDayPairs(hours []int) int {
	var result int
	for i, hour1 := range hours {
		for j := i + 1; j < len(hours); j++ {
			if isCompleteDay(hour1, hours[j]) {
				result++
			}
		}
	}

	return result
}

func countCompleteDayPairs2(hours []int) int {
	var result int
	for i, hour1 := range hours {
		for j := i + 1; j < len(hours); j++ {
			if (hour1+hours[j])%24 == 0 {
				result++
			}
		}
	}

	return result
}

func isCompleteDay(i, j int) bool {
	total := (float64(i) + float64(j)) / 24
	return total == math.Floor(total)
}
