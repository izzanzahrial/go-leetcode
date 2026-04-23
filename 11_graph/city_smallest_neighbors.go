package graph

import (
	"container/heap"
	"math"
)

type CityVertex struct {
	City   int
	Weight int
}

type CityPriorityQueue []*CityVertex

func (pq CityPriorityQueue) Len() int {
	return len(pq)
}

func (pq CityPriorityQueue) Less(i, j int) bool {
	return pq[i].Weight < pq[j].Weight
}

func (pq CityPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *CityPriorityQueue) Push(x any) {
	CityVertex := x.(*CityVertex)
	*pq = append(*pq, CityVertex)
}

func (pq *CityPriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	CityVertex := old[n-1]
	*pq = old[0 : n-1]
	return CityVertex
}

func dijkstra(n int, graph map[int][]CityVertex, start int, distanceThreshold int) int {
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt
	}
	dist[start] = 0

	pq := &CityPriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &CityVertex{City: start, Weight: 0})
	for pq.Len() > 0 {
		current := heap.Pop(pq).(*CityVertex)

		if current.Weight > distanceThreshold {
			continue
		}

		for _, vertex := range graph[current.City] {
			newWeight := current.Weight + vertex.Weight
			if newWeight < dist[vertex.City] {
				dist[vertex.City] = newWeight
				heap.Push(pq, &CityVertex{City: vertex.City, Weight: newWeight})
			}
		}
	}

	// Count reachable cities within threshold (excluding start)
	count := 0
	for i, d := range dist {
		if i != start && d <= distanceThreshold {
			count++
		}
	}
	return count
}

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	graph := make(map[int][]CityVertex, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], CityVertex{edge[1], edge[2]})
		graph[edge[1]] = append(graph[edge[1]], CityVertex{edge[0], edge[2]})
	}

	minCount := math.MaxInt
	city := -1
	for i := 0; i < n; i++ {
		count := dijkstra(n, graph, i, distanceThreshold)
		if count < minCount || (count == minCount && i > city) {
			minCount = count
			city = i
		}
	}
	return city
}
