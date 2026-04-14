package stacknqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateParenthesis(t *testing.T) {
	subTest := []struct {
		name        string
		n           int
		expectedRes []string
	}{
		{
			name:        "generate parenthesis true",
			n:           3,
			expectedRes: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name:        "generate parenthesis true 2",
			n:           1,
			expectedRes: []string{"()"},
		},
	}

	for _, test := range subTest {
		t.Run(test.name, func(t *testing.T) {
			result := generateParenthesis(test.n)
			assert.Equal(t, test.expectedRes, result)
		})
	}
}
