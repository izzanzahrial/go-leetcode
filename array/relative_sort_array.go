package array

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	result := make([]int, len(arr1))
	hashMap := make(map[int]int, len(arr1))

	// count the total for each num in array 1
	for _, num := range arr1 {
		hashMap[num] += 1
	}

	// check if the num of array 2 is located in hashMap
	// add into result if it's located
	index := 0
	for _, num := range arr2 {
		if _, ok := hashMap[num]; ok {
			for i := 0; i < hashMap[num]; i++ {
				result[index] = num
				index += 1
			}

			hashMap[num] = -1
		}
	}

	// create the non counted value from hashMap
	// sort it
	var notCounted []int
	for num, val := range hashMap {
		if val > -1 {
			notCounted = append(notCounted, num)
		}
	}

	// add the rest of non counted value into result
	sort.Ints(notCounted)
	for _, num := range notCounted {
		for i := 0; i < hashMap[num]; i++ {
			result[index] = num
			index++
		}
	}

	return result
}
