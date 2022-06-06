package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLastStoneWeight(t *testing.T) {
	subTest := []struct {
		name        string
		stones      []int
		expectedRes int
	}{
		{
			name:        "last stone weight true",
			stones:      []int{2, 7, 4, 1, 8, 1},
			expectedRes: 1,
		},
		{
			name:        "last stone weight true",
			stones:      []int{1},
			expectedRes: 1,
		},
	}

	for _, test := range subTest {
		t.Run(test.name, func(t *testing.T) {
			result := lastStoneWeight(test.stones)
			assert.Equal(t, test.expectedRes, result)
		})
	}
}
