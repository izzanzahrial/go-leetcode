package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpiralOrder(t *testing.T) {
	subTest := []struct {
		name        string
		matrix      [][]int
		expectedRes []int
	}{
		{
			name:        "matrix 3 by 3",
			matrix:      [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expectedRes: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name:        "matrix 4 by 4",
			matrix:      [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			expectedRes: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}

	for _, test := range subTest {
		t.Run(test.name, func(t *testing.T) {
			result := spiralOrder(test.matrix)
			assert.Equal(t, test.expectedRes, result)
		})
	}
}
