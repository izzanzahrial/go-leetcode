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

	for i := 0; i < len(a); i++ {
		ok := func() bool {
			for _, num := range prime {
				if num > a[i] || num > b[i] {
					break
				}
				divA := checkDivisor(a[i], num)
				divB := checkDivisor(b[i], num)
				if divA != divB {
					return false
				}
			}
			return true
		}()
		if ok {
			result += 1
		}
	}

	return result
}

func checkDivisor(i, j int) bool {
	return i%j == 0
}

func checkPrime(num int) bool {
	for i := int(math.Floor(float64(num) / 2)); i > 1; i-- {
		if num%i == 0 {
			return false
		}
	}
	return true
}
