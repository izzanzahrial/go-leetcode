package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchMatrix(t *testing.T) {
	subTest := []struct {
		name        string
		target      int
		matrix      [][]int
		expectedRes bool
	}{
		{
			name:        "search matrix true",
			target:      3,
			matrix:      [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			expectedRes: true,
		},
		{
			name:        "search matrix false",
			target:      13,
			matrix:      [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			expectedRes: false,
		},
	}

	for _, test := range subTest {
		t.Run(test.name, func(t *testing.T) {
			result := searchMatrix(test.matrix, test.target)
			assert.Equal(t, test.expectedRes, result)
		})
	}
}
