package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumIslands(t *testing.T) {
	subTest := []struct {
		name        string
		grid        [][]byte
		expectedRes int
	}{
		{
			name: "number of island 1",
			grid: [][]byte{{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'}},
			expectedRes: 1,
		},
		{
			name: "number of island 3",
			grid: [][]byte{{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'}},
			expectedRes: 3,
		},
	}

	for _, test := range subTest {
		t.Run(test.name, func(t *testing.T) {
			result := numIslands(test.grid)
			assert.Equal(t, test.expectedRes, result)
		})
	}
}
