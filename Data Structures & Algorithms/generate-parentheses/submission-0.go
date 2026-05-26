func generateParenthesis(n int) []string {
	result := []string{}
	var dfs func (string, int, int)
	dfs = func(s string, open, close int) {
		if len(s) == 2 * n {
			result = append(result, s)
		}
		if open < n {
			dfs(s+"(", open + 1, close)
		}
		if close < open {
			dfs(s + ")", open, close + 1)
		}
	}
	dfs("", 0, 0)
	return result
}
