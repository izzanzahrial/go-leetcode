package stacknqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvalRPN(t *testing.T) {
	subTest := []struct {
		name        string
		tokens      []string
		expectedRes int
	}{
		{
			name:        "eval RPN true",
			tokens:      []string{"4", "13", "5", "/", "+"},
			expectedRes: 6,
		},
		{
			name:        "eval RPN true",
			tokens:      []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			expectedRes: 22,
		},
	}

	for _, test := range subTest {
		t.Run(test.name, func(t *testing.T) {
			result := evalRPN(test.tokens)
			assert.Equal(t, test.expectedRes, result)
		})
	}
}
