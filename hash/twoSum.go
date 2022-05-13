package hash

// https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)

	for key, val := range nums {
		diff := target - val
		if idx, ok := hash[diff]; ok {
			return []int{idx, key}
		}
		hash[val] = key
	}

	return []int{}
}
