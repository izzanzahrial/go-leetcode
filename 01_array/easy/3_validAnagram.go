package array

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

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := make(map[rune]int, len(s))
	tMap := make(map[rune]int, len(t))

	for _, char := range s {
		if _, exists := sMap[char]; exists {
			sMap[char]++
		} else {
			sMap[char] = 1
		}
	}

	for _, char := range t {
		if _, exists := tMap[char]; exists {
			tMap[char]++
		} else {
			tMap[char] = 1
		}
	}

	for key, sVal := range sMap {
		if tVal := tMap[key]; tVal != sVal {
			return false
		}
	}

	return true
}
