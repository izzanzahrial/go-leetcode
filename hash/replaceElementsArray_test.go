package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayChange(t *testing.T) {
	subTest := []struct {
		name        string
		nums        []int
		operations  [][]int
		expectedRes []int
	}{
		{
			name:        "array length of 4",
			nums:        []int{1, 2, 4, 6},
			operations:  [][]int{{1, 3}, {4, 7}, {6, 1}},
			expectedRes: []int{3, 2, 7, 1},
		},
		{
			name:        "array length of 2",
			nums:        []int{1, 2},
			operations:  [][]int{{1, 3}, {2, 1}, {3, 2}},
			expectedRes: []int{2, 1},
		},
	}

	for _, test := range subTest {
		t.Run(test.name, func(t *testing.T) {
			result := arrayChange(test.nums, test.operations)
			assert.Equal(t, test.expectedRes, result)
		})
	}
}
