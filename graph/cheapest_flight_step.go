package graph

import (
	"container/heap"
	"math"
)

type Flight struct {
	City  int
	Price int
}

type FlightVertex struct {
	City         int
	CurrentPrice int
	CurrStop     int
}

type FlightPriorityQueue []*FlightVertex

func (pq FlightPriorityQueue) Len() int {
	return len(pq)
}

func (pq FlightPriorityQueue) Less(i, j int) bool {
	return pq[i].CurrentPrice < pq[j].CurrentPrice
}

func (pq FlightPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *FlightPriorityQueue) Push(x any) {
	FlightVertex := x.(*FlightVertex)
	*pq = append(*pq, FlightVertex)
}

func (pq *FlightPriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	FlightVertex := old[n-1]
	*pq = old[0 : n-1]
	return FlightVertex
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	// create graph map
	flightMap := make(map[int][]Flight)

	for _, flight := range flights {
		flightMap[flight[0]] = append(flightMap[flight[0]], Flight{flight[1], flight[2]})
	}

	// create slice of min price for the city
	minPriceAtStops := make([][]int, n)
	for i := 0; i < n; i++ {
		minPriceAtStops[i] = make([]int, k+2)
		for j := 0; j <= k+1; j++ {
			minPriceAtStops[i][j] = math.MaxInt
		}
	}
	minPriceAtStops[src][0] = 0

	pq := &FlightPriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &FlightVertex{City: src, CurrentPrice: 0, CurrStop: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*FlightVertex)
		currentCity := current.City
		currentPrice := current.CurrentPrice

		if currentCity == dst {
			return currentPrice
		}

		if current.CurrStop > k {
			continue
		}

		for _, flight := range flightMap[currentCity] {
			newPrice := currentPrice + flight.Price
			if newPrice < minPriceAtStops[flight.City][current.CurrStop+1] {
				minPriceAtStops[flight.City][current.CurrStop+1] = newPrice
				heap.Push(pq, &FlightVertex{
					City:         flight.City,
					CurrentPrice: newPrice,
					CurrStop:     current.CurrStop + 1,
				})
			}
		}
	}

	return -1
}
