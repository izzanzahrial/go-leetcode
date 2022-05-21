package sieveoferatosthenes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountNonDivisible(t *testing.T) {
	subTest := []struct {
		name   string
		input  []int
		result []int
	}{
		{
			name:   "count non divisible true",
			input:  []int{3, 1, 2, 3, 6},
			result: []int{2, 4, 3, 2, 0},
		},
	}

	for _, test := range subTest {
		t.Run(test.name, func(t *testing.T) {
			result := countNonDivisible(test.input)
			assert.Equal(t, test.result, result)
		})
	}
}
