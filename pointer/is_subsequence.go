package pointer

func IsSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	pointS := 0
	for _, val := range t {
		if string(val) == string(s[pointS]) {
			pointS++
			if pointS == len(s) {
				return true
			}
		}
	}

	return false
}
