package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNStraightHand(t *testing.T) {
	subTest := []struct {
		name        string
		hand        []int
		groupSize   int
		expectedRes bool
	}{
		{
			name:        "groupsize 3 true",
			hand:        []int{1, 2, 3, 6, 2, 3, 4, 7, 8},
			groupSize:   3,
			expectedRes: true,
		},
		{
			name:        "groupsize 4 false",
			hand:        []int{1, 2, 3, 4, 5},
			groupSize:   4,
			expectedRes: false,
		},
	}

	for _, test := range subTest {
		t.Run(test.name, func(t *testing.T) {
			result := isNStraightHand(test.hand, test.groupSize)
			assert.Equal(t, test.expectedRes, result)
		})
	}
}
