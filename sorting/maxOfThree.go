package sorting

import "sort"

// https://app.codility.com/programmers/lessons/6-sorting/max_product_of_three/

func maxOfThree(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)-1] * nums[len(nums)-2] * nums[len(nums)-3]
}
