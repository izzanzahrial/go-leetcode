package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	subTest := []struct {
		name   string
		nums   []int
		target int
		result []int
	}{
		{
			name:   "is two sum true",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			result: []int{0, 1},
		},
		{
			name:   "is two sum true",
			nums:   []int{3, 2, 4},
			target: 6,
			result: []int{1, 2},
		},
		{
			name:   "is two sum true",
			nums:   []int{3, 3},
			target: 6,
			result: []int{0, 1},
		},
	}

	for _, test := range subTest {
		t.Run(test.name, func(t *testing.T) {
			result := twoSum(test.nums, test.target)
			assert.Equal(t, test.result, result)
		})
	}
}
