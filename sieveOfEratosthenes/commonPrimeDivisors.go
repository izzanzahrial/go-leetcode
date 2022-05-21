package sieveoferatosthenes

func commonPrimeDivisors(a, b []int) int {
	var result int
	var prime []int

	// TODO: create the prime number

	for idx, val := range a {
		for _, num := range prime {
			i := checkDivisor(val, num)
			j := checkDivisor(b[idx], num)
			if i != j {
				break
			}
		}
		result += 1
	}

	return result
}

func checkDivisor(i, j int) bool {
	if i%j == 0 {
		return true
	}
	return false
}
