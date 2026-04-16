package heap

import (
	"container/heap"
	"math"

	internalheap "github.com/izzanzahrial/go-leetcode/08_heap"
)

// https://leetcode.com/problems/last-stone-weight/
func lastStoneWeight(stones []int) int {
	heap := internalheap.NewHeap()

	for _, val := range stones {
		heap.Add(val)
	}

	for heap.Length() != 1 {
		pop1 := heap.Pop()
		pop2 := heap.Pop()

		val := math.Abs(float64(pop1 - pop2))

		heap.Add(int(val))
	}

	return heap.Peek()
}

// version 2
type nums2 []int

func (n nums2) Len() int {
	return len(n)
}

func (n nums2) Less(i, j int) bool {
	return n[i] > n[j]
}

func (n nums2) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n *nums2) Push(i any) {
	*n = append(*n, i.(int))
}

func (n *nums2) Pop() any {
	old := *n
	i := old[len(old)-1]
	*n = old[:len(old)-1]
	return i
}

func lastStoneWeight2(stones []int) int {
	arr := &nums2{}
	heap.Init(arr)

	for _, stone := range stones {
		heap.Push(arr, stone)
	}

	for arr.Len() > 1 {
		stone1 := heap.Pop(arr).(int)
		stone2 := heap.Pop(arr).(int)

		result := stone1 - stone2
		if result != 0 {
			heap.Push(arr, result)
		}
	}

	if arr.Len() == 0 {
		return 0
	}

	return arr.Pop().(int)
}
