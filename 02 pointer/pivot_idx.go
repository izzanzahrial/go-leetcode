package pointer

// https://leetcode.com/problems/find-pivot-index/description/?envType=study-plan-v2&envId=leetcode-75
func PivotIndex(nums []int) int {
	leftSum := 0
	rightSum := total(nums)

	for idx, num := range nums {
		rightSum -= num

		if leftSum == rightSum {
			return idx
		}

		leftSum += num
	}

	return -1
}

func total(nums []int) int {
	var result int

	for _, num := range nums {
		result += num
	}

	return result
}
