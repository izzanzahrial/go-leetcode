package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {
	subTest := []struct {
		name     string
		testCase []string
		result   bool
	}{
		{
			name:     "test is anagram true",
			testCase: []string{"anagram", "nagaram"},
			result:   true,
		},
		{
			name:     "test is anagram false",
			testCase: []string{"rat", "car"},
			result:   false,
		},
	}

	for _, test := range subTest {
		t.Run(test.name, func(t *testing.T) {
			result := isAnagram(test.testCase[0], test.testCase[1])
			assert.Equal(t, test.result, result)
		})
	}
}
