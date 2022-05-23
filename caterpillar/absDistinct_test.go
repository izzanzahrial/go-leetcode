package caterpillar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbsDistinct(t *testing.T) {
	nums := []int{-5, -3, -1, 0, 3, 6}
	result := absDistinct(nums)
	expectedRes := 5
	assert.Equal(t, expectedRes, result)
}
