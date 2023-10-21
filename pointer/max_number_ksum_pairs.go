package pointer

// https://leetcode.com/problems/max-number-of-k-sum-pairs/solutions/3545683/golang-sort-vs-dictionary/?envType=study-plan-v2&envId=leetcode-75
func MaxOperations(nums []int, k int) int {
	var result int
	hash := make(map[int]int)

	for _, num := range nums {
		if val, ok := hash[num]; ok {
			result++
			if val > 1 {
				hash[num] -= 1
			} else {
				delete(hash, num)
			}
		} else {
			hash[k-num] += 1
		}
	}

	return result
}
