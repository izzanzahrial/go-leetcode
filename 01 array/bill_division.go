package array

import "fmt"

// https://www.hackerrank.com/challenges/bon-appetit/problem?isFullScreen=true
func BonAppetit(bill []int32, k int32, b int32) {
	var annasEat int32
	for i, food := range bill {
		if i == int(k) {
			continue
		}

		annasEat += food
	}

	annasActualBill := annasEat / 2
	diff := b - annasActualBill
	if diff != 0 {
		fmt.Print(diff)
		return
	}
	fmt.Print("Bon Appetit")
}
