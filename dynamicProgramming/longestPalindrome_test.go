package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindrom(t *testing.T) {
	subTest := []struct {
		name        string
		str         string
		expectedRes string
	}{
		{
			name:        "palindrome aba",
			str:         "aba",
			expectedRes: "aba",
		},
		{
			name:        "palindrome adffaaaffqwerq",
			str:         "adffaaaffqwerq",
			expectedRes: "ffaaaff",
		},
	}

	for _, test := range subTest {
		t.Run(test.name, func(t *testing.T) {
			result := palindrome(test.str)
			assert.Equal(t, test.expectedRes, result)
		})
	}
}
