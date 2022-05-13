package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	subtest := []struct {
		name     string
		testCase []int
		result   bool
	}{
		{
			name:     "contains duplicate true",
			testCase: []int{1, 2, 3, 1},
			result:   true,
		},
		{
			name:     "contains duplicate false",
			testCase: []int{1, 2, 3, 4},
			result:   false,
		},
		{
			name:     "contains duplicate true",
			testCase: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			result:   true,
		},
	}

	for _, test := range subtest {
		t.Run(test.name, func(t *testing.T) {
			result := containsDuplicate(test.testCase)
			assert.Equal(t, test.result, result)
		})
	}
}
