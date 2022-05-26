package dynamicprogramming

func palindrome(s string) string {
	var longest string
	var longestLen int

	for i := 0; i < len(s); i++ {
		// odd length
		left, right := i, i
		for left >= 0 && right < len(s) && s[left] == s[right] {
			currLen := right - left + 1
			if currLen > longestLen {
				longestLen = currLen
				longest = s[left : right+1]
			}
			left -= 1
			right += 1
		}

		// even length
		left, right = i, i+1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			currLen := right - left + 1
			if currLen > longestLen {
				longestLen = currLen
				longest = s[left : right+1]
			}
			left -= 1
			right += 1
		}
	}

	return longest
}
