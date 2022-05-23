package caterpillar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinAbsSumOfTwo(t *testing.T) {
	nums := []int{-8, 4, 5, -10, 3}
	expectedRes := 3
	result := minAbsSumOfTwo(nums)
	assert.Equal(t, expectedRes, result)
}
