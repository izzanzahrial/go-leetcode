package slidingwindow

// https://leetcode.com/problems/longest-repeating-character-replacement/
func characterReplacement(s string, k int) int {
	freqMap := make(map[byte]int)
	result, leftIdx, maxFreq := 0, 0, 0

	for rightIdx := 0; rightIdx < len(s); rightIdx++ {
		freqMap[s[rightIdx]]++
		maxFreq = max(maxFreq, freqMap[s[rightIdx]])

		// (rightIdx - leftIdx + 1) - maxFreq = count of character that need to be replaced
		for (rightIdx-leftIdx+1)-maxFreq > k {
			freqMap[s[leftIdx]]--
			leftIdx++
		}

		result = max(result, (rightIdx - leftIdx + 1))
	}

	return result
}
