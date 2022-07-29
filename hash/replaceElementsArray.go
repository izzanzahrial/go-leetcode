package hash

// https://leetcode.com/problems/replace-elements-in-an-array/

func arrayChange(nums []int, operations [][]int) []int {
	hashMap := make(map[int]int)

	for idx, num := range nums {
		hashMap[num] = idx
	}

	for _, array := range operations {
		idx := hashMap[array[0]]
		delete(hashMap, array[0])
		hashMap[array[1]] = idx
		nums[idx] = array[1]
	}

	return nums
}
