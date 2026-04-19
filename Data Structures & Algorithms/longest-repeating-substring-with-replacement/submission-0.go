func characterReplacement(s string, k int) int {
	maxLen := 0
	start := 0
	chrs := []rune(s)
	count := make(map[rune]int)
	maxFreq := 0

	for end, c := range chrs {
		count[c]++
		maxFreq = max(maxFreq, count[c])

		if end - start + 1 > maxFreq + k {
			count[chrs[start]]--
			start++
		}
		maxLen = max(maxLen, end - start + 1)
	}
	return maxLen
}
