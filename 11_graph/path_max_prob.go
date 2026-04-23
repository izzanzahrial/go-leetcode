package graph

import "container/heap"

// https://leetcode.com/problems/path-with-maximum-probability/
type ProbEdge struct {
	Node int
	Prob float64 // probability to go to the current node
}

type Vertex struct {
	Node int
	Prob float64 // probability from the start until the current node
}

type ProbPriorityQueue []*Vertex

func (pq ProbPriorityQueue) Len() int {
	return len(pq)
}

func (pq ProbPriorityQueue) Less(i, j int) bool {
	return pq[i].Prob > pq[j].Prob
}

func (pq ProbPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *ProbPriorityQueue) Push(x any) {
	vertex := x.(*Vertex)
	*pq = append(*pq, vertex)
}

func (pq *ProbPriorityQueue) Pop() any {
	old := *pq
	lastIdx := pq.Len() - 1
	vertex := old[lastIdx]
	*pq = old[:lastIdx]
	return vertex
}

func maxProbability(n int, ProbEdges [][]int, succProb []float64, start_node int, end_node int) float64 {
	// since the graph map for the neirhboring nodes is not provided
	// we create map for node and it's neighboring nodes
	nodeNeighbors := make(map[int][]ProbEdge)

	for i := range ProbEdges {
		// first node = u
		// second node = v
		// probability to go to u -> v or v -> u = w
		u, v, w := ProbEdges[i][0], ProbEdges[i][1], succProb[i]

		// add both of them to the nodeNeighbors map
		// but if the value of the probability is 0 we continue, meaning we cant go there
		if w == 0 {
			continue
		}

		nodeNeighbors[u] = append(nodeNeighbors[u], ProbEdge{
			Node: v,
			Prob: w,
		})

		nodeNeighbors[v] = append(nodeNeighbors[v], ProbEdge{
			Node: u,
			Prob: w,
		})
	}

	// since this is undirected graph, we can keep coming back the previous node,
	// we handle this by creating visited nodes
	visited := make([]bool, n)

	// we also create a slice check wether if there are multiple ways to the node,
	// if the previous prob is greater than the current prob
	maxProb := make([]float64, n)

	// the prob start at 1 because if we start at the starting node,
	// it will always success
	pq := &ProbPriorityQueue{{Node: start_node, Prob: 1}}
	heap.Init(pq)

	for pq.Len() > 0 {
		vertex := heap.Pop(pq).(*Vertex)

		if vertex.Node == end_node {
			return vertex.Prob
		}

		if visited[vertex.Node] {
			continue
		}
		visited[vertex.Node] = true

		// check all it's neighboring nodes
		for _, e := range nodeNeighbors[vertex.Node] {
			// now check visited for the neighboring node
			if visited[e.Node] {
				continue
			}

			newProb := vertex.Prob * e.Prob
			if newProb > maxProb[e.Node] {
				maxProb[e.Node] = newProb
				// we push the node into the heap, but with the probability value,
				// that have been added from the start
				heap.Push(pq, &Vertex{Node: e.Node, Prob: newProb})
			}
		}
	}

	return 0
}
