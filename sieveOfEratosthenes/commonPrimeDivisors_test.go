package sieveoferatosthenes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommonPrimeDivisors(t *testing.T) {
	subTest := []struct {
		name   string
		inputA []int
		inputB []int
		result int
	}{
		{
			name:   "common prime divisors true",
			inputA: []int{15, 10, 3},
			inputB: []int{75, 30, 5},
			result: 1,
		},
	}

	for _, test := range subTest {
		t.Run(test.name, func(t *testing.T) {
			result := commonPrimeDivisors(test.inputA, test.inputB)
			assert.Equal(t, test.result, result)
		})
	}
}
