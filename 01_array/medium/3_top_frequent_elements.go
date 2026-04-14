package array

import "sort"

// https://leetcode.com/problems/top-k-frequent-elements
func topKFrequent(nums []int, k int) []int {
	var result []int

	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num] += 1
	}

	// For each number and its count in our map
	// Add the number to the array at index (count-1)
	// Example: if 1 appears 3 times, add 1 to frequency[2]
	frequency := make([][]int, len(nums))
	for num, frequent := range numMap {
		frequency[frequent-1] = append(frequency[frequent-1], num)
	}

	// We loop backwards from the highest frequency number
	for i := len(frequency) - 1; i >= 0; i-- {
		// Add all the numbers within the current frequency list
		// Example : in 3 frequency meaning frequency[2] there is only one number which is 1
		// so add 1 into the result
		for _, num := range frequency[i] {
			result = append(result, num)

			// If the length of the result already the same as k, return
			if len(result) == k {
				return result
			}
		}
	}

	return result
}

// or you can use heap
// if we are using min heap, pop until the size of the heap == k
// if we are using max heap, add the value of the pop into an array until the size of the array == k
func topKFrequent2(nums []int, k int) []int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	arr := make([][2]int, 0, len(count))
	for num, cnt := range count {
		arr = append(arr, [2]int{cnt, num})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] > arr[j][0]
	})

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = arr[i][1]
	}
	return res
}
