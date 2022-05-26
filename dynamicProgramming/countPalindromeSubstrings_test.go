package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindromeSubstrings(t *testing.T) {
	str := "aaa"
	expectedRes := 6
	result := palindromeSubstrings(str)
	assert.Equal(t, expectedRes, result)
}
