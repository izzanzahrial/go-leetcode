package array

// https://leetcode.com/problems/sort-colors/description/
func sortColors(nums []int) {
	// Create bucket sort
	// Since we already know the value only containts 0,1,2 we create 3 buckets
	bucketMap := make(map[int]int, 3)
	for _, num := range nums {
		bucketMap[num] += 1
	}

	index := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < bucketMap[i]; j++ {
			nums[index] = i
			index += 1
		}
	}
}
