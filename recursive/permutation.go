package recursive

// https://leetcode.com/problems/permutations/submissions/
func Permute(nums []int) [][]int {
	var result [][]int

	recursive([]int{}, nums, &result)

	return result
}

func recursive(curr, nums []int, result *[][]int) {
	if len(nums) == 0 {
		*result = append(*result, curr)
	}

	for i, val := range nums {
		curr1 := nums[:i]
		curr2 := nums[i+1:]
		newCurr := append(curr, val)
		newNums := append(append([]int{}, curr1...), curr2...)
		recursive(newCurr, newNums, result)
	}
}
