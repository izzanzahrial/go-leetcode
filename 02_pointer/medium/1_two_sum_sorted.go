package pointer

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		currResult := numbers[left] + numbers[right]
		if currResult == target {
			// add +1 because 1-indexed array, which why tho, just use 0
			return []int{left + 1, right + 1}
		}

		if currResult > target {
			right--
		}

		if currResult < target {
			left++
		}
	}

	return []int{}
}
