func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := [26]int{}
	for i := 0; i < len(s); i++ {
		count[rune(s[i]) - 'a']++
		count[rune(t[i]) - 'a']--
	}

	for _, val := range count {
		if val != 0 {
			return false
		}
	}
	return true
}

