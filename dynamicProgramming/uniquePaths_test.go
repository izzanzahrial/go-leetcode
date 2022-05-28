package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePaths(t *testing.T) {
	m := 3
	n := 7
	res := 28
	result := uniquePaths(m, n)
	assert.Equal(t, res, result)
}
