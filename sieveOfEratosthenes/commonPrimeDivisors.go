package sieveoferatosthenes

import (
	"math"
)

func commonPrimeDivisors(a, b []int) int {
	var result int
	var prime []int

	for i := 2; i < 6001; i++ {
		if ok := checkPrime(i); ok {
			prime = append(prime, i)
		}
	}

	// TODO: this part still wrong
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

func checkPrime(num int) bool {
	for i := int(math.Floor(float64(num) / 2)); i > 1; i-- {
		if num%i == 0 {
			return false
		}
	}
	return true
}
