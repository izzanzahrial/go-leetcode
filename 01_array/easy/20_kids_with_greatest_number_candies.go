package array

// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/description/?envType=study-plan-v2&envId=leetcode-75
func KidsWithCandies(candies []int, extraCandies int) []bool {
	var max int
	var result []bool

	for _, val := range candies {
		if max < val {
			max = val
		}
	}

	for _, val := range candies {
		if val+extraCandies >= max {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}

	return result
}
