package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCheapestPrice(t *testing.T) {
	tests := []struct {
		cities      int
		flights     [][]int
		src         int
		dst         int
		stop        int
		expectedVal int
	}{
		{
			cities: 4,
			flights: [][]int{
				{0, 1, 100},
				{1, 2, 200},
				{2, 0, 100},
				{1, 3, 600},
			},
			src:         0,
			dst:         3,
			stop:        1,
			expectedVal: 700,
		},
		{
			cities: 3,
			flights: [][]int{
				{0, 1, 100},
				{1, 2, 100},
				{0, 2, 500},
			},
			src:         0,
			dst:         2,
			stop:        1,
			expectedVal: 200,
		},
		{
			cities: 3,
			flights: [][]int{
				{0, 1, 100},
				{1, 2, 100},
				{0, 2, 500},
			},
			src:         0,
			dst:         2,
			stop:        0,
			expectedVal: 500,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("find cheapest price with stop %d", i+1), func(t *testing.T) {
			result := findCheapestPrice(test.cities, test.flights, test.src, test.dst, test.stop)
			assert.Equal(t, test.expectedVal, result)
		})
	}
}
