package slidingwindow

// https://leetcode.com/problems/next-greater-element-ii/
func NextGreaterElements(nums []int) []int {
	var result []int

	for i, num := range nums {
		j := i + 1
		if j == len(nums) {
			j = 0
		}

		for j != i {
			if nums[j] > num {
				result = append(result, nums[j])
				break
			}

			j++
			if j == len(nums) {
				j = 0
			}
		}

		if j == i {
			result = append(result, -1)
		}
	}

	return result
}
