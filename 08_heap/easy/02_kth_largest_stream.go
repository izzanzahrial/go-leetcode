package heap

import "container/heap"

type KthLargest struct {
	heaps nums
	k     int
}

type nums []int

func (n nums) Len() int {
	return len(n)
}

// min heap
func (n nums) Less(i, j int) bool {
	return n[i] < n[j]
}

func (n nums) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n *nums) Push(num any) {
	*n = append(*n, num.(int))
}

func (n *nums) Pop() any {
	old := *n
	num := old[len(old)-1]
	*n = old[:len(old)-1]
	return num
}

func (n nums) Peak() int {
	return n[n.Len()-1]
}

func Constructor(k int, arr []int) KthLargest {
	n := &nums{}
	heap.Init(n)

	for _, num := range arr {
		heap.Push(n, num)
	}

	if n.Len() > k {
		heap.Pop(n)
	}

	return KthLargest{
		heaps: *n,
		k:     k,
	}
}

func (this *KthLargest) Add(val int) int {
	// peak
	if this.heaps.Peak() < val {
		if this.heaps.Len() > this.k {
			heap.Pop(&this.heaps)
		}

		heap.Push(&this.heaps, val)
	}

	return this.heaps.Peak()
}
