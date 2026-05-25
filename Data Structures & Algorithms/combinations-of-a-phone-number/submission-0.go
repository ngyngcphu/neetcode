func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	phone := map[byte]string {
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno",
		'7': "pqrs", '8': "tuv", '9': "wxyz",
	}
	result := []string{}

	var dfs func(string, int)
	dfs = func(path string, index int) {
		if index == len(digits) {
			result = append(result, path)
			return
		}
		for _, p := range phone[digits[index]] {
			dfs(path + string(p), index + 1)
		}
	}

	dfs("", 0)
	return result
}
