package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTargetSum(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	expectedRes := 5
	result := targetSum(nums, target)
	assert.Equal(t, expectedRes, result)
}
