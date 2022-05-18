package sorting

import "sort"

// https://app.codility.com/programmers/lessons/6-sorting/distinct/

func distinct(nums []int) int {
	var total int
	sort.Ints(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			total += 1
		}
	}

	return total
}

func distinct2(nums []int) int {
	var total int
	hash := make(map[int]int)

	for _, val := range nums {
		hash[val] += 1
	}

	for _, val := range hash {
		if val == 1 {
			total += 1
		}
	}

	return total
}
