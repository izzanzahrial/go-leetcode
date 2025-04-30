package permutation

import "strings"

func CheckInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var permutations []string
	runes := []rune(s1)
	visited := make([]bool, len(runes))
	generatePermutationsFixed(runes, visited, "", &permutations)

	for _, perm := range permutations {
		if strings.Contains(s2, perm) {
			return true
		}
	}

	return false
}

func generatePermutationsFixed(runes []rune, visited []bool, currentPerm string, result *[]string) {
	// Base Case: The permutation has the same length as the original string
	if len(currentPerm) == len(runes) {
		*result = append(*result, currentPerm)
		return
	}

	// Recursive Step: Try adding each *unused* character
	for i := 0; i < len(runes); i++ {
		// If the character at index 'i' hasn't been visited/used in this path yet
		if !visited[i] {
			// Mark it as visited
			visited[i] = true

			// Recursively call, adding the character runes[i] to currentPerm
			// Note: No index 'idx' is needed here, we always check all possibilities
			generatePermutationsFixed(runes, visited, currentPerm+string(runes[i]), result)

			// Backtrack: Unmark the character so it can be used in other paths
			visited[i] = false
		}
	}
}

// ANOTHER APPROACH BUT SLOWER
func CheckInclusion2(s1 string, s2 string) bool {
	// early exit, because if s1 lenght greater than s2
	// the s2 will be never containing substring of s1
	if len(s1) > len(s2) {
		return false
	}

	var permutations []string
	createPermutations(s1, "", &permutations)
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
func createPermutations(str, currStr string, result *[]string) {
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
		createPermutations(newStr, currStr, result)

		// backtrack, delete the previous value of str[i]
		currStr = currStr[:len(currStr)-1]
	}
}
