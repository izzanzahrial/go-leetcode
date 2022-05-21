package sieveoferatosthenes

func countNonDivisible(nums []int) []int {
	var result []int

	for _, num := range nums {
		var current int

		for _, val := range nums {
			if val == num {
				continue
			}
			if num%val != 0 {
				current += 1
			}
		}
		result = append(result, current)
	}

	return result
}
