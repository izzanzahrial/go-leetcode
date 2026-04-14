package pointer

import (
	"container/heap"
	"math"
)

// https://leetcode.com/problems/sliding-window-maximum/
func maxSlidingWindow(nums []int, k int) []int {
	output := []int{}
	q := []int{}
	l, r := 0, 0

	for r < len(nums) {
		for len(q) > 0 && nums[q[len(q)-1]] < nums[r] {
			q = q[:len(q)-1]
		}
		q = append(q, r)

		if l > q[0] {
			q = q[1:]
		}

		if (r + 1) >= k {
			output = append(output, nums[q[0]])
			l += 1
		}
		r += 1
	}

	return output
}

type Pair struct {
	index, value int
}

type PairHeap []Pair

func (h PairHeap) Len() int { return len(h) }
func (h PairHeap) Less(i, j int) bool {
	return h[i].value > h[j].value || (h[i].value == h[j].value && h[i].index < h[j].index)
}
func (h PairHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *PairHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *PairHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxSlidingWindow2(nums []int, k int) []int {
	res := make([]int, len(nums)-k+1)

	i, j, index := 0, 0, 0
	pq := &PairHeap{}
	heap.Init(pq)

	for j < len(nums) {
		pair := Pair{j, nums[j]}
		heap.Push(pq, pair)
		if j-i+1 < k {
			j++
		} else if j-i+1 == k {
			for (*pq)[0].index < i {
				heap.Pop(pq)
			}
			first := (*pq)[0]
			res[index] = first.value
			if first.index < i+1 {
				heap.Pop(pq)
			}
			i++
			j++
			index++
		}
	}

	return res
}

// TIME LIMIT EXCEEDED
func maxSlidingWindow3(nums []int, k int) []int {
	var result []int

	for i := 0; i < len(nums)-k+1; i++ {
		currMax := math.MinInt
		for j := i; j < i+k; j++ {
			if nums[j] > currMax {
				currMax = nums[j]
			}
		}

		result = append(result, currMax)
	}

	return result
}
