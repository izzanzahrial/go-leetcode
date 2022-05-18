package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxOfThree(t *testing.T) {
	subtest := []struct {
		name   string
		input  []int
		result int
	}{
		{
			name:   "max of three true",
			input:  []int{-1, 3, -1, 3, 1, 3},
			result: 27,
		},
		{
			name:   "max of three true",
			input:  []int{2, 1, 1, 2, 2},
			result: 8,
		},
	}
	for _, test := range subtest {
		t.Run(test.name, func(t *testing.T) {
			result := maxOfThree(test.input)
			assert.Equal(t, test.result, result)
		})
	}
}
