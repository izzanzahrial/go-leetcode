package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	expectedRes := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	result := threeSum(nums)

	assert.Equal(t, expectedRes, result)
}
