package pointer

// https://leetcode.com/problems/reverse-string/description/
func reverseString(s []byte) {
	// i starts at the beginning, j starts at the end
	// since in go you can't for loop with multiple conditional updates like i++, j--
	// move the j-- inside the loop
	for i, j := 0, len(s)-1; i < j; i++ {
		// Swap the elements at positions i and j
		s[i], s[j] = s[j], s[i]
		j--
	}
	// No explicit return is needed because slices are reference types,
	// and the modification is done in-place.
}
