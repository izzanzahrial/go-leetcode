package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoOverLapp(t *testing.T) {
	arr1 := []int{1, 3, 7, 9, 9, 11}
	arr2 := []int{5, 6, 8, 9, 10, 13}
	expectedRes := 2
	result := noOverlapp(arr1, arr2)
	assert.Equal(t, expectedRes, result)
}
