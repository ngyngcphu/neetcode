func restoreIpAddresses(s string) []string {
	path := []string{}
	result := []string{}
	var dfs func(int)
	dfs = func(start int) {
		if start == len(s) && len(path) == 4 {
			result = append(result, strings.Join(path, "."))
			return
		}
		for end := start; end < len(s); end++ {
			segment := s[start:end+1]
			if len(segment) > 3 {
				break
			}
			if len(segment) > 1 && segment[0] == '0' {
				break
			}
			val, _ := strconv.Atoi(segment)
			if val > 255 {
				break
			}
			path = append(path, segment)
			dfs(end + 1)
			path = path[:len(path) - 1]
		}
	}
	dfs(0)
	return result
}
