package stacknqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBracket(t *testing.T) {
	subtest := []struct {
		name   string
		input  string
		result bool
	}{
		{
			name:   "bracket true",
			input:  "{[()()]}",
			result: true,
		},
		{
			name:   "bracket false",
			input:  "{}([]]",
			result: false,
		},
	}

	for _, test := range subtest {
		t.Run(test.name, func(t *testing.T) {
			result := brackets(test.input)
			assert.Equal(t, test.result, result)
		})
	}
}
