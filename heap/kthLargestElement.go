package heap

import (
	"sort"

	"github.com/izzanzahrial/go-leetcode/datastructure"
)

// https://leetcode.com/problems/kth-largest-element-in-an-array/submissions/

func findKthLargest(nums []int, k int) int {
	heap := datastructure.NewHeap()

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
