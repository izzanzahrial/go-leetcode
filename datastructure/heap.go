package datastructure

type Heap struct {
	array []int
}

func NewHeap() *Heap {
	return &Heap{}
}

func (h *Heap) Add(i int) {
	h.array = append(h.array, i)
	h.heapifyUp(len(h.array) - 1)
}

func (h *Heap) heapifyUp(idx int) {
	for h.array[parent(idx)] < h.array[idx] {
		h.swap(parent(idx), idx)
		idx = parent(idx)
	}
}

func (h *Heap) swap(idx1, idx2 int) {
	h.array[idx1], h.array[idx2] = h.array[idx2], h.array[idx1]
}

func (h *Heap) Pop() int {
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

func (h *Heap) heapifyDown(idx int) {
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

func (h *Heap) Peek() int {
	return h.array[0]
}

func (h *Heap) Length() int {
	return len(h.array)
}

func (h *Heap) PopLast() int {
	val := h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]
	return val
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
