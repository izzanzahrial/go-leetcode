package sieveoferatosthenes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSemiPrime(t *testing.T) {
	subTest := []struct {
		name   string
		inputN int
		inputA []int
		inputB []int
		result []int
	}{
		{
			name:   "count semi prime true",
			inputN: 26,
			inputA: []int{1, 4, 16},
			inputB: []int{26, 10, 20},
			result: []int{10, 4, 0},
		},
	}

	for _, test := range subTest {
		t.Run(test.name, func(t *testing.T) {
			result := countSemiPrime(test.inputN, test.inputA, test.inputB)
			assert.Equal(t, test.result, result)
		})
	}
}
