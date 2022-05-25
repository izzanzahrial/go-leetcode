package dynamicprogramming

func climbingStairs(steps int) int {
	if steps == 1 {
		return 1
	}

	distinctSteps := []int{1, 2}

	for i := 2; i < steps; i++ {
		distinctSteps = append(distinctSteps, distinctSteps[i-1]+distinctSteps[i-2])
	}

	return distinctSteps[steps-1]
}
