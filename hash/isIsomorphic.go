package hash

func IsIsomorphic(s string, t string) bool {
	sMap := make(map[byte]byte, len(s))
	tMap := make(map[byte]byte, len(s))

	for i, _ := range s {
		_, sOK := sMap[s[i]]
		_, tOK := tMap[t[i]]

		if !sOK && !tOK {
			// if char isn't found in both maps
			sMap[s[i]] = t[i]
			tMap[t[i]] = s[i]

		} else if sOK {
			//if found in sMap but the char isn't same as existing
			if sMap[s[i]] != t[i] {
				return false
			}

		} else if tOK {
			//if found in tMap but the char isn't same as existing
			if tMap[t[i]] != s[i] {
				return false
			}
		}
	}

	return true
}
