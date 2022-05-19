package primencomposite

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlags(t *testing.T) {
	subTest := []struct {
		name   string
		input  []int
		result int
	}{
		{
			name:   "flags true",
			input:  []int{1, 5, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2},
			result: 3,
		},
	}

	for _, test := range subTest {
		t.Run(test.name, func(t *testing.T) {
			result := flags(test.input)
			assert.Equal(t, test.result, result)
		})
	}
}
