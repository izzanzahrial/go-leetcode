package array

// https://leetcode.com/problems/majority-element/description/
func majorityElement(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	numsMap := make(map[int]int)
	for _, num := range nums {
		numsMap[num]++
	}

	var result []int
	k := len(nums) / 3
	for key, num := range numsMap {
		if num > k {
			result = append(result, key)
		}
	}

	return result
}
