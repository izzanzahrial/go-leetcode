package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Distinctfunc func([]int) int

func TestDistinct(t *testing.T) {
	subTest := []struct {
		name     string
		function Distinctfunc
		testCase []int
		result   int
	}{
		{
			name:     "test distinct true",
			function: distinct,
			testCase: []int{1, 2, 3, 3, 3},
			result:   2,
		},
		{
			name:     "test distinct2 true",
			function: distinct2,
			testCase: []int{3, 3, 3, 4, 7},
			result:   2,
		},
	}

	for _, test := range subTest {
		t.Run(test.name, func(t *testing.T) {
			result := test.function(test.testCase)
			assert.Equal(t, test.result, result)
		})
	}
}
