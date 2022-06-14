package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonSubsequence(t *testing.T) {
	subTest := []struct {
		name        string
		text1       string
		text2       string
		expectedRes int
	}{
		{
			name:        "longest common subsequence 3",
			text1:       "abcde",
			text2:       "ace",
			expectedRes: 3,
		},
		{
			name:        "longest common subsequence 0",
			text1:       "abc",
			text2:       "def",
			expectedRes: 0,
		},
	}

	for _, test := range subTest {
		t.Run(test.name, func(t *testing.T) {
			result := longestCommonSubsequence(test.text1, test.text2)
			assert.Equal(t, test.expectedRes, result)
		})
	}
}
