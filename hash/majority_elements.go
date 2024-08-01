package hash

// https://leetcode.com/problems/majority-element/description/
func majorityElement(nums []int) int {
	numMap := make(map[int]int)

	for _, num := range nums {
		numMap[num]++
	}

	var result int
	var max int
	for k, num := range numMap {
		if num > max {
			max = num
			result = k
		}
	}

	return result
}
