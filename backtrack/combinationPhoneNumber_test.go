package backtrack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetterCombinations(t *testing.T) {
	subTest := []struct {
		name        string
		digits      string
		expectedRes []string
	}{
		{
			name:        "letter combination two number",
			digits:      "23",
			expectedRes: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name:        "letter combination empty",
			digits:      "",
			expectedRes: []string{},
		},
		{
			name:        "letter combination one number",
			digits:      "2",
			expectedRes: []string{"a", "b", "c"},
		},
	}

	for _, test := range subTest {
		t.Run(test.name, func(t *testing.T) {
			result := letterCombinations(test.digits)
			assert.Equal(t, test.expectedRes, result)
		})
	}
}
