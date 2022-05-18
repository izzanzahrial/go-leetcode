package mod

// https://app.codility.com/programmers/lessons/5-prefix_sums/count_div/

func countDiv(a, b, c int) int {
	var total int

	for i := a; i <= b; i++ {
		if i%c == 0 {
			total += 1
		}
	}

	return total
}
