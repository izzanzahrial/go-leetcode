package sieveoferatosthenes

import (
	"fmt"
	"math"
)

func countSemiPrime(n int, a, b []int) []int {
	var result []int
	var prime []int
	// var semiPrime []int

	asd
	// TODO: still error finding the prime
	for i := 2; i <= n; i++ {
		for j := 2; j <= int(math.Floor(float64(i)/2)); j++ {
			if i%j == 0 {
				break
			}
			prime = append(prime, i)
		}
	}

	fmt.Println(prime)

	// TODO: create semi prime from prime number

	// for idx, val := range a {
	// 	var current int

	// 	for i := val; i <= b[idx]; i++ {
	// check if the number in semi prime
	// 	}

	// 	result = append(result, current)
	// }

	return result
}
