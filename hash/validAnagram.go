package hash

// https://leetcode.com/problems/valid-anagram/

func isAnagram(first string, second string) bool {
	if len(first) != len(second) {
		return false
	}

	hash := make(map[rune]int)
	hash2 := make(map[rune]int)

	for _, val := range first {
		hash[val] += 1
	}

	for _, val := range second {
		if _, ok := hash[val]; !ok {
			return false
		}
		hash2[val] += 1
	}

	for _, val := range first {
		if hash[val] != hash2[val] {
			return false
		}
	}

	return true
}
