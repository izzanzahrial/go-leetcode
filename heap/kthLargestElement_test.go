package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindKthLargest(t *testing.T) {
	subTest := []struct {
		name        string
		function    func(nums []int, k int) int
		nums        []int
		k           int
		expectedRes int
	}{
		{
			name:        "find kth largest element using heap",
			function:    findKthLargest,
			nums:        []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:           4,
			expectedRes: 4,
		},
		{
			name:        "find kth largest element using sort",
			function:    findKthLargest2,
			nums:        []int{3, 2, 1, 5, 6, 4},
			k:           2,
			expectedRes: 5,
		},
	}

	for _, test := range subTest {
		t.Run(test.name, func(t *testing.T) {
			result := test.function(test.nums, test.k)
			assert.Equal(t, test.expectedRes, result)
		})
	}
}
