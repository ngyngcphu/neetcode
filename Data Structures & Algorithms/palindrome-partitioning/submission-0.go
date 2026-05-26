func partition(s string) [][]string {
	result := [][]string{}
	path := []string{}
	var dfs func(int)
	dfs = func(start int) {
		if start == len(s) {
			result = append(result, append([]string{}, path...))
			return
		}
		for end := start; end < len(s); end++ {
			if isPalindrome(s, start, end) {
				path = append(path, s[start:end+1])
				dfs(end + 1)
				path = path[:len(path) - 1]
			}
		}
	}
	dfs(0)
	return result
}

func isPalindrome(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
