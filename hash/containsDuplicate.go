package hash

// https://leetcode.com/problems/contains-duplicate/

func containsDuplicate(nums []int) bool {
	hash := make(map[int]int)

	for _, val := range nums {
		if _, ok := hash[val]; ok {
			return true
		}
		hash[val] = val
	}

	return false
}

// func containsDuplicate2(nums []int) bool {
// 	hash := make(map[int]int)

// 	for _, val := range nums {
// 		if hash[val] > 0 {
// 			return true
// 		}
// 		hash[val] = 1
// 	}

// 	return false
// }
