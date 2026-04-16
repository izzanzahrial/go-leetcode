package heap

import "container/heap"

type Element struct {
	nums []int
	seen map[int]struct{}
}

func (e Element) Len() int {
	return len(e.nums)
}

// max heap
func (e Element) Less(i, j int) bool {
	return e.nums[i] > e.nums[j]
}

func (e Element) Swap(i, j int) {
	e.nums[i], e.nums[j] = e.nums[j], e.nums[i]
}

func (e *Element) Push(x any) {
	intX := x.(int)
	if _, exists := e.seen[intX]; exists {
		return
	}

	e.nums = append(e.nums, intX)
}

func (e *Element) Pop() any {
	old := e.nums
	num := old[len(old)-1]
	e.nums = old[:len(old)-1]
	return num
}

// https://leetcode.com/problems/kth-largest-element-in-an-array/
func findKthLargest(nums []int, k int) int {
	n := &Element{}
	heap.Init(n)
	for _, num := range nums {
		heap.Push(n, num)
	}

	var result int
	for range k {
		result = heap.Pop(n).(int)
	}

	return result
}

// version 2
type heaps struct {
	array []int
}

func (h *heaps) add(i int) {
	h.array = append(h.array, i)
	h.heapifyUp(len(h.array) - 1)
}

func (h *heaps) heapifyUp(idx int) {
	for h.array[idx] > h.array[parent(idx)] {
		h.swap(idx, parent(idx))
		idx = parent(idx)
	}
}

func (h *heaps) swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func (h *heaps) pop() int {
	val := h.array[0]
	lastIdx := len(h.array) - 1

	if len(h.array) == 0 {
		return -1
	}

	h.array[0] = h.array[lastIdx]
	h.array = h.array[:lastIdx]
	h.heapifyDown(0)

	return val
}

func (h *heaps) heapifyDown(idx int) {
	lastIdx := len(h.array) - 1
	l, r := left(idx), right(idx)
	var childToCompare int

	for l <= lastIdx {
		if h.array[l] == h.array[lastIdx] {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.array[idx] < h.array[childToCompare] {
			h.swap(idx, childToCompare)
			idx = childToCompare
			l, r = left(idx), right(idx)
		} else {
			return
		}
	}
}

func parent(idx int) int {
	return (idx - 1) / 2
}

func left(idx int) int {
	return (idx * 2) + 1
}

func right(idx int) int {
	return (idx * 2) + 2
}

func findKthLargest2(nums []int, k int) int {
	heap := &heaps{}

	for _, num := range nums {
		heap.add(num)
	}

	for i := 0; i < k-1; i++ {
		heap.pop()
	}

	return heap.pop()
}
