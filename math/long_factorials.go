package math

import (
	"fmt"
	"math/big"
)

// https://www.hackerrank.com/challenges/extra-long-factorials/problem?isFullScreen=true
func ExtraLongFactorials(n int32) {
	facMap := map[int]*big.Int{
		1: big.NewInt(1),
		2: big.NewInt(2),
	}

	result := checkFactorials(int64(n), facMap)
	fmt.Printf("%d", result)
}

func checkFactorials(num int64, facMap map[int]*big.Int) *big.Int {
	if facNum, exists := facMap[int(num)]; exists {
		return facNum
	}

	bigNum := big.NewInt(num)
	prevFac := checkFactorials(num-1, facMap)

	result := new(big.Int).Mul(bigNum, prevFac)
	facMap[int(num)] = result
	return result
}
