func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	
	memo := make(map[int]int)
	var helper func(idx int) int
	helper = func(idx int) int {
		if idx == 0 || idx == 1 {
			return 1
		}
		if val, ok := memo[idx]; ok {
			return val
		}
		if s[idx - 1] != '0' {
			memo[idx] = helper(idx - 1)
		}
		twoDigit, _ := strconv.Atoi(s[idx-2:idx])
		if twoDigit >= 10 && twoDigit <= 26 {
			memo[idx] += helper(idx - 2)
		}
		return memo[idx]
	}
	return helper(len(s))
}
