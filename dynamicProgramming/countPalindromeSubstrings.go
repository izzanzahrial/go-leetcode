package dynamicprogramming

import "fmt"

func palindromeSubstrings(s string) int {
	var result int

	for i := 0; i < len(s); i++ {
		left, right := i, i
		for left >= 0 && right < len(s) && s[left] == s[right] {
			result += 1
			left -= 1
			right += 1
		}

		fmt.Println(result)

		left, right = i, i+1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			result += 1
			left -= 1
			right += 1
		}

		fmt.Println(result)
	}

	return result
}
