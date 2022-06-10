package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWallsAndGates(t *testing.T) {
	subTest := []struct {
		name        string
		rooms       [][]int
		expectedRes [][]int
	}{
		{
			name: "walls and gates 4 by 4 room",
			rooms: [][]int{{2147483647, -1, 0, 2147483647},
				{2147483647, 2147483647, 2147483647, -1},
				{2147483647, -1, 2147483647, -1},
				{0, -1, 2147483647, 2147483647}},
			expectedRes: [][]int{{3, -1, 0, 1},
				{2, 2, 1, -1},
				{1, -1, 2, -1},
				{0, -1, 3, 4}},
		},
		{
			name:        "walls and gates 2 by 2 room",
			rooms:       [][]int{{0, -1}, {2147483647, 2147483647}},
			expectedRes: [][]int{{0, 1}, {1, 2}},
		},
	}

	for _, test := range subTest {
		t.Run(test.name, func(t *testing.T) {
			result := WallsAndGates2(test.rooms)
			assert.Equal(t, test.expectedRes, result)
		})
	}
}
