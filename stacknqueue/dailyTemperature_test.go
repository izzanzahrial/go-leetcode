package stacknqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDailyTemperatures(t *testing.T) {
	subTest := []struct {
		name         string
		temperatures []int
		expectedRes  []int
	}{
		{
			name:         "daily temperatures true",
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			expectedRes:  []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			name:         "daily temperatures true",
			temperatures: []int{30, 40, 50, 60},
			expectedRes:  []int{1, 1, 1, 0},
		},
	}

	for _, test := range subTest {
		t.Run(test.name, func(t *testing.T) {
			result := dailyTemperatures(test.temperatures)
			assert.Equal(t, test.expectedRes, result)
		})
	}
}
