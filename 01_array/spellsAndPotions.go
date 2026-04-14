package array

// https://leetcode.com/problems/successful-pairs-of-spells-and-potions/?envType=study-plan-v2&envId=leetcode-75
func SuccessfulPairs(spells []int, potions []int, success int64) []int {
	var successRate []int

	for _, spell := range spells {
		var currRate int
		for _, potion := range potions {
			if int64(spell*potion) >= success {
				currRate++
			}
		}
		successRate = append(successRate, currRate)
	}

	return successRate
}
