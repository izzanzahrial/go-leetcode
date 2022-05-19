package leader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDominator(t *testing.T) {
	subTest := []struct {
		name   string
		nums   []int
		result int
	}{
		{
			name:   "dominator true",
			nums:   []int{1, 2, 3, 3, 4, 5, 3, 5, 1, 1, 1},
			result: 1,
		},
	}

	for _, test := range subTest {
		t.Run(test.name, func(t *testing.T) {
			result := dominator(test.nums)
			assert.Equal(t, test.result, result)
		})
	}
}
