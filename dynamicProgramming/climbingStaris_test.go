package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbingStairs(t *testing.T) {
	subTest := []struct {
		name        string
		steps       int
		expectedRes int
	}{
		{
			name:        "climbing stairs 1",
			steps:       1,
			expectedRes: 1,
		},
		{
			name:        "climbing stairs 2",
			steps:       2,
			expectedRes: 2,
		},
		{
			name:        "climbing stairs many",
			steps:       3,
			expectedRes: 3,
		},
	}

	for _, test := range subTest {
		result := climbingStairs(test.steps)
		assert.Equal(t, test.expectedRes, result)
	}
}
