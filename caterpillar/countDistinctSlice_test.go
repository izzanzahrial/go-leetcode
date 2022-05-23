package caterpillar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountDistinctSlice(t *testing.T) {
	m := 6
	nums := []int{3, 4, 5, 5, 2}
	result := countDistinctSlice(m, nums)
	assert.Equal(t, 7, result)
}
