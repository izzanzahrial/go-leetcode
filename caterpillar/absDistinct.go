package caterpillar

import (
	"math"
)

func absDistinct(nums []int) int {
	var newNums = make(map[float64]int)

	for _, num := range nums {
		absNums := math.Abs(float64(num))
		newNums[absNums] = num
	}

	return len(newNums)
}
