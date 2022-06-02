package stacknqueue

func dailyTemperatures(temperatures []int) []int {
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
