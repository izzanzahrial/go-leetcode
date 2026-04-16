package heap

import (
	"sort"

	heap "github.com/izzanzahrial/go-leetcode/08_heap"
)

// https://leetcode.com/problems/kth-largest-element-in-an-array/submissions/

func findKthLargest(nums []int, k int) int {
	heap := heap.NewHeap()

	for _, num := range nums {
		heap.Add(num)
	}

	for i := 0; i < k-1; i++ {
		heap.Pop()
	}

	val := heap.Pop()

	return val
}

func findKthLargest2(nums []int, k int) int {
	sort.Ints(nums)

	return nums[len(nums)-k]
}
