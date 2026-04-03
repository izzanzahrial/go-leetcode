package pointer

import "sort"

// https://leetcode.com/problems/boats-to-save-people/description/
func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	var result int

	start, end := 0, len(people)-1
	for start < end {
		// since the maximum number of people is 2
		// we search for perfect/almost perfect amount of weight
		// by totaling the weight of the lightest and heaviest people
		if people[start]+people[end] <= limit {
			result++
			start++
			// end--
		} else {
			// if the weight limit for two poeple doesn't account for it
			// only add the heaviest poeple, meaning the end of the line
			result++
			// end--
		}

		// move the end-- to here for less code
		// but for readibility use the inside if clause
		end--
	}

	// if there is one left people, because the start is equal to the end
	// add it to the result
	if start == end {
		result++
	}

	return result
}
