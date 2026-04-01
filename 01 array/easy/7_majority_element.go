package array

import "math"

// https://leetcode.com/problems/majority-element/description/
func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	numsMap := make(map[int]int)
	for _, num := range nums {
		numsMap[num]++
	}

	var result int
	var currHighest int
	k := len(nums) / 3
	for key, num := range numsMap {
		if num > k && num > currHighest {
			result = key
			currHighest = num
		}
	}

	return result
}

func majorityElement2(nums []int) int {
	currNum := nums[0]
	currResult := 1

	// start at index 1 because we already initialize the first num in currNum
	for i := 1; i < len(nums); i++ {
		if currNum != nums[i] {
			currResult--

			// if the current result equal to 0, we change the currNum
			if currResult == 0 {
				currResult = 1
				currNum = nums[i]
			}
		} else {
			currResult++
		}
	}

	return currNum
}

func majorityElement3(nums []int) int {
	numsMap := make(map[int]int)

	for _, num := range nums {
		numsMap[num] += 1
	}

	majority := math.MinInt
	var majorityNum int
	n := len(nums) / 2
	for key, count := range numsMap {
		if count > n {
			if count > majority {
				majority = count
				majorityNum = key
			}
		}
	}

	return majorityNum
}
