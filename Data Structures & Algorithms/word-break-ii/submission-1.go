func wordBreak(s string, wordDict []string) []string {
	result := []string{}
	path := []string{}
	var dfs func(int)
	dfs = func(start int) {
		if start == len(s) {
			result = append(result, strings.Join(path, " "))
			return
		}
		for end := start; end < len(s); end++ {
			for _, word := range wordDict {
				diff := end + 1 - len(word)
				if diff == start && s[diff:end + 1] == word {
					path = append(path, word)
					dfs(end + 1)
					path = path[:len(path) - 1]
				}
			}
		}
	}
	dfs(0)
	return result
}
