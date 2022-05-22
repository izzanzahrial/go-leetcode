package primencomposite

import "fmt"

func flags(array []int) int {
	var max int
	var hash []int

	for idx, val := range array {
		if idx == 0 {
			continue
		}
		if array[idx-1] < val && val > array[idx+1] {
			hash = append(hash, idx)
		}
		fmt.Println(val)
	}

	fmt.Println(hash)

	// TODO: still doesnt know how to find the max K flags
	for i := 0; i < len(hash)-1; i++ {
		if hash[i+1]-len(hash) >= hash[i+1]-hash[i] {
			max += 1
		}
		fmt.Println(hash[i])
	}

	return max
}
