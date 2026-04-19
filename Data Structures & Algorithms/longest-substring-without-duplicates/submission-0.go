func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	start := 0
	count := make(map[rune]int)
	characters := []rune(s)

	for end, c := range characters {
		count[c]++
		for count[c] > 1 {
			count[characters[start]]--
			start++
		}
		maxLen = max(maxLen, end - start + 1)
	}
	return maxLen
}
