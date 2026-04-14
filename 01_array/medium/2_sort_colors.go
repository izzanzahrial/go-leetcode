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

func sortColors2(nums []int) {
	var bucket0, bucket1, bucket2 int
	for _, num := range nums {
		switch num {
		case 0:
			bucket0++
		case 1:
			bucket1++
		case 2:
			bucket2++
		}
	}

	for i := 0; i < bucket0; i++ {
		nums[i] = 0
	}

	for i := bucket0; i < bucket0+bucket1; i++ {
		nums[i] = 1
	}

	for i := bucket0 + bucket1; i < bucket0+bucket1+bucket2; i++ {
		nums[i] = 2
	}
}
