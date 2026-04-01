package array

// https://leetcode.com/problems/intersection-of-two-arrays/description/
func Intersection(nums1 []int, nums2 []int) []int {
	single := make(map[int]bool)
	for _, num := range nums1 {
		single[num] = true
	}

	var result []int
	for _, num := range nums2 {
		if single[num] {
			result = append(result, num)
			single[num] = false
		}
	}

	return result
}
