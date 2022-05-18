package mod

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountDiv(t *testing.T) {
	subTest := []struct {
		name   string
		num1   int
		num2   int
		modby  int
		result int
	}{
		{
			name:   "count div true",
			num1:   1,
			num2:   3,
			modby:  2,
			result: 1,
		}, {
			name:   "count div true",
			num1:   2,
			num2:   6,
			modby:  3,
			result: 2,
		},
	}

	for _, test := range subTest {
		t.Run(test.name, func(t *testing.T) {
			result := countDiv(test.num1, test.num2, test.modby)
			assert.Equal(t, test.result, result)
		})
	}
}
