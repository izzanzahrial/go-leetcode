package heap

import (
	"container/heap"
	"sort"
)

// https://leetcode.com/problems/car-pooling/
func carPooling(trips [][]int, capacity int) bool {
	sort.Slice(trips, func(i, j int) bool {
		return trips[i][1] < trips[j][1]
	})

	h := &MinHeap{}
	heap.Init(h)
	curPass := 0

	for _, trip := range trips {
		numPass, start, end := trip[0], trip[1], trip[2]

		for h.Len() > 0 && (*h)[0][0] <= start {
			curPass -= heap.Pop(h).([2]int)[1]
		}

		curPass += numPass
		if curPass > capacity {
			return false
		}

		heap.Push(h, [2]int{end, numPass})
	}

	return true
}

type MinHeap [][2]int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.([2]int)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
