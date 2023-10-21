package slidingwindow

func MaxVowels(s string, k int) int {
	var max int

	for i := 0; i < len(s)-k+1; i++ {
		var curr int
		j := i
		l := k

		for l > 0 {
			if isVowel(string(s[j])) {
				curr++
			}

			j++
			l--
		}

		if max < curr {
			max = curr
		}
	}

	return max
}

func isVowel(s string) bool {
	switch s {
	case "a":
		return true
	case "i":
		return true
	case "u":
		return true
	case "e":
		return true
	case "o":
		return true
	default:
		return false
	}
}
