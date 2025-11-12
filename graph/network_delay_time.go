package graph

import (
	"container/heap"
	"math"
)

type NetworkEdge struct {
	Node int
	Time int
}

type NetworkVertex struct {
	Node int
	Time int
}

type NetworkPriorityQueue []*NetworkVertex

func (pq NetworkPriorityQueue) Len() int {
	return len(pq)
}

func (pq NetworkPriorityQueue) Less(i, j int) bool {
	return pq[i].Time < pq[j].Time
}

func (pq NetworkPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *NetworkPriorityQueue) Push(x any) {
	NetworkVertex := x.(*NetworkVertex)
	*pq = append(*pq, NetworkVertex)
}

func (pq *NetworkPriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	NetworkVertex := old[n-1]
	*pq = old[0 : n-1]
	return NetworkVertex
}

func networkDelayTime(times [][]int, n int, k int) int {
	// create the graph map
	networkMap := make(map[int][]NetworkEdge)
	for _, time := range times {
		networkMap[time[0]] = append(networkMap[time[0]], NetworkEdge{time[1], time[2]})
	}

	// map of the node and the time it was visited
	visitedAt := make(map[int]int, n)
	for i := 1; i <= n; i++ {
		visitedAt[i] = math.MaxInt
	}
	visitedAt[k] = 0

	pq := &NetworkPriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &NetworkVertex{Node: k, Time: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*NetworkVertex)

		// this is to make sure it doesn't loop forever
		if current.Time > visitedAt[current.Node] {
			continue
		}

		for _, node := range networkMap[current.Node] {
			newTime := current.Time + node.Time
			if newTime < visitedAt[node.Node] {
				visitedAt[node.Node] = newTime
				heap.Push(pq, &NetworkVertex{Node: node.Node, Time: newTime})
			}
		}
	}

	maxTime := 0
	// search the max value within visited map
	// if one of them value is still math.MaxInt, then it means it wasn't visited return -1
	for _, t := range visitedAt {
		if t == math.MaxInt {
			return -1
		}
		if t > maxTime {
			maxTime = t
		}
	}

	return maxTime
}
