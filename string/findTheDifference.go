package string

// https://leetcode.com/problems/find-the-difference/
func FindTheDifference(s string, t string) byte {
	var sumS, sumT int

	for _, chr := range s {
		sumS += int(chr)
	}

	for _, chr := range t {
		sumT += int(chr)
	}

	return byte(sumT - sumS)
}
