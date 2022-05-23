package caterpillar

import (
	"fmt"
	"strconv"
)

func countDistinctSlice(m int, nums []int) int {
	var distinct = make(map[int]string)

	for idx, num := range nums {
		for idx2, num2 := range nums {
			if num <= m {
				dis := num + num2
				str := strconv.Itoa(idx) + " " + strconv.Itoa(idx2)
				if _, ok := distinct[dis]; !ok {
					distinct[dis] = str
				}
			}
		}
	}

	fmt.Println(distinct)

	if len(distinct) > 1000000000 {
		return 1000000000
	}

	return len(distinct)
}
