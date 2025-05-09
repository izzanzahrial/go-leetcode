package array

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
