package array

// https://leetcode.com/problems/contains-duplicate/description/
// https://neetcode.io/problems/duplicate-integer
func HasDuplicate(nums []int) bool {
	numsMap := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		if _, exists := numsMap[num]; exists {
			return true
		}
		numsMap[num] = struct{}{}
	}

	return false
}
