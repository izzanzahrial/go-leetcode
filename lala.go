package main

type minHeap struct {
	array [][]int
}

func (m *minHeap) put(target, weight int) {
	m.array = append(m.array, []int{weight, target})
	m.heapifyUp(len(m.array) - 1)
}

func (m *minHeap) heapifyUp(idx int) {
	for m.array[parent(idx)][0] > m.array[idx][0] {
		m.swap(parent(idx), idx)
		idx = parent(idx)
	}
}

func (m *minHeap) pop() []int {
	if len(m.array) == 0 {
		return []int{}
	}

	result := m.array[0]
	lastIdx := len(m.array) - 1

	m.array[0] = m.array[lastIdx]
	m.array = m.array[:lastIdx]
	m.heapifyDown(0)

	return result
}

func (m *minHeap) heapifyDown(idx int) {
	lastIdx := len(m.array) - 1
	l, r := left(idx), right(idx)
	var childToCompare int

	for l <= lastIdx {
		if l == lastIdx {
			childToCompare = l
		} else if m.array[l][0] < m.array[r][0] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if m.array[idx][0] > m.array[childToCompare][0] {
			m.swap(idx, childToCompare)
			idx = childToCompare
			l, r = left(idx), right(idx)
		} else {
			return
		}
	}
}

func (m *minHeap) swap(i, j int) {
	m.array[i], m.array[j] = m.array[j], m.array[i]
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

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func networkDelayTime(times [][]int, n, k int) int {
	var result int
	nodes := make(map[int][][]int)

	for _, node := range times {
		currNode := node[0]
		nodes[currNode] = append(nodes[currNode], []int{node[1], node[2]})
	}

	minHeap := minHeap{
		array: [][]int{{0, k}},
	}
	visited := make(map[int]bool)

	for minHeap.array != nil {
		currNode := minHeap.pop()
		w1 := currNode[0]
		n1 := currNode[1]
		if _, ok := visited[n1]; ok {
			continue
		}
		visited[n1] = true
		result = max(result, w1)

		for _, val := range nodes[n1] {
			if _, ok := visited[val[0]]; !ok {
				minHeap.put(val[0], val[1])
			}
		}
	}

	if len(visited) == n {
		return result
	}

	return -1
}
