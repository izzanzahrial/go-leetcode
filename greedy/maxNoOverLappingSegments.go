package greedy

import "fmt"

func noOverlapp(arr1, arr2 []int) int {
	var result int

	for i := 0; i < len(arr1); i++ {
		fmt.Println(i)
		// TODO: still look ugly
		if i == 0 {
			if arr2[i] >= arr1[i+1] {
				continue
			}
		} else if i == len(arr1)-1 {
			if arr2[i-1] >= arr1[i] {
				continue
			}
		} else {
			if arr2[i-1] >= arr1[i] || arr2[i] >= arr1[i+1] {
				continue
			}
		}
		result += 1
	}

	return result
}
