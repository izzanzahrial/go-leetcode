package pointer

// https://leetcode.com/problems/find-the-highest-altitude/description/?envType=study-plan-v2&envId=leetcode-75
func LargestAltitude(gain []int) int {
	maxAlt := 0
	currAlt := 0

	for _, val := range gain {
		currAlt += val

		if currAlt > maxAlt {
			maxAlt = currAlt
		}
	}

	return maxAlt
}
