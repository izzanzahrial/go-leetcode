package heap

import "container/heap"

type nums [][]int

func (n nums) Len() int {
	return len(n)
}

func (n nums) Less(i, j int) bool {
	// i'll put the distance in index 2
	return n[i][2] < n[j][2]
}

func (n nums) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n *nums) Push(x any) {
	*n = append(*n, x.([]int))
}

func (n *nums) Pop() any {
	old := *n
	i := old[len(old)-1]
	*n = old[:len(old)-1]
	return i
}

// https://leetcode.com/problems/k-closest-points-to-origin/
func kClosest(points [][]int, k int) [][]int {
	n := &nums{}
	heap.Init(n)
	for _, point := range points {
		distance := point[0]*point[0] + point[1]*point[1]
		p := []int{point[0], point[1], distance}
		heap.Push(n, p)
	}

	result := [][]int{}
	for i := 0; i < k; i++ {
		point := heap.Pop(n).([]int)
		result = append(result, point[:2])
	}

	return result
}
