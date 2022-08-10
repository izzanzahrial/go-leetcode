package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfArithmeticSlices(t *testing.T) {
	subTest := []struct {
		name        string
		nums        []int
		expectedRes int
	}{
		{
			name:        "5 length arithmetic slices",
			nums:        []int{1, 2, 3, 4},
			expectedRes: 3,
		},
		{
			name:        "same number arithmetic slices",
			nums:        []int{7, 7, 7, 7},
			expectedRes: 3,
		},
		{
			name:        "1 length arithmetic slices",
			nums:        []int{1},
			expectedRes: 0,
		},
	}

	for _, test := range subTest {
		t.Run(test.name, func(t *testing.T) {
			result := numberOfArithmeticSlices(test.nums)
			assert.Equal(t, test.expectedRes, result)
		})
	}
}
