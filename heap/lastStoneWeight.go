package heap

import (
	"math"

	"github.com/izzanzahrial/go-leetcode/datastructure"
)

// https://leetcode.com/problems/last-stone-weight/
func lastStoneWeight(stones []int) int {
	heap := datastructure.NewHeap()

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
