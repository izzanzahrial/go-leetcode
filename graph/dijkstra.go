package graph

import (
	"container/heap"
	"math"
)

// Edge represents a connection between two nodes and its cost.
type Edge struct {
	to     string
	weight int
}

// Graph maps each node to a list of its edge connections.
type Graph map[string][]Edge

// Item represents a node in the priority queue.
type Item struct {
	node     string
	priority int
	index    int
}

// PriorityQueue implements heap.Interface for djikstra's algorithm.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Djikstra computes the shortest distance from start to every node
// if the distance within the graph is infinity, it means that the node is unreachable.
func Dijkstra(g Graph, start string) map[string]int {
	dist := make(map[string]int)
	for node := range g {
		dist[node] = math.MaxInt // initialize all distances to infinity
	}
	dist[start] = 0 // set the distance of the start node to 0, because the distance from start to start is 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{node: start, priority: 0}) // set the priority to 0 for the start node

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item)
		currentNode := current.node
		currentDist := current.priority

		// skip if we already have a shorter path
		if dist[currentNode] < currentDist {
			continue
		}

		// explore the neighbors of the current node
		for _, edge := range g[currentNode] {
			newDist := currentDist + edge.weight // the current weight to the current node + the weight of to visit the neighbor node
			if dist[edge.to] > newDist {         // if the neighbor node is already in the queue with a greater distance
				dist[edge.to] = newDist
				heap.Push(pq, &Item{node: edge.to, priority: newDist})
			}
		}
	}

	return dist
}
