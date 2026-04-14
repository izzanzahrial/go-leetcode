package array

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

func twoSum2(nums []int, target int) []int {
	needMap := make(map[int]int, len(nums))

	for idx, val := range nums {
		if mapIdx, ok := needMap[val]; ok {
			return []int{mapIdx, idx}
		}

		needMap[target-val] = idx
	}

	return nil
}

func twoSum3(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}
