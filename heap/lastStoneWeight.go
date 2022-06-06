package heap

import "math"

// https://leetcode.com/problems/last-stone-weight/

type heap struct {
	array []int
}

func (h *heap) add(i int) {
	h.array = append(h.array, i)
	h.heapifyUp(len(h.array) - 1)
}

func (h *heap) heapifyUp(idx int) {
	for h.array[parent(idx)] < h.array[idx] {
		h.swap(parent(idx), idx)
		idx = parent(idx)
	}
}

func (h *heap) swap(idx1, idx2 int) {
	h.array[idx1], h.array[idx2] = h.array[idx2], h.array[idx1]
}

func (h *heap) pop() int {
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

func (h *heap) heapifyDown(idx int) {
	lastIdx := len(h.array) - 1
	l, r := left(idx), right(idx)
	var childToCompare int

	for l <= lastIdx {
		if l == lastIdx {
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

func lastStoneWeight(stones []int) int {
	heap := &heap{}

	for _, val := range stones {
		heap.add(val)
	}

	for len(heap.array) != 1 {
		pop1 := heap.pop()
		pop2 := heap.pop()

		val := math.Abs(float64(pop1 - pop2))

		heap.add(int(val))
	}

	return heap.array[0]
}
