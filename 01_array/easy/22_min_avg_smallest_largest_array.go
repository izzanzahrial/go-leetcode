package array

import (
	"math"
	"sort"
)

// https://leetcode.com/problems/minimum-average-of-smallest-and-largest-elements/description/
func minimumAverage(nums []int) float64 {
	sort.Ints(nums)
	var avg []float64
	left, right := 0, len(nums)-1
	for left < right {
		avgNum := (float64(nums[left]) + float64(nums[right])) / 2
		avg = append(avg, avgNum)
		left++
		right--
	}

	min := math.MaxFloat64
	var result float64
	for _, num := range avg {
		if num < min {
			min = num
			result = num
		}
	}

	return result
}
