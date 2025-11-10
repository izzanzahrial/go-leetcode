package graph

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDijkstra(t *testing.T) {
	tests := []struct {
		name         string
		graph        Graph
		start        string
		notReachAble int
	}{
		{
			name: "every note is reachable",
			graph: Graph{
				"A": {{"B", 4}, {"C", 2}},
				"B": {{"C", 3}, {"D", 2}, {"E", 3}},
				"C": {{"B", 1}, {"D", 4}, {"E", 5}},
				"D": {},
				"E": {{"D", 1}},
			},
			start:        "A",
			notReachAble: 0,
		},
		{
			name: "note E is not reachable",
			graph: Graph{
				"A": {{"B", 2}},
				"B": {{"C", 3}},
				"C": {{"B", 1}, {"D", 4}},
				"D": {},
				"E": {{"D", 1}},
			},
			start:        "A",
			notReachAble: 1,
		},
		{
			name: "every note is not reachable except the starting node",
			graph: Graph{
				"A": {},
				"B": {},
				"C": {},
				"D": {},
				"E": {},
			},
			start:        "A",
			notReachAble: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dist := Dijkstra(test.graph, test.start)

			var notReachAble int
			for _, val := range dist {
				if val == math.MaxInt {
					notReachAble++
				}
			}

			assert.Equal(t, test.notReachAble, notReachAble, "not reachable node is not equal")
		})
	}
}
