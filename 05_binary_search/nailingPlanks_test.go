package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNailingPlanks(t *testing.T) {
	subTest := []struct {
		name   string
		plankA []int
		plankB []int
		nails  []int
		result int
	}{
		{
			name:   "nailing planks true",
			plankA: []int{1, 4, 5, 8},
			plankB: []int{4, 5, 9, 10},
			nails:  []int{4, 6, 7, 10},
			result: 4,
		},
	}

	for _, test := range subTest {
		result := nailingPlanks(test.plankA, test.plankB, test.nails)
		assert.Equal(t, test.result, result)
	}
}
