package string

// https://www.hackerrank.com/challenges/repeated-string/problem?isFullScreen=true
func RepeatedString(s string, n int64) int64 {
	sLen := int64(len(s))
	if sLen == 0 { // Handle empty string case
		return 0
	}

	var result int64
	for i := 0; i < int(sLen); i++ {
		if s[i] == 'a' {
			result++
		}
	}

	// in go divisor will floor the result
	fullRepeats := n / sLen
	result = result * fullRepeats

	// Calculate the length of the remaining substring
	remainderLength := n % sLen
	for i := int64(0); i < remainderLength; i++ {
		if s[i] == 'a' {
			result++
		}
	}

	return result
}
