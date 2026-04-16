package heap

import (
	"container/heap"
)

type MinCapitalHeap [][]int

func (h MinCapitalHeap) Len() int           { return len(h) }
func (h MinCapitalHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h MinCapitalHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinCapitalHeap) Push(x any)        { *h = append(*h, x.([]int)) }
func (h *MinCapitalHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MaxProfitHeap []int

func (h MaxProfitHeap) Len() int           { return len(h) }
func (h MaxProfitHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxProfitHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxProfitHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MaxProfitHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	minCapital := &MinCapitalHeap{}
	maxProfit := &MaxProfitHeap{}

	for i := 0; i < len(capital); i++ {
		heap.Push(minCapital, []int{capital[i], profits[i]})
	}

	for i := 0; i < k; i++ {
		for minCapital.Len() > 0 && (*minCapital)[0][0] <= w {
			item := heap.Pop(minCapital).([]int)
			heap.Push(maxProfit, item[1])
		}
		if maxProfit.Len() == 0 {
			break
		}
		w += heap.Pop(maxProfit).(int)
	}

	return w
}
