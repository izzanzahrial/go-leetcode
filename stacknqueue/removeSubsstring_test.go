package stacknqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveOccurrences(t *testing.T) {
	subTest := []struct {
		name        string
		s           string
		part        string
		expectedRes string
	}{
		{
			name:        "remove abc",
			s:           "daabcbaabcbc",
			part:        "abc",
			expectedRes: "dab",
		},
		{
			name:        "remove xy",
			s:           "axxxxyyyyb",
			part:        "xy",
			expectedRes: "ab",
		},
	}

	for _, test := range subTest {
		t.Run(test.name, func(t *testing.T) {
			result := removeOccurrences(test.s, test.part)
			assert.Equal(t, test.expectedRes, result)
		})
	}
}
