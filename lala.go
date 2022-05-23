package main

func noOverlapp(arr1, arr2 []int) int {
	var result int

	for i := 1; i < len(arr1); i++ {
		if arr2[i-1] <= arr1[i] || arr2[i] >= arr1[i+1] {
			continue
		}
		result += 1
	}

	return result
}
