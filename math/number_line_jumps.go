package math

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// If they start at the same position, they meet at time 0.
	if x1 == x2 {
		return "YES"
	}

	if x1 > x2 && v1 >= v2 { // If K1 is ahead and not slower than K2
		return "NO" // K2 cannot catch up
	}
	// K2 is ahead (x2 > x1) but K1 is not faster (v1 <= v2)
	if x2 > x1 && v2 >= v1 { // If K2 is ahead and not slower than K1
		return "NO" // K1 cannot catch up
	}

	// We need to check if the distance difference is exactly divisible by the speed difference.
	// Equation: x1 + k*v1 = x2 + k*v2 => k * (v1 - v2) = x2 - x1
	// We need (x2 - x1) to be divisible by (v1 - v2).

	// Using the original equation variables for clarity in the modulo check
	// num = x2 - x1 (distance between them)
	// den = v1 - v2 (relative speed of K1 towards K2)

	// Note: Because of the earlier checks, if x1 != x2, we know num and den will have the same sign:
	// If x1 < x2, num > 0. We must have v1 > v2 (otherwise returned NO), so den > 0. num/den > 0.
	// If x2 < x1, num < 0. We must have v2 > v1 (otherwise returned NO), so v1 - v2 < 0, den < 0. num/den > 0.
	// So, if the remainder is 0, k = num / den will be a positive integer.
	// If x1 == x2, num = 0, 0 % den == 0 for any non-zero den. If den is 0 (v1 == v2), x1==x2 was YES.

	num := x2 - x1
	den := v1 - v2

	// Check for divisibility
	if num%den == 0 {
		return "YES"
	} else {
		return "NO"
	}
}
