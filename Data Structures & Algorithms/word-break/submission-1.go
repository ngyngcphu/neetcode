func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s) + 1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for _, word := range wordDict {
			j := i - len(word)
			if j >= 0 && dp[j] && s[j:i] == word {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
