package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxAreaOfIsland(t *testing.T) {
	subTest := []struct {
		name        string
		grid        [][]int
		expectedRes int
	}{
		{
			name: "max area of island 6",
			grid: [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}},
			expectedRes: 6,
		},
		{
			name:        "max area of island 0",
			grid:        [][]int{{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			expectedRes: 0,
		},
	}

	for _, test := range subTest {
		t.Run(test.name, func(t *testing.T) {
			result := maxAareaOfIsland(test.grid)
			assert.Equal(t, test.expectedRes, result)
		})
	}
}
