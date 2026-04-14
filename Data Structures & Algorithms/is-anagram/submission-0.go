func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sRunes, tRunes := []rune(s), []rune(t)
	sort.Slice(sRunes, func(i, j int) bool {
		return sRunes[i] < sRunes[j]
	})
	sort.Slice(tRunes, func(i, j int) bool {
		return tRunes[i] < tRunes[j]
	})

	for i := 0; i < len(s); i++ {
		if sRunes[i] != tRunes[i] {
			return false
		}
	}
	return true
}
