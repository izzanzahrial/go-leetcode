package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	subTest := []struct {
		name        string
		nums        []int
		target      int
		expectedRes int
	}{
		{
			name:        "search rotated array true",
			nums:        []int{4, 5, 6, 7, 0, 1, 2},
			target:      0,
			expectedRes: 4,
		},
		{
			name:        "search rotated array false",
			nums:        []int{4, 5, 6, 7, 0, 1, 2},
			target:      3,
			expectedRes: -1,
		},
	}

	for _, test := range subTest {
		t.Run(test.name, func(t *testing.T) {
			result := search(test.nums, test.target)
			assert.Equal(t, test.expectedRes, result)
		})
	}
}
