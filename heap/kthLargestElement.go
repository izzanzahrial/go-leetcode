package heap

import "sort"

// https://leetcode.com/problems/kth-largest-element-in-an-array/submissions/

func findKthLargest(nums []int, k int) int {
	heap := &heap{}

	for _, num := range nums {
		heap.add(num)
	}

	for i := 0; i < k-1; i++ {
		heap.pop()
	}

	val := heap.pop()

	return val
}

func findKthLargest2(nums []int, k int) int {
	sort.Ints(nums)

	return nums[len(nums)-k]
}
