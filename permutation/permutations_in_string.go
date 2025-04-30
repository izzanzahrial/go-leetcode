package permutation

import "strings"

// FAST USING SLIDING WINDOW IN S2
func CheckInclusion(s1 string, s2 string) bool {
	// create slice of 26 integers based on 26 characters in the alphabet
	s1CharSlice := [26]int{}
	for i := range s1 {
		// -97 for the index of the character in the alphabet
		// because of 'a' has index 97, 'b' has index 98 and so on
		// add the value of that character index += 1
		s1CharSlice[s1[i]-97]++
	}

	var left int
	for right := range s2 {
		// remove the value of that character index -= 1
		s1CharSlice[s2[right]-97]--

		if s1CharSlice == [26]int{} {
			return true
		}

		// if right already equal or more than len(s1)
		// we need to add the value of that character index to remove it from the slice
		// because previously we minused it in `s1CharSlice[s2[right]-97]--`
		// then we move the most left pointer of the s2
		if right+1 >= len(s1) {
			s1CharSlice[s2[left]-97]++
			left++
		}
	}

	return false
}

// STILL SLOW
func CheckInclusion2(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var permutations []string
	runes := []rune(s1)
	visited := make([]bool, len(runes))
	generatePermutationsFixed2(runes, visited, "", &permutations)

	for _, perm := range permutations {
		if strings.Contains(s2, perm) {
			return true
		}
	}

	return false
}

func generatePermutationsFixed2(runes []rune, visited []bool, currentPerm string, result *[]string) {
	if len(currentPerm) == len(runes) {
		*result = append(*result, currentPerm)
		return
	}

	for i := 0; i < len(runes); i++ {
		if !visited[i] {
			visited[i] = true
			generatePermutationsFixed2(runes, visited, currentPerm+string(runes[i]), result)

			// Backtrack: Unmark the character so it can be used in other paths
			visited[i] = false
		}
	}
}

// ANOTHER APPROACH BUT SLOWER
func CheckInclusion3(s1 string, s2 string) bool {
	// early exit, because if s1 lenght greater than s2
	// the s2 will be never containing substring of s1
	if len(s1) > len(s2) {
		return false
	}

	var permutations []string
	createPermutations3(s1, "", &permutations)
	for _, perm := range permutations {
		if contain := strings.Contains(s2, perm); contain {
			return true
		}
	}

	return false
}

// This function is slow because it's relies heavily on string slicing
// and string concatenation. In 'go' strings are immutable
// so each of these operation creates entierly new string
// this lead to a lot of memory allocation and copying
func createPermutations3(str, currStr string, result *[]string) {
	// if the length of the string already 0
	// means we already use all of the characters for the current permutation
	// add the current permutation str into the result
	// then return backtrack
	if len(str) == 0 {
		*result = append(*result, currStr)
		return
	}

	for i, char := range str {
		currStr = currStr + string(char)

		// create new str but removing the current character
		newStr := str[:i] + str[i+1:] // THIS STRING SLICING AND CONCATENATION IS SLOW!!!

		// recursive for the next step of decision tree
		createPermutations3(newStr, currStr, result)

		// backtrack, delete the previous value of str[i]
		currStr = currStr[:len(currStr)-1]
	}
}
