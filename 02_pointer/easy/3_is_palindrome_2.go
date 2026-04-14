package pointer

func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			// We delete one of them for each test case
			// meaning we have 2 test case, if one of the test case return true, we return true else false
			return isPalindrome2(left+1, right, s) || isPalindrome2(left, right-1, s)
		}

		left++
		right--
	}

	return true
}

func isPalindrome2(left, right int, s string) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func validPalindrome2(s string) bool {
	for idx := range s {
		currStr := s[:idx] + s[idx+1:]

		if result := isPalindrome3(currStr); result {
			return true
		}
	}

	return false
}

func isPalindrome3(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}
