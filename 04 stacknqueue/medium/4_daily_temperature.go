package stacknqueue

// https://leetcode.com/problems/daily-temperatures/description/
func dailyTemperatures(temperatures []int) []int {
	warmerStack := make([]int, len(temperatures))
	var indexStack []int

	for i, temp := range temperatures {
		// the logic is we will keep poping up all temperatures from the index stack
		// to be placed in certain index within warmer stack
		// e.g. we have temp [78, 77, 76, 80]
		// ======first loop======
		// warmerStack = [0, 0, 0, 0]
		// add 78 index to the index stack '0'
		// indexStack = [0]
		// ======second loop======
		// warmerStack = [0, 0, 0, 0]
		// add 77 index to the index stack '1'
		// indexStack = [0, 1]
		// ======third loop======
		// warmerStack = [0, 0, 0, 0]
		// add 76 index to the index stack '2'
		// indexStack = [0, 1, 2]
		// ======forth loop======
		// warmerStack = [0, 0, 0, 0]
		// indexStack = [0, 1, 2]
		// since the last temp meaning (80) is bigger than the top of indexstack
		// we will proceed to the inner loop
		// we get the top of the indexStack meaning '2'
		// that's the location of the latest temperature that's lower than the current temperature
		// we will pop '2' from the indexStack
		// then add the value to the warmerStack for the current index warmerStack[2]
		// the value is calculated by determined the range of the current temperature to the lower temperature index
		// (i - index) meaning (3 - 1)
		// keep looping so on until the indexStack is 0 or the top of indexStack temperature is higher than the current temperature
		// meaning we have to add the current temperature index to the indexStack
		for len(indexStack) > 0 && temperatures[indexStack[len(indexStack)-1]] < temp {
			index := indexStack[len(indexStack)-1]
			indexStack = indexStack[:len(indexStack)-1]
			warmerStack[index] = i - index
		}

		indexStack = append(indexStack, i)
	}

	return warmerStack
}

func dailyTemperatures2(temperatures []int) []int {
	var result []int

	for i := 0; i < len(temperatures); i++ {
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] {
				result = append(result, j-i)
				break
			}
			if j == len(temperatures)-1 {
				result = append(result, 0)
			}
		}
		if i == len(temperatures)-1 {
			result = append(result, 0)
		}
	}

	return result
}
