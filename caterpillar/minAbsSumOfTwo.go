package caterpillar

import (
	"math"
)

func minAbsSumOfTwo(nums []int) int {
	min := math.MaxInt

	for _, num := range nums {
		for _, num2 := range nums {
			currMin := num + num2
			absMin := int(math.Abs(float64(currMin)))
			if min > absMin {
				min = int(absMin)
			}
		}
	}

	return min
}
