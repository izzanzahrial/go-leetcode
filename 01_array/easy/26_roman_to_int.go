package array

// https://leetcode.com/problems/roman-to-integer/description/
func romanToInt(s string) int {
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var result int
	for i := 0; i < len(s); i++ {
		// if the next char is bigger than the current char
		// then the current char is subtracted from the result
		// e.g. 'I' is the current char then 'V' is the next char
		// where you need to account for the number '4'
		// so the result will be we minus 1 first from the result
		// and then add the value of 'V' (5) to the result, so we get '4'
		if i+1 < len(s) && romanMap[string(s[i])] < romanMap[string(s[i+1])] {
			result = result - romanMap[string(s[i])]
		} else {
			result = result + romanMap[string(s[i])]
		}
	}

	return result
}
