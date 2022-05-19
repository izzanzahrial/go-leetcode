package leader

func dominator(nums []int) int {
	var result int
	var max int
	hash := make(map[int]int)

	for _, num := range nums {
		hash[num] += 1
		if hash[num] > max {
			max = hash[num]
			result = num
		}
	}

	return result
}
