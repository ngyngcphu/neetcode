func wordBreak(s string, wordDict []string) []string {
	result := []string{}
	path := []string{}
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}
	var	dfs func(int)
	dfs = func(start int) {
		if start == len(s) {
			result = append(result, strings.Join(path, " "))
			return
		}
		for end := start; end < len(s); end++ {
			if wordSet[s[start:end+1]] {
				path = append(path, s[start:end+1])
				dfs(end + 1)
				path = path[:len(path) - 1]
			}
		}
	}
	dfs(0)
	return result
}
